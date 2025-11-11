package cache

import (
	"log"
)

var cache map[string]interface{}

func Set(key string, value interface{}) {
	cache[key] = value
}

func Get(key string) interface{} {
	return cache[key]
}

func Delete(key string) {
	cache[key] = nil
	log.Printf("deletion of %s - successful", key)
}
