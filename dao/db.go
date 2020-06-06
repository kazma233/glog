package dao

import (
	"context"
	"fmt"
	"glog/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = "blog"
var articleCollectionName = "article"
var articleCategoryCollectionName = "article_category"

var (
	// ArticleCollection article collection
	ArticleCollection *mongo.Collection
	// ArticleCategoryCollection article category collection
	ArticleCategoryCollection *mongo.Collection
)

func init() {
	mongoConf := config.Conf().MongoConfig

	mongoURL := fmt.Sprintf("%s:%d", mongoConf.Host, mongoConf.Port)

	clientOptions := options.
		Client().
		SetConnectTimeout(3 * time.Second).
		SetHosts([]string{mongoURL}).
		SetAuth(options.Credential{
			Username: mongoConf.Username,
			Password: mongoConf.Password,
		})

	client, err := mongo.Connect(Create3SCtx(), clientOptions)

	if err != nil {
		panic(err)
	}

	dbClient := client.Database(dbName)

	ArticleCollection = dbClient.Collection(articleCollectionName)
	ArticleCategoryCollection = dbClient.Collection(articleCategoryCollectionName)
}

// Create3SCtx 获取3s的context
func Create3SCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx
}
