package model

import "gorm.io/gorm"

type Tag struct {
	ID        uint        `gorm:"primary_key" json:"id"`

	Name      string

	gorm.Model
}

type NiceTag struct {
	ID        uint        `gorm:"primary_key" json:"id"`

	TagID     uint32 `json:"tag_id"`

	ArticleID uint32 `json:"article_id"`

	*gorm.Model
}


func (t Tag) TableName() string {
	return "tags"
}

func (t NiceTag) TableName() string {
	return "nice_tags"
}