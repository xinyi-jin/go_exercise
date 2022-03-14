package alg599

/* 599. 两个列表的最小索引总和
假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。

你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。



示例 1:

输入: list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
输出: ["Shogun"]
解释: 他们唯一共同喜爱的餐厅是“Shogun”。
示例 2:

输入:list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["KFC", "Shogun", "Burger King"]
输出: ["Shogun"]
解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。


提示:

1 <= list1.length, list2.length <= 1000
1 <= list1[i].length, list2[i].length <= 30
list1[i] 和 list2[i] 由空格 ' ' 和英文字母组成。
list1 的所有字符串都是 唯一 的。
list2 中的所有字符串都是 唯一 的。 */

// 思路：存储较短list的哈希表，另一个list直接查表
func findRestaurant(list1 []string, list2 []string) []string {
	max := len(list1) + len(list2)
	res := []string{}
	tempMap := make(map[string]int)
	for i, v := range list1 {
		tempMap[v] = i
	}
	for i, v := range list2 {
		if n, ok := tempMap[v]; ok {
			if n+i < max {
				max = n + i
				res = []string{v}
			} else if n+i == max {
				res = append(res, v)
			}
		}
	}
	return res
}

// 特定条件下，性能会更好
func findRestaurantEx(list1 []string, list2 []string) (res []string) {
	if len(list1) <= len(list2) {
		res = genResStr(list2, genMap(list1))
	} else {
		res = genResStr(list1, genMap(list2))
	}
	return
}

func genMap(list []string) map[string]int {
	res := make(map[string]int)
	for i, v := range list {
		res[v] = i
	}
	return res
}

func genResStr(list []string, tempMap map[string]int) (res []string) {
	max := len(list) + len(tempMap)
	for i, v := range list {
		if n, ok := tempMap[v]; ok {
			if n+i < max {
				max = n + i
				res = []string{v}
			} else if n+i == max {
				res = append(res, v)
			}
		}
	}
	return res
}
