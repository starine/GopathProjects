package _chan

import (
	"testing"
)

func ExampleWorker() {
	for i := 0; i < 10; i++ {
		go Worker()
		wg.Add(1)
	}
	wg.Wait()
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

// 读nil管道会阻塞，不能添加期望值
func TestChanReadNil(t *testing.T) {
	ChanReadNil() //fatal error: all goroutines are asleep - deadlock!
}

// 写nil管道会阻塞，不能添加期望值
func TestChanWriteNil(t *testing.T) {
	ChanWriteNil() //fatal error: all goroutines are asleep - deadlock!
}

func ExampleChanReadClosed() {
	ChanReadClosed()
	// Output:
	// 0
}

func TestChanWriteClosed(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("expected panic")
		}
	}()

	ChanWriteClosed()
}

func ExampleChanCap() {
	ChanCap()
	// Output:
	// 2
	// 10

}
