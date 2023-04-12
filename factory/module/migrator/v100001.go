package migrator

import (
	"github.com/open-tdp/go-helper/strutil"

	"tdp-aiart/cmd/args"
	"tdp-aiart/model/user"
)

func v100001() error {

	if isMigrated("v100001") {
		return nil
	}

	if err := v100001AddUser(); err != nil {
		return err
	}

	return addMigration("v100001", "添加默认账号")

}

func v100001AddUser() error {

	_, err := user.Create(&user.CreateParam{
		Username: "admin",
		Password: "123456",
		Level:    1,
		AppKey:   strutil.Rand(32),
		Email:    "admin@opentdp.org",
		StoreKey: args.Dataset.Secret,
	})

	return err

}
