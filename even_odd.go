package main

import (
	"log"
)

func EvenOdd() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	go generateEven(evenChan)
	go generateOdd(oddChan)

	for {
		select {
		case even, ok := <-evenChan:
			if ok {
				log.Printf("even: %d", even)
			} else {
				evenChan = nil
			}
		case odd, ok := <-oddChan:
			if ok {
				log.Printf("odd: %d", odd)
			} else {
				oddChan = nil
			}
		}

		if evenChan == nil && oddChan == nil {
			return
		}
	}
}

func generateEven(evenChan chan int) {
	for i := 0; i <= 10; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func generateOdd(oddChan chan int) {
	for i := 1; i <= 10; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}
