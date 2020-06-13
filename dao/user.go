package dao

import (
	"glog/models"

	"go.mongodb.org/mongo-driver/bson"
)

type userDao struct {
}

// UserDao 用户数据层
var UserDao = userDao{}

// FindByUsername 通过用户名查找用户
func (userDao) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := userColl.FindOne(create3SCtx(), bson.M{"username": username}).Decode(user)

	return user, err
}

// Insert 新增用户
func (userDao) Insert(userRegister *models.UserRegister) error {
	_, err := userColl.InsertOne(create3SCtx(), userRegister)

	return err
}
