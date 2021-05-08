package service

import (
	"errors"
	"nicetry/internal/model"
	"strings"
)

func (s *Service) GetTag(id uint) (model.Tag, error) {
	t := model.Tag{ID: id}
	if err := t.Get(s.Dao.DB.Where("id = ?", id)); err != nil {
		return model.Tag{}, err
	}
	return t, nil
}

func (s *Service) AddTag(name string, pid uint) error {
	t := model.Tag{Name: name, ParentId: pid}

	if err := t.Add(s.Dao.DB); err != nil {

		if strings.Contains(err.Error(), "Duplicate") {
			return errors.New("标签已存在")
		}
		return err
	}

	return nil
}
