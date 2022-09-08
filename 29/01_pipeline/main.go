package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	for {
		var wg sync.WaitGroup
		wg.Add(2)

		fmt.Print("Ввод: ")
		var input string
		fmt.Scan(&input)

		if input == "стоп" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите число еще раз")
			continue
		}

		squareChan := calcSquare(num, &wg)
		doubleNumber(squareChan, &wg)
		wg.Wait()
	}
}

// calcSquare - run a goroutine
// that calculates the square of a number and
// sends the result to a channel called out and
// displays the result on the screen
func calcSquare(num int, wg *sync.WaitGroup) chan int {
	outChan := make(chan int)

	go func() {
		defer close(outChan)
		defer wg.Done()

		square := num * num
		outChan <- square
		fmt.Println("Квадрат:", square)
	}()

	return outChan
}

// doubleNumber - run a goroutine that receives data from the "in" channel and
// multiplies the number by two and
// displays the result on the screen
func doubleNumber(inChan chan int, wg *sync.WaitGroup) {

	go func() {
		defer wg.Done()

		result := <-inChan
		result *= 2
		fmt.Println("Произведение:", result)
	}()
}
