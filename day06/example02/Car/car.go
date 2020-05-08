package main

import "fmt"

type Car interface {
	run()
	drink()
	getTotal()
}
type BWM struct {
	Name  string
	Total int
}

func (b *BWM) run() {
	fmt.Println("running...", b.Name)
}

func (b *BWM) drink() {
	fmt.Println("this car is drinking...", b.Name)
}
func (b *BWM) getTotal() {
	fmt.Println("this car has ", b.Total)
}

func main() {
	bwm := &BWM{
		Name:  "bwm",
		Total: 4,
	}
	bwm.run()
	bwm.drink()
	bwm.getTotal()
}
