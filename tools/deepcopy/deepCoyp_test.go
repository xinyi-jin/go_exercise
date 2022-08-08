package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @Description: 压测深拷贝 -gob
 */
func BenchmarkDeepCopy_Gob(b *testing.B) {
	src := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByGob(src, new(BidRequest))
	}
}

/**
 * @Description: 压测深拷贝 -json
 */
func BenchmarkDeepCopy_Json(b *testing.B) {
	src := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByJson(src, new(BidRequest))
	}
}

/**
 * @Description: 压测深拷贝 -custom
 */
func BenchmarkDeepCopy_custom(b *testing.B) {
	src := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByCustom(src, new(BidRequest))
	}
}

/**
 * @Description: 测试拷贝是否ok
 */
func TestCpoyIsOk(t *testing.T) {
	src := initData()
	//1.gob
	dst01 := new(BidRequest)
	DeepCopyByGob(src, dst01)
	bs01, _ := json.Marshal(dst01)
	fmt.Printf("%v\n", string(bs01))

	//2.json
	dst02 := new(BidRequest)
	DeepCopyByJson(src, dst02)
	bs02, _ := json.Marshal(dst02)
	fmt.Printf("%v\n", string(bs02))

	//3.custom
	dst03 := new(BidRequest)
	DeepCopyByCustom(src, dst03)
	bs03, _ := json.Marshal(dst02)
	fmt.Printf("%v\n", string(bs03))
}
