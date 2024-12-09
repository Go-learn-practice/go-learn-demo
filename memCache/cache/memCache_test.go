package cache

import (
	"testing"
	"time"
)

func TestCacheOP(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"junjie0", 789, time.Second * 10},
		{"junjie1", false, time.Second * 12},
		{"junjie2", true, time.Second * 14},
		{"junjie3", map[string]interface{}{"a": 3, "b": false}, time.Second * 16},
		{"junjie4", "abcdefghijklmn", time.Second * 18},
		{"junjie5", "你好欢迎来到我的世界", time.Second * 20},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, d := range testData {
		c.Set(d.key, d.val, d.expire)
		val, ok := c.Get(d.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		if d.key != "junjie3" && val != d.val {
			t.Errorf("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if d.key == "junjie3" && !ok1 {
			t.Errorf("缓存取值数据与预期不一致")
		}
	}

	if int64(len(testData)) != c.Keys() {
		t.Error("缓存数量不一致")
	}
	c.Del(testData[0].key)
	c.Del(testData[1].key)
	if int64(len(testData)) != c.Keys()+2 {
		t.Error("缓存数量不一致")
	}

	time.Sleep(time.Second * 20)
	if c.Keys() != 0 {
		t.Error("过期缓存清空失败")
	}
}
