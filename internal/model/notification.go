package model

import "gorm.io/gorm"

type Notification struct {

	ID      		uint 			`gorm:"primarykey"`
	
	PostId			uint			`json:"nice_id"`				// 1 niceId 2 commentId

	LikeType		uint			`json:"like_type"`				// 1 nice 2 comment

	UserId          uint 			`json:"user_id"`				
	
	Content         string			`json:"content"`
	
	Readed          bool			`json:"readed"`
	
	*gorm.Model
}
