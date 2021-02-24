package main

import (
	"fmt"
	"github.com/kkkunny/GoMy/pool"
	"time"
)

func main() {
	pl := pool.NewGoroutinePool(10)
	pl.Do(func() {
		for {
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	})
	fmt.Println(pl.GetBusyNumber())
}
