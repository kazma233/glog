package models

import (
	"time"
)

// ArticleStatus 文章状态
type ArticleStatus string

const (
	// ArticleShow show
	ArticleShow ArticleStatus = "SHOW"
	// ArticleHidden hidden
	ArticleHidden ArticleStatus = "HIDDEN"
)

// Article 文章实体
type Article struct {
	ArticleID   string        `bson:"articleId" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Tags        []string      `bson:"tags" json:"tags"`
	Category    string        `bson:"category" json:"category"`
	SubTitle    string        `bson:"subTitle" json:"subtitle"`
	Content     string        `bson:"content" json:"content"`
	CreateTime  LocalTime     `bson:"createTime" json:"latestTime"`
	UpdateTime  LocalTime     `bson:"updateTime" json:"-"`
	Visit       int64         `bson:"visit" json:"visit"`
	Status      ArticleStatus `bson:"status" json:"-"`
	ArchiveDate string        `bson:"archiveDate" json:"-"`
}

// ArticleCategory 文章分类
type ArticleCategory struct {
	ArticleCategoryID string    `bson:"articleCategoryID"`
	Name              string    `bson:"name"`
	CreateTime        time.Time `bson:"createTime"`
}

type (
	// ArticleQuery 文章查询
	ArticleQuery struct {
		Query
	}

	// ManageArticleQuery 管理页面的文章查询对象
	ManageArticleQuery struct {
		Query
		ArticleID string        `form:"id"`
		Title     string        `form:"title"`
		Category  string        `form:"category"`
		Tags      string        `form:"tags"`
		Status    ArticleStatus `form:"status"`
	}

	// ArticleSave 保存文章
	ArticleSave struct {
		Title    string        `json:"title" binding:"required"`
		Tags     string        `json:"tags" binding:"required"`
		Category string        `json:"category" binding:"required"`
		SubTitle string        `json:"subTitle" binding:"required"`
		Content  string        `json:"content" binding:"required"`
		Status   ArticleStatus `json:"status" binding:"required"`
	}

	// ArticleUpdate 更新文章
	ArticleUpdate struct {
		ArticleID string        `json:"id" binding:"required"`
		Title     string        `json:"title" binding:"required"`
		Tags      string        `json:"tags" binding:"required"`
		Category  string        `json:"category" binding:"required"`
		SubTitle  string        `json:"subTitle" binding:"required"`
		Content   string        `json:"content" binding:"required"`
		Status    ArticleStatus `json:"status" binding:"required"`
	}

	// ArticleUpdateDTO 存入数据库对象
	ArticleUpdateDTO struct {
		ArticleID string
		Title     string
		Tags      []string
		Category  string
		SubTitle  string
		Content   string
		Status    ArticleStatus
	}
)

type (
	// ArticleSimple 没有content的Article实体
	ArticleSimple struct {
		ArticleID  string    `bson:"articleId" json:"id"`
		Title      string    `bson:"title" json:"title"`
		Tags       []string  `bson:"tags" json:"tags"`
		Category   string    `bson:"category" json:"category"`
		SubTitle   string    `bson:"subTitle" json:"subTitle"`
		CreateTime LocalTime `bson:"createTime" json:"latestTime"`
	}

	// ArticleDetail Article实体
	ArticleDetail struct {
		ArticleID  string        `bson:"articleId" json:"id"`
		Title      string        `bson:"title" json:"title"`
		Tags       []string      `bson:"tags" json:"tags"`
		Category   string        `bson:"category" json:"category"`
		SubTitle   string        `bson:"subTitle" json:"subTitle"`
		Content    string        `bson:"content" json:"content"`
		Visit      int64         `bson:"visit" json:"visit"`
		Status     ArticleStatus `bson:"status" json:"status"`
		CreateTime LocalTime     `bson:"createTime" json:"latestTime"`
	}

	// ArticleGroupSimple 归档
	ArticleGroupSimple struct {
		ArticleID string `bson:"articleId" json:"id"`
		Title     string `bson:"title" json:"title"`
		Visit     int64  `bson:"visit" json:"visit"`
	}

	// ArticleGroup 文章归档日期聚合
	ArticleGroup struct {
		ArchiveDate string               `bson:"_id" json:"archiveDate"`
		Articles    []ArticleGroupSimple `bson:"articles" json:"articles"`
	}

	// ArticleManageDetail 管理文章实体类
	ArticleManageDetail struct {
		ArticleID  string        `bson:"articleId" json:"id"`
		Title      string        `bson:"title" json:"title"`
		Tags       []string      `bson:"tags" json:"tags"`
		Category   string        `bson:"category" json:"category"`
		SubTitle   string        `bson:"subTitle" json:"subTitle"`
		Visit      int64         `bson:"visit" json:"visit"`
		CreateTime LocalTime     `bson:"createTime" json:"latestTime"`
		Status     ArticleStatus `bson:"status" json:"status"`
	}
)
