package dao

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"nicetry/internal/model"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (d *Dao) Create(nickname, password, mail, avatar, desc string, recommendBy uint, link string, points float64) (model.User,error) {

	user := model.User{
		Nickname: nickname,
		Password: password,
		Mail: mail,
		Avatar: avatar,
		Desc: desc,
		RecommendBy: recommendBy,
		Link: link,
		Points: points,
	}

	count, err := user.CheckColumn(d.DB, "mail", user.Mail)

	if err != nil {
		return model.User{}, err
	}

	if count != 0 {
		return model.User{}, errors.New(e.GetMsg(e.INVALID_PARAMS))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return model.User{}, errors.New(e.BcryptPasswordError.Msg())
	}
	user.Password = string(hash)
	if err := user.Create(d.DB); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (d *Dao) Login(mail, password string) (string, bool) {

	user := model.User{Mail: mail, Password: password}

	pwd := user.Password

	if err := user.Get(d.DB.Where("mail = ?", mail)); err != nil {
		return "", false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))

	if err != nil {
		return  "", false
	}

	token, err := utils.GenerateToken(user.Mail, user.Password, user.ID)

	if err != nil {
		return  "", false
	}

	if err := user.Token(d.Cache, token); err != nil {
		return "", false
	}

	return token, true
}

func (d *Dao) Query(mail string) (u model.User,err error){

	u.Mail = mail

	if err := u.Get(d.DB); err != nil {
		return u, err
	}

	return
}