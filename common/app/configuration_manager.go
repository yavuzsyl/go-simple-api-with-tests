package app

import "go-product-app/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{
		PostgreSqlConfig: ConfigPostgreSql(),
	}
}

func ConfigPostgreSql() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Database:              "productapp",
		User:                  "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
