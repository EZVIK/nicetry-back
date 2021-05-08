package model

import "gorm.io/gorm"

type ITag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID uint `gorm:"primarykey"`

	Name string `gorm:"varchar(50);unique" json:"name"`

	ParentId uint

	*gorm.Model
}

func (ITag) TableName() string {
	return "tags"
}
func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) Get(db *gorm.DB) error {
	return db.Model(&t).Error
}

func (t *Tag) Add(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t *Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}

func GetTags(db *gorm.DB, ids []uint) (t []Tag, err error) {

	if err = db.Find(&t, ids).Error; err != nil {
		return t, err
	}

	return t, nil
}

type NiceTag struct {
	ID uint `gorm:"primary_key" json:"id"`

	TagID uint `json:"tag_id"`

	NiceID uint `json:"nice_id"`

	*gorm.Model
}
