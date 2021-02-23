package main

import (
	"fmt"
	"github.com/kkkunny/GoMy/container"
)

func main() {
	set1 := container.NewSet()
	for i := 0; i < 10; i++ {
		set1.Add(i)
	}
	set2 := container.NewSet()
	for i := 5; i < 15; i++ {
		set2.Add(i)
	}
	set1.Differed(set2)
	fmt.Println(set1.GetString())
}
