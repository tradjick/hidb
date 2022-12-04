package hidb

import (
	"github.com/tradjick/hiconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connect(configPath string, configLoader hiconfig.ConfigLoader) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fetchConfig(configPath, configLoader).AsDSN(), // data source name
		DefaultStringSize:         256,                                           // default size for string fields
		DisableDatetimePrecision:  true,                                          // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                          // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                          // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                         // autoconfigure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("db connection failed")
	}
}
