package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	DbUser string
	DbName string
	DbPass string
	DbAddr string
}

type ManageArticalConfig struct {
	ServerConfig ServerConfig
	DbConfig     DbConfig
}

func LoadConfig() *ManageArticalConfig {

	var manager ManageArticalConfig
	cfg, err := ini.Load("config.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	manager.ServerConfig.Port = cfg.Section("server").Key("port").String()

	//load database info
	manager.DbConfig.DbName = cfg.Section("database").Key("db_name").String()
	manager.DbConfig.DbPass = cfg.Section("database").Key("db_password").String()
	manager.DbConfig.DbUser = cfg.Section("database").Key("db_username").String()
	manager.DbConfig.DbAddr = cfg.Section("database").Key("db_address").String()

	return &manager
}
