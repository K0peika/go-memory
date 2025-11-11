package cache

import (
	"log"
)

var cache map[string]any

func Set(key string, value any) {
	cache = map[string]any{
		key: value,
	}
}

func Get(key string) any {
	if _, ok := cache[key]; ok {
		return cache[key]
	}
	return "key NOT found"
}

func Delete(key string) {
	cache[key] = nil
	log.Printf("deletion of %s - successful", key)
}
