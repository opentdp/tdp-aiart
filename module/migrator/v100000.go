package migrator

import (
	"github.com/open-tdp/go-helper/dborm"

	"tdp-aiart/model"
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
