package dao

import (
	"glog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type articleDao struct{}

// ArticleDao 控制获取Article的doa
var ArticleDao = articleDao{}

// FindArticle 分页查找文章
func (articleDao) FindArticle(articleQuery *models.ArticleQuery) ([]models.ArticleSimple, error) {
	skip := (articleQuery.PageNo - 1) * articleQuery.PageSize

	cursor, err := articleColl.Find(
		create3SCtx(),
		bson.M{"status": models.ArticleShow},
		options.
			Find().
			SetProjection(bson.M{"articleId": 1, "title": 1, "tags": 1, "category": 1, "subTitle": 1, "createTime": 1}).
			SetSkip(skip).
			SetLimit(articleQuery.PageSize).
			SetSort(bson.M{"createTime": -1}),
	)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(create3SCtx())

	articles := []models.ArticleSimple{}
	err = cursor.All(create3SCtx(), &articles)

	return articles, err
}

func (articleDao) CountArticle(articleQuery *models.ArticleQuery) (int64, error) {
	return articleColl.CountDocuments(create3SCtx(), bson.M{"status": models.ArticleShow})
}

func (articleDao) Article(id string) (*models.ArticleDetail, error) {
	articleDetail := &models.ArticleDetail{}
	sr := articleColl.FindOne(
		create3SCtx(),
		bson.M{"status": models.ArticleShow, "articleId": id},
		options.
			FindOne().
			SetProjection(bson.M{"articleId": 1, "title": 1, "tags": 1, "category": 1, "subTitle": 1, "createTime": 1, "content": 1, "visit": 1, "status": 1}),
	)
	if err := sr.Err(); err != nil {
		return nil, err
	}

	err := sr.Decode(articleDetail)

	return articleDetail, err
}

// FindArticleGroupByArchive 按归档日期分组查询
func (articleDao) FindArticleGroupByArchive() ([]models.ArticleGroup, error) {
	result := []models.ArticleGroup{}

	cursor, err := articleColl.Aggregate(create3SCtx(), mongo.Pipeline{
		{{"$match", bson.D{{"status", models.ArticleShow}}}},
		{{"$group", bson.D{{"_id", "$archiveDate"}, {"articles", bson.D{{"$push", "$$ROOT"}}}}}},
		{{"$project", bson.D{{"articles.articleId", 1}, {"articles.title", 1}, {"articles.visit", 1}}}},
		{{"$sort", bson.M{"_id": -1}}},
	})

	if err != nil {
		return nil, err
	}

	err = cursor.All(create3SCtx(), &result)

	return result, err
}

// SaveArticle 保存
func (articleDao) SaveArticle(article *models.Article) error {
	_, err := articleColl.InsertOne(create3SCtx(), article)

	return err
}

func (articleDao) Update(articleUpdate *models.ArticleUpdateDTO) error {
	_, err := articleColl.UpdateOne(create3SCtx(), bson.M{
		"articleId": articleUpdate.ArticleID,
	}, bson.M{
		"$set": bson.M{
			"title":    articleUpdate.Title,
			"subTitle": articleUpdate.SubTitle,
			"tags":     articleUpdate.Tags,
			"category": articleUpdate.Category,
			"content":  articleUpdate.Content,
			"status":   articleUpdate.Status,
		},
	})

	return err
}

// AllArticle 按条件查找
func (articleDao) AllArticle(articleQuery *models.ManageArticleQuery) ([]models.ArticleManageDetail, error) {
	skip := (articleQuery.PageNo - 1) * articleQuery.PageSize

	cursor, err := articleColl.Find(create3SCtx(), getAllArticleQueryFilter(articleQuery), options.
		Find().
		SetProjection(bson.M{"articleId": 1, "title": 1, "tags": 1, "category": 1, "subTitle": 1, "createTime": 1, "visit": 1, "status": 1}).
		SetSkip(skip).
		SetLimit(articleQuery.PageSize).
		SetSort(bson.M{"createTime": -1}),
	)

	if err != nil {
		return nil, err
	}

	result := []models.ArticleManageDetail{}
	err = cursor.All(create3SCtx(), &result)

	return result, err
}

// AllArticle 按条件查找
func (articleDao) AllArticleSize(articleQuery *models.ManageArticleQuery) (int64, error) {
	return articleColl.CountDocuments(create3SCtx(), getAllArticleQueryFilter(articleQuery))
}

// 浏览量+1
func (articleDao) VisitInc(id string) {
	_, _ = articleColl.UpdateOne(create3SCtx(), bson.M{"articleId": id}, bson.M{"$inc": bson.M{"visit": 1}})
}

func getAllArticleQueryFilter(articleQuery *models.ManageArticleQuery) bson.M {
	filter := bson.M{}

	if articleQuery.ArticleID != "" {
		filter["articleId"] = articleQuery.ArticleID
	}

	if articleQuery.Title != "" {
		filter["title"] = articleQuery.Title
	}

	if articleQuery.Tags != "" {
		filter["tags"] = articleQuery.Tags
	}

	if articleQuery.Category != "" {
		filter["category"] = articleQuery.Category
	}

	if articleQuery.Status != "" {
		filter["status"] = articleQuery.Status
	}

	return filter
}
