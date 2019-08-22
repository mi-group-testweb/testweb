package main

import (
	"time"
)

func main() {
	var num int64
	for num <= 10{
		num++
		timeInt := time.Now().Unix()
		timestamp := (timeInt - timeInt%60) * 1000
		WriteLogToElk(timestamp, num)
		time.Sleep(time.Minute)
	}
	return
}