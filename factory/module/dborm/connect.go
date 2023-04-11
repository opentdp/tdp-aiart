package dborm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"tdp-aiart/cmd/args"
	"tdp-aiart/helper/logman"
)

var Db *gorm.DB

func Connect() {

	config := &gorm.Config{
		Logger: newLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if db, err := gorm.Open(dialector(), config); err != nil {
		logman.Fatal("Connect to databse failed", "error", err)
	} else {
		Db = db
	}

}

func dialector() gorm.Dialector {

	switch args.Database.Type {
	case "sqlite":
		return useSqlite()
	case "mysql":
		return useMysql()
	default:
		logman.Fatal("Database type error", "type", args.Database.Type)
	}

	return nil

}
