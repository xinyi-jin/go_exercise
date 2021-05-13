package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Reflect struct {
	name    string
	age     int
	address string
}

func TestPKGReflect(t *testing.T) {
	var roll interface{}
	r := &Reflect{"name", 24, "zhengzhou,henan"}
	roll = r
	nt := reflect.TypeOf(roll)
	nv := reflect.ValueOf(roll)
	fmt.Printf("type %v \nvalue %v \n", nt, nv)
	res := roll.(*Reflect)
	fmt.Printf("res %v \n", res)
}
