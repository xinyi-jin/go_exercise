package main

import "fmt"

func main() {
	var c = &Car{
		LunTai: "block",
		LunZi:  "two",
		Sorts:  "bike",
	}

	var b = &Bike{
		Color: "red",
	}
	b.Car = *c

	fmt.Printf("%s", b.string())

}

type Car struct {
	LunZi  string
	LunTai string
	Sorts  string
}

type Bike struct {
	Car
	Color string
}

func (b *Bike) string() string {
	return fmt.Sprintf("%s车是%s颜色的", b.Sorts, b.Color)
}
