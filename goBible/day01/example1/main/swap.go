package main

import "fmt"

func main() {

	fmt.Println("swap before:")
	fmt.Println("A:", A)
	fmt.Println("B:", B)
	//swap(&A,&B)
	swap1(A, B)

	fmt.Println(swap1(A, B))
	fmt.Println("swap after:")
	fmt.Println("A:", A)
	fmt.Println("B:", B)
}

type Swap int

var (
	A Swap = 100
	B Swap = 200
)

func swap(a *Swap, b *Swap) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func swap1(a Swap, b Swap) (c Swap, d Swap) {

	return b, a
}
