package main

import "fmt"

/* 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 简单实现
var (
	rs1, rs2, rs3, rs4, rs5 int
)

type LRUCache struct {
	cacheMaps  map[int]int //存储缓存键值数据，values 标识返回状态
	timeWeight []int       //时间长短标识，存储map key 标识时间权重 0 ----> 变大
	cacheSize  int
}

func Constructor(capacity int) *LRUCache {
	cm := new(LRUCache)
	cm.cacheMaps = make(map[int]int, capacity)
	cm.timeWeight = make([]int, 0)
	cm.cacheSize = capacity - 1
	return cm
}

func (this *LRUCache) get(key int) int {
	if _, ok := this.cacheMaps[key]; ok {
		// 去除原先存在的权重值
		for i, v := range this.timeWeight {
			if v == key {
				this.timeWeight = append(this.timeWeight[:i], this.timeWeight[i+1:]...)
				break
			}
		}
		this.timeWeight = append(this.timeWeight[:], key)
		return this.cacheMaps[key]
	}

	return -1
}

func (this *LRUCache) put(key int, value int) {
	if _, ok := this.cacheMaps[key]; ok {
		this.cacheMaps[key] = value
		// 去除原先存在的权重值
		for i, v := range this.timeWeight {
			if v == key {
				this.timeWeight = append(this.timeWeight[:i], this.timeWeight[i+1:]...)
				break
			}
		}
	} else if len(this.cacheMaps) <= this.cacheSize {
		this.cacheMaps[key] = value
	} else {
		delete(this.cacheMaps, this.timeWeight[0])
		this.timeWeight = this.timeWeight[1:]
		this.cacheMaps[key] = value
	}
	this.timeWeight = append(this.timeWeight[:], key)
}

func main() {

	cache := Constructor(2)
	cache.put(1, 1)
	cache.put(2, 2)
	rs1 = cache.get(1) // 返回  1
	cache.put(3, 3)    // 该操作会使得密钥 2 作废
	rs2 = cache.get(2) // 返回 -1 (未找到)
	cache.put(4, 4)    // 该操作会使得密钥 1 作废

	rs3 = cache.get(1) // 返回 -1 (未找到)
	rs4 = cache.get(3) // 返回  3
	rs5 = cache.get(4) // 返回  4

	fmt.Printf("result: %d  %d  %d  %d  %d", rs1, rs2, rs3, rs4, rs5)
}
