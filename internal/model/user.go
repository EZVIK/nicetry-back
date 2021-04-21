package model

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"nicetry/global"
	"nicetry/pkg/utils"
	"time"
)

type User struct {

	ID        			uint 		`gorm:"primarykey"`
	Nickname  			string 		`gorm:"varchar;" json:"nickname"`
	Mail 				string 		`gorm:"varchar;unique" json:"mail"`
	Password     		string 		`gorm:"varchar;" json:"password"`
	Avatar              string 		`gorm:"varchar;" json:"avatar"`
	Desc 				string 		`gorm:"varchar;" json:"desc"`
	RecommendBy  		uint 		`json:"recommend_by"`
	Link 				string 		`gorm:"varchar;" json:"link"`
	Points 				float64 	`gorm:"default:0.0" json:"points"`
	*gorm.Model
}

func (u *User) GetCachePrefix() string {
	return fmt.Sprintf("user: %v", u.ID)
}

func (u *User) Create(db *gorm.DB,) error {
	return db.Create(&u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Updates(&u).Error
}

func (u *User) Delete(db *gorm.DB,) error {
	return db.Delete(&u).Error
}

func (u *User) Get(db *gorm.DB) (error) {
	return db.Debug().First(&u).Error
}

func (u *User) Login(db *gorm.DB) (error) {
	return db.Where( "mail = ? ", u.Mail).Find(&u).Error
}

func (u *User) CheckColumn(db *gorm.DB, columnName, value string) (c int64, err error) {

	err = db.Debug().Model(&User{}).Where("(?) = ?", columnName, value).Count(&c).Error
	//err = db.Debug().Model(&User{}).QueryFields()


	if err != nil {
		return 0, err
	}
	return  c, nil
}

func (u *User) Token(r *redis.Client, token string) error {
	return r.Set(GetCachePreName(u), token, time.Hour * 36).Err()
}

func (u *User) GetLikes(db *gorm.DB) (ls []Like, err error) {
	return nil, nil
}

func (u *User) GetPoints(db *gorm.DB) (p int64, err error){
	return 0, nil
}

func (u *User) GetReferCodes(db *gorm.DB) (codes ReferralCode, err error) {
	err = db.Debug().Model(&ReferralCode{}).Where("from = ?", u.ID).Find(&codes).Error
	return codes, err
}

func (u *User) CreateReferCode(db *gorm.DB) error {

	code := utils.RandStringBytesMask(5)

	r := ReferralCode{Code: code, From: u.ID, Used: false}

	err := db.Create(&r).Error

	if err != nil {
		global.Logger.Info(err)
	}

	return err
}