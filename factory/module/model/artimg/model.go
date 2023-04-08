package artimg

import (
	"tdp-aiart/module/dborm"
	"tdp-aiart/module/model"
)

// 创建配置

type CreateParam struct {
	UserId         uint `binding:"required"`
	Subject        string
	Prompt         string
	NegativePrompt string
	Styles         string
	Strength       float64
	InputFile      string
	OutputFile     string
	Status         string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Artimg{
		UserId:         data.UserId,
		Subject:        data.Subject,
		Prompt:         data.Prompt,
		NegativePrompt: data.NegativePrompt,
		Styles:         data.Styles,
		Strength:       data.Strength,
		InputFile:      data.InputFile,
		OutputFile:     data.OutputFile,
		Status:         data.Status,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新配置

type UpdateParam struct {
	Id      uint
	UserId  uint
	Subject string
	Status  string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Artimg{
			Id: data.Id,
		}).
		Updates(model.Artimg{
			UserId:  data.UserId,
			Subject: data.Subject,
			Status:  data.Status,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Artimg{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Artimg{})

	return result.Error

}

// 获取配置

type FetchParam struct {
	Id     uint
	UserId uint
	Status string
}

func Fetch(data *FetchParam) (*model.Artimg, error) {

	var item *model.Artimg

	result := dborm.Db.
		Where(&model.Artimg{
			Id:     data.Id,
			UserId: data.UserId,
			Status: data.Status,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	UserId uint
	Status string
}

func FetchAll(data *FetchAllParam) ([]*model.Artimg, error) {

	var items []*model.Artimg

	result := dborm.Db.
		Where(&model.Artimg{
			UserId: data.UserId,
			Status: data.Status,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Artimg{}).
		Where(&model.Artimg{
			UserId: data.UserId,
		}).
		Count(&count)

	return count, result.Error

}
