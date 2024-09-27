package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

// 从神策分析配置页面中获取数据接收的 URL
const (
	LOG_ADDRESS  = "./access.log"
	HOUR_ROTATE  = false
	PROJECT_NAME = "test"
	TIME_FREE    = false
)

type Properties struct {
	Price int       `json:"price,omitempty"`
	Name  string    `json:"name,omitempty"`
	Money float64   `json:"money,omitempty"`
	Isold bool      `json:"isold,omitempty"`
	Ts    time.Time `json:"ts,omitempty"`
	List  []string  `json:"list,omitempty"`
}

type SAParamsBatch []*SAParams

type SAParams struct {
	DistinctId string
	Event      string
	Properties map[string]interface{}
	IsLoginId  bool
}

func main() {
	/* properties := map[string]interface{}{
		"price": 12,
		"name":  "apple",
	} */
	temp := &Properties{
		Price: 9000000000000000,
		Name:  "apple",
		Money: 12.06,
		Isold: true,
		Ts:    time.Now(),
		List:  []string{"apple", "banana", "orange"},
	}
	properties, err := ToMap(temp, "json")
	if err != nil {
		panic(err)
	}
	for k, v := range properties {
		fmt.Println(k, v, reflect.ValueOf(v).Kind())
	}

	// 神策分析数据接收的 URL
	// SA_SERVER_URL := "SA_SERVER_URL"
	// // 发送数据的超时时间，单位毫秒
	// SA_REQUEST_TIMEOUT := 100000
	// // Debug 模式下，是否将数据导入神策分析
	// //   true - 校验数据，并将数据导入到神策分析中
	// //   false - 校验数据，但不进行数据导入
	// SA_DEBUG_WRITE_DATA := true

	// // 初始化 Debug Consumer
	// consumer, err := sdk.InitDebugConsumer(SA_SERVER_URL, SA_DEBUG_WRITE_DATA, SA_REQUEST_TIMEOUT)
	// if err != nil {
	// 	panic(err)
	// }
	// // ...
	// // 使用 Consumer 来构造 SensorsAnalytics 对象
	// sa := sdk.InitSensorsAnalytics(consumer, "default", false)
	// defer sa.Close()

	// data := SAParams{
	// 	DistinctId: "60771863",
	// 	Event:      "GameEvent",
	// 	Properties: properties,
	// 	IsLoginId:  true,
	// }
	// err = sa.Track(data.DistinctId, data.Event, data.Properties, data.IsLoginId)

	// // 初始化一个 Consumer，用于数据发送
	// c, err := sdk.InitConcurrentLoggingConsumer(SA_SERVER_URL, HOUR_ROTATE)
	// if err != nil {
	// 	panic(err)
	// }

	// // 使用 Consumer 来构造 SensorsAnalytics 对象
	// sa := sdk.InitSensorsAnalytics(c, PROJECT_NAME, TIME_FREE)
	// defer sa.Close()

	// identity := sdk.Identity{
	// 	Identities: map[string]string{
	// 		sdk.LOGIN_ID: "60771863",
	// 		sdk.MOBILE:   "15226048650",
	// 		sdk.EMAIL:    "test@sensorsdata.cn",
	// 	},
	// }
	// sa.Bind(identity)

	// err = sa.Track(data.DistinctId, data.Event, data.Properties, data.IsLoginId)
	// sa.Track("1", data.Event, data.Properties, data.IsLoginId)
	// sa.Track("2", data.Event, data.Properties, data.IsLoginId)
	// sa.Track("3", data.Event, data.Properties, data.IsLoginId)
	sap := NewSAParams("95466200", "GameEvent", properties, false)
	err = sap.ReportOne()

	// var saps SAParamsBatch
	// saps = append(saps, NewSAParams("1", "GameEvent", properties, true))
	// saps = append(saps, NewSAParams("2", "GameEvent", properties, true))
	// saps = append(saps, NewSAParams("3", "GameEvent", properties, true))
	// saps.ReportBatch()
	if err != nil {
		fmt.Println("track failed", err)
		return
	}
}

func NewSAParams(distinctId, event string, properties map[string]interface{}, isLoginId bool) *SAParams {
	return &SAParams{
		DistinctId: distinctId,
		Event:      event,
		Properties: properties,
		IsLoginId:  isLoginId,
	}
}

func (s *SAParams) ReportOne() error {
	// 初始化一个 Consumer，用于数据发送
	consumer, err := sdk.InitConcurrentLoggingConsumer(LOG_ADDRESS, HOUR_ROTATE)
	if err != nil {
		panic(err)
	}

	// 使用 Consumer 来构造 SensorsAnalytics 对象
	sa := sdk.InitSensorsAnalytics(consumer, PROJECT_NAME, TIME_FREE)
	defer sa.Close()

	return sa.Track(s.DistinctId, s.Event, s.Properties, s.IsLoginId)
}

// 报405 方法不被允许
func (s *SAParams) ReportDefault() error {
	// 神策分析数据接收的 URL
	// SA_SERVER_URL := "SA_SERVER_URL"
	// // 发送数据的超时时间，单位毫秒
	// SA_REQUEST_TIMEOUT := 100000
	// // Debug 模式下，是否将数据导入神策分析
	// //   true - 校验数据，并将数据导入到神策分析中
	// //   false - 校验数据，但不进行数据导入
	// // SA_DEBUG_WRITE_DATA := true

	// // 初始化 Debug Consumer
	// consumer, err := sdk.InitDefaultConsumer(SA_SERVER_URL, SA_REQUEST_TIMEOUT)
	// if err != nil {
	// 	return err
	// }

	// // 使用 Consumer 来构造 SensorsAnalytics 对象
	// sa := sdk.InitSensorsAnalytics(consumer, PROJECT_NAME, TIME_FREE)
	// defer sa.Close()

	// 神策分析数据接收的 URL
	SA_SERVER_URL := "SA_SERVER_URL"
	// 发送数据的超时时间，单位毫秒
	SA_REQUEST_TIMEOUT := 100000

	// 初始化 Default Consumer
	consumer, err := sdk.InitDefaultConsumer(SA_SERVER_URL, SA_REQUEST_TIMEOUT)
	if err != nil {
		panic(err)
	}
	// 使用 Consumer 来构造 SensorsAnalytics 对象
	sa := sdk.InitSensorsAnalytics(consumer, PROJECT_NAME, TIME_FREE)
	defer sa.Close()

	return sa.Track(s.DistinctId, s.Event, s.Properties, s.IsLoginId)
}

func (s SAParamsBatch) ReportBatch() error {
	// 初始化一个 Consumer，用于数据发送
	consumer, err := sdk.InitConcurrentLoggingConsumer(LOG_ADDRESS, HOUR_ROTATE)
	if err != nil {
		panic(err)
	}

	// 使用 Consumer 来构造 SensorsAnalytics 对象
	sa := sdk.InitSensorsAnalytics(consumer, PROJECT_NAME, TIME_FREE)
	defer sa.Close()

	for _, v := range s {
		err := sa.Track(v.DistinctId, v.Event, v.Properties, v.IsLoginId)
		if err != nil {
			return err
		}
	}

	return nil
}

// ToMap 结构体转为Map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			if tag, _ := parseTag(tagValue); tag != "" {
				out[tag] = v.Field(i).Interface()
			}
		}
	}
	return out, nil
}

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}
