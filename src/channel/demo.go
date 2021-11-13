package _chan

import (
	"fmt"
	"time"
)

//单向管道
func readChan(chanName <-chan int) {
	<-chanName
}
func writeChan(chanName chan<- int) {
	chanName <- 1
}

//select
func addNumToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

//for-range
func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get elemet from chanName: %d\n", e)
	}
}
