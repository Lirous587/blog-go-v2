package service

import (
	"blog/cache"
	"blog/models"
	"blog/pkg/snowflake"
	"blog/repository"
	"blog/utils"
	"fmt"
)

const (
	userExisted = "用户已存在"
)

type UsrService interface {
	Signup(data *models.UserSignupParams) error
	Login(data *models.UserLoginParams) (*models.UserData, error)
	Update(data *models.UserUpdateParams) error
	Logout(data *models.UserLogoutParams) error
}

type UserCacheService struct {
	repo repository.UserRepo
	cch  cache.UserCache
}

func NewUserRepoService(cch cache.UserCache, repo repository.UserRepo) *UserCacheService {
	return &UserCacheService{
		cch:  cch,
		repo: repo,
	}
}

func (s *UserCacheService) Signup(data *models.UserSignupParams) error {
	//1.判断用户是否存在
	ok, err := s.repo.CheckExist(data)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf(userExisted)
	}
	//2.生成uid并保存相关信息
	uid := snowflake.GenID()
	encryptedPassword := utils.EncryptPassword(data.Password)
	user := &models.UserData{
		UID:      uid,
		Name:     data.Name,
		Email:    data.Email,
		Password: encryptedPassword,
	}
	//3.将用户储存在数据库
	return s.repo.Save(user)
}

func (s *UserCacheService) Login(data *models.UserLoginParams) (*models.UserData, error) {
	encryptedPassword := utils.EncryptPassword(data.Password)
	data.Password = encryptedPassword
	// 身份校验
	ret, err := s.repo.Validate(data)
	if err != nil {
		return nil, err
	}
	var token string
	// 生成token 使用redis或jwt
	if token, err = s.cch.GenToken(ret.UID); err != nil {
		return nil, err
	}
	//将token保存
	ret.Token = token
	return ret, nil
}

func (s *UserCacheService) Update(user *models.UserUpdateParams) error {
	return nil
}

func (s *UserCacheService) Logout(token *models.UserLogoutParams) error {
	return nil
}
