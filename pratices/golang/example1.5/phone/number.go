package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xuri/excelize"
)

type rst struct {
	location string `json:"location"`
}

func main() {
	file, err := excelize.OpenFile("F:\\go\\src\\phones\\number\\001.xlsx")

	cell, err := file.GetCellValue("Sheet1", "B1")
	/* rows, _ := file.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	} */
	if err != nil {
		fmt.Println("getCellValue err", err)
		return
	}
	f, _ := http.Get("http://mobsec-dianhua.baidu.com/dianhua_api/open/location?tel=" + string(cell))

	source, _ := ioutil.ReadAll(f.Body)

	s := rst{}
	json.Unmarshal([]byte(source), &s)

	fmt.Println("location:", s)
	fmt.Println(string(source))
	fmt.Println(cell)
}
