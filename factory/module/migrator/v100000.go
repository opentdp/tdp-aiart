package migrator

import (
	"tdp-aiart/module/dborm"
)

func v100000() error {

	if err := v100000AutoMigrate(); err != nil {
		return err
	}

	return nil

}

func v100000AutoMigrate() error {

	return dborm.Db.AutoMigrate(
		&dborm.Artimg{},
		&dborm.Config{},
		&dborm.Migration{},
		&dborm.User{},
	)

}
