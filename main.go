package main

import (
	"fmt"
	"hw1/cache"
)

func main() {
	cache.Set("UserID", 12)

	fmt.Println(cache.Get("User1"))
	fmt.Println(cache.Get("UserID"))

	cache.Delete("UserID")

	fmt.Println(cache.Get("UserID"))
}
