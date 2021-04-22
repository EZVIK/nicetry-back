package model

import "gorm.io/gorm"

type ThumbsUp struct {

	ID      		uint 			`gorm:"primarykey"`

	PostId			uint			`json:"post_id"`

	LikeType		uint			`json:"like_type"`				// 1 nice 2 comment
	
	UserId          uint 			`json:"user_id"`

	*gorm.Model
}

type ThumbsUp_ interface {
	LikeAdd(db *gorm.DB) error
}

func (l *ThumbsUp) GetLikes()  {

}

func (l *ThumbsUp) Add(db *gorm.DB) error {
	return db.Create(&l).Error
}

func (l *ThumbsUp) Delete(db *gorm.DB) error {
	return db.Delete(&l).Error
}