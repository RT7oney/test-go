package main

import (
	"fmt"
)

// var ch = make(chan int, 100)

func main() {
	go f1()
	fmt.Println("我是main")
	// return
	// res := <-ch
	// fmt.Println("结果", res)
}

func f1() {
	// i := 1
	// j := 2
	// k := i + j
	// ch <- k
	fmt.Println("我是f1")
}
