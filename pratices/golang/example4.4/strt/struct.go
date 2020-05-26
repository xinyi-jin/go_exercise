package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 定义结构体
type Employee struct {
	Id        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 声明结构体变量
// var dilbert Emploryee

func NewEmployee(id int, name string, address string, dob time.Time, pos string, sal int, mangerId int) *Employee {
	self := &Employee{
		Id:        id,
		Name:      name,
		Address:   address,
		DoB:       dob,
		Position:  pos,
		Salary:    sal,
		ManagerID: mangerId,
	}
	return self
}

func (e *Employee) Print() {
	fmt.Println("Employee:", *e)
}

func main() {
	var em Employee
	emp := NewEmployee(1, "小笨猪", "ChinaZhengzhou", time.Now(), "老板", 10000, 2)

	// em = append(em[0:], *emp)
	data, err := json.MarshalIndent(*emp, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("emp:", *emp)

	fmt.Printf("data emp :%s\n", data)

	err = json.Unmarshal(data, &em)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("emp data:%s", em)
}
