package dao

import (
	"glog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type articleDaoctr struct{}

// ArticleDao 控制获取Article的doa
var ArticleDao = articleDaoctr{}

// FindArticle 分页查找文章
func (articleDaoctr) FindArticle(articleQuery *models.ArticleQuery) ([]models.ArticleSimple, error) {
	skip := (articleQuery.PageNo - 1) * articleQuery.PageSize

	cursor, err := articleColl.Find(
		create3SCtx(),
		bson.M{"status": models.ArticleShow},
		options.
			Find().
			SetProjection(bson.M{"articleId": 1, "title": 1, "tags": 1, "category": 1, "subTitle": 1, "updateTime": 1}).
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

func (articleDaoctr) CountArticle(articleQuery *models.ArticleQuery) (int64, error) {
	return articleColl.CountDocuments(create3SCtx(), bson.M{"status": models.ArticleShow})
}

func (articleDaoctr) Article(id string) (*models.ArticleDetail, error) {
	articleDetail := &models.ArticleDetail{}
	sr := articleColl.FindOne(
		create3SCtx(),
		bson.M{"status": models.ArticleShow, "articleId": id},
		options.
			FindOne().
			SetProjection(bson.M{"articleId": 1, "title": 1, "tags": 1, "category": 1, "subTitle": 1, "updateTime": 1, "content": 1, "visit": 1}),
	)
	if err := sr.Err(); err != nil {
		return nil, err
	}

	err := sr.Decode(articleDetail)

	return articleDetail, err
}

// SaveArticle 保存
func (articleDaoctr) SaveArticle(article *models.Article) error {
	_, err := articleColl.InsertOne(create3SCtx(), article)

	return err
}

// FindArticleGroupByArchive 按归档日期分组查询
func (articleDaoctr) FindArticleGroupByArchive() ([]models.ArticleGroup, error) {
	result := []models.ArticleGroup{}

	cursor, err := articleColl.Aggregate(create3SCtx(), mongo.Pipeline{
		{{"$match", bson.D{{"status", models.ArticleShow}}}},
		{{"$group", bson.D{{"_id", "archiveDate"}, {"articles", bson.D{{"$push", "$$ROOT"}}}}}},
		{{"$project", bson.D{{"articles.articleId", 1}, {"articles.title", 1}, {"articles.visit", 1}}}},
	})

	if err != nil {
		return nil, err
	}

	err = cursor.All(create3SCtx(), &result)

	return result, err
}
