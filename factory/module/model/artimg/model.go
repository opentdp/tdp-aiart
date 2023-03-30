package artimg

import (
	"tdp-aiart/module/dborm"
)

// 创建配置

type CreateParam struct {
	UserId     uint `binding:"required"`
	Subject    any  `binding:"required"`
	OutputFile string
	Status     string
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Artimg{
		UserId:     data.UserId,
		Subject:    data.Subject,
		OutputFile: data.OutputFile,
		Status:     data.Status,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新配置

type UpdateParam struct {
	Id         uint
	UserId     uint
	Subject    any
	OutputFile string
	Status     string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Artimg{
			Id: data.Id,
		}).
		Updates(dborm.Artimg{
			UserId:     data.UserId,
			Subject:    data.Subject,
			OutputFile: data.OutputFile,
			Status:     data.Status,
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
		Where(&dborm.Artimg{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Artimg{})

	return result.Error

}

// 获取配置

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Artimg, error) {

	var item *dborm.Artimg

	result := dborm.Db.
		Where(&dborm.Artimg{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	UserId uint
}

func FetchAll(data *FetchAllParam) ([]*dborm.Artimg, error) {

	var items []*dborm.Artimg

	result := dborm.Db.
		Where(&dborm.Artimg{
			UserId: data.UserId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Artimg{}).
		Where(&dborm.Artimg{
			UserId: data.UserId,
		}).
		Count(&count)

	return count, result.Error

}
