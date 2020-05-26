package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/xuri/excelize"
)

type Res struct {
	Response map[string]In1
}

type In1 struct {
	Location string
}

func main() {
	readFile()
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
