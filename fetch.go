package hidb

import (
	"github.com/tradjick/hiconfig"
	"gorm.io/gorm"
)

var dbConfigPath string
var dbConfigLoader hiconfig.ConfigLoader

func SetDBConfigPath(configPath string, configLoader hiconfig.ConfigLoader) {
	dbConfigPath = configPath
	dbConfigLoader = configLoader
}

func FetchDB() *gorm.DB { //cached connection
	if db == nil {
		if dbConfigPath == "" {
			panic("db config path not set")
		}
		connect(dbConfigPath, dbConfigLoader)
	}
	return db
}
