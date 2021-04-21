package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"nicetry/internal/model"
	"nicetry/pkg/utils"
	"time"
)

func (s *Service) Login(mail, password string) (string, error) {

	// 查询用户信息
	user := model.User{Mail: mail, Password: password}
	if err := user.Login(s.Dao.DB); err != nil {
		return "", err
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码有误")
	}

	// 生成token
	token, err := utils.GenerateToken(mail, user.Password, user.ID)
	if err != nil {
		return "", err
	}

	// 保存token
	key := fmt.Sprintf("%v:", user.ID)
	if err := s.Dao.Cache.Set(key, token, time.Hour * 24).Err(); err != nil {
		return "", err
	}

	return token, nil
}


func (s *Service) Register(referralCode,nickname, password, mail, avatar, link, desc string) error {

	// 开启事务
	tx := s.Dao.BeginTx()

	defer func() {
		tx.Commit()
	}()

	// 1 验证推荐码
	rc := model.ReferralCode{Code: referralCode, Used: true}

	rcNum := rc.ConsumeCode(tx.Where("code = ? AND used = ?", referralCode, 0))

	if rcNum == 0 {
		tx.Rollback()
		return errors.New("推荐码无效")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 	//加密处理
	user := model.User{
		Nickname: nickname,
		Password: string(hash),
		Mail: mail,
		Avatar: avatar,
		Desc: desc,
		RecommendBy: 0,
		Link: link,
		Points: 0.0,
	}

	err = user.Create(tx)		// 注册用户

	if err != nil {
		tx.Rollback()
		return err
	}

	rc.To = user.ID
	rcNum = rc.ConsumeCode(tx.Where("code = ? AND used = ?", referralCode, 1))

	if rcNum == 0 {
		tx.Rollback()
		return errors.New("推荐码绑定用户失败")
	}

	// 异步添加
	for i:=0;i<2;i++ {
		go user.CreateReferCode(s.Dao.DB)
	}

	return nil
}