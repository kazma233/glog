package service

import (
	"glog/config"
	"glog/dao"
	"glog/models"
	"glog/utils/logx"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userSer struct {
}

// UserService 用户service
var UserService = userSer{}

func (userSer) Login(login *models.UserLogin) (string, error) {
	user, err := dao.UserDao.FindByUsername(login.Username)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logx.Error("无此用户: %s", login.Username)
			return "", models.ErrUserNotExits
		}

		logx.Error("查询用户出错: %v", err)
		return "", models.ErrUnknow
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)) != nil {
		return "", models.ErrUserLogin
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        user.UserID,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	toeknString, err := token.SignedString([]byte(config.Conf().JwtKey))
	if err != nil {
		logx.Error("生成jwt token错误: %v", err)
		return "", models.ErrUnknow
	}

	return toeknString, nil
}

func (userSer) Register(userRegister *models.UserRegister) error {
	user, _ := dao.UserDao.FindByUsername(userRegister.Username)

	if user != nil && user.UserID != "" {
		return models.ErrUserExits
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		logx.Error("密码加密错误: %v", err)
		return models.ErrUnknow
	}

	userRegister.Password = string(pw)

	err = dao.UserDao.Insert(userRegister)

	if err != nil {
		logx.Error("注册时发生异常: %v", err)

		return models.ErrRegister
	}

	return nil
}
