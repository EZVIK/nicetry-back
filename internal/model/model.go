package model

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"nicetry/global"
	"nicetry/pkg/setting"
	"nicetry/pkg/utils"
	"strconv"
	"time"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
	)

	gormClient, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	//if err := gormClient.AutoMigrate(Nice{}, NiceTag{}, Tag{}, User{}, ReferralCode{}, Comment{}, Notification{}, PointLog{}, ThumbsUp{}, Tag{}); err != nil {
	//	log.Fatalf("models.AutoMigrate err: %v", err)
	//}

	sqlDB, err := gormClient.DB()             // 使用 database/sql 维护连接池
	sqlDB.SetConnMaxIdleTime(5 * time.Second) // 设置 空闲连接的存活时间
	sqlDB.SetConnMaxLifetime(3 * time.Second) // 设置 连接可复用的最大时间
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return gormClient.Debug(), nil
}

func NewCacheEngine(cacheSetting *setting.CacheSettingS) (*redis.Client, error) {

	rClient := redis.NewClient(&redis.Options{
		Addr:     cacheSetting.Host,
		Password: cacheSetting.Password, // no password set
		DB:       0,                     // use default DB
	})

	str, err := rClient.Ping().Result()

	if err != nil {
		global.Logger.Info(str)
		panic(err)
	}

	return rClient, nil
}

func PageChecker(pageIndex string, pageSize string) (pi int, ps int) {

	pi, err1 := strconv.Atoi(pageIndex)
	ps, err2 := strconv.Atoi(pageSize)

	if err1 != nil {
		pi = 1
	}

	if err2 != nil {
		ps = 10
	}

	return pi, ps
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func NewModel() gorm.Model {
	t, _ := utils.GetNowTimeCST()
	return gorm.Model{CreatedAt: t, UpdatedAt: t}
}

func IModel(id uint, createdTime, updatedTime time.Time) *gorm.Model {
	t, _ := utils.GetNowTimeCST()

	return &gorm.Model{ID: id, CreatedAt: t, UpdatedAt: t}

}

type Order struct {
	gorm.Model
	FuckerID uint
	Fucker   Fucker
	Price    float64
}

type Fucker struct {
	gorm.Model
	Username string
	//Orders   []Order
}

func (o *Order) Get(db *gorm.DB) (os []Order, err error) {

	if err = db.Preload("Fucker").Find(&os).Error; err != nil {
		return
	}

	return
}
