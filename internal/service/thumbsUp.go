package service

import (
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
	thumbsUp := model.ThumbsUp{PostId: postId, LikeType: likeType, UserId: userId}

	thu, err := thumbsUp.GetPost(d)

	if err != nil {
		return err
	}

	// 点赞
	if err := thumbsUp.Like(thu, d, cache); err != nil {
		return err
	}

	if err := thumbsUp.Create(d); err != nil {
		return err
	}

	// SAVE User likes in redis

	return nil
}
