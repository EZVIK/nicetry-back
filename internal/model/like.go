package model

import "gorm.io/gorm"

type Like struct {

	ID      		uint 			`gorm:"primarykey"`

	PostId			uint			`json:"post_id"`

	LikeType		uint			`json:"like_type"`				// 1 nice 2 comment
	
	UserId          uint 			`json:"user_id"`

	*gorm.Model
}

type ThumbsUp interface {
	LikeAdd(db *gorm.DB) error
}

func (l *Like) GetLikes()  {

}

func (l *Like) Add(db *gorm.DB) error {
	return db.Create(&l).Error
}

func (l *Like) Delete(db *gorm.DB) error {
	return db.Delete(&l).Error
}