package reback

import (
	"fmt"
	"testing"
)

func Test_getKanHuxi(t *testing.T) {
	type args struct {
		pool []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_getKanHuxi_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}}, 0},
		{"Test_getKanHuxi_02", args{[]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, 0},
		{"Test_getKanHuxi_03", args{[]int{
			3, 2, 1, 1, 1, 1, 1, 1, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 1},
		{"Test_getKanHuxi_04", args{[]int{
			3, 3, 0, 1, 1, 1, 1, 1, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 2},
		{"Test_getKanHuxi_05", args{[]int{
			3, 3, 3, 0, 0, 0, 1, 1, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 3},
		{"Test_getKanHuxi_06", args{[]int{
			3, 3, 3, 0, 0, 0, 1, 1, 1, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 4},
		{"Test_getKanHuxi_07", args{[]int{
			3, 3, 3, 0, 3, 0, 1, 1, 1, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 5},
		{"Test_getKanHuxi_08", args{[]int{
			3, 3, 3, 0, 3, 3, 1, 1, 1, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 6},
		{"Test_getKanHuxi_09", args{[]int{
			3, 3, 3, 0, 3, 3, 1, 3, 0, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL * 7},
		{"Test_getKanHuxi_10", args{[]int{
			0, 3, 3, 0, 3, 3, 1, 0, 0, 3,
			0, 0, 0, 0, 3, 0, 0, 0, 0, 0,
		}}, HUXI_KAN_SMALL*5 + HUXI_KAN_BIG*1},
		{"Test_getKanHuxi_11", args{[]int{
			0, 3, 3, 0, 3, 3, 1, 0, 0, 3,
			0, 0, 0, 0, 3, 0, 0, 3, 0, 0,
		}}, HUXI_KAN_SMALL*5 + HUXI_KAN_BIG*2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getKanHuxi(tt.args.pool); got != tt.want {
				t.Errorf("getKanHuxi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShunHuXiBigNot27A(t *testing.T) {
	type args struct {
		pool []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_getShunHuXiBigNot27A_01", args{[]int{
			0, 0, 0, 1, 1, 1, 1, 1, 1, 0,
		}, 0}, HUXI_123_27A_BIG * 0},
		{"Test_getShunHuXiBigNot27A_02", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
		}, 1}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiBigNot27A_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 0}, HUXI_123_27A_BIG * 0},
		{"Test_getShunHuXiBigNot27A_04", args{[]int{
			0, 0, 0, 2, 2, 2, 0, 0, 0, 0,
		}, 0}, HUXI_123_27A_BIG * 0},
		{"Test_getShunHuXiBigNot27A_05", args{[]int{
			0, 2, 2, 2, 2, 0, 2, 0, 0, 2,
		}, 2}, HUXI_123_27A_BIG * 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getShunHuXiBigNot27A(tt.args.pool, tt.args.num); got != tt.want {
				t.Errorf("getShunHuXiBigNot27A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShunHuXiBig(t *testing.T) {
	type args struct {
		pool []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_getShunHuXiBig_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiBig_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiBig_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiBig_04", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 2, 1, 0, 0, 0, 1, 0, 0, 1,
		}}, HUXI_123_27A_BIG * 2},
		{"Test_getShunHuXiBig_05", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 2, 1, 1, 1, 1, 1, 0, 0, 1,
		}}, HUXI_123_27A_BIG * 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getShunHuXiBig(tt.args.pool); got != tt.want {
				t.Errorf("getShunHuXiBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShunHuXiSmall(t *testing.T) {
	type args struct {
		pool []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_getShunHuXiSmall_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiSmall_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiSmall_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiSmall_04", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_123_27A_BIG * 1},
		{"Test_getShunHuXiSmall_05", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},
		{"Test_getShunHuXiSmall_06", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 0,
		}}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2},
		{"Test_getShunHuXiSmall_07", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getShunHuXiSmall(tt.args.pool); got != tt.want {
				t.Errorf("getShunHuXiSmall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcHuXi(t *testing.T) {
	type args struct {
		pool     []int
		handsNum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 特殊胡牌
		{"TestCalcHuXi_Special_01", args{[]int{
			2, 2, 2, 0, 0, 0, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 1, 0,
		}, 21}, 18}, // 21张一般胡
		{"TestCalcHuXi_Special_02", args{[]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 1, 0,
		}, 21}, 15}, // 21张一般胡
		{"TestCalcHuXi_Special_03", args{[]int{
			1, 1, 1, 0, 2, 2, 2, 0, 0, 0,
			2, 2, 2, 0, 3, 0, 3, 0, 0, 0,
		}, 21}, 15}, // 21张一般胡
		{"TestCalcHuXi_Special_04", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, 0}, // 单吊将胡
		{"TestCalcHuXi_Special_05", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, 0}, // 单吊将胡
		{"TestCalcHuXi_Special_06", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, 0}, // 单吊坎胡
		{"TestCalcHuXi_Special_07", args{[]int{
			0, 1, 0, 0, 0, 0, 1, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_SMALL * 1}, // 单吊27A胡
		{"TestCalcHuXi_Special_08", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 1, 0, 0, 1,
		}, 3}, HUXI_123_27A_BIG * 1}, // 单吊27A胡
		{"TestCalcHuXi_Special_09", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_SMALL * 1}, // 单吊123胡
		{"TestCalcHuXi_Special_0A", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_BIG * 1}, // 单吊123胡
		{"TestCalcHuXi_Special_0B", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 2, 2, 2, 2, 0, 2, 0, 0, 2,
		}, 12}, HUXI_123_27A_BIG * 2}, // 大字既能27A 又能顺子胡 好像没有这种组合

		// 特殊牌型 不胡
		{"TestCalcHuXi_Special_11", args{[]int{
			2, 2, 2, 0, 0, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 0, 1, 1, 1, 0,
		}, 21}, -1}, // 21张一般
		{"TestCalcHuXi_Special_12", args{[]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 0, 1,
		}, 21}, -1}, // 21张一般
		{"TestCalcHuXi_Special_13", args{[]int{
			1, 1, 1, 0, 2, 2, 2, 0, 0, 0,
			2, 2, 2, 0, 3, 0, 2, 1, 0, 0,
		}, 21}, -1}, // 21张一般
		{"TestCalcHuXi_Special_14", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, -1}, // 单吊将
		{"TestCalcHuXi_Special_15", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		}, 2}, -1}, // 单吊将
		{"TestCalcHuXi_Special_16", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 2,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1}, // 单吊坎
		{"TestCalcHuXi_Special_17", args{[]int{
			0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1}, // 单吊27A
		{"TestCalcHuXi_Special_18", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 2, 0, 0, 0,
		}, 3}, -1}, // 单吊27A
		{"TestCalcHuXi_Special_19", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1}, // 单吊123
		{"TestCalcHuXi_Special_1A", args{[]int{
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1}, // 单吊123

		// 不带将牌胡
		{"TestCalcHuXi_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 9}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 6}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_04", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_05", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 9}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},
		{"TestCalcHuXi_06", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 0,
		}, 15}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2},
		{"TestCalcHuXi_07", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},

		// 带将牌胡
		{"TestCalcHuXi_11", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 11}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_12", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 8}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_13", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 2, 0, 0, 0, 0,
		}, 5}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_14", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 8}, HUXI_123_27A_BIG * 1},
		{"TestCalcHuXi_15", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},
		{"TestCalcHuXi_16", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 2, 1, 1, 1, 0, 0,
		}, 17}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2},
		{"TestCalcHuXi_17", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 0,
		}, 8}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1},

		// 带将牌胡+带坎
		{"TestCalcHuXi_21", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_22", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 3, 0, 1, 1, 1,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_23", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 3, 0, 2, 0, 0, 0, 0,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_24", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 3, 0, 3,
		}, 17}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*2 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_25", args{[]int{
			3, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_26", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 3, 2, 1, 1, 1, 0, 0,
		}, 20}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_27", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 3, 3,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_BIG*2},

		// 不带将牌胡+带坎
		{"TestCalcHuXi_31", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_32", args{[]int{
			3, 0, 0, 0, 3, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*2},
		{"TestCalcHuXi_33", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 3,
		}, 6}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_34", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 3, 0, 0, 3, 0, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*2},
		{"TestCalcHuXi_35", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 3, 0, 0, 0, 0, 0, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_36", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 3,
		}, 18}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_37", args{[]int{
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
		}, 18}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_SMALL*2 + HUXI_KAN_BIG*2},
		{"TestCalcHuXi_38", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 3, 3,
		}, 21}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*2},

		// 不带将牌不胡
		{"TestCalcHuXi_41", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
		}, 9}, -1},
		{"TestCalcHuXi_42", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 0,
		}, 6}, -1},
		{"TestCalcHuXi_43", args{[]int{
			0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1},
		{"TestCalcHuXi_44", args{[]int{
			0, 0, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1},
		{"TestCalcHuXi_45", args{[]int{
			0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 1, 0, 0, 0, 0,
		}, 9}, -1},
		{"TestCalcHuXi_46", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 1, 0, 1, 1, 0, 0,
		}, 15}, -1},
		{"TestCalcHuXi_47", args{[]int{
			1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1},

		// 带将牌不胡
		{"TestCalcHuXi_51", args{[]int{
			0, 0, 0, 0, 0, 0, 1, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 0, 1, 1, 0,
		}, 11}, -1},
		{"TestCalcHuXi_52", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 1, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 1, 1,
		}, 8}, -1},
		{"TestCalcHuXi_53", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 1, 0, 2, 0, 0, 0, 0,
		}, 5}, -1},
		{"TestCalcHuXi_54", args{[]int{
			2, 0, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 8}, -1},
		{"TestCalcHuXi_55", args{[]int{
			0, 1, 0, 1, 0, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 1, 0, 0, 0, 0, 2,
		}, 11}, -1},
		{"TestCalcHuXi_56", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 0,
			1, 1, 1, 0, 2, 1, 1, 1, 0, 2,
		}, 17}, -1},
		{"TestCalcHuXi_57", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 0,
		}, 8}, -1},
		{"TestCalcHuXi_58", args{[]int{
			0, 2, 3, 1, 1, 1, 2, 0, 0, 0,
			1, 0, 0, 0, 2, 0, 1, 1, 0, 2,
		}, 17}, -1},
		{"TestCalcHuXi_59", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 3,
		}, 14}, -1},

		// 不带将牌不胡
		{"TestCalcHuXi_61", args{[]int{
			0, 0, 0, 0, 0, 0, 1, 0, 0, 2,
			1, 1, 1, 1, 0, 0, 0, 1, 1, 0,
		}, 9}, -1},
		{"TestCalcHuXi_62", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 1, 0, 0,
			1, 2, 1, 0, 0, 0, 0, 0, 1, 1,
		}, 9}, -1},
		{"TestCalcHuXi_63", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		}, 3}, -1},
		{"TestCalcHuXi_64", args{[]int{
			2, 0, 1, 0, 0, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1},
		{"TestCalcHuXi_65", args{[]int{
			0, 1, 0, 1, 0, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 9}, -1},
		{"TestCalcHuXi_66", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 0,
			1, 1, 1, 1, 2, 1, 1, 1, 0, 2,
		}, 18}, -1},
		{"TestCalcHuXi_67", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1},
		{"TestCalcHuXi_68", args{[]int{
			0, 2, 3, 1, 1, 1, 2, 0, 0, 0,
			1, 0, 0, 0, 2, 1, 1, 1, 0, 2,
		}, 18}, -1},
		{"TestCalcHuXi_69", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 2, 1, 3,
		}, 15}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CalcHuXi(tt.args.pool, tt.args.handsNum); got != tt.want {
				t.Errorf("CalcHuXi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcHuXiHuHandResult(t *testing.T) {
	type args struct {
		pool     []int
		handsNum int
	}

	// huHandResult := &HuHandResult{}
	// huHandResult.init()

	tests := []struct {
		name string
		args args
		huxi int
		want *HuHandResult
	}{
		// 特殊胡牌
		{"TestCalcHuXiHuHandResult_Special_01", args{[]int{
			2, 2, 2, 0, 0, 0, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 1, 0,
		}, 21}, 18, nil}, // 21张一般胡
		{"TestCalcHuXiHuHandResult_Special_02", args{[]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 1, 0,
		}, 21}, 15, nil}, // 21张一般胡
		{"TestCalcHuXiHuHandResult_Special_03", args{[]int{
			1, 1, 1, 0, 2, 2, 2, 0, 0, 0,
			2, 2, 2, 0, 3, 0, 3, 0, 0, 0,
		}, 21}, 15, nil}, // 21张一般胡
		{"TestCalcHuXiHuHandResult_Special_04", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, 0, nil}, // 单吊将胡
		{"TestCalcHuXiHuHandResult_Special_05", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, 0, nil}, // 单吊将胡
		{"TestCalcHuXiHuHandResult_Special_06", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, 0, nil}, // 单吊坎胡
		{"TestCalcHuXiHuHandResult_Special_07", args{[]int{
			0, 1, 0, 0, 0, 0, 1, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_SMALL * 1, nil}, // 单吊27A胡
		{"TestCalcHuXiHuHandResult_Special_08", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 1, 0, 0, 1,
		}, 3}, HUXI_123_27A_BIG * 1, nil}, // 单吊27A胡
		{"TestCalcHuXiHuHandResult_Special_09", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_SMALL * 1, nil}, // 单吊123胡
		{"TestCalcHuXiHuHandResult_Special_0A", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_BIG * 1, nil}, // 单吊123胡
		{"TestCalcHuXiHuHandResult_Special_0B", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 2, 2, 2, 2, 0, 2, 0, 0, 2,
		}, 12}, HUXI_123_27A_BIG * 2, nil}, // 大字既能27A 又能顺子胡 好像没有这种组合

		/* // 特殊牌型 不胡
		{"TestCalcHuXiHuHandResult_Special_11", args{[]int{
			2, 2, 2, 0, 0, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 0, 1, 1, 1, 0,
		}, 21}, -1, nil}, // 21张一般
		{"TestCalcHuXiHuHandResult_Special_12", args{[]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
			2, 2, 2, 1, 1, 1, 1, 1, 0, 1,
		}, 21}, -1, nil}, // 21张一般
		{"TestCalcHuXiHuHandResult_Special_13", args{[]int{
			1, 1, 1, 0, 2, 2, 2, 0, 0, 0,
			2, 2, 2, 0, 3, 0, 2, 1, 0, 0,
		}, 21}, -1, nil}, // 21张一般
		{"TestCalcHuXiHuHandResult_Special_14", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 2}, -1, nil}, // 单吊将
		{"TestCalcHuXiHuHandResult_Special_15", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		}, 2}, -1, nil}, // 单吊将
		{"TestCalcHuXiHuHandResult_Special_16", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 2,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil}, // 单吊坎
		{"TestCalcHuXiHuHandResult_Special_17", args{[]int{
			0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil}, // 单吊27A
		{"TestCalcHuXiHuHandResult_Special_18", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 2, 0, 0, 0,
		}, 3}, -1, nil}, // 单吊27A
		{"TestCalcHuXiHuHandResult_Special_19", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil}, // 单吊123
		{"TestCalcHuXiHuHandResult_Special_1A", args{[]int{
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil}, // 单吊123 */

		// 不带将牌胡
		{"TestCalcHuXiHuHandResult_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 9}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 6}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_04", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_05", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 9}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_06", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 0,
		}, 15}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2, nil},
		{"TestCalcHuXiHuHandResult_07", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1, nil},

		// 带将牌胡
		{"TestCalcHuXiHuHandResult_11", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 11}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_12", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 8}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_13", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 2, 0, 0, 0, 0,
		}, 5}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_14", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 8}, HUXI_123_27A_BIG * 1, nil},
		{"TestCalcHuXiHuHandResult_15", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_16", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 2, 1, 1, 1, 0, 0,
		}, 17}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2, nil},
		{"TestCalcHuXiHuHandResult_17", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 0,
		}, 8}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1, nil},

		// 带将牌胡+带坎
		{"TestCalcHuXiHuHandResult_21", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_22", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 3, 0, 1, 1, 1,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1, nil},
		{"TestCalcHuXiHuHandResult_23", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 3, 0, 2, 0, 0, 0, 0,
		}, 11}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1 + HUXI_KAN_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_24", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 3, 0, 3,
		}, 17}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*2 + HUXI_KAN_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_25", args{[]int{
			3, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_26", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 3, 2, 1, 1, 1, 0, 0,
		}, 20}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*1, nil},
		{"TestCalcHuXiHuHandResult_27", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 3, 3,
		}, 14}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_BIG*2, nil},

		// 不带将牌胡+带坎
		{"TestCalcHuXiHuHandResult_31", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*1, nil},
		{"TestCalcHuXiHuHandResult_32", args{[]int{
			3, 0, 0, 0, 3, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_SMALL*2, nil},
		{"TestCalcHuXiHuHandResult_33", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 3,
		}, 6}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*1, nil},
		{"TestCalcHuXiHuHandResult_34", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 3, 0, 0, 3, 0, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_KAN_BIG*2, nil},
		{"TestCalcHuXiHuHandResult_35", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 3, 0, 0, 0, 0, 0, 0,
		}, 12}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_BIG*1, nil},
		{"TestCalcHuXiHuHandResult_36", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 3,
		}, 18}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*1, nil},
		{"TestCalcHuXiHuHandResult_37", args{[]int{
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
		}, 18}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*1 + HUXI_KAN_SMALL*2 + HUXI_KAN_BIG*2, nil},
		{"TestCalcHuXiHuHandResult_38", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 3, 3,
		}, 21}, HUXI_123_27A_BIG*1 + HUXI_123_27A_SMALL*2 + HUXI_KAN_BIG*2, nil},

		// 不带将牌不胡
		{"TestCalcHuXiHuHandResult_41", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
		}, 9}, -1, nil},
		{"TestCalcHuXiHuHandResult_42", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 0,
		}, 6}, -1, nil},
		{"TestCalcHuXiHuHandResult_43", args{[]int{
			0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil},
		{"TestCalcHuXiHuHandResult_44", args{[]int{
			0, 0, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1, nil},
		{"TestCalcHuXiHuHandResult_45", args{[]int{
			0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 1, 0, 0, 0, 0,
		}, 9}, -1, nil},
		{"TestCalcHuXiHuHandResult_46", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 1, 0, 1, 1, 0, 0,
		}, 15}, -1, nil},
		{"TestCalcHuXiHuHandResult_47", args{[]int{
			1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1, nil},

		// 带将牌不胡
		{"TestCalcHuXiHuHandResult_51", args{[]int{
			0, 0, 0, 0, 0, 0, 1, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 0, 1, 1, 0,
		}, 11}, -1, nil},
		{"TestCalcHuXiHuHandResult_52", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 1, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 1, 1,
		}, 8}, -1, nil},
		{"TestCalcHuXiHuHandResult_53", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 1, 0, 2, 0, 0, 0, 0,
		}, 5}, -1, nil},
		{"TestCalcHuXiHuHandResult_54", args{[]int{
			2, 0, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 8}, -1, nil},
		{"TestCalcHuXiHuHandResult_55", args{[]int{
			0, 1, 0, 1, 0, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 1, 0, 0, 0, 0, 2,
		}, 11}, -1, nil},
		{"TestCalcHuXiHuHandResult_56", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 0,
			1, 1, 1, 0, 2, 1, 1, 1, 0, 2,
		}, 17}, -1, nil},
		{"TestCalcHuXiHuHandResult_57", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 0,
		}, 8}, -1, nil},
		{"TestCalcHuXiHuHandResult_58", args{[]int{
			0, 2, 3, 1, 1, 1, 2, 0, 0, 0,
			1, 0, 0, 0, 2, 0, 1, 1, 0, 2,
		}, 17}, -1, nil},
		{"TestCalcHuXiHuHandResult_59", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 3,
		}, 14}, -1, nil},

		// 不带将牌不胡
		{"TestCalcHuXiHuHandResult_61", args{[]int{
			0, 0, 0, 0, 0, 0, 1, 0, 0, 2,
			1, 1, 1, 1, 0, 0, 0, 1, 1, 0,
		}, 9}, -1, nil},
		{"TestCalcHuXiHuHandResult_62", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 1, 0, 0,
			1, 2, 1, 0, 0, 0, 0, 0, 1, 1,
		}, 9}, -1, nil},
		{"TestCalcHuXiHuHandResult_63", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		}, 3}, -1, nil},
		{"TestCalcHuXiHuHandResult_64", args{[]int{
			2, 0, 1, 0, 0, 1, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1, nil},
		{"TestCalcHuXiHuHandResult_65", args{[]int{
			0, 1, 0, 1, 0, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 9}, -1, nil},
		{"TestCalcHuXiHuHandResult_66", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 0,
			1, 1, 1, 1, 2, 1, 1, 1, 0, 2,
		}, 18}, -1, nil},
		{"TestCalcHuXiHuHandResult_67", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, -1, nil},
		{"TestCalcHuXiHuHandResult_68", args{[]int{
			0, 2, 3, 1, 1, 1, 2, 0, 0, 0,
			1, 0, 0, 0, 2, 1, 1, 1, 0, 2,
		}, 18}, -1, nil},
		{"TestCalcHuXiHuHandResult_69", args{[]int{
			1, 1, 0, 1, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 2, 1, 3,
		}, 15}, -1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.huxi != -1 {
				fmt.Println(tt.args.pool)
				huxi, got := CalcHuXi(tt.args.pool, tt.args.handsNum)
				/* if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CalcHuXi() got1 = %v, want %v", got, tt.want)
				} */
				t.Log(tt.args.pool)
				if huxi != -1 && got != nil {
					cardsLogStr, flagLogStr := "CardsLog ", "FlagLog "
					for i := 0; i <= got.SeqCnt; i++ {
						cardsLogStr += CardsLog(got.Sequence[i].Cards[:])
						flagLogStr += FlagLog(got.Sequence[i].Flag)
						// t.Logf("Group:%v\n", got.Sequence[i])
						// t.Logf("Cards:%v\n", got.Sequence[i].Cards)
						// t.Logf("Flag:%v\n", got.Sequence[i].Flag)
					}
					fmt.Println(cardsLogStr)
					fmt.Println(flagLogStr)
				}
			}
		})
	}
}
