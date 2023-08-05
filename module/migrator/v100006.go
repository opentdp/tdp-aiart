package migrator

import (
	"tdp-aiart/cmd/args"

	"github.com/opentdp/go-helper/dborm"
)

func v100006() error {

	if isMigrated("v100006") {
		return nil
	}

	if err := v100006Update(); err != nil {
		return err
	}

	return addMigration("v100006", "更新AiArt配置")

}

func v100006Update() error {

	if args.Database.Type == "mysql" {
		v100006Mysql()
	} else if args.Database.Type == "sqlite" {
		v100006Sqlite()
	}

	return nil

}

func v100006Mysql() {

	dborm.Db.Exec("UPDATE artwork SET input_file = CONCAT('/upload/', input_file) WHERE input_file != '' AND input_file NOT LIKE '/upload/%'")
	dborm.Db.Exec("UPDATE artwork SET output_file = CONCAT('/upload/', output_file) WHERE output_file != '' AND output_file NOT LIKE '/upload/%'")

}

func v100006Sqlite() {

	dborm.Db.Exec("UPDATE artwork SET input_file = '/upload/'|| input_file WHERE input_file != '' AND input_file NOT LIKE '/upload/%'")
	dborm.Db.Exec("UPDATE artwork SET output_file = '/upload/' || output_file WHERE output_file != '' AND output_file NOT LIKE '/upload/%'")

}
