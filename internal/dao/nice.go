package dao

import (
	"gorm.io/gorm"
	"nicetry/internal/model"
)


func (d *Dao) CreateNice(title, desc, content string, niceType uint, tags []model.Tag) (model.Nice,error) {
	nice := model.Nice{
		Title: title,
		Desc: desc,
		Content: content,
		NiceType: niceType,
		Model: model.NewModel(),
	}
	if err := nice.Create(d.DB); err != nil {
		return model.Nice{}, err
	}
	return nice, nil
}


func (d *Dao) DeleteNice(id uint) error {
	//t, _ := utils.GetNowTimeCST()
	nice := model.Nice{ Model: &gorm.Model{ID: id}}

	return nice.Delete(d.DB)
}

func (d *Dao) GetNice(id uint) (n model.Nice, err error) {
	n.ID = id
	if err := n.Get(d.DB); err != nil {
		return n, err
	}

	tags, err := n.GetNiceTags(d.DB)

	if err != nil {
	    return
	}

	n.Tags = tags

	return n, err
}
