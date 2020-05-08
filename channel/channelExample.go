package main

import "fmt"

func main() {
	number := 123
	squareChan := make(chan int)
	cubeChan := make(chan int)

	go squares(number, squareChan)
	go cubes(number, cubeChan)

	sq := <-squareChan
	cu := <-cubeChan

	fmt.Println("total :", sq+cu)

}

/*求每个数字的值*/
func digits(number int, ch chan int) {
	for number != 0 {
		num := number % 10
		ch <- num

		number /= 10
	}

	close(ch)
}

/*求此数字的平方和*/
func squares(number int, ch chan int) {
	sum := 0

	chl := make(chan int)
	go digits(number, chl)
	for digit := range chl {
		sum += digit * digit
	}
	ch <- sum

}

/*求此数的立方和*/
func cubes(number int, ch chan int) {
	sum := 0
	chl := make(chan int)

	go digits(number, chl)
	for digit := range chl {
		sum += digit * digit * digit
	}

	ch <- sum
}
