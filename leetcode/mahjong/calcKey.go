package mahjong

func calcKey(index []int16, pos []int16) int32 {
	p := -1
	x := 0
	posP := 0
	b := false
	//数牌
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if index[i*9+j] == 0 {
				if b {
					b = false
					x |= 0x1 << uint(p)
					p++
				}
			} else {
				p++
				b = true
				pos[posP] = int16(i*9 + j)
				posP++
				switch index[i*9+j] {
				case 2:
					x |= 0x3 << uint(p)
					p += 2
				case 3:
					x |= 0xF << uint(p)
					p += 4
				case 4:
					x |= 0x3F << uint(p)
					p += 6
				}
			}
		}
		if b {
			b = false
			x |= 0x1 << uint(p)
			p++
		}
	}

	//字牌
	for i := 27; i < int(33); i++ {
		if index[i] > 0 {
			p++
			pos[posP] = int16(i)
			posP++
			switch index[i] {
			case 2:
				x |= 0x3 << uint(p)
				p += 2
			case 3:
				x |= 0xF << uint(p)
				p += 4
			case 4:
				x |= 0x3F << uint(p)
				p += 6
			}
			x |= 0x1 << uint(p)
			p++
		}
	}

	return int32(x)
}
