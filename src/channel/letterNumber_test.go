package _chan

import (
	"fmt"
	"sync"
	"testing"
)

func TestLetterNumber(t *testing.T) {
	letterNumber()
}

func letterNumber() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d%d", i, i+1)
				i += 2
				letter <- true
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= len(str) {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+2])
				i += 2
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}
