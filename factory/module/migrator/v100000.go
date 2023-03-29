package migrator

import (
	"tdp-aiart/module/dborm"
)

func v100000() error {

	return dborm.Db.AutoMigrate(
		&dborm.Config{},
		&dborm.User{},
	)

}
