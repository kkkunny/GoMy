package main

import (
	"GoMy/worker"
	"fmt"
	"time"
)

func main() {
	wok := worker.NewWorker(1, Test, true)
	wok.Start()
	wok.Wait()
}

func Test() {
	fmt.Println(11111)
	time.Sleep(1 * time.Second)
}
