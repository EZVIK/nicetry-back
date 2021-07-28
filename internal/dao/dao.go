package dao

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

type ThumbsUp_ interface {
	LikeAdd(db *gorm.DB) error
}

type Repository interface {
	Create(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
	Get(db *gorm.DB) error

	//Gets(db *gorm.DB, column, value string, pageSize int, pageIndex int) (nl []model.NiceList, err error)
	//ViewAdd(db *gorm.DB) error
	//GetComments(db *gorm.DB) (cos []model.Comment, err error)
	//GetNiceTags(db *gorm.DB) (tags []model.Tag, err error)
}

type Dao struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func New(engine *gorm.DB, r *redis.Client) *Dao {
	return &Dao{DB: engine, Cache: r}
}

func (d *Dao) BeginTx() *gorm.DB {
	return d.DB.Begin()
}

func (d *Dao) CommitTx(db *gorm.DB) {
	defer func() {
		db.Commit()
	}()
}
