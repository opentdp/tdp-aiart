package migrator

import (
	"tdp-aiart/module/dborm"
	"tdp-aiart/module/model"
)

func v100000() error {

	return v100000AutoMigrate()

}

func v100000AutoMigrate() error {

	return dborm.Db.AutoMigrate(
		&model.Artimg{},
		&model.Config{},
		&model.Migration{},
		&model.User{},
	)

}
