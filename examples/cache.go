package main

import (
	"fmt"
	"time"

	"github.com/xuender/toolkit"
)

func main() {
	// LRU
	cache := toolkit.NewCache(time.Second*3, true)
	cache.Set("key1", "value1")
	cache.SetByDuration("key2", "value2", time.Second)
	cache.Set("key3", "value3")

	fmt.Println("init size:", cache.Size())
	time.Sleep(time.Second * 2)
	cache.Get("key3") // reset expire time.
	fmt.Println("2 Second:", cache.Size())
	time.Sleep(time.Second * 2)
	fmt.Println("4 Second:", cache.Size())
}
