package migrator

import (
	"tdp-aiart/model/config"
)

func v100004() error {

	if isMigrated("v100004") {
		return nil
	}

	if err := v100004AddConfig(); err != nil {
		return err
	}

	return addMigration("v100004", "添加绘画参数")

}

func v100004AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "SecretId",
			Value:       "",
			Type:        "string",
			Module:      "tencent",
			Description: "腾讯云 SecretId",
		},
		{
			Name:        "SecretKey",
			Value:       "",
			Type:        "string",
			Module:      "tencent",
			Description: "腾讯云 SecretKey",
		},
		{
			Name:        "ArtworkQuota",
			Value:       "10",
			Type:        "number",
			Module:      "system",
			Description: "绘画初始配额",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}
