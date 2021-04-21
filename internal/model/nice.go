package model

import (
	"gorm.io/gorm"
	"nicetry/global"
)

type Nice struct {

	ID      		uint 			`gorm:"primarykey"`

	Title   		string 			`gorm:"varchar(50);unique" json:"title"` 		// 标题

	Desc    		string 			`gorm:"varchar(200);" json:"desc"`				// 详情

	Content 		string 			`gorm:"type:text;" json:"content"`				// 内容

	View    		int64  			`gorm:"default:0;" json:"view"`					// 被查看数量

	Like    		int64  			`gorm:"default:0;" json:"like"`					// 被点赞数量

	UserId  		uint  			 `gorm:"index" json:"user_id"`					// 创建人

	NiceType		uint 			 `gorm:"tinyint(100);" json:"nice_type"`								// 类型 0 challenge 1 happy hour

	Tags    		[]Tag  			 `gorm:"many2many:nice_tags;"`

	*gorm.Model

}

type NiceList struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	//Avatar  string `json:"avatar"`
	//Tags    []Tag  `json:"tags"`
}

func (n *Nice) TableName() string {
	return "nice"
}

func (n *Nice) Create(db *gorm.DB,) error {
	return db.Create(&n).Error
}

func (n *Nice) Update(db *gorm.DB,) error {
	return db.Updates(&n).Error
}

func (n *Nice) Delete(db *gorm.DB,) error {
	return db.Delete(&n).Error
}

func (n *Nice) Get(db *gorm.DB) error {

	err := db.Debug().First(&n, n.ID).Error

	global.Logger.Info(err)

	return err
}

func (n *Nice) Gets(db *gorm.DB, column , value string, pageSize int, pageIndex int) (nl []NiceList, err error) {

	// db.Model(n).Scopes(Paginate(pageIndex, pageSize)).Where(column + "? = ", value)
	err = db.Debug().Scopes(Paginate(pageIndex, pageSize)).Model(&Nice{}).
		Where(column + "? = ", value).
		Scan(&nl).
		Order("view desc").Error

	if err != nil {
		return  nil, err
	}
	return
}

func (n *Nice) ViewAdd(db *gorm.DB) error {
	return db.Model(&Nice{}).Where("id = ?", n.ID).Update("view", gorm.Expr("view + ?", 1)).Error
}

func (n *Nice) LikeAdd(db *gorm.DB) error {
	return db.Model(&Nice{}).Where("id = ?", n.ID).Update("like", gorm.Expr("? + ?", "like",1)).Error
}

func (n *Nice) GetComments(db *gorm.DB) (cos []Comment,err error) {

	if err := db.Model(&Comment{}).Where("nice_id = ?", n.ID).Find(&cos).Error; err != nil {
		return cos, err
	}
	return cos ,err
}

func (n *Nice) GetNiceTags(db *gorm.DB) (tags []Tag, err error) {

	err = db.Raw("SELECT t.id, t.name FROM tags t LEFT JOIN nice_tags nt ON t.id = nt.tag_id WHERE (nt.nice_id = @niceId)", map[string]interface{}{"niceId": n.ID}).Find(&tags).Error

	if err != nil {
	    return tags, err
	}

	return tags, err
}