package service

import (
	"glog/dao"
	"glog/models"
	"glog/utils/logx"
	"glog/utils/pageable"
	"glog/utils/tools"
	"strings"
	"time"

	"github.com/rs/xid"
)

type articleServiceCtr struct{}

// ArticleService article service
var ArticleService = articleServiceCtr{}

// Article 文章列表
func (articleServiceCtr) Article(query *models.ArticleQuery) *models.Page {
	query.InitDate()

	total, err := dao.ArticleDao.CountArticle(query)
	if err != nil {
		logx.Error("查询文章数量失败: %v", err)

		return pageable.Empty(query.PageSize)
	}

	val, err := dao.ArticleDao.FindArticle(query)
	if err != nil {
		logx.Error("查询文章列表失败: %v", err)

		return pageable.Empty(query.PageSize)
	}

	return pageable.Result(query.PageNo, query.PageSize, total, val)
}

func (articleServiceCtr) Detail(id string) *models.ArticleDetail {
	articleDetail, err := dao.ArticleDao.Article(id)

	if err != nil {
		logx.Error("查询id： %v的文章出错: %v", id, err)
		return nil
	}

	return articleDetail
}

// Save 保存文章
func (articleServiceCtr) Save(articleSave *models.ArticleSave) {
	now := time.Now().In(tools.SHLoc)
	archiveDate := now.Format("2006-01-02")
	article := &models.Article{
		ArticleID:   xid.New().String(),
		Title:       articleSave.Title,
		Tags:        strings.Split(articleSave.Tags, ","),
		Category:    articleSave.Category,
		SubTitle:    articleSave.SubTitle,
		Content:     articleSave.Content,
		CreateTime:  models.LocalTime{Time: now},
		UpdateTime:  models.LocalTime{Time: now},
		Visit:       0,
		Status:      articleSave.Status,
		ArchiveDate: archiveDate,
	}

	err := dao.ArticleDao.SaveArticle(article)

	if err != nil {
		logx.Error("新增文章失败: %v", err)
	}
}

// FindArticleGroupByArchive 按照归档日期分组
func (articleServiceCtr) FindArticleGroupByArchive() []models.ArticleGroup {

	result, err := dao.ArticleDao.FindArticleGroupByArchive()

	if err != nil {
		logx.Error("新增文章分许信息失败: %v", err)
	}

	return result
}
