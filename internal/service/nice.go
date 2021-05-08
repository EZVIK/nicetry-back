package service

import (
	"errors"
	"nicetry/internal/model"
	"nicetry/pkg/utils"
)

func (s *Service) GetNice(id string) (model.Nice, error) {
	d := s.Dao.DB
	nice := model.Nice{NoNumber: id}

	//if err := nice.Get(d.Preload("User").Preload("Tag")); err != nil {

	if err := nice.Get(d.Where("no_number = ?", id).Preload("User").Preload("Tags")); err != nil {
		return model.Nice{}, err
	}

	comm, err := nice.GetComments(d)
	if err != nil {
		return model.Nice{}, err
	}

	nice.Comments = comm
	go nice.ViewAdd(s.Dao.DB)
	return nice, err
}

func (s *Service) GetNiceList(column, value string, pageSize int, pageIndex int) (ns []model.Nice, err error) {

	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	n := model.Nice{}

	ns, err = n.Gets(s.Dao.DB, column, value, pageSize, pageIndex)

	if err != nil {
		return ns, err
	}

	return
}

func (s *Service) AddNice(Title, Desc, Content string, UserId, NiceType uint, tags []uint) (model.Nice, error) {

	tx := s.Dao.BeginTx()
	defer func() {
		tx.Commit()
	}()

	nice_check := model.Nice{
		Title: Title,
	}

	if err := nice_check.Get(tx.Where("title = ?", nice_check.Title)); err != nil && err.Error() != "record not found" {
		return model.Nice{}, err
	} else if nice_check.ID != 0 {
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

	if err := nice.Create(tx); err != nil {
		tx.Rollback()
		return model.Nice{}, err
	}

	nts := []model.NiceTag{}
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

	return like.Add(d)
}
