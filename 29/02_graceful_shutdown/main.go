package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	down := make(chan os.Signal, 1)
	signal.Notify(down, os.Interrupt)

	numChan := make(chan int)

	go startNumberGenerator(numChan)

	for {
		select {
		case <-down:
			fmt.Println("Выхожу из программы")
			return

		default:
			squareChan := calcSquare(numChan)
			fmt.Println(<-squareChan)
		}
	}
}

// calcSquare - create a goroutine
// that receives a number from the "in" channel,
// calculates the square of the number and
// returns the "out" channel
func calcSquare(in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		num := <-in
		out <- num * num
	}()

	return out
}

// startNumberGenerator - Infinitely circularly generate
// natural numbers from 1 to maxInt.
// Generate at time intervals of half a second
// Send the generated number to the "out" channel
func startNumberGenerator(out chan int) {
	num := 1

	for {
		time.Sleep(time.Second / 2)
		out <- num
		num++
	}
}
