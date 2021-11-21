package chinachess

import (
	"log"
)

// 中国象棋着法实现
type ChinaChessRoomData struct {
	pos             int                      // 当前操作玩家位置 0 红方 1 黑方
	chessBoardIndex [CHESSBOARD_NUM]int      // 棋盘索引数组
	chessManIndex   [CHESSMAN_NUM]int        // 棋子索引数组
	chessBoardMan   [CHESSBOARD_NUM]int      // 初始棋盘棋子位置 可用于客户端初始化使用，后续便于残局功能开发
	seats           [2]*ChinaChessPlayerData // 座位信息
	gameRecord      []*GameRecord            // 对局纪录
}

// 游戏记录
type GameRecord struct {
	pos            int  // 玩家位置 0 红方 1 黑方
	cmType         int  // 操作棋子类型
	cmPos          int  // 棋子起始位置
	cmDropPointPos int  // 棋子落点位置
	cmEat          bool // 是否吃子
	cmEated        int  //  被吃的棋子
}

func NewChessMan() *ChinaChessRoomData {
	crd := new(ChinaChessRoomData)
	crd.init()
	return crd
}

func (crd *ChinaChessRoomData) init() {
	crd.pos = 0
	crd.chessBoardIndex = CHESSBOARD_INDEX
	crd.chessManIndex = CHESSBMAN_INDEX
	crd.chessBoardMan = CHESSBOARDMAN
	for i := 0; i < CHINACHESS_PLAYERNUM; i++ {
		crd.seats[i] = NewChinaChessPlayerData(i, crd.chessBoardIndex, crd.chessManIndex)
	}
}

/*
	参数说明：
	1.chessMan 棋子类型 对应单个棋子，不区分棋子颜色
	2.pos 当前操作玩家位置
	3.nextPos 棋子着之后的棋盘位置
*/
func (crd *ChinaChessRoomData) move(chessMan, pos, nextPos int) bool {
	// 操作玩家不匹配
	if crd.pos != pos {
		log.Panicf("move curpos(%v) not %v\n", crd.pos, pos)
		return false
	}
	// 获取当前操作玩家
	player := crd.seats[pos]
	// 重置玩家数据
	player.clear()

	if ok := crd.checkParams(chessMan, pos, nextPos); !ok {
		log.Panicf("ChinaChessRoomData checkParams error\n")
		return false
	}
	player.cmDropPoints = player.searchDropPoints()
	if player.inDropPoints(nextPos) {
		if crd.cmExist(nextPos) { // 吃子
			player.cmEat = true
			player.cmEated = crd.removeCm(nextPos) // 更新棋子数据
		}
		// 更新棋盘数据
		crd.chessBoardIndex[nextPos], crd.chessBoardIndex[player.cmPos] = crd.chessBoardIndex[player.cmPos], 0
		// 游戏记录
		crd.gameRecord = append(crd.gameRecord, &GameRecord{
			pos:            player.pos,
			cmType:         player.cmType,
			cmPos:          player.cmPos,
			cmDropPointPos: nextPos,
			cmEat:          player.cmEat,
			cmEated:        player.cmEated,
		})
		log.Printf("move sucess\n")
		// 广播玩家操作
		crd.nextPlayer()
		return true
	}
	return false
}

/*
	请求悔棋
*/
func (crd *ChinaChessRoomData) concelMove(chessMan, pos, nextPos int) bool {
	// 操作玩家不匹配
	if crd.pos == pos {
		log.Panicf("concelMove curpos(%v) not %v\n", crd.pos, pos)
		return false
	}
	// 获取当前操作玩家
	// player := crd.seats[pos]
	// 通知当前操作玩家悔棋

	return false
}

/*
	同意悔棋
*/
func (crd *ChinaChessRoomData) agreeConcelMove(chessMan, pos, nextPos int) bool {
	// 操作玩家不匹配
	if crd.pos != pos {
		log.Panicf("agreeConcelMove curpos(%v) not %v\n", crd.pos, pos)
		return false
	}
	// 还原棋盘和操作玩家位置，通知玩家着棋
	crd.nextPlayer()

	// 获取悔棋玩家
	p := crd.seats[crd.pos]

	// 获取最后一次操作记录
	n := len(crd.gameRecord)
	lastGameRecord := crd.gameRecord[n-1]
	crd.gameRecord = append(crd.gameRecord[0:n-1], crd.gameRecord[:]...)
	if lastGameRecord.cmEat { // 有吃子操作
		crd.chessManIndex[lastGameRecord.cmEated] = lastGameRecord.cmDropPointPos // 还原被吃棋子数据
	}
	// 更新棋盘数据
	crd.chessBoardIndex[lastGameRecord.cmDropPointPos], crd.chessBoardIndex[lastGameRecord.cmPos] = 0, crd.chessBoardIndex[lastGameRecord.cmPos]
	return false
}

// 校验参数
func (crd *ChinaChessRoomData) checkParams(chessMan, pos, nextPos int) bool {
	// 获取当前操作玩家
	player := crd.seats[pos]
	// 更新玩家棋盘棋子数据
	player.chessBoardIndex = crd.chessBoardIndex
	player.chessManIndex = crd.chessManIndex

	if !player.checkChessMan(chessMan) {
		log.Panicf("chessman(%v) invalid\n", chessMan)
		return false
	}
	if !player.checkPlayerChessMan(chessMan) {
		log.Panicf("chessman(%v) player(%v) not match\n", chessMan, crd.pos)
		return false
	}
	if !player.checkChessManPos(chessMan) {
		log.Panicf("chessman(%v) can`t found\n", chessMan)
		return false
	}

	player.nextCmType = crd.chessBoardIndex[nextPos]
	if !player.checkDropPoint(nextPos) {
		log.Panicf("checkPoint failed\n")
	}
	// 更新玩家当前棋子信息
	player.cm = chessMan
	player.cmPos = crd.chessManIndex[chessMan]
	player.cmType = crd.chessBoardIndex[player.cmPos]
	log.Printf("checkParams pass\n")
	return true
}

// 棋子是否存在
func (crd *ChinaChessRoomData) cmExist(pos int) bool {
	return crd.chessBoardIndex[pos] != 0
}

// 轮转玩家
func (crd *ChinaChessRoomData) nextPlayer() {
	crd.pos = (crd.pos + 1) % 2
}

// 去除棋子
func (crd *ChinaChessRoomData) removeCm(pos int) int {
	cmEated := 0
	for i := 0; i < CHESSMAN_NUM; i++ {
		if crd.chessManIndex[i] == pos {
			crd.chessManIndex[i] = -1
			cmEated = i
			break
		}
	}
	return cmEated
}
