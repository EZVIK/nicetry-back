package dao

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)


type Dao struct {
	DB *gorm.DB
	Cache *redis.Client
}

func New(engine *gorm.DB, r *redis.Client) *Dao {
	return &Dao{DB: engine, Cache: r}
}


func (d *Dao) BeginTx() (*gorm.DB) {
	return d.DB.Begin()
}

func (d *Dao) CommitTx(db *gorm.DB) {
	defer func() {
		db.Commit()
	}()
}