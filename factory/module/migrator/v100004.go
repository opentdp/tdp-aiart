package migrator

import (
	"tdp-aiart/module/model/config"
)

func v100004() error {

	if isMigrated("v100004") {
		return nil
	}

	if err := v100004AddConfig(); err != nil {
		return err
	}

	return addMigration("v100004", "添加配额参数")

}

func v100004AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "ArtworkQuota",
			Value:       "10",
			Type:        "number",
			Module:      "system",
			Description: "初始绘画配额",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}
