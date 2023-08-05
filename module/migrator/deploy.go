package migrator

import (
	"github.com/opentdp/go-helper/logman"

	"tdp-aiart/model/migration"
)

func Deploy() {

	if err := doMigrate(); err != nil {
		logman.Fatal("migrate database failed", "error", err)
	}

}

func addMigration(k, v string) error {

	_, err := migration.Create(&migration.CreateParam{
		Version: k, Description: v,
	})

	return err

}

func isMigrated(k string) bool {

	rq := &migration.FetchParam{Version: k}

	if rs, err := migration.Fetch(rq); err == nil {
		return rs.Id > 0
	}

	return false

}

func doMigrate() error {

	funcs := []func() error{
		v100000,
		v100001,
		v100002,
		v100003,
		v100004,
		v100005,
		v100006,
	}

	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil

}
