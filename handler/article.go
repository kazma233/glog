package handler

import (
	"glog/models"
	"glog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ctrArticle struct {
}

// ArticleCtr 文章ctr
var ArticleCtr = &ctrArticle{}

// Articles 文章列表
func (*ctrArticle) Articles(c *gin.Context) {
	articleQuery := &models.ArticleQuery{}
	if err := c.ShouldBindQuery(articleQuery); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.PARAM_BIND_ERROR)
		return
	}

	c.JSON(http.StatusOK, models.Success(service.Article(articleQuery)))
}

// Save 保存文章
func (*ctrArticle) Save(c *gin.Context) {
	articleSave := &models.ArticleSave{}
	if err := c.BindJSON(articleSave); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.PARAM_BIND_ERROR)
		return
	}

	service.Save(articleSave)

	c.JSON(http.StatusOK, models.Success(nil))
}

// Group 分组查询
func (*ctrArticle) Group(c *gin.Context) {

	c.JSON(http.StatusOK, models.Success(service.FindArticleGroupByArchive()))
}
