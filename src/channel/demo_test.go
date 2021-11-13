package _chan

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteRead(t *testing.T) {
	var mychan = make(chan int, 10)
	writeChan(mychan)
	readChan(mychan)
}

//select
func TestChanSelect(t *testing.T) {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)

	go addNumToChan(chan1)
	go addNumToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(time.Second)
		}
	}
}

//for-range
func TestChanRange(t *testing.T) {
	mychan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		mychan <- i
	}
	/*for e := range mychan {
		fmt.Printf("Get elemet from chanName: %d\n", e)
	}*/
	for i := 0; i < 11; i++ {
		e := <-mychan
		fmt.Printf("Get elemet from chanName: %d\n", e)
	}
}
