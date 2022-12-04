package hidb

import (
	"fmt"
	"github.com/tradjick/hiconfig"
)

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Schema   string
}

func (dbc *dbConfig) AsDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%01d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbc.Username,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.Schema)
}

func fetchConfig(p string, l hiconfig.ConfigLoader) *dbConfig {
	c := new(dbConfig)
	l(p, c)
	return c
}
