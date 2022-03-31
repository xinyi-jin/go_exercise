package reback

const (
	// 小牌
	BEARD_SMALL_1 int64 = iota
	BEARD_SMALL_2
	BEARD_SMALL_3
	BEARD_SMALL_4
	BEARD_SMALL_5
	BEARD_SMALL_6
	BEARD_SMALL_7
	BEARD_SMALL_8
	BEARD_SMALL_9
	BEARD_SMALL_A

	// 大牌
	BEARD_BIG_1
	BEARD_BIG_2
	BEARD_BIG_3
	BEARD_BIG_4
	BEARD_BIG_5
	BEARD_BIG_6
	BEARD_BIG_7
	BEARD_BIG_8
	BEARD_BIG_9
	BEARD_BIG_A

	// 最大20
	BEARD_MAX
)

const (
	HUXI_KAN_SMALL     = 0 // 小坎
	HUXI_KAN_BIG       = 0 // 大坎
	HUXI_SHUN_SMALL    = 0 // 除123之外的小顺子
	HUXI_SHUN_BIG      = 0 // 除123之外的打顺子
	HUXI_PENG_SMALL    = 1 // 小碰
	HUXI_PENG_BIG      = 3 // 大碰
	HUXI_WEI_SMALL     = 3 // 小偎
	HUXI_WEI_BIG       = 6 // 大偎
	HUXI_PAO_SMALL     = 6 // 小跑
	HUXI_PAO_BIG       = 9 // 大跑
	HUXI_TI_SMALL      = 9 // 小提
	HUXI_TI_BIG        = 9 // 大提
	HUXI_123_27A_SMALL = 3 // 小123/小27A
	HUXI_123_27A_BIG   = 6 // 大123/大27A
)

const (
	GROUPTYPE_123_27A_SMALL  int64 = iota // 小123/小27A
	GROUPTYPE_123_27A_BIG                 // 大123/大27A
	GROUPTYPE_SHUN_SMALL                  // 除123之外的小顺子
	GROUPTYPE_SHUN_BIG                    // 除123之外的大顺子
	GROUPTYPE_JIAO_aaA                    // 绞牌 两小一大
	GROUPTYPE_JIAO_aAA                    // 绞牌 一小两大
	GROUPTYPE_KAN_SMALL                   // 小坎
	GROUPTYPE_KAN_BIG                     // 大坎
	GROUPTYPE_PENG_SMALL                  // 小碰
	GROUPTYPE_PENG_BIG                    // 大碰
	GROUPTYPE_WEI_SMALL                   // 小偎
	GROUPTYPE_WEI_BIG                     // 大偎
	GROUPTYPE_CHOUWEI_SMALL               // 小臭偎
	GROUPTYPE_CHOUWEI_BIG                 // 大臭偎
	GROUPTYPE_PAO_SMALL                   // 小跑
	GROUPTYPE_PAO_BIG                     // 大跑
	GROUPTYPE_WEIPAO_SMALL                // 小偎后跑
	GROUPTYPE_WEIPAO_BIG                  // 大偎后跑
	GROUPTYPE_PENGPAO_SMALL               // 小碰后跑
	GROUPTYPE_PENGPAO_BIG                 // 大碰后跑
	GROUPTYPE_TI_SMALL                    // 小提
	GROUPTYPE_TI_BIG                      // 大提
	GROUPTYPE_WEITI_SMALL                 // 小偎后提
	GROUPTYPE_WEITI_BIG                   // 大偎后提
	GROUPTYPE_JIANG                       // 将牌
	GROUPTYPE_QISHOUTI_SMALL              // 小起手提
	GROUPTYPE_QISHOUTI_BIG                // 大起手提
	GROUPTYPE_MAX                         // 最大临界值
)
