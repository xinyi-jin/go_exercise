package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "xiaoshazhu"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func Insert(id int, name string, sex string) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx failed")
		return false
	}
	// 预编译SQL语句
	stmt, err := tx.Prepare("INSERT INTO student(`id`,`name`,`sex`)VALUES(?,?,?)")
	if err != nil {
		fmt.Println("ex failed")
		return false
	}
	// 传参
	res, err := stmt.Exec(id, name, sex)
	if err != nil {
		fmt.Println("res failed")
		return false
	}
	// 提交事务
	tx.Commit()

	// 获取上一个插入自增的ID
	fmt.Println(res.LastInsertId())

	return true
}

func Delete(id int) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		return false
	}

	// 预编译
	stmt, err := tx.Prepare("DELETE FROM student WHERE id = ?")
	if err != nil {
		return false
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false
	}
	fmt.Println(res)

	tx.Commit()
	return true
}

func Update(id int, name string) {

}
