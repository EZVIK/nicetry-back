package service

import (
	"nicetry/internal/model"
)

func (s *Service) GetNice(id uint) (model.Nice, error) {

	nice, err := s.Dao.GetNice(id)

	if err == nil && nice.ID != 0 {
		go nice.ViewAdd(s.Dao.DB)
	}

	n:= model.Nice{ID: nice.ID}
	comm, err := n.GetComments(s.Dao.DB)
	nice.Comments = comm
	return nice, err
}


func (s *Service) GetNiceList(column , value string, pageSize int, pageIndex int) (ns []model.Nice,err error) {


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
	nice := model.Nice{
		Title: Title,
		Desc: Desc,
		Content: Content,
		NiceType: NiceType,
		UserId: UserId,
		Model: model.NewModel(),
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

	like := model.ThumbsUp{ PostId: postId, LikeType: likeType, UserId: userId}

	return like.Add(d)
}