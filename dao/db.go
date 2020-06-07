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
var userCollectionName = "user"

var (
	// articleColl article collection
	articleColl *mongo.Collection
	// articleCategoryColl article category collection
	articleCategoryColl *mongo.Collection
	// userColl user collection
	userColl *mongo.Collection
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

	client, err := mongo.Connect(create3SCtx(), clientOptions)

	if err != nil {
		panic(err)
	}

	dbClient := client.Database(dbName)

	articleColl = dbClient.Collection(articleCollectionName)
	articleCategoryColl = dbClient.Collection(articleCategoryCollectionName)
	userColl = dbClient.Collection(userCollectionName)
}

// create3SCtx 获取3s的context
func create3SCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx
}
