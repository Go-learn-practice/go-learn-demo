package cache_server

import (
	"memCache/cache"
	"time"
)

type CacheServer struct {
	memCache cache.Cache
}

func NewMemCache() *CacheServer {
	//memCache的值cache.NewMemCache()必须实现了cache.Cache的所有接口
	return &CacheServer{
		memCache: cache.NewMemCache(),
	}
}

// size: 1KB 100KB 1MB 2MB 1GB
func (cs *CacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}

// 将value写入缓存
func (cs *CacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := time.Second * 0
	if len(expire) > 0 {
		expireTs = expire[0]
	}
	return cs.memCache.Set(key, val, expireTs)
}

// 根据key值获取value
func (cs *CacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

// 删除key值
func (cs *CacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

// 判断key是否存在
func (cs *CacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

// 清空所有key
func (cs *CacheServer) Flush() bool {
	return cs.memCache.Flush()
}

// 获取缓存中所有key的数量
func (cs *CacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
