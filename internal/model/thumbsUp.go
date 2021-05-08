package model

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"nicetry/global"
)

type ThumbsUp struct {
	ID uint `gorm:"primarykey"`

	PostId uint `json:"post_id"`

	LikeType uint `json:"like_type"` // 1 nice 2 comment

	UserId uint `json:"user_id"`

	*gorm.Model
}

type ThumbsUp_ interface {
	LikeAdd(db *gorm.DB) error
}

func (l *ThumbsUp) Like(t ThumbsUp_, db *gorm.DB, cache *redis.Client) error {

	key := ""
	if l.LikeType == 1 {
		key = fmt.Sprintf("{USER_COMMENT_LIKES}:%v", l.UserId)
	} else {
		key = fmt.Sprintf("{USER_NICE_LIKES}:%v", l.UserId)
	}

	if err := l.Check(cache, key); err != nil {
		return err
	}

	if err := t.LikeAdd(db); err != nil {
		return err
	}

	if err := l.Save(cache, key); err != nil {
		return err
	}

	return nil
}

// 获取POST 对象
func (l *ThumbsUp) GetPost(db *gorm.DB) (tu ThumbsUp_, err error) {

	if l.LikeType == 1 {
		s := &Nice{}
		err = db.Model(s).First(&s, l.PostId).Error
		return s, err
	} else if l.LikeType == 2 {
		s := &Comment{}
		err = db.Model(s).First(&s, l.PostId).Error
		return s, err
	}

	return
}

func (l *ThumbsUp) Add(db *gorm.DB) error {
	return db.Create(&l).Error
}

func (l *ThumbsUp) Delete(db *gorm.DB) error {
	return db.Delete(&l).Error
}

func (l *ThumbsUp) Check(cache *redis.Client, userLikeKey string) error {

	if !global.AppSetting.IfCheckLike {
		return nil
	}

	// 判断 是否已经点赞
	if isMember, _ := cache.SIsMember(userLikeKey, l.PostId).Result(); isMember {
		return errors.New("已经进行过点赞👍")
	}

	return nil
}

func (l *ThumbsUp) Save(cache *redis.Client, key string) error {

	if err := cache.SAdd(key, l.PostId).Err(); err != nil {
		return err
	}

	return nil
}
