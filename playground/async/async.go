package async

import (
	"fmt"
	"time"
)

func WaitTest(text string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow running message", text)
}

func WaitTestChannel(text string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow running message with channel", text)
	done <- true
	close(done)
}

func GoroutineWaitTest(text string) {
	go WaitTest(text)
}

func ChannelWaitTest(text string) {
	done := make(chan bool)
	go WaitTestChannel("Hello!", done)

	fmt.Println("Channel result", <-done)
}

func ChannelWaitTestMultiple(text string) {
	waits := make([]chan bool, 4)

	waits[0] = make(chan bool)
	go WaitTestChannel("Hello!", waits[0])

	waits[1] = make(chan bool)
	go WaitTestChannel("Hello 2!", waits[1])

	waits[2] = make(chan bool)
	go WaitTestChannel("Hello 3!", waits[2])

	waits[3] = make(chan bool)
	go WaitTestChannel("Hello 4!", waits[3])

	for _, done := range waits {
		<-done
	}
}
