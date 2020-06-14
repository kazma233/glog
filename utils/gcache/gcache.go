package gcache

import (
	"encoding/json"
	"glog/models"
	"glog/utils/logx"
	"time"

	"github.com/patrickmn/go-cache"
)

type articleCache struct{}

// ArticleCacheHandler 文章缓存控制
var ArticleCacheHandler = articleCache{}
var globalArticleCache = cache.New(time.Hour, 10*time.Second)

func (articleCache) CacheHome(query *models.ArticleQuery, page *models.Page) {
	cacheKey, err := json.Marshal(query)
	if err != nil {
		logx.Error("缓存首页数据时，获得缓存的key失败： %v", err)
		return
	}

	globalArticleCache.SetDefault(string(cacheKey), page)
}

func (articleCache) GetCacheHome(query *models.ArticleQuery) (*models.Page, bool) {
	cacheKey, err := json.Marshal(query)
	if err != nil {
		logx.Error("获取首页缓存数据时，获得缓存的key失败： %v", err)
		return nil, false
	}

	result, ok := globalArticleCache.Get(string(cacheKey))
	if ok {
		return result.(*models.Page), true
	}

	return nil, false
}

func (articleCache) Get(id string) (*models.ArticleDetail, bool) {
	val, ok := globalArticleCache.Get(id)
	if ok {
		return val.(*models.ArticleDetail), true
	}

	return nil, false
}

func (articleCache) Put(key string, val *models.ArticleDetail) {
	globalArticleCache.SetDefault(key, val)
}

func (a articleCache) Update(key string, val *models.ArticleDetail) {
	a.Delete(key)
	a.Put(key, val)
}

func (articleCache) Delete(key string) {
	globalArticleCache.Delete(key)
}
