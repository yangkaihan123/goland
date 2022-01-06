package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	//t := time.Now()
	//fmt.Println(t)                                               // 2022-01-06 10:23:25.290428 +0800 CST m=+0.003196501
	//fmt.Printf("%02d.%02d.%02d\n", t.Year(), t.Month(), t.Day()) //06.01.2022
	t := time.Now().UTC()
	//fmt.Println(t)
	//fmt.Println(time.Now())
	// calculating times 计算次数
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosecond
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // 2022-01-13 02:30:53.4314787 +0000 UTC
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) //06 Jan 22 02:31 UTC
	fmt.Println(t.Format(time.ANSIC))  //Thu Jan  6 02:32:30 2022
}
