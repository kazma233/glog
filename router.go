package main

import (
	"glog/handler"
	"glog/models"
	"glog/utils/logx"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router 设置路由和中间件
func Router() *gin.Engine {
	engine := gin.Default()

	engine.HandleMethodNotAllowed = true
	engine.Use(logsFilter)
	engine.Use(crosFilter)

	articleGroup := engine.Group("/articles")
	articleGroup.GET("", handler.ArticleCtr.Articles)
	articleGroup.POST("", handler.ArticleCtr.Save)
	articleGroup.GET("/group", handler.ArticleCtr.Group)

	engine.Use(errorFilter)
	return engine
}

// 异常拦截器
func errorFilter(c *gin.Context) {
	if c.Writer.Status() == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, models.RESOURCE_NOT_EXITS)
		return
	}
	if c.Writer.Status() == http.StatusMethodNotAllowed {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, models.METHOD_NOT_ALLOWED)
		return
	}
	if c.Writer.Status() == http.StatusUnauthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NOT_AUTH)
		return
	}

	uri := c.Request.RequestURI
	logx.Error("find error: uri=%v, params=%v, errors=%v", uri, c.Params, c.Errors)
	c.AbortWithStatusJSON(http.StatusInternalServerError, models.UNKNOW_ERROR)
}

// crosFilter 跨越中间件
func crosFilter(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}

	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Set-Cookie, Accept, Origin, User-Agent, x-requested-with, access-control-allow-origin")
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
