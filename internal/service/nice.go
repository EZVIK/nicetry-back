package service

import (
	"errors"
	"nicetry/internal/model"
	"nicetry/pkg/utils"
	"strconv"
)

func (s *Service) GetNice(id string) (map[string]interface{}, error) {
	d := s.Dao.DB
	nice := model.Nice{NoNumber: id}

	if err := nice.Get(d.Where("no_number = ?", id).Preload("User").Preload("Tags")); err != nil {
		return map[string]interface{}{}, err
	}

	// 1. Get comments
	comm, err := nice.GetComments(d)
	if err != nil {
		return map[string]interface{}{}, err
	}
	nice.Comments = comm

	// 2. view +1
	go nice.ViewAdd(s.Dao.DB)

	// 3. get comments create user id and avatar uri
	ids := []uint{}
	for _, comment := range nice.Comments {
		ids = append(ids, comment.UserId)
	}

	iUsers, _ := s.GetUsersAvatar(ids)

	// 4. result map
	userInfo := make(map[string]interface{})

	// 5. turn struct to map[string]interface
	st := utils.ToMap(nice)
	for _, k := range iUsers {
		userInfo[strconv.Itoa(int(k.ID))] = k
	}

	st["userInfo"] = userInfo

	return st, err
}

func (s *Service) GetNiceList(column, value string, pageSize int, pageIndex int) (os []model.Nice, err error) {

	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	n := model.Nice{}
	os, err = n.Gets(s.Dao.DB, column, value, pageSize, pageIndex)
	//o := model.Order{}
	//os, err = o.Get(s.Dao.DB)

	if err != nil {
		return os, err
	}

	return
}

func (s *Service) AddNice(Title, Desc, Content string, UserId, NiceType uint, tags []uint) (model.Nice, error) {

	// start a transactions
	tx := s.Dao.BeginTx()
	defer func() {
		tx.Commit()
	}()

	niceCheck := model.Nice{
		Title: Title,
	}

	if err := niceCheck.Get(tx.Where("title = ?", niceCheck.Title)); err != nil && err.Error() != "record not found" {
		return model.Nice{}, err
	} else if niceCheck.ID != 0 {
		return model.Nice{}, errors.New("标题该已存在")
	}

	nice := model.Nice{
		Title:    Title,
		Desc:     Desc,
		Content:  Content,
		NiceType: NiceType,
		UserId:   UserId,
		Model:    model.NewModel(),
		NoNumber: utils.GetNoNumber(Title),
	}

	// TODO middleware
	if err := nice.Create(tx); err != nil {
		tx.Rollback()
		return model.Nice{}, err
	}

	var nts []model.NiceTag = make([]model.NiceTag, len(tags))
	for _, k := range tags {
		nts = append(nts, model.NiceTag{TagID: k, NiceID: nice.ID})
	}

	if err := tx.Create(&nts).Error; err != nil {
		tx.Rollback()
		return model.Nice{}, err
	}

	return nice, nil
}

func (s *Service) LikeNice(postId, likeType, userId uint) error {
	d := s.Dao.DB

	user := model.User{ID: userId}
	if err := user.Get(d); err != nil && user.ID != 0 {
		return err
	}

	if likeType == 1 {
		// nice
		nice := model.Nice{ID: postId}

		if err := nice.LikeAdd(d); err != nil {
			return err
		}

	} else if likeType == 2 {
		// comment

		comm := model.Comment{ID: postId}

		if err := comm.LikeAdd(d); err != nil {
			return err
		}
	}

	like := model.ThumbsUp{PostId: postId, LikeType: likeType, UserId: userId}

	return like.Create(d)
}
