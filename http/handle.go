package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var WinRate float64
var WinCnt, Total int64

type GDYPlayers struct {
	Snid           int32
	Phone          string
	WinScore       int64
	SurplusCardCnt int
	IsBanker       bool
}
type GDYData struct {
	Players []GDYPlayers
}

type TestData struct {
	PlayerWinRate
	Players     map[int32]*PlayersData
	SeatWinCnt  [4]int64
	SeatWinRate [4]float64
}

type PlayersData struct {
	PlayerWinRate
	GDYPlayers
}

type PlayerWinRate struct {
	WinRate float64
	WinCnt  int64
	Total   int64
}

var testData = TestData{
	Players: make(map[int32]*PlayersData),
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
	h3 := func(w http.ResponseWriter, _ *http.Request) {
		str := fmt.Sprintf("gdy:\nWinCnt:%v Total:%v WinRate:%v%%\n", testData.WinCnt, testData.Total, testData.WinRate*100)

		str += "=================================\n"
		for i, v := range testData.SeatWinRate {
			str += fmt.Sprintf("Seat:%v, WinCnt:%v Total:%v WinRate:%v%%\n", i, testData.SeatWinCnt[i], testData.Total, v*100)
		}

		str += "=================================\n"
		for _, v := range testData.Players {
			str += fmt.Sprintf("SnID:%v, WinCnt:%v Total:%v WinRate:%v%%\n", v.Snid, v.WinCnt, v.Total, v.WinRate*100)
		}
		io.WriteString(w, str)
	}
	h4 := func(w http.ResponseWriter, r *http.Request) {
		var temp = GDYData{
			Players: make([]GDYPlayers, 0),
		}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, "io error")
		}
		json.Unmarshal(data, &temp)

		for i, v := range temp.Players {
			if _, ok := testData.Players[v.Snid]; !ok {
				testData.Players[v.Snid] = &PlayersData{}
			}
			p := testData.Players[v.Snid]
			if v.WinScore > 0 {
				testData.SeatWinCnt[i]++
				p.WinCnt++
				if v.IsBanker {
					testData.WinCnt++
				}
			}
			p.GDYPlayers = v
			p.Total++
			p.WinRate = float64(p.WinCnt) / float64(p.Total)
			fmt.Printf("player:%+v\n", v)
		}
		testData.Total++
		testData.WinRate = float64(testData.WinCnt) / float64(testData.Total)
		for i, v := range testData.SeatWinCnt {
			testData.SeatWinRate[i] = float64(v) / float64(testData.Total)
		}
		fmt.Printf("WinCnt:%v Total:%v WinRate:%v%%\n\n", testData.WinCnt, testData.Total, testData.WinRate*100)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	http.HandleFunc("/gdy-ai", h3)
	http.HandleFunc("/gdy-ai-data", h4)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
