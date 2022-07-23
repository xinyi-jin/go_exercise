package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/xuri/excelize"
)

const (
	fileUrl      = "F:\\go\\src\\phone\\Number\\50021.xls"                          //源文件路径
	apiUrl       = "http://mobsec-dianhua.baidu.com/dianhua_api/open/location?tel=" //调用接口地址
	numbers      = 50022                                                            //读取的数据行数
	aimSheetName = "Sheet1"
	aimFileName  = "50021fb.xls"
)

type Res struct {
	Response map[string]In1
}

type In1 struct {
	Location string
}

type mes struct {
	Name   string
	Number string
}

func readFile() {
	f, err := excelize.OpenFile("F:\\go\\src\\phone\\number\\50021.xls")
	if err != nil {
		fmt.Println(err)
		return
	}
	phoneNums := []string{}
	for i := 1; i < 50022; i++ {
		b1 := "B" + strconv.Itoa(i)
		cell, err := f.GetCellValue("Sheet1", b1)
		if err != nil {
			fmt.Println(err)
			return
		}
		phoneNums = append(phoneNums, cell)
	}

	file := excelize.NewFile()
	index := file.NewSheet("Sheet1")
	row := 1
	for _, num := range phoneNums {
		url := "http://mobsec-dianhua.baidu.com/dianhua_api/open/location?tel=" + num
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		bodyString, err := ioutil.ReadAll(resp.Body)
		str := string(bodyString)
		res := Res{}
		json.Unmarshal([]byte(str), &res)
		loca := res.Response[num].Location

		a1 := "A" + strconv.Itoa(row)
		file.SetCellValue("Sheet1", a1, loca)
		row++
		fmt.Println(row)
	}
	file.SetActiveSheet(index)
	errs := file.SaveAs("50021fb.xls")
	if err != nil {
		fmt.Println(errs)
	}
}

func writeFile(mes []mes) {
	file := excelize.NewFile()
	index := file.NewSheet(aimSheetName)
	row := 1
	for _, mesData := range mes {
		// 先写入原有数据，避免数据丢失
		a1 := "A" + strconv.Itoa(row)
		b1 := "B" + strconv.Itoa(row)
		c1 := "C" + strconv.Itoa(row)
		file.SetCellValue("Sheet1", a1, mesData.Name)
		file.SetCellValue("Sheet1", b1, mesData.Number)
		// 跳过非法的电话号码
		if len(mesData.Number) != 11 {
			row++
			continue
		}
		url := apiUrl + mesData.Number
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		bodyString, err := ioutil.ReadAll(resp.Body)
		str := string(bodyString)
		res := Res{}
		json.Unmarshal([]byte(str), &res)
		loca := res.Response[mesData.Number].Location

		// 写入接口返回的数据
		file.SetCellValue("Sheet1", c1, loca)
		fmt.Println(row)
		row++
	}
	file.SetActiveSheet(index)
	errs := file.SaveAs(aimFileName)
	if errs != nil {
		fmt.Println(errs)
	}
}

func readFileEx() []mes {
	f, err := excelize.OpenFile(fileUrl)
	if err != nil {
		fmt.Println(err)
	}
	mesData := []mes{}
	for i := 1; i < numbers; i++ {
		mes := mes{}
		a1 := "A" + strconv.Itoa(i)
		b1 := "B" + strconv.Itoa(i)
		name, err := f.GetCellValue("Sheet1", a1)
		number, err := f.GetCellValue("Sheet1", b1)
		mes.Name = name
		mes.Number = number
		if err != nil {
			fmt.Println(err)
		}
		mesData = append(mesData[0:], mes)
	}
	return mesData

}

/*
func write() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Hello world.")
	f.SetCellValue("Sheet1", "B1", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := f.SaveAs("14fb.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
*/
