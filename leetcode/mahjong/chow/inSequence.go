package chow

func color(c int64) int64 {
	if c < 0 {
		return -1
	}
	return c / 9
}

func inSequence(card int64, pool [34]int64) bool {
	cards := []int64{card - 3, card - 2, card - 1}
	for i := 0; i < 3; i++ {
		cards[0]++
		cards[1]++
		cards[2]++

		if color(cards[0]) == color(cards[1]) && color(cards[0]) == color(cards[2]) &&
			pool[cards[0]] > 0 && pool[cards[1]] > 0 && pool[cards[2]] > 0 {
			return true
		}
	}
	return false
}
