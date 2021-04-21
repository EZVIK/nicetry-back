package service

import "nicetry/internal/model"

func (s *Service) GetTag(id uint) (model.Tag, error) {
	t := model.Tag{ID: id}
	if err := t.Get(s.Dao.DB.Where("id = ?", id)); err != nil {
		return model.Tag{}, err
	}
	return t, nil
}

func (s *Service) AddTag(name string, pid uint) (error) {
	t := model.Tag{Name: name, ParentId: pid}
	return t.Add(s.Dao.DB)
}
