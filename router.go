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
	g := gin.Default()

	g.HandleMethodNotAllowed = true
	g.Use(logsFilter)
	g.Use(corsFilter)

	articleGroup := g.Group("/articles")
	articleGroup.GET("", handler.ArticleCtr.Articles)
	articleGroup.GET("/:first", handlerArticleRouter)
	articleGroup.POST("", handler.ArticleCtr.Save)

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

// corsFilter 跨域中间件
func corsFilter(c *gin.Context) {
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
