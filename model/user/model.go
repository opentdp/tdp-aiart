package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/open-tdp/go-helper/secure"

	"tdp-aiart/model"
	"tdp-aiart/module/dborm"
)

// 创建用户

type CreateParam struct {
	Username     string `binding:"required"`
	Password     string `binding:"required"`
	Level        uint
	AppKey       string
	Email        string `binding:"required"`
	Description  string
	ArtworkQuota uint
	StoreKey     string // 存储密钥
}

func Create(data *CreateParam) (uint, error) {

	if data.Password != "" {
		pw, err := CreateSecret(data.Password)
		if err != nil {
			return 0, err
		}
		data.Password = pw
	}

	if data.AppKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.AppKey, data.StoreKey)
		if err == nil {
			data.AppKey = secret
		}
	}

	item := &model.User{
		Username:     data.Username,
		Password:     data.Password,
		Level:        data.Level,
		AppId:        uuid.NewString(),
		AppKey:       data.AppKey,
		Email:        data.Email,
		Description:  data.Description,
		ArtworkQuota: data.ArtworkQuota,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新用户

type UpdateParam struct {
	Id           uint
	Username     string
	Password     string
	Level        uint
	Email        string
	AppKey       string
	Description  string
	ArtworkQuota uint
	StoreKey     string // 存储密钥
}

func Update(data *UpdateParam) error {

	if data.Password != "" {
		pw, err := CreateSecret(data.Password)
		if err != nil {
			return err
		}
		data.Password = pw
	}

	if data.AppKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.AppKey, data.StoreKey)
		if err == nil {
			data.AppKey = secret
		}
	}

	result := dborm.Db.
		Where(&model.User{
			Id: data.Id,
		}).
		Updates(model.User{
			Username:     data.Username,
			Password:     data.Password,
			Level:        data.Level,
			Email:        data.Email,
			AppKey:       data.AppKey,
			Description:  data.Description,
			ArtworkQuota: data.ArtworkQuota,
		})

	return result.Error

}

// 更新用户配额

type UpdateQuotaParam struct {
	Id           uint
	ArtworkQuota int
}

func UpdateQuota(data *UpdateQuotaParam) error {

	update := map[string]any{}

	if data.ArtworkQuota > 0 {
		update["artwork_quota"] = gorm.Expr("artwork_quota + ?", data.ArtworkQuota)
	} else if data.ArtworkQuota < 0 {
		update["artwork_quota"] = gorm.Expr("artwork_quota - ?", -data.ArtworkQuota)
	}

	result := dborm.Db.
		Model(&model.User{}).
		Where(&model.User{
			Id: data.Id,
		}).
		Updates(update)

	return result.Error

}

// 删除用户

type DeleteParam struct {
	Id       uint
	Username string
}

func Delete(data *DeleteParam) error {

	var item *model.User

	result := dborm.Db.
		Where(&model.User{
			Id:       data.Id,
			Username: data.Username,
		}).
		Delete(&item)

	return result.Error

}

// 获取用户

type FetchParam struct {
	Id       uint
	Username string
	AppId    string
	Email    string
	StoreKey string // 存储密钥
}

func Fetch(data *FetchParam) (*model.User, error) {

	var item *model.User

	result := dborm.Db.
		Where(&model.User{
			Id:       data.Id,
			Username: data.Username,
			AppId:    data.AppId,
			Email:    data.Email,
		}).
		First(&item)

	if item.AppKey != "" && data.StoreKey != "" {
		item.AppKey, _ = secure.Des3Decrypt(item.AppKey, data.StoreKey)
	}

	return item, result.Error

}

// 获取用户列表

type FetchAllParam struct {
	Level uint
}

func FetchAll(data *FetchAllParam) ([]*model.User, error) {

	var items []*model.User

	result := dborm.Db.
		Where(&model.User{
			Level: data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取用户总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.User{}).
		Where(&model.User{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
