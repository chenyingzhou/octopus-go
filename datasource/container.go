package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/chenyingzhou/octopus-go/config"
	"github.com/elastic/go-elasticsearch/v7"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type container struct {
	mysql         map[string]connectionPair
	elasticSearch map[string]connectionPair
}

type connectionPair struct {
	connection interface{}
	config     config.ConnectionConfig
}

var c = new(container)
var initLck sync.Mutex

func GetMysqlClient(name string) (*sql.DB, error) {
	var err error
	_, ok := c.mysql[name]
	if !ok {
		err = initMysqlClient(name)
	}
	cp, ok := c.mysql[name]
	conn := cp.connection
	return conn.(*sql.DB), err
}

func initMysqlClient(name string) error {
	initLck.Lock()
	defer initLck.Unlock()
	_, ok := c.mysql[name]
	if ok {
		return nil
	}
	connConfig, ok := config.Config.DataSourceConfigMap.Mysql[name]
	if !ok {
		return errors.New("Mysql config not exists, name: " + name)
	}
	dataSourceUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s", connConfig.Username, connConfig.Password, connConfig.Host, connConfig.Database)
	DB, err := sql.Open("mysql", dataSourceUrl)
	if err != nil {
		return err
	}
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		return errors.New("Mysql connect fail, name: " + name)
	}
	if c.mysql == nil {
		c.mysql = make(map[string]connectionPair)
	}
	c.mysql[name] = connectionPair{connection: DB, config: connConfig}
	return nil
}

func GetESClient(name string) (*elasticsearch.Client, error) {
	// TODO:
	return nil, nil
}

func initESClient(name string) error {
	// TODO:
	return nil
}
