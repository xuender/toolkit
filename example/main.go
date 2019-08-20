package main

import (
	"fmt"
	"time"

	"github.com/xuender/tookit"
)

func main() {
	cache := tookit.NewCache(1 * time.Second)
	defer cache.Close()
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	fmt.Println(cache.Get("key1"))
	fmt.Println(cache.Count())
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Count())
}
