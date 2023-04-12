package model

// 图片

type Artwork struct {
	Id             uint   `gorm:"primaryKey"`
	UserId         uint   `gorm:"index"`
	Subject        string `gorm:"size:128"`
	Prompt         string `gorm:"size:1024"`
	NegativePrompt string `gorm:"size:1024"`
	Styles         string `gorm:"size:256"`
	Strength       float64
	InputFile      string `gorm:"size:256"`
	OutputFile     string `gorm:"size:256"`
	Favorites      uint
	Recommends     uint
	Status         string `gorm:"size:32"`
	CreatedAt      int64
	UpdatedAt      int64
}

// 系统配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:32;uniqueIndex:idx_config"`
	Value       string `gorm:"size:1024"`
	Type        string `gorm:"size:32;default:string"`
	Module      string `gorm:"size:32;uniqueIndex:idx_config"`
	Description string `gorm:"size:1024"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 迁移记录

type Migration struct {
	Id          uint   `gorm:"primaryKey"`
	Version     string `gorm:"size:64;uniqueIndex"`
	Description string `gorm:"size:1024"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 用户

type User struct {
	Id           uint   `gorm:"primaryKey"`
	Username     string `gorm:"size:64;uniqueIndex"`
	Password     string `gorm:"size:128" json:"-"`
	Level        uint   `gorm:"default:5"`
	AppId        string `gorm:"size:128;uniqueIndex"`
	AppKey       string `gorm:"size:128" json:"-"`
	Email        string `gorm:"size:255;uniqueIndex"`
	Description  string `gorm:"size:1024;default:挥一挥手"`
	ArtworkQuota uint   `gorm:"size:32;default:3"`
	CreatedAt    int64
	UpdatedAt    int64
}
