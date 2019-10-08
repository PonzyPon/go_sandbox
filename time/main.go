package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())       // 2019
	fmt.Println(now.Month())      // September
	fmt.Println(now.Day())        // 26
	fmt.Println(now.Hour())       // 15
	fmt.Println(now.Minute())     // 20
	fmt.Println(now.Second())     // 10
	fmt.Println(now.Nanosecond()) // 790612000
	fmt.Println(now.Unix())       // 1569478810
}
