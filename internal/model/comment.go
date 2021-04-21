package model

import "gorm.io/gorm"

type Comment struct {
	ID      		uint 			`gorm:"primarykey"`
	NiceId			uint			`json:"nice_id"`
	UserId			uint			`json:"user_id"`
	Content			string			`json:"content"`
	Like			int64 			`json:"like"`
	*gorm.Model

}

func (n *Comment) LikeAdd(db *gorm.DB) error {
	return db.Model(&n).Update("like", gorm.Expr("? + ?", "like", 1)).Error
}

func (n *Comment) Add(db *gorm.DB) error {
	return db.Create(&n).Error
}