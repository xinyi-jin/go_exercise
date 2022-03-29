package reback

import "testing"

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
			if got := getKanHuxi(tt.args.pool); got != tt.want {
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
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 0}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiBigNot27A_02", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
		}, 1}, HUXI_SHUN_BIG * 0},
		{"Test_getShunHuXiBigNot27A_03", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 0}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiBigNot27A_04", args{[]int{
			2, 2, 2, 2, 2, 2, 0, 0, 0, 0,
		}, 0}, HUXI_SHUN_BIG * 2},
		{"Test_getShunHuXiBigNot27A_05", args{[]int{
			0, 2, 2, 2, 2, 0, 2, 0, 0, 2,
		}, 2}, HUXI_SHUN_BIG * 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShunHuXiBigNot27A(tt.args.pool, tt.args.num); got != tt.want {
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
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiBig_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiBig_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_SHUN_BIG * 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShunHuXiBig(tt.args.pool); got != tt.want {
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
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiSmall_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiSmall_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiSmall_04", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_SHUN_BIG * 1},
		{"Test_getShunHuXiSmall_05", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},
		{"Test_getShunHuXiSmall_06", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 0,
		}}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2},
		{"Test_getShunHuXiSmall_07", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShunHuXiSmall(tt.args.pool); got != tt.want {
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
		// 不带将牌胡
		{"TestCalcHuXi_01", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 9}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_02", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 6}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_03", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 3}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_04", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_05", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 9}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},
		{"TestCalcHuXi_06", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 0,
		}, 15}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2},
		{"TestCalcHuXi_07", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 6}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},

		// 带将牌胡
		{"TestCalcHuXi_11", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 11}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_12", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 8}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_13", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 2, 0, 0, 0, 0,
		}, 5}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_14", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		}, 8}, HUXI_SHUN_BIG * 1},
		{"TestCalcHuXi_15", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 11}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},
		{"TestCalcHuXi_16", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 2, 1, 1, 1, 0, 0,
		}, 17}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2},
		{"TestCalcHuXi_17", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 0, 0,
		}, 8}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1},

		// 带将牌胡+带坎
		{"TestCalcHuXi_21", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 2,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 14}, HUXI_SHUN_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_22", args{[]int{
			0, 0, 0, 0, 2, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 3, 0, 1, 1, 1,
		}, 11}, HUXI_SHUN_BIG*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_23", args{[]int{
			0, 0, 0, 3, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 3, 0, 2, 0, 0, 0, 0,
		}, 11}, HUXI_SHUN_BIG*1 + HUXI_KAN_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_24", args{[]int{
			2, 0, 0, 1, 1, 1, 0, 0, 0, 3,
			1, 1, 1, 0, 0, 0, 0, 3, 0, 3,
		}, 17}, HUXI_SHUN_BIG*1 + HUXI_KAN_BIG*2 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_25", args{[]int{
			3, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 2,
		}, 14}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_26", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 3, 2, 1, 1, 1, 0, 0,
		}, 20}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_27", args{[]int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 2, 3, 3,
		}, 14}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1 + HUXI_KAN_BIG*2},

		// 不带将牌胡+带坎
		{"TestCalcHuXi_31", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		}, 12}, HUXI_SHUN_BIG*1 + HUXI_KAN_SMALL*1},
		{"TestCalcHuXi_32", args{[]int{
			3, 0, 0, 0, 3, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 1, 1, 1,
		}, 12}, HUXI_SHUN_BIG*1 + HUXI_KAN_SMALL*2},
		{"TestCalcHuXi_33", args{[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 3,
		}, 6}, HUXI_SHUN_BIG*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_34", args{[]int{
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 3, 0, 0, 3, 0, 0,
		}, 12}, HUXI_SHUN_BIG*1 + HUXI_KAN_BIG*2},
		{"TestCalcHuXi_35", args{[]int{
			0, 1, 0, 1, 1, 1, 1, 0, 0, 1,
			1, 1, 1, 3, 0, 0, 0, 0, 0, 0,
		}, 12}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_36", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 0, 3,
		}, 18}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2 + HUXI_KAN_BIG*1},
		{"TestCalcHuXi_37", args{[]int{
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
			1, 1, 1, 0, 3, 0, 0, 0, 0, 3,
		}, 18}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*1 + HUXI_KAN_SMALL*2 + HUXI_KAN_BIG*2},
		{"TestCalcHuXi_38", args{[]int{
			0, 2, 0, 1, 1, 1, 2, 0, 0, 2,
			1, 1, 1, 0, 0, 1, 1, 1, 3, 3,
		}, 21}, HUXI_SHUN_BIG*1 + HUXI_SHUN_SMALL*2 + HUXI_KAN_BIG*2},

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcHuXi(tt.args.pool, tt.args.handsNum); got != tt.want {
				t.Errorf("CalcHuXi() = %v, want %v", got, tt.want)
			}
		})
	}
}
