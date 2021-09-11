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
	gameRecord      string                   // 对局纪录
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

// pos             int   // 起始棋子位置
// nextPos         int   // 落子位置

/*
	参数说明：
	1.chessMan 棋子类型 对应单个棋子，不区分棋子颜色
	2.pos 当前操作玩家位置
	3.nextPos 棋子着之后的棋盘位置
*/
func (crd *ChinaChessRoomData) move(chessMan, pos, nextPos int) bool {
	if ok := crd.checkParams(chessMan, pos, nextPos); !ok {
		log.Panicf("ChinaChessRoomData checkParams error\n")
		return false
	}
	player := crd.seats[pos]
	player.cmDropPoints = player.searchDropPoints()
	if player.inDropPoints(nextPos) {
		log.Printf("move sucess\n")
		return true
	}
	return false
}

// 校验参数
func (crd *ChinaChessRoomData) checkParams(chessMan, pos, nextPos int) bool {
	// 操作玩家不匹配
	if crd.pos != pos {
		log.Panicf("curpos(%v) not %v\n", crd.pos, pos)
		return false
	}
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
