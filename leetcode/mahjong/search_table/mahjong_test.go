package searchtable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMahjongTable(t *testing.T) {
	/* pairs := getPairs()
	groups := getGroups()

	allCombine := getAllHuCombine(pairs, groups)
	for _, v := range allCombine {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcKey(nums)
			if _, ok := mahjongTable[k]; !ok {
				fmt.Println("mahjongTable not found ", k)
			}
		}
	}
	for _, v := range allCombine {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcCardsKeyBase(v)
			if _, ok := mahjongTable1[k]; !ok {
				fmt.Println("mahjongTable1 not found ", k)
			}
		}
	}
	for _, v := range allCombine {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcKeyBase(nums)
			if _, ok := mahjongTable2[k]; !ok {
				fmt.Println("mahjongTable2 not found ", k)
			}
		}
	} */

	nohu := getAllNotHuCombine()
	for _, v := range nohu {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcKey(nums)
			if _, ok := mahjongTable[k]; ok {
				fmt.Println("mahjongTable nohu ", k)
			}
		}
	}
	/* for _, v := range nohu {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcCardsKeyBase(v)
			if _, ok := mahjongTable1[k]; ok {
				fmt.Println("mahjongTable1 nohu ", k)
			}
		}
	}
	for _, v := range nohu {
		nums := count(v)
		if checkIsValid(nums) {
			k := calcKeyBase(nums)
			if _, ok := mahjongTable2[k]; ok {
				fmt.Println("mahjongTable2 nohu ", k)
			}
		}
	} */
}

var mahjongTable map[int]int
var mahjongTable1 map[int]int
var mahjongTable2 map[int]int

func init() {
	f, err := os.Open("../huCards.json")
	if err != nil {
		log.Fatal("Open", err)
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	json.Unmarshal([]byte(bytes), &mahjongTable)
	fmt.Println("mahjongTable len ", len(mahjongTable))

	f1, err := os.Open("../huCards1.json")
	if err != nil {
		log.Fatal("Open", err)
	}
	defer f1.Close()
	bytes1, _ := ioutil.ReadAll(f1)
	json.Unmarshal([]byte(bytes1), &mahjongTable1)
	fmt.Println("mahjongTable 1 len ", len(mahjongTable1))

	f2, err := os.Open("../huCards2.json")
	if err != nil {
		log.Fatal("Open", err)
	}
	defer f2.Close()
	bytes2, _ := ioutil.ReadAll(f2)
	json.Unmarshal([]byte(bytes2), &mahjongTable2)
	fmt.Println("mahjongTable 2 len ", len(mahjongTable2))
}
