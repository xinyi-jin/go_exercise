package chow

import (
	"hnmatch/gamerule/beard"
	"log"
	"math/rand"
)

type Node struct {
	Cards    []int64
	Children []*Node
}

func NewNode(cards []int64) *Node {
	return &Node{
		Cards:    cards,
		Children: []*Node{},
	}
}

func (n *Node) AddNode(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) Traver() {
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		for _, v := range n.Children {
			log.Printf("cards %v \n", v.Cards)
			dfs(v)
		}
	}
	dfs(n)
}

func GetChowCards(pool [beard.BEARD_MAX]int, c int64) []int64 {
	cards := make([]int64, 0)
	cnt := 1
	for {
		node := searchChowCards(pool, c)
		if node != nil {
			n := len(node.Children)
			if n == 0 {
				break
			}
			pos := rand.Intn(n)
			cards = append(cards, node.Children[pos].Cards...)
			if pool[c] == 0 {
				break
			}
			pool[cards[len(cards)-3]]--
			pool[cards[len(cards)-2]]--
			if cnt != 1 {
				pool[cards[len(cards)-1]]--
			}
			if pool[c] == 0 {
				break
			}
			cnt++
			pool[c]--
		}
	}
	return cards
}

func searchChowCards(pool [beard.BEARD_MAX]int, c int64) *Node {
	res := NewNode([]int64{})
	if c < 0 || c >= beard.BEARD_MAX {
		return res
	}
	// 移除3张牌
	for i := beard.BEARD_SMALL_1; i < beard.BEARD_MAX; i++ {
		if pool[i] > 2 {
			pool[i] = 0
		}
	}

	if isJiao(pool, c) {
		otherCard := c
		if IsBigCard(c) {
			otherCard -= 10
		} else {
			otherCard += 10
		}
		if pool[otherCard] >= 2 {
			cards := []int64{otherCard, otherCard, c}
			res.Children = append(res.Children, NewNode(cards))
		}
		if pool[otherCard] >= 1 && pool[c] >= 1 {
			cards := []int64{otherCard, c, c}
			res.Children = append(res.Children, NewNode(cards))
		}
	}
	if is27A(pool, c) {
		cards := []int64{beard.BEARD_INVAILD, beard.BEARD_INVAILD, c}
		switch c % 10 {
		case 1:
			cards[0] = c + 5
			cards[1] = c + 8
		case 6:
			cards[0] = c - 5
			cards[1] = c + 3
		case 9:
			cards[0] = c - 8
			cards[1] = c - 3
		}
		if cards[0] != beard.BEARD_INVAILD {
			res.Children = append(res.Children, NewNode(cards))
		}
	}
	if inSequence(pool, c) {
		tempPool := pool
		cards := []int64{c - 3, c - 2, c - 1}
		tempPool[c]++
		for i := 0; i < 3; i++ {
			cards[0]++
			cards[1]++
			cards[2]++
			temp := make([]int64, 3)
			copy(temp, cards)
			if IsSameType(temp[0], temp[1]) && IsSameType(temp[0], temp[2]) &&
				tempPool[temp[0]] > 0 && tempPool[temp[1]] > 0 && tempPool[temp[2]] > 0 {
				if temp[0] == c {
					temp[0], temp[1], temp[2] = temp[1], temp[2], temp[0]
				}
				if temp[1] == c {
					temp[1], temp[2] = temp[2], temp[1]
				}
				res.Children = append(res.Children, NewNode(temp))
			}
		}
	}
	return res
}

// 手牌能不能吃
func CanChow(pool [beard.BEARD_MAX]int, c int64) bool {
	if c < 0 || c >= beard.BEARD_MAX {
		return false
	}

	// 移除3张牌
	for i := beard.BEARD_SMALL_1; i < beard.BEARD_MAX; i++ {
		if pool[i] > 2 {
			pool[i] = 0
		}
	}

	if isJiao(pool, c) {
		return true
	}

	if is27A(pool, c) {
		return true
	}

	return inSequence(pool, c)
}

func isJiao(pool [beard.BEARD_MAX]int, c int64) bool {
	if pool[c] == 0 {
		return false
	}
	otherCard := c
	if IsBigCard(c) {
		otherCard -= 10
	} else {
		otherCard += 10
	}

	switch pool[c] {
	case 0:
		if pool[otherCard] == 2 {
			return true
		}
	case 1, 2:
		if pool[otherCard] > 0 && pool[otherCard] < 3 {
			return true
		}
	}
	return false
}

func is27A(pool [beard.BEARD_MAX]int, c int64) bool {
	checkCardNum := func(a, b, d int64) bool {
		pool[c]++
		return pool[a] > 0 && pool[a] < 4 &&
			pool[b] > 0 && pool[b] < 4 &&
			pool[d] > 0 && pool[d] < 4
	}

	if IsBigCard(c) {
		if c != beard.BEARD_BIG_2 && c != beard.BEARD_BIG_7 && c != beard.BEARD_BIG_A {
			return false
		}
		return checkCardNum(beard.BEARD_BIG_2, beard.BEARD_BIG_7, beard.BEARD_BIG_A)
	} else {
		if c != beard.BEARD_SMALL_2 && c != beard.BEARD_SMALL_7 && c != beard.BEARD_SMALL_A {
			return false
		}
		return checkCardNum(beard.BEARD_SMALL_2, beard.BEARD_SMALL_7, beard.BEARD_SMALL_A)
	}
}

func inSequence(pool [beard.BEARD_MAX]int, c int64) bool {
	cards := []int64{c - 3, c - 2, c - 1}
	pool[c]++
	for i := 0; i < 3; i++ {
		cards[0]++
		cards[1]++
		cards[2]++
		if IsSameType(cards[0], cards[1]) && IsSameType(cards[0], cards[2]) &&
			pool[cards[0]] > 0 && pool[cards[1]] > 0 && pool[cards[2]] > 0 {
			return true
		}
	}
	return false
}

func IsBigCard(card int64) bool {
	return card >= beard.BEARD_BIG_1 && card < beard.BEARD_MAX
}

func IsSameType(c1, c2 int64) bool {
	if c1 < 0 || c2 < 0 ||
		c1 >= beard.BEARD_MAX || c2 >= beard.BEARD_MAX {
		return false
	}
	return c1/10 == c2/10
}
