package db

import (
	"GOPROJECT/dbServer/bean"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接配置
const (
	userName = "root"
	password = "xiaoshazhu"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

// 数据库连接池
var DB *sql.DB

// 初始化数据库连接池
func InitDB() {
	// 拼接数据库连接信息
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)

	// 设置最大连接数
	DB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(10)

	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database failed")
		return
	}

	fmt.Println("connect database success")
}

// 插入数据
func Insert(id int, name string) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx failed")
		return false
	}

	// 预编译SQL
	stmt, err := DB.Prepare("INSERT INTO student (ID,NAME)VALUES(?,?)")
	if err != nil {
		fmt.Println("stmt failed")
		return false
	}

	// 传递参数执行
	res, err := stmt.Exec(id, name)
	if err != nil {
		fmt.Println("result failed")
		return false
	}

	res.LastInsertId()
	// 提交事务
	tx.Commit()

	return true
}

// 查询
func Query(id int) bean.Student {
	var stu bean.Student
	err := DB.QueryRow("SELECT * FROM student WHERE id = ?", id).Scan(&stu.Id, &stu.Name)
	if err != nil {
		fmt.Println("query failed")
	}
	return stu
}

// 查询所有信息
func QueryAll() []bean.Student {
	var stus []bean.Student
	rows, err := DB.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println(" queryAll failed")
	}

	for rows.Next() {
		var stu bean.Student
		rows.Scan(&stu.Id, &stu.Name)
		stus = append(stus, stu)
	}
	return stus

}
