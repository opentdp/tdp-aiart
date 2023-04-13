package migrator

import (
	"tdp-aiart/model/config"
)

func v100002() error {

	if isMigrated("v100002") {
		return nil
	}

	if err := v100002AddConfig(); err != nil {
		return err
	}

	return addMigration("v100002", "添加系统参数")

}

func v100002AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "Registrable",
			Value:       "true",
			Type:        "bool",
			Module:      "system",
			Description: "允许注册",
		},
		{
			Name:        "SecretId",
			Value:       "",
			Type:        "string",
			Module:      "Tencent",
			Description: "腾讯云 SecretId",
		},
		{
			Name:        "SecretKey",
			Value:       "",
			Type:        "string",
			Module:      "Tencent",
			Description: "腾讯云 SecretKey",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}