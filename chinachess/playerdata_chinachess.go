package chinachess

import "log"

type ChinaChessPlayerData struct {
	pos             int                 // 玩家座位
	cm              int                 // 当前棋子
	cmPos           int                 // 当前棋子位置
	cmType          int                 // 当前棋子类型
	cmDropPoints    []int               // 当前棋子可落子点
	cmEat           bool                // 是否吃子
	cmEated         int                 // 被吃的棋子
	nextCmType      int                 // 落点棋子类型
	chessBoardIndex [CHESSBOARD_NUM]int // 棋盘索引数组
	chessManIndex   [CHESSMAN_NUM]int   // 棋子索引数组
}

func NewChinaChessPlayerData(pos int, chessBoardIndex [CHESSBOARD_NUM]int, chessManIndex [CHESSMAN_NUM]int) *ChinaChessPlayerData {
	return &ChinaChessPlayerData{
		pos:             pos,
		chessBoardIndex: chessBoardIndex,
		chessManIndex:   chessManIndex,
	}
}

func (cpd *ChinaChessPlayerData) clear() {
	cpd.cm = -1
	cpd.cmPos = -1
	cpd.cmType = 0
	cpd.cmDropPoints = nil
	cpd.cmEat = false
	cpd.cmEated = 0
	cpd.nextCmType = 0
	cpd.chessBoardIndex = [CHESSBOARD_NUM]int{}
	cpd.chessManIndex = [CHESSMAN_NUM]int{}
}

// 搜索落子点 每种棋子对应不同的落子点数量
func (cpd *ChinaChessPlayerData) searchDropPoints() []int {
	switch cpd.cmType {
	case RED_GENERALS, BLACK_GENERALS: // 将帅
		return cpd.searchDropPointsGenerals()
	case RED_ROOKS, BLACK_ROOKS: // 车
		return cpd.searchDropPointsRooks()
	case RED_HORSES, BLACK_HORSES: // 马
		return cpd.searchDropPointsHorses()
	case RED_CONNONS, BLACK_CONNONS: // 炮
		return cpd.searchDropPointsConnons()
	case RED_ELEPHANTS, BLACK_ELEPHANTS: // 象
		return cpd.searchDropPointsElephants()
	case RED_GUARDS, BLACK_GUARDS: // 士
		return cpd.searchDropPointsGuards()
	case RED_SOLDIERS, BLACK_SOLDIERS: // 兵
		return cpd.searchDropPointsSoldiers()
	}
	return nil
}

// 搜索将帅可落点
func (cpd *ChinaChessPlayerData) searchDropPointsGenerals() []int {
	x, y := getxyByCmPos(cpd.cmPos)
	if !checkxy(x, y) {
		log.Panicf("ChinaChessPlayerData pos(%v) cmPos(%v) invalid\n", cpd.pos, cpd.cmPos)
		return nil
	}
	dropPoints := make([]int, 0)
	// 上
	nx, ny := x+1, y
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}
	// 下
	nx, ny = x-1, y
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}
	// 左
	nx, ny = x, y-1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}
	// 右
	nx, ny = x, y+1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}
	return dropPoints
}

// 搜索车可落点
func (cpd *ChinaChessPlayerData) searchDropPointsRooks() []int {

	return nil
}

// 搜索马可落点
func (cpd *ChinaChessPlayerData) searchDropPointsHorses() []int {
	x, y := getxyByCmPos(cpd.cmPos)
	if !checkxy(x, y) {
		log.Panicf("ChinaChessPlayerData pos(%v) cmPos(%v) invalid\n", cpd.pos, cpd.cmPos)
		return nil
	}
	dropPoints := make([]int, 0)

	nx, ny := x+1, y-2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x, y-1) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x+1, y+2
	if pos := cpd.CanDropPoint(nx, y+1); pos != -1 && !cpd.IsBanFoot(x, y+1) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x+2, y-1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x+1, y) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x+2, y+1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x+1, y) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-1, y-2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x, y-1) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-1, y+2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x, y+1) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-2, y-1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x-1, y) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-2, y+1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x-1, y) {
		dropPoints = append(dropPoints, pos)
	}
	return nil
}

// 搜索炮可落点
func (cpd *ChinaChessPlayerData) searchDropPointsConnons() []int {

	return nil
}

// 搜索象可落点
func (cpd *ChinaChessPlayerData) searchDropPointsElephants() []int {
	x, y := getxyByCmPos(cpd.cmPos)
	if !checkxy(x, y) {
		log.Panicf("ChinaChessPlayerData pos(%v) cmPos(%v) invalid\n", cpd.pos, cpd.cmPos)
		return nil
	}
	dropPoints := make([]int, 0)

	nx, ny := x+2, y+2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x+1, y+1) && !cpd.IsCrossRiver(nx, ny) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x+2, y-2
	if pos := cpd.CanDropPoint(nx, y+1); pos != -1 && !cpd.IsBanFoot(x+1, y-1) && !cpd.IsCrossRiver(nx, ny) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-2, y+2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x-1, y+2) && !cpd.IsCrossRiver(nx, ny) {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-2, y-2
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 && !cpd.IsBanFoot(x-1, y-1) && !cpd.IsCrossRiver(nx, ny) {
		dropPoints = append(dropPoints, pos)
	}
	return nil
}

// 搜索士可落点
func (cpd *ChinaChessPlayerData) searchDropPointsGuards() []int {
	x, y := getxyByCmPos(cpd.cmPos)
	if !checkxy(x, y) {
		log.Panicf("ChinaChessPlayerData pos(%v) cmPos(%v) invalid\n", cpd.pos, cpd.cmPos)
		return nil
	}
	dropPoints := make([]int, 0)

	nx, ny := x+1, y+1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x+1, y-1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-1, y+1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}

	nx, ny = x-1, y-1
	if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
		dropPoints = append(dropPoints, pos)
	}
	return nil
}

// 搜索兵可落点
func (cpd *ChinaChessPlayerData) searchDropPointsSoldiers() []int {
	x, y := getxyByCmPos(cpd.cmPos)
	if !checkxy(x, y) {
		log.Panicf("ChinaChessPlayerData pos(%v) cmPos(%v) invalid\n", cpd.pos, cpd.cmPos)
		return nil
	}
	dropPoints := make([]int, 0)

	if cpd.IsCrossRiver(x, y) { // 已过河
		nx, ny := x+1, y
		if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
			dropPoints = append(dropPoints, pos)
		}
		nx, ny = x, y-1
		if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
			dropPoints = append(dropPoints, pos)
		}
		nx, ny = x, y+1
		if pos := cpd.CanDropPoint(nx, ny); pos != -1 {
			dropPoints = append(dropPoints, pos)
		}
	} else { // 未过河
		if pos := cpd.CanDropPoint(x+1, y); pos != -1 {
			dropPoints = append(dropPoints, pos)
		}
	}

	return nil
}

// 通过棋子位置获取棋子类型
func (cpd *ChinaChessPlayerData) getCmTypeByCmPos(cmPos int) int {
	return cpd.chessBoardIndex[cmPos]
}

// 校验落点
func (cpd *ChinaChessPlayerData) checkDropPoint(dropPoint int) bool {
	switch {
	case cpd.nextCmType > 0:
		if cpd.pos == 1 {
			log.Panicf("nextPos(%v) player(%v) not match\n", dropPoint, cpd.pos)
			return false
		}
	case cpd.nextCmType < 0:
		if cpd.pos == 0 {
			log.Panicf("nextPos(%v) player(%v) not match\n", dropPoint, cpd.pos)
			return false
		}
	case cpd.nextCmType == 0:
		log.Printf("nextPos(%v) not chessman\n", dropPoint)
		return true
	}
	return false
}

// 校验棋子是否和玩家相匹配
func (cpd *ChinaChessPlayerData) checkPlayerChessMan(chessMan int) bool {
	return cpd.pos == 1 && chessMan < 17
}

// 校验棋子类型是否和玩家相匹配
func (cpd *ChinaChessPlayerData) checkPlayerChessManType(chessManPos int) bool {
	return cpd.pos == 0 && cpd.chessBoardIndex[chessManPos] > 0
}

// 校验棋子
func (cpd *ChinaChessPlayerData) checkChessMan(chessMan int) bool {
	return chessMan > 0 && chessMan < 34
}

// 校验棋子位置
func (cpd *ChinaChessPlayerData) checkChessManPos(chessMan int) bool {
	return cpd.chessManIndex[chessMan] != -1
}

// 校验棋子类型 -1不存在 0红方 1 黑方
func (cpd *ChinaChessPlayerData) checkChessManType(chessManPos int) int {
	switch {
	case cpd.chessBoardIndex[chessManPos] == 0:
		return -1
	case cpd.chessBoardIndex[chessManPos] > 0:
		return 0
	case cpd.chessBoardIndex[chessManPos] < 0:
		return 1
	}
	return -1
}

// 校验落子位置
func (cpd *ChinaChessPlayerData) inDropPoints(nextPos int) bool {
	for _, v := range cpd.cmDropPoints {
		if v == nextPos {
			return true
		}
	}
	return false
}

// 是否可以落子 -1 不可以落子
func (cpd *ChinaChessPlayerData) CanDropPoint(x, y int) int {
	if checkxy(x, y) {
		dropPoint := getCmPosByxy(x, y)
		if cpd.chessBoardIndex[dropPoint] == 0 { // 落点无棋子
			return dropPoint
		} else if cpd.chessBoardIndex[dropPoint] > 0 && cpd.pos == 1 || cpd.chessBoardIndex[dropPoint] < 0 && cpd.pos == 0 { // 吃子
			return dropPoint
		}
	}
	return -1
}

// 拌马腿/堵象眼
func (cpd *ChinaChessPlayerData) IsBanFoot(x, y int) bool {
	if checkxy(x, y) {
		dropPoint := getCmPosByxy(x, y)
		if cpd.chessBoardIndex[dropPoint] == 0 { // 落点无棋子
			return false
		}
	}
	return true
}

// 过河
func (cpd *ChinaChessPlayerData) IsCrossRiver(x, y int) bool {
	if checkxy(x, y) {
		if cpd.pos == 0 && x < 5 || cpd.pos == 1 && x > 4 {
			return true
		}
	}
	return false
}

// 校验坐标
func checkxy(x, y int) bool {
	return x >= 0 && x <= X_MAX && y >= 0 && x <= Y_MAX
}

// 通过坐标获取棋盘位置
func getCmPosByxy(x, y int) int {
	return x*10 + y
}

// 通过坐标获取棋盘位置
func getxyByCmPos(cmpos int) (int, int) {
	return cmpos % 9, cmpos % 10
}
