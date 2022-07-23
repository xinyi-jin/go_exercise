package logic

import "fmt"

type GameArr [4][4]int

var row int = 4

func (g *GameArr) Right90() {
	var tempArr = new(GameArr)
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			tempArr[i][j] = g[row-j-1][i]
		}
	}
	*g = *tempArr
}

func (g *GameArr) Left90() {
	var tempArr = new(GameArr)
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			tempArr[i][j] = g[j][row-i-1]
		}
	}
	*g = *tempArr
}

func (g *GameArr) Up() {
	for i := 0; i < row; i++ {
		for j := 0; j < row; {
			add := false
			if g[i][j] != 0 {
				for x := 1; x < row; x++ {
					if g[i][j] == 0 {
						continue
					}
					if g[i][j] == g[x][j] {
						g[i][j] = 2 * g[x][j]
						g[x][j] = 0
						add = true
						j = x + 1
					}
					break
				}
			}
			if !add {
				j++
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			if g[i][j] == 0 {
				for x := 1; x < row; x++ {
					if g[x][j] == 0 {
						continue
					}
					g[i][j] = g[x][j]
					g[x][j] = 0
					break
				}

			}
		}
	}

}

func (g *GameArr) UpT() bool {
	GSIZE := 4
	var ok bool
	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; {
			add := false
			if g[row][col] != 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					if g[row][col] == g[r][col] {
						g[row][col] += g[r][col]
						g[r][col] = 0
						ok = true
						row = r + 1
						add = true
					}

					break
				}
			}

			if !add {
				row++
			}
		}
	}

	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; row++ {
			if g[row][col] == 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					g[row][col] = g[r][col]
					g[r][col] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameArr) Print() {
	for _, v := range g {
		for _, vv := range v {
			fmt.Printf("%d ,", vv)
		}
		fmt.Println()
	}
}

func main() {
	var arr GameArr
	// arr = make([4][4]int)
	arr = [4][4]int{
		{0, 2, 0, 0},
		{4, 0, 0, 0},
		{8, 2, 4, 0},
		{8, 4, 2, 0},

		/* {0, 2, 4, 0},
		{2, 0, 0, 0},
		{0, 2, 0, 2},
		{0, 2, 2, 0}, */
	}
	arr.Print()

	// arr.Right90()
	// arr.Right90()
	// fmt.Println("xuanzanyhou:")
	// arr.Print()

	fmt.Println("add after:")
	arr.Up()
	arr.Print()

	/* fmt.Println("左旋转:")
	arr.Left90()
	arr.Print()

	fmt.Println("右旋转:")
	arr.Right90()
	arr.Print() */
}
