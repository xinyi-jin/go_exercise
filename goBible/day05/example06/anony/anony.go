package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name  string
	sorts string
}

type Trans struct {
	Cart
	int
	time time.Time
}

/* 匿名结构体参数 参数名出现冲突的时候，按照tran.name 和 tran.Cart.name 进行区分使用
当tran.Cart.name 和tran.Bike.name 同时存在的时候，不能使用tran.name ，因为程序也找不到是指的哪一个name，必须进行指定 */
func main() {

	var car = new(Cart)
	var tran = new(Trans)
	car.name = "朱贺"
	tran.name = "zhuhe"
	tran.sorts = "huoche"

	tran.int = 108
	tran.time = time.Now()

	fmt.Println(car)
	fmt.Println(tran)
}
