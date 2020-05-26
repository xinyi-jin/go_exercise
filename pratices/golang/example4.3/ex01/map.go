package main

import (
	"fmt"
	"sort"
)

func main() {
	/* names := map[string]int{
		"name1": 1,
		"name2": 2,
		"name3": 3,
		"name4": 4,
		"name5": 5,
		"name6": 6,
		"name7": 7,
		"name8": 8,
	}

	for v := range names {
		fmt.Println("v", v)
	} */

	var nst map[string]int
	var sl []string
	nst = make(map[string]int)
	sl = make([]string, len(nst))

	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("name%d", i)
		nst[str] = i
	}

	for v := range nst {
		sl = append(sl, v)
	}

	sort.Strings(sl)

	delete(nst, "name6")

	if re, ok := nst["name6"]; !ok {
		fmt.Println("error not found", re)
	}
	for _, v := range sl {
		fmt.Printf("key:%s, value:%d\n", v, nst[v])
	}
}
