package main

import (
	"GOPROJECT/dbServer/db"
	"fmt"
	"os"
)

func main() {
	db.InitDB()

	// if res := db.Insert(1, "benzhu"); res {
	// 	fmt.Println("insert sucess")
	// }

	// stu := db.Query(1)
	// fmt.Println(stu)

	stus := db.QueryAll()
	file, err := os.OpenFile("f:/write_log_test.log", os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("openFile err:", err)
	}
	for _, v := range stus {
		// fmt.Println("stu:", v)
		fmt.Fprintln(file, "write test ...")
		fmt.Fprintln(file, v)
	}

}
