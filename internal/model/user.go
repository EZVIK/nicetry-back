package model

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"nicetry/global"
	"nicetry/pkg/utils"
	"time"
)

type LoginLog struct {
	ID     uint `gorm:"primarykey"`
	UserId uint `json:"user_id"`
	*gorm.Model
}

type IUser struct {
	ID       uint   `gorm:"primarykey"`
	Mail     string `gorm:"varchar;unique" json:"mail"`
	Nickname string `gorm:"varchar;" json:"nickname"`
	Avatar   string `gorm:"varchar;" json:"avatar"`
}

type User struct {
	ID           uint    `gorm:"primarykey"`
	Nickname     string  `gorm:"varchar;" json:"nickname"`
	Mail         string  `gorm:"varchar;unique" json:"mail"`
	Password     string  `gorm:"varchar;" json:"password"`
	Avatar       string  `gorm:"varchar;" json:"avatar"`
	Desc         string  `gorm:"varchar;" json:"desc"`
	RecommendBy  uint    `json:"recommend_by"`
	Link         string  `gorm:"varchar;" json:"link"`
	Points       float64 `gorm:"default:0.0" json:"points"`
	EnableStatus bool    `gorm:"default:false" json:"enable_status"`
	*gorm.Model
}

func (IUser) TableName() string {
	return "users"
}
func (User) TableName() string {
	return "users"
}

func (u *User) GetCachePrefix() string {
	return fmt.Sprintf("user: %v", u.ID)
}

// create
func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

// update
func (u *User) Update(db *gorm.DB) error {
	return db.Updates(&u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(&u).Error
}

func (u *User) Get(db *gorm.DB) error {
	return db.Debug().First(&u).Error
}

// login
func (u *User) Login(db *gorm.DB) error {
	return db.Where("mail = ? ", u.Mail).Find(&u).Error
}

//
func (u *User) CheckColumn(db *gorm.DB, columnName, value string) (c int64, err error) {
	err = db.Debug().Model(&User{}).Where("(?) = ?", columnName, value).Count(&c).Error
	//err = db.Debug().Model(&User{}).QueryFields()

	if err != nil {
		return 0, err
	}
	return c, nil
}

//
func (u *User) Token(r *redis.Client, token string) error {
	return r.Set(GetCachePreName(u), token, time.Hour*36).Err()
}

// get likes record
func (u *User) GetLikes(db *gorm.DB) (ls []ThumbsUp, err error) {
	return nil, nil
}

// get points record
func (u *User) GetPoints(db *gorm.DB) (p int64, err error) {
	return 0, nil
}

// get referral_code list
func (u *User) GetReferCodes(db *gorm.DB) (codes []ReferralCode, err error) {
	err = db.Debug().Model(&ReferralCode{}).Where("`from` = ? ", u.ID).Find(&codes).Error
	return codes, err
}

// generate referral code
func (u *User) CreateReferCode(db *gorm.DB) error {

	code := utils.RandStringBytesMask(5)

	r := ReferralCode{Code: code, From: u.ID, Used: false}

	err := db.Create(&r).Error

	if err != nil {
		global.Logger.Info(err)
	}

	return err
}

// get users avatar by ids
func (u *User) GetUsersAvatar(db *gorm.DB, ids []uint) []IUser {
	iu := []IUser{}

	db.Find(&iu, "id in ?", ids)

	return iu
}
