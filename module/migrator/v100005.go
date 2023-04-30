package migrator

import (
	"tdp-aiart/model/config"
)

func v100005() error {

	if isMigrated("v100005") {
		return nil
	}

	if err := v100005AddConfig(); err != nil {
		return err
	}

	return addMigration("v100005", "添加OpenAI参数")

}

func v100005AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "ApiKey",
			Value:       "",
			Type:        "string",
			Module:      "openai",
			Description: "OpenAI 授权密钥",
		},
		{
			Name:        "ApiUrl",
			Value:       "",
			Type:        "string",
			Module:      "openai",
			Description: "OpenAI 接口地址",
		},
		{
			Name:        "ProxyUrl",
			Value:       "",
			Type:        "string",
			Module:      "openai",
			Description: "OpenAI 接口代理",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}
