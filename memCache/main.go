package main

import _cache "memCache/cache"

func main() {
	cache := _cache.NewMemCache()
	cache.SetMaxMemory("100MB")
}
