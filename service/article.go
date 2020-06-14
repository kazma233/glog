package service

import (
	"glog/dao"
	"glog/models"
	"glog/utils/gcache"
	"glog/utils/logx"
	"glog/utils/pageable"
	"glog/utils/tools"
	"strings"
	"time"

	"github.com/rs/xid"
)

type articleSer struct{}
type manageArticleSer struct{}

// ArticleService article service
var ArticleService = articleSer{}
var ManageArticleService = manageArticleSer{}

// Article 文章列表
func (articleSer) Article(query *models.ArticleQuery) *models.Page {
	query.InitDate()

	result, ok := gcache.ArticleCacheHandler.GetCacheHome(query)
	if ok {
		return result
	}

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

	pageResult := pageable.Result(query.PageNo, query.PageSize, total, val)

	gcache.ArticleCacheHandler.CacheHome(query, pageResult)

	return pageResult
}

func (articleSer) Detail(id string) *models.ArticleDetail {
	dao.ArticleDao.VisitInc(id)

	result, ok := gcache.ArticleCacheHandler.Get(id)

	if ok {
		return result
	}

	articleDetail, err := dao.ArticleDao.Article(id)

	if err != nil {
		logx.Error("查询id： %v的文章出错: %v", id, err)
		return nil
	}

	gcache.ArticleCacheHandler.Put(id, articleDetail)

	return articleDetail
}

// FindArticleGroupByArchive 按照归档日期分组
func (articleSer) FindArticleGroupByArchive() []models.ArticleGroup {

	result, err := dao.ArticleDao.FindArticleGroupByArchive()

	if err != nil {
		logx.Error("新增文章分许信息失败: %v", err)
	}

	return result
}

// Save 保存文章
func (manageArticleSer) Save(articleSave *models.ArticleSave) {
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

// 更新文章
func (manageArticleSer) Update(articleUpdate *models.ArticleUpdate) error {
	articleUpdateDTO := &models.ArticleUpdateDTO{
		ArticleID: articleUpdate.ArticleID,
		Title:     articleUpdate.Title,
		SubTitle:  articleUpdate.SubTitle,
		Content:   articleUpdate.Content,
		Status:    articleUpdate.Status,
		Tags:      strings.Split(articleUpdate.Tags, ","),
		Category:  articleUpdate.Category,
	}

	gcache.ArticleCacheHandler.Delete(articleUpdate.ArticleID)

	return dao.ArticleDao.Update(articleUpdateDTO)
}

// AllArticle 通过条件查询所有的文章
func (manageArticleSer) AllArticle(articleQuery *models.ManageArticleQuery) *models.Page {
	articleQuery.InitDate()

	total, err := dao.ArticleDao.AllArticleSize(articleQuery)
	if err != nil {
		logx.Error("查询管理文章总数失败: %v", err)
		return pageable.Empty(articleQuery.PageSize)
	}

	result, err := dao.ArticleDao.AllArticle(articleQuery)
	if err != nil {
		logx.Error("查询管理文章列表失败: %v", err)
		return pageable.Empty(articleQuery.PageSize)
	}

	return pageable.Result(articleQuery.PageNo, articleQuery.PageSize, total, result)
}
