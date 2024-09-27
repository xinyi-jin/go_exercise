package rldqn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	/*0~8(1到9筒)*/
	MJ_DOTS_1 int64 = iota
	MJ_DOTS_2
	MJ_DOTS_3
	MJ_DOTS_4
	MJ_DOTS_5
	MJ_DOTS_6
	MJ_DOTS_7
	MJ_DOTS_8
	MJ_DOTS_9
	/*9~17(1到9索)*/
	MJ_BAMBOO_1
	MJ_BAMBOO_2
	MJ_BAMBOO_3
	MJ_BAMBOO_4
	MJ_BAMBOO_5
	MJ_BAMBOO_6
	MJ_BAMBOO_7
	MJ_BAMBOO_8
	MJ_BAMBOO_9
	/*18~26(1到9万)*/
	MJ_CHARACTERS_1
	MJ_CHARACTERS_2
	MJ_CHARACTERS_3
	MJ_CHARACTERS_4
	MJ_CHARACTERS_5
	MJ_CHARACTERS_6
	MJ_CHARACTERS_7
	MJ_CHARACTERS_8
	MJ_CHARACTERS_9
	/*27~30(东南西北风)*/
	MJ_WIND_EAST
	MJ_WIND_SOUTH
	MJ_WIND_WEST
	MJ_WIND_NORTH
	/*31~33(中发白)*/
	MJ_DRAGON_RED
	MJ_DRAGON_GREEN
	MJ_DRAGON_WHITE
	/* 34~41春夏秋冬，梅兰竹菊 */
	MJ_FLOWER_SPRING
	MJ_FLOWER_SUMMER
	MJ_FLOWER_AUTUMN
	MJ_FLOWER_WINTER
	MJ_FLOWER_PLUM
	MJ_FLOWER_ORCHID
	MJ_FLOWER_BAMBOO
	MJ_FLOWER_CHRYSANTHEMUN
	/*最大42*/
	MJ_MAX
)

var Card2actionstr = map[int64]string{
	/*0~8(1到9筒)*/
	MJ_DOTS_1: "dots-1",
	MJ_DOTS_2: "dots-2",
	MJ_DOTS_3: "dots-3",
	MJ_DOTS_4: "dots-4",
	MJ_DOTS_5: "dots-5",
	MJ_DOTS_6: "dots-6",
	MJ_DOTS_7: "dots-7",
	MJ_DOTS_8: "dots-8",
	MJ_DOTS_9: "dots-9",
	/*9~17(1到9索)*/
	MJ_BAMBOO_1: "bamboo-1",
	MJ_BAMBOO_2: "bamboo-2",
	MJ_BAMBOO_3: "bamboo-3",
	MJ_BAMBOO_4: "bamboo-4",
	MJ_BAMBOO_5: "bamboo-5",
	MJ_BAMBOO_6: "bamboo-6",
	MJ_BAMBOO_7: "bamboo-7",
	MJ_BAMBOO_8: "bamboo-8",
	MJ_BAMBOO_9: "bamboo-9",
	/*18~26(1到9万)*/
	MJ_CHARACTERS_1: "characters-1",
	MJ_CHARACTERS_2: "characters-2",
	MJ_CHARACTERS_3: "characters-3",
	MJ_CHARACTERS_4: "characters-4",
	MJ_CHARACTERS_5: "characters-5",
	MJ_CHARACTERS_6: "characters-6",
	MJ_CHARACTERS_7: "characters-7",
	MJ_CHARACTERS_8: "characters-8",
	MJ_CHARACTERS_9: "characters-9",
	/*27~30(东南西北风)*/
	MJ_WIND_EAST:  "winds-east",
	MJ_WIND_SOUTH: "winds-south",
	MJ_WIND_WEST:  "winds-west",
	MJ_WIND_NORTH: "winds-north",
	/*31~33(中发白)*/
	MJ_DRAGON_RED:   "dragons-red",
	MJ_DRAGON_GREEN: "dragons-green",
	MJ_DRAGON_WHITE: "dragons-white",
}
var Actionstr2card = map[string]int64{
	/*0~8(1到9筒)*/
	"dots-1": MJ_DOTS_1,
	"dots-2": MJ_DOTS_2,
	"dots-3": MJ_DOTS_3,
	"dots-4": MJ_DOTS_4,
	"dots-5": MJ_DOTS_5,
	"dots-6": MJ_DOTS_6,
	"dots-7": MJ_DOTS_7,
	"dots-8": MJ_DOTS_8,
	"dots-9": MJ_DOTS_9,
	/*9~17(1到9索)*/
	"bamboo-1": MJ_BAMBOO_1,
	"bamboo-2": MJ_BAMBOO_2,
	"bamboo-3": MJ_BAMBOO_3,
	"bamboo-4": MJ_BAMBOO_4,
	"bamboo-5": MJ_BAMBOO_5,
	"bamboo-6": MJ_BAMBOO_6,
	"bamboo-7": MJ_BAMBOO_7,
	"bamboo-8": MJ_BAMBOO_8,
	"bamboo-9": MJ_BAMBOO_9,
	/*18~26(1到9万)*/
	"characters-1": MJ_CHARACTERS_1,
	"characters-2": MJ_CHARACTERS_2,
	"characters-3": MJ_CHARACTERS_3,
	"characters-4": MJ_CHARACTERS_4,
	"characters-5": MJ_CHARACTERS_5,
	"characters-6": MJ_CHARACTERS_6,
	"characters-7": MJ_CHARACTERS_7,
	"characters-8": MJ_CHARACTERS_8,
	"characters-9": MJ_CHARACTERS_9,
	/*27~30(东南西北风)*/
	"winds-east":  MJ_WIND_EAST,
	"winds-south": MJ_WIND_SOUTH,
	"winds-west":  MJ_WIND_WEST,
	"winds-north": MJ_WIND_NORTH,
	/*31~33(中发白)*/
	"dragons-red":   MJ_DRAGON_RED,
	"dragons-green": MJ_DRAGON_GREEN,
	"dragons-white": MJ_DRAGON_WHITE,
}

type MahjongData struct {
	Hands     []int64
	Discards  []int64
	Piles     [][]int64
	ActionStr string
}

type MahjonParams struct {
	Hand  []string   `json:"hand,omitempty"`
	Piles [][]string `json:"piles,omitempty"`
	Table []string   `json:"table,omitempty"`
}

func NewMahjongData() *MahjongData {
	return &MahjongData{
		Hands:    make([]int64, 0),
		Discards: make([]int64, 0),
		Piles:    make([][]int64, 4),
	}
}

func Card2str(card int64) string {
	return Card2actionstr[card]
}

func Str2card(str string) int64 {
	return Actionstr2card[str]
}

func Cards2str(cards []int64) []string {
	res := make([]string, 0)
	for _, v := range cards {
		res = append(res, Card2str(v))
	}
	return res
}

func Evaluation(mahjongData *MahjongData) {
	// mahjongData := MahjongData{
	// 	Hands: []int64{
	// 		MJ_BAMBOO_1,
	// 		MJ_BAMBOO_2,
	// 		MJ_BAMBOO_3,
	// 		MJ_CHARACTERS_9,
	// 		MJ_CHARACTERS_9,
	// 	},
	// 	Piles: [][]int64{
	// 		{
	// 			MJ_BAMBOO_5,
	// 			MJ_BAMBOO_5,
	// 			MJ_BAMBOO_5,
	// 		},
	// 		{},
	// 		{},
	// 		{},
	// 	},
	// 	Discards: []int64{
	// 		MJ_CHARACTERS_5,
	// 		MJ_CHARACTERS_5,
	// 		MJ_CHARACTERS_5,
	// 	},
	// }
	data := MahjonParams{
		Hand: Cards2str(mahjongData.Hands),
		Piles: [][]string{
			Cards2str(mahjongData.Piles[0]),
			Cards2str(mahjongData.Piles[1]),
			Cards2str(mahjongData.Piles[2]),
			Cards2str(mahjongData.Piles[3]),
		},
		Table: Cards2str(mahjongData.Discards),
	}

	/* data := MahjonParams{
		Hand: []string{
			"bamboo-1",
			"bamboo-2",
			"bamboo-3",
			"characters-9",
			"characters-9",
		},
		Piles: [][]string{
			{
				"bamboo-5",
				"bamboo-5",
				"bamboo-5",
			},
			{},
			{},
			{},
		},
		Table: []string{
			"characters-5",
			"characters-5",
			"characters-5",
		},
	} */

	byteBody, err := json.Marshal(data)
	fmt.Println("json:", string(byteBody))
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8888", bytes.NewBuffer(byteBody))

	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	mahjongData.ActionStr = string(resBody)
	fmt.Println("resBody", string(resBody))
}
