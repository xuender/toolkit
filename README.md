# toolkit

golang toolkit.

## Usage

### cache
```go
package main

import (
	"fmt"
	"time"

	"github.com/xuender/toolkit"
)

func main() {
	cache := toolkit.NewCache(1 * time.Second)
	defer cache.Close()
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	fmt.Println(cache.Get("key1"))
	fmt.Println(cache.Count())
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Count())
}
```

### chmap
```go
package main

import (
	"fmt"
	"time"

	"github.com/xuender/toolkit"
)

func main() {
	cache := toolkit.NewCache(1 * time.Second)
	defer cache.Close()
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	fmt.Println(cache.Get("key1"))
	fmt.Println(cache.Count())
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Count())
}
```
