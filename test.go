package main

import (
	"fmt"
	// "strconv"
	"time"
)

func main() {
	start := 13
	all := 200
	for i := start; i <= all; i++ {
		str := "[" + bar(i, all) + "] " + fmt.Sprintf("%f", (float64(i)/float64(all))*100) + "%"
		fmt.Printf("\r%s", str)
		time.Sleep(1 * time.Second)
	}
}

func bar(now, all int) string {
	// 40是进度条总长度
	size := float64(100) / float64(40)
	str := ""
	k := int(((float64(now) / float64(all)) * 100) / size)
	for i := 0; i < 40; i++ {
		if i <= k {
			str += "="
		} else {
			str += " "
		}
	}
	return str
}
