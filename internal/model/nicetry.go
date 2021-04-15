package model

import (
	"gorm.io/gorm"

)

type Nice struct {

	Title   string `gorm:"varchar(50);unique" json:"title"` 		// 标题

	Desc    string `gorm:"varchar(200)" json:"desc"`				// 详情

	Content string `gorm:"type:text" json:"content"`				// 内容

	View    int64  `gorm:"default:0" json:"view"`					// 被查看数量

	Like    int64  `gorm:"default:0" json:"like"`					// 被点赞数量

	UserId  uint    `gorm:"index" json:"user_id"`					// 创建人

	NiceType uint 	`json:"nice_type`								// 类型 0 challenge 1 happy hour

	Tags    []Tag  `gorm:"many2many:nice_tags;"`

	gorm.Model

}

func (n Nice) TableName() string {
	return "nice"
}