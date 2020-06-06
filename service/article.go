package service

import (
	"glog/dao"
	"glog/models"
	"glog/utils/logx"
	"glog/utils/tools"
	"time"

	"github.com/rs/xid"
)

// Article 文章列表
func Article(query *models.ArticleQuery) []models.Article {
	val, err := dao.FindArticle(query)
	if err != nil {
		logx.Error("查询文章列表失败: %v", err)

		return []models.Article{}
	}

	return val
}

// Save 保存文章
func Save(articleSave *models.ArticleSave) {
	now := time.Now().In(tools.SHLoc)
	archiveDate := now.Format("2006-01-02")
	article := &models.Article{
		ArticleID:   xid.New().String(),
		Title:       articleSave.Title,
		Tags:        articleSave.Tags,
		Category:    articleSave.Category,
		SubTitle:    articleSave.SubTitle,
		Content:     articleSave.Content,
		CreateTime:  models.LocalTime{Time: now},
		UpdateTime:  models.LocalTime{Time: now},
		Visit:       0,
		Status:      articleSave.Status,
		ArchiveDate: archiveDate,
	}

	err := dao.SaveArticle(article)

	if err != nil {
		logx.Error("新增文章失败: %v", err)
	}
}

// FindArticleGroupByArchive 按照归档日期分组
func FindArticleGroupByArchive() []models.ArticleGroup {

	result, err := dao.FindArticleGroupByArchive()

	if err != nil {
		logx.Error("新增文章分许信息失败: %v", err)
	}

	return result
}
