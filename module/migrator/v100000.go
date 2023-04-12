package migrator

import (
	"tdp-aiart/model"
	"tdp-aiart/module/dborm"
)

func v100000() error {

	return v100000AutoMigrate()

}

func v100000AutoMigrate() error {

	return dborm.Db.AutoMigrate(
		&model.Artwork{},
		&model.Config{},
		&model.Migration{},
		&model.User{},
	)

}
