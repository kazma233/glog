package handler

import (
	"glog/models"
	"glog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleCtr struct{}

type manageArticleCtr struct{}

// ArticleCtr 文章ctr
var ArticleCtr = &articleCtr{}

// ManageArticleCtr 管理文章ctr
var ManageArticleCtr = &manageArticleCtr{}

// Articles 文章列表
func (*articleCtr) Articles(c *gin.Context) {
	articleQuery := &models.ArticleQuery{}

	if err := c.ShouldBindQuery(articleQuery); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	c.JSON(http.StatusOK, models.Success(service.ArticleService.Article(articleQuery)))
}

func (*articleCtr) Detail(c *gin.Context) {
	id := c.Param("first")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	c.JSON(http.StatusOK, models.Success(service.ArticleService.Detail(id)))
}

// Group 分组查询
func (*articleCtr) Group(c *gin.Context) {

	c.JSON(http.StatusOK, models.Success(service.ArticleService.FindArticleGroupByArchive()))
}

// Save 保存文章
func (*manageArticleCtr) Save(c *gin.Context) {
	articleSave := &models.ArticleSave{}
	if err := c.BindJSON(articleSave); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	service.ManageArticleService.Save(articleSave)

	c.JSON(http.StatusOK, models.Success(nil))
}

// Update 更新文章
func (*manageArticleCtr) Update(c *gin.Context) {
	articleUpdate := &models.ArticleUpdate{}
	if err := c.BindJSON(articleUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorParamBind)
		return
	}

	err := service.ManageArticleService.Update(articleUpdate)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error500)
		return
	}

	c.JSON(http.StatusOK, models.Success(nil))
}

func (*manageArticleCtr) AllArticle(c *gin.Context) {
	manageArticleQuery := &models.ManageArticleQuery{}
	c.ShouldBindQuery(manageArticleQuery)

	c.JSON(http.StatusOK, models.Success(service.ManageArticleService.AllArticle(manageArticleQuery)))
}
