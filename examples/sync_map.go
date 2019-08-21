package main

import (
	"fmt"

	"github.com/xuender/toolkit"
)

func main() {
	m := toolkit.NewSyncMap()
	m.Set("key1", "value1")
	m.Set("key2", "value2")

	fmt.Println(m.Get("key1"))
	fmt.Println(m.Size())
}
