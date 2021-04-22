package service

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"nicetry/internal/model"
)

// 点赞
func (s *Service) Like(postId, likeType, userId uint) error {
	d := s.Dao.DB
	cache := s.Dao.Cache

	// check user info
	user := model.User{ID: userId}
	if err := user.Get(d.Where("id = ?", userId)); err != nil || user.ID == 0 {
		return err
	}

	// using cache check user if like this post
	like := model.ThumbsUp{ PostId: postId, LikeType: likeType, UserId: userId}
	userLikeKey := ""


	if likeType == 1 {  // nice
		userLikeKey = fmt.Sprintf("{USER_NICE_LIKES}:%v", userId)
		if err := Check(cache, userLikeKey, postId); err != nil {
			return err
		}

		nice := model.Nice{ID: postId}
		if err := nice.LikeAdd(d); err != nil {
			return err
		}
	} else if likeType == 2 { // comment
		userLikeKey = fmt.Sprintf("{USER_COMMNET_LIKES}:%v", userId)
		if err := Check(cache, userLikeKey, postId); err != nil {
			return err
		}

		comm := model.Comment{ID: postId}
		if err := comm.LikeAdd(d); err != nil {
			return err
		}
	}

	// like record
	err := like.Add(d)

	if err != nil {
	    return err
	}

	// SAVE User likes in redis
	if err := cache.SAdd(userLikeKey, postId).Err(); err != nil {
		return err
	}

	return nil
}

func Check(cache *redis.Client, userLikeKey string, postId uint) error {

	// 判断 是否已经点赞
	if isMember, _ := cache.SIsMember(userLikeKey, postId).Result(); isMember {
		return errors.New("已经进行过点赞👍")
	}

	return nil
}