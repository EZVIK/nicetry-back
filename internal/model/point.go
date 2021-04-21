package model

import "gorm.io/gorm"

type PointLog struct {

	ID        uint 	`gorm:"primarykey"`

	UserId    uint

	Consume   int64

	*gorm.Model
}