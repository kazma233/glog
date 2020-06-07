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
	CreateTime  LocalTime     `bson:"createTime" json:"-"`
	UpdateTime  LocalTime     `bson:"updateTime" json:"latestTime"`
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

	// ArticleSave 保存文章
	ArticleSave struct {
		Title    string        `json:"title" binding:"required"`
		Tags     []string      `json:"tags" binding:"required"`
		Category string        `json:"category" binding:"required"`
		SubTitle string        `json:"subTitle" binding:"required"`
		Content  string        `json:"content" binding:"required"`
		Status   ArticleStatus `json:"status" binding:"required"`
	}
)

type (
	// ArticleSimple 没有content的Article实体
	ArticleSimple struct {
		ArticleID  string    `bson:"articleId" json:"id"`
		Title      string    `bson:"title" json:"title"`
		Tags       []string  `bson:"tags" json:"tags"`
		Category   string    `bson:"category" json:"category"`
		SubTitle   string    `bson:"subTitle" json:"subtitle"`
		UpdateTime LocalTime `bson:"updateTime" json:"latestTime"`
	}

	// ArticleDetail Article实体
	ArticleDetail struct {
		ArticleID  string    `bson:"articleId" json:"id"`
		Title      string    `bson:"title" json:"title"`
		Tags       []string  `bson:"tags" json:"tags"`
		Category   string    `bson:"category" json:"category"`
		SubTitle   string    `bson:"subTitle" json:"subtitle"`
		Content    string    `bson:"content" json:"content"`
		Visit      int64     `bson:"visit" json:"visit"`
		UpdateTime LocalTime `bson:"updateTime" json:"latestTime"`
	}

	// ArticleGroupSimple 归档
	ArticleGroupSimple struct {
		ArticleID string `bson:"articleId" json:"id"`
		Title     string `bson:"title" json:"title"`
		Visit     int64  `bson:"visit" json:"visit"`
	}

	// ArticleGroup 文章归档日期聚合
	ArticleGroup struct {
		Articles []ArticleGroupSimple `bson:"articles" json:"articles"`
	}
)
