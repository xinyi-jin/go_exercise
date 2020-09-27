package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User
type User struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {
	// 建立mysql连接
	db, err := gorm.Open("mysql", "root:xiaoshazhu@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("db init err:%v", err)
	}
	defer db.Close()

	// 设置自动迁移
	db.AutoMigrate(&User{})
	// u1 := &User{
	// 	ID:      1,
	// 	Name:    "zhuzhu",
	// 	Age:     23,
	// 	Address: "zhumadian",
	// }
	// // 创建用户
	// db.Create(u1)

	// // 查询用户
	// var u = new(User)
	// db.First(u)
	// fmt.Printf("%#v\n", u)

	// var uu User
	// db.Find(&uu, "Address=?", "zhumadian")
	// fmt.Printf("%#v\n", uu)

	// // 更新
	var u User
	// db.Model(&u).Update("Address", "zhengzhou")
	// // 删除
	db.Delete(&u)
}
