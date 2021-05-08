package model

import "gorm.io/gorm"

type ReferralCode struct {
	ID uint `gorm:"primarykey"`

	From uint `json:"from"`

	To uint `json:"to"`

	Code string `gorm:"varchar;unique" json:"code"`

	Used bool `json:"used"`

	*gorm.Model
}

func (rc *ReferralCode) Get(db *gorm.DB) error {
	return db.Where("code = ?", rc.Code).First(&rc).Error
}

func (rc *ReferralCode) ConsumeCode(db *gorm.DB) int64 {
	return db.Model(&rc).Updates(rc).RowsAffected
}
