package model

import "gorm.io/gorm"

type Comment struct {
	ID       uint   `gorm:"primarykey"`
	NiceId   uint   `json:"nice_id"`
	UserId   uint   `json:"user_id"`
	Content  string `json:"content"`
	ThumbsUp int64  `json:"thumbs_up"`
	*gorm.Model
}

func (n *Comment) LikeAdd(db *gorm.DB) error {
	return db.Model(&n).Update("thumbs_up", gorm.Expr("thumbs_up + ?", "", 1)).Error
}

func (n *Comment) Add(db *gorm.DB) error {
	return db.Create(&n).Error
}
