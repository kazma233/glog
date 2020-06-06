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
	ArticleID   string        `bson:"articleID" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Tags        []string      `bson:"tags" json:"tags"`
	Category    string        `bson:"category" json:"category"`
	SubTitle    string        `bson:"subTitle" json:"subTitle"`
	Content     string        `baon:"content" json:"content"`
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
	// ArticleGroup 文章归档日期聚合
	ArticleGroup struct {
		Articles []Article `bson:"articles" json:"articles"`
	}
)
