package dao

import (
	"glog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindArticle 分页查找文章
func FindArticle(articleQuery *models.ArticleQuery) ([]models.Article, error) {
	skip := (articleQuery.PageNo - 1) * articleQuery.PageSize

	cursor, err := ArticleCollection.Find(
		Create3SCtx(),
		bson.M{"status": models.ArticleShow},
		options.
			Find().
			SetSkip(skip).
			SetLimit(articleQuery.PageSize).
			SetSort(bson.M{"createTime": -1}),
	)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(Create3SCtx())

	articles := []models.Article{}
	err = cursor.All(Create3SCtx(), &articles)

	return articles, err
}

// SaveArticle 保存
func SaveArticle(article *models.Article) error {
	_, err := ArticleCollection.InsertOne(Create3SCtx(), article)

	return err
}

// FindArticleGroupByArchive 按归档日期分组查询
func FindArticleGroupByArchive() ([]models.ArticleGroup, error) {
	result := []models.ArticleGroup{}

	cursor, err := ArticleCollection.Aggregate(Create3SCtx(), mongo.Pipeline{
		{{"$match", bson.D{{"status", models.ArticleShow}}}},
		{{"$group", bson.D{{"_id", "archiveDate"}, {"articles", bson.D{{"$push", "$$ROOT"}}}}}},
	})

	if err != nil {
		return nil, err
	}

	err = cursor.All(Create3SCtx(), &result)

	return result, err
}
