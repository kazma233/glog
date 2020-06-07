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
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	c.JSON(http.StatusOK, models.Success(service.ArticleService.Article(articleQuery)))
}

func (*ctrArticle) Detail(c *gin.Context) {
	id := c.Param("first")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	c.JSON(http.StatusOK, models.Success(service.ArticleService.Detail(id)))
}

// Save 保存文章
func (*ctrArticle) Save(c *gin.Context) {
	articleSave := &models.ArticleSave{}
	if err := c.BindJSON(articleSave); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	service.ArticleService.Save(articleSave)

	c.JSON(http.StatusOK, models.Success(nil))
}

// Group 分组查询
func (*ctrArticle) Group(c *gin.Context) {

	c.JSON(http.StatusOK, models.Success(service.ArticleService.FindArticleGroupByArchive()))
}
