package driver

import (
	"fmt"
	"github.com/chenyingzhou/octopus-go/claw"
	"github.com/chenyingzhou/octopus-go/config"
	"github.com/chenyingzhou/octopus-go/consts"
	"github.com/chenyingzhou/octopus-go/plate"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/siddontang/go-log/log"
	"strconv"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	if e.Action != "update" && e.Action != "insert" {
		return nil
	}
	ignore := true
	database := e.Table.Schema
	for _, cfg := range config.Config.DataSourceConfigMap.Mysql {
		if cfg.Listen && cfg.Database == database {
			ignore = false
		}
	}
	if ignore {
		return nil
	}

	table := e.Table.Name
	ids := make([]int32, 0, 2)
	for _, row := range e.Rows {
		id, err := strconv.ParseInt(fmt.Sprintf("%v", row[e.Table.PKColumns[0]]), 10, 64)
		if err == nil {
			ids = append(ids, int32(id))
		}
	}

	food := plate.Food{
		SourceType: consts.SourceTypeMysql,
		DataSource: database,
		DataSet:    table,
		Ids:        ids,
	}
	for _, c := range claw.FindConfigs(database, table) {
		c.Handle(food)
	}
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func CanalStart() {
	for _, connectionConfig := range config.Config.DataSourceConfigMap.Mysql {
		if connectionConfig.Listen {
			cfg := canal.NewDefaultConfig()
			cfg.Addr = connectionConfig.Host
			cfg.User = connectionConfig.Username
			cfg.Password = connectionConfig.Password
			cfg.Dump.TableDB = connectionConfig.Database

			c, err := canal.NewCanal(cfg)
			if err != nil {
				log.Fatal(err)
			}
			handler := MyEventHandler{}
			c.SetEventHandler(&handler)
			go c.Run()
		}
	}
}
