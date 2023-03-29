package migrator

import (
	"tdp-aiart/module/model/config"
)

func v100002() error {

	if isMigrated("v100002") {
		return nil
	}

	if err := v100002AddConfig(); err != nil {
		return err
	}

	return addMigration("v100002", "添加默认参数")

}

func v100002AddConfig() error {

	list := []config.CreateParam{
		{
			Name:        "secretId",
			Value:       "",
			Module:      "Tencent",
			Description: "腾讯云 SecretId",
		},
		{
			Name:        "secretKey",
			Value:       "",
			Module:      "Tencent",
			Description: "腾讯云 SecretKey",
		},
	}

	for _, item := range list {
		_, err := config.Create(&item)
		if err != nil {
			return err
		}
	}

	return nil

}
