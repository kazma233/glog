package service

import (
	"glog/config"
	"glog/dao"
	"glog/models"
	"glog/utils/logx"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
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

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        user.UserID,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	toeknString, err := token.SignedString(config.Conf().JwtKey)
	if err != nil {
		logx.Error("生成jwt token错误: %v", err)
		return "", models.ErrUnknow
	}

	return toeknString, nil
}
