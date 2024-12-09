package cache

import (
	"fmt"
	"time"
)

type memCache struct {
	//最大内存
	maxMemorySize int64
	//最大内存字符串表示
	maxMemorySizeStr string
	//当前已使用的内存
	currentMemorySize int64

	values map[string]memCacheValue
}

type memCacheValue struct {
	//value值
	val interface{}
	//过期时间
	expireTime time.Time
	//value大小
	size int64
}

func NewMemCache() Cache {
	return &memCache{}
}

// size: 1KB 100KB 1MB 2MB 1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)

	return true
}

// 将value写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	fmt.Println("called set")
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		size:       GetValueSize(val),
	}
	mc.values[key] = v
	return true
}

// 根据Key值获取value
func (mc *memCache) Get(key string) (interface{}, bool) {
	return nil, false
}

// 删除key值
func (mc *memCache) Del(key string) bool {
	return true
}

// 判断key是否存在
func (mc *memCache) Exists(key string) bool {
	return true
}

// 清空所有key
func (mc *memCache) Flush() bool {
	return true
}

// 获取缓存中所有key的数量
func (mc *memCache) Keys() int64 {
	return 0
}
