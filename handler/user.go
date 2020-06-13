package handler

import (
	"glog/models"
	"glog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userCtr struct{}

// UserCtr 用户模块
var UserCtr = userCtr{}

// Login 登录
func (*userCtr) Login(c *gin.Context) {
	userLogin := &models.UserLogin{}
	if err := c.BindJSON(userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	token, err := service.UserService.Login(userLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Faild(err))
		return
	}

	c.JSON(http.StatusOK, models.Success(token))
}

// Register 注册
func (*userCtr) Register(c *gin.Context) {
	userRegister := &models.UserRegister{}
	if err := c.BindJSON(userRegister); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	err := service.UserService.Register(userRegister)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Faild(err))
		return
	}

	c.JSON(http.StatusOK, models.Success(nil))
}

func (*userCtr) PasswordChange(c *gin.Context) {

}
