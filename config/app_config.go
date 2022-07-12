package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	DataSourceConfigMap DataSourceConfigMap `yaml:"datasource"`
}

type DataSourceConfigMap struct {
	Mysql         map[string]ConnectionConfig `yaml:"mysql"`
	ElasticSearch map[string]ConnectionConfig `yaml:"elastic-search"`
}

type ConnectionConfig struct {
	Listen    bool     `yaml:"listen"`
	Host      string   `yaml:"host"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	Database  string   `yaml:"database"`
	Addresses []string `yaml:"addresses"`
}

var Config = new(AppConfig)

func init() {
	yamlFile, err := ioutil.ReadFile("config/app.yaml")
	if err != nil {
		log.Fatalf("Read config file fail #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("Unmarshal config fail #%v", err)
	}
}
