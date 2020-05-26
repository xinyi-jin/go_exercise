package main

import "GO_WORK/pratices/example05/db"

func main() {

	/* config_path := flag.String("Config", "./bin/db.cfg", "config file path!")

	flag.Parse()
	fmt.Println("config", *config_path)

	var a []int

	s := append(a, 1) */

	db.InitDB()
	db.Delete(100)

}
