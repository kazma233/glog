package main

import (
	"glog/config"
	"glog/handler"
	"glog/models"
	"glog/utils/logx"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Router 设置路由和中间件
func Router() *gin.Engine {
	g := gin.Default()

	g.HandleMethodNotAllowed = true
	g.Use(corsFilter)
	g.Use(logsFilter)

	articleGroup := g.Group("/articles")
	articleGroup.GET("", handler.ArticleCtr.Articles)
	articleGroup.GET("/:first", handlerArticleRouter)

	userGroup := g.Group("/users")
	userGroup.POST("/login", handler.UserCtr.Login)
	userGroup.POST("/register", handler.UserCtr.Register)

	manageGroup := g.Group("/manage", authRouter)
	manageGroup.POST("/articles", handler.ArticleCtr.Save)

	g.Use(errorFilter)
	return g
}

func handlerArticleRouter(c *gin.Context) {
	one := c.Param("first")
	if "group" == one {
		handler.ArticleCtr.Group(c)
		return
	}

	handler.ArticleCtr.Detail(c)
}

// 异常拦截器
func errorFilter(c *gin.Context) {
	if c.Writer.Status() == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Error404)
		return
	}
	if c.Writer.Status() == http.StatusMethodNotAllowed {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, models.Error405)
		return
	}
	if c.Writer.Status() == http.StatusUnauthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Error401)
		return
	}

	uri := c.Request.RequestURI
	logx.Error("find error: uri=%v, params=%v, errors=%v", uri, c.Params, c.Errors)
	c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error500)
}

// authRouter 认证拦截器
func authRouter(c *gin.Context) {
	auth := c.GetHeader("auth")
	token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf().JwtKey), nil
	})

	if err != nil {
		logx.Error("jwt 解析失败: %v", err)

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	standardClaims := token.Claims.(*jwt.StandardClaims)
	c.Set(models.ConsUID, standardClaims.Id)
	c.Next()
}

// corsFilter 跨域中间件
func corsFilter(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}

	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, User-Agent, auth")
	c.Writer.Header().Set("Access-Control-Max-Age", "36000")

	// 如果是options方法，直接返回状态200就行了
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

func logsFilter(c *gin.Context) {
	logx.Info("%v use %v request %v", c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI)
	c.Next()
}
