package model

import "gorm.io/gorm"

// Questionnaire

type Questionnaire struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	NoNumber  string
	UserId    uint
	User      IUser
	Questions []Question
	Tags      []Tag
	*gorm.Model
}

type Question struct {
	ID            uint `gorm:"primarykey"`
	Title         string
	Desc          string
	Selection     []Selection
	SelectionType uint   // 0 single 1 muti 3 bool
	CorrectAnswer []uint //
	UserId        uint
	*gorm.Model
}

type Selection struct {
	ID    uint `gorm:"primarykey"`
	Title string
	*gorm.Model
}
