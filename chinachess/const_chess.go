package chinachess

const (
	CHINACHESS_PLAYERNUM = 2 // 游戏人数
)
const (
	CHESSBOARD_NUM = 90 // 棋盘落点数量
	CHESSMAN_NUM   = 33 // 棋子数量
	X_MAX          = 9  // 横坐标最大值
	Y_MAX          = 8  // 纵坐标最大值
)

// 棋盘
var CHESSBOARD = [CHESSBOARD_NUM]int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, // 车 马 相 仕 帅 仕 相 马 车
	9, 10, 11, 12, 13, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 23, 24, 25, 26, // * 炮 * * * * * 炮 *
	27, 28, 29, 30, 31, 32, 33, 34, 35, // 兵 * 兵 * 兵 * 兵 * 兵
	36, 37, 38, 39, 40, 41, 42, 43, 44,
	// 楚河          汉界
	45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, // 卒 * 卒 * 卒 * 卒 * 卒
	63, 64, 65, 66, 67, 68, 69, 70, 71, // * 炮 * * * * * 炮 *
	72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, // 车 马 象 士 将 士 象 马 车
}

// 初始棋盘棋子位置
var CHESSBOARDMAN = [CHESSBOARD_NUM]int{
	2, 4, 10, 8, 1, 9, 11, 5, 3, // 车 马 相 仕 帅 仕 相 马 车
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 6, 0, 0, 0, 0, 0, 7, 0, // * 炮 * * * * * 炮 *
	12, 0, 13, 0, 14, 0, 15, 0, 16, // 兵 * 兵 * 兵 * 兵 * 兵
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 楚河          汉界
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	28, 0, 29, 0, 30, 0, 31, 0, 32, // 卒 * 卒 * 卒 * 卒 * 卒
	0, 22, 0, 0, 0, 0, 0, 23, 0, // * 炮 * * * * * 炮 *
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	18, 20, 26, 24, 17, 25, 27, 21, 19, // 车 马 象 士 将 士 象 马 车
}

// 棋子
const (
	NOTCHESSMAN   = 0 // 无棋子
	RED_GENERALS  = 1 // 红方帅
	RED_ROOKS     = 2 // 红方车
	RED_HORSES    = 3 // 红方马
	RED_CONNONS   = 4 // 红方炮
	RED_ELEPHANTS = 5 // 红方相
	RED_GUARDS    = 6 // 红方士
	RED_SOLDIERS  = 7 // 红方兵

	BLACK_GENERALS  = -1 // 黑方将
	BLACK_ROOKS     = -2 // 黑方车
	BLACK_HORSES    = -3 // 黑方马
	BLACK_CONNONS   = -4 // 黑方炮
	BLACK_ELEPHANTS = -5 // 黑方象
	BLACK_GUARDS    = -6 // 黑方仕
	BLACK_SOLDIERS  = -7 // 黑方卒
)

/*
无棋子：0
黑方将：1
黑方车：2，3
黑方马：4，5
黑方炮：6，7
黑方仕：8，9
黑方象：10，11
黑方卒：12，13，14，15，16
红方帅：17
红方车：18，19
红方马：20，21
红方炮：22，23
红方士：24，25
红方相：26，27
红方兵：28，29，30，31，32
*/
// 棋盘
var CHESSMAN = [CHESSMAN_NUM]int{
	0,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
}

// 棋盘索引数组
var CHESSBOARD_INDEX = [CHESSBOARD_NUM]int{
	-2, -3, -5, -6, -1, -6, -5, -3, -2, // 车 马 相 仕 帅 仕 相 马 车
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -4, 0, 0, 0, 0, 0, -4, 0, // * 炮 * * * * * 炮 *
	-7, 0, -7, 0, -7, 0, -7, 0, -7, // 兵 * 兵 * 兵 * 兵 * 兵
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 楚河          汉界
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	7, 0, 7, 0, 7, 0, 7, 0, 7, // 卒 * 卒 * 卒 * 卒 * 卒
	0, 4, 0, 0, 0, 0, 0, 4, 0, // * 炮 * * * * * 炮 *
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	2, 3, 5, 6, 1, 6, 5, 3, 2, // 车 马 象 士 将 士 象 马 车
}

// 棋子索引数组 -1无效棋子
var CHESSBMAN_INDEX = [CHESSMAN_NUM]int{-1, 4, 0, 8, 1, 7, 19, 25, 3, 5, 2, 6, 27, 29, 31, 33, 35, 85, 81, 89, 82, 88, 64, 70, 84, 86, 83, 87, 54, 56, 58, 60, 62}