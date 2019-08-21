# toolkit

golang toolkit.

## Usage

### Cache
```go
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
```

### SyncMap
```go
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

```

### ChMap
```go
package main

import (
	"fmt"

	"github.com/xuender/toolkit"
)

func main() {
	chMap := toolkit.NewChMap()
	defer chMap.Close()
	chMap.Set("key1", "value1")
	chMap.Set("key2", "value2")

	fmt.Println(chMap.Get("key1"))
	fmt.Println(chMap.Size())
}

```
