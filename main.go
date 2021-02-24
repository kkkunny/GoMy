package main

import (
	"fmt"
	"github.com/kkkunny/GoMy/pool"
	"time"
)

func main() {
	pl := pool.NewGoroutinePool(2)
	pl.Do(func() {
		for i := 0; i < 5; i++ {
			fmt.Println(1)
			time.Sleep(1 * time.Second)
		}
	})
	pl.Do(func() {
		for i := 0; i < 5; i++ {
			fmt.Println(2)
			time.Sleep(1 * time.Second)
		}
	})
	pl.Do(func() {
		for {
			fmt.Println(3)
			time.Sleep(1 * time.Second)
		}
	})
	time.Sleep(2 * time.Second)
}
