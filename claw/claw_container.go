package claw

import (
	"database/sql"
	"encoding/json"
	_ "github.com/chenyingzhou/octopus-go/config"
	"github.com/chenyingzhou/octopus-go/consts"
	"github.com/chenyingzhou/octopus-go/datasource"
	"log"
	"time"
)

var Container = make(map[int32]Config)

func init() {
	initContainer()
	go func() {
		for {
			initContainer()
			time.Sleep(5 * 60 * time.Second)
		}
	}()
}

func initContainer() {
	db, err := datasource.GetMysqlClient("app")
	if err != nil {
		log.Fatalf("Connect to database fail")
	}
	rows, err := db.Query("SELECT * FROM `claw_config`")
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	if err != nil {
		log.Printf("Read claw config fail, #%v", err)
		return
	}
	var id int32
	var name string
	var sourceTree string
	var targetType string
	var targetSource string
	var targetSet string
	var comments string
	var deleted bool
	var createdTime time.Time
	var updatedTime time.Time
	for rows.Next() {
		err = rows.Scan(&id, &name, &sourceTree, &targetType, &targetSource, &targetSet, &comments, &deleted, &createdTime, &updatedTime)
		if err != nil {
			log.Printf("解析配置失败, #%v", err)
			continue
		}
		st := new(SourceTree)
		err = json.Unmarshal([]byte(sourceTree), st)
		if err != nil {
			log.Printf("解析配置失败, #%v", err)
			continue
		}
		config := Config{id, name, st, consts.SourceType(targetType), targetSource, targetSet, comments, deleted, createdTime, updatedTime}
		Container[id] = config
	}
}
