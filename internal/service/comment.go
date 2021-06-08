package service

import "nicetry/internal/model"

func (s *Service) AddComment(niceId, userid uint, content string) error {
	d := s.Dao.DB
	comm := model.Comment{NiceId: niceId, UserId: userid, Content: content}
	return comm.Add(d)
}

func (s *Service) GetComments(niceId uint) ([]model.Comment, error) {

	d := s.Dao.DB
	nice := model.Nice{ID: niceId}

	return nice.GetComments(d.Order("created_at asc"))
}
