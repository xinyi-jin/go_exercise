package json

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

/**
 * @Description: 请求信息
 */
type BidRequest struct {
	ID     string  `json:"id"`
	Imp    []*Imp  `json:"imp"`
	Device *Device `json:"device"`
}

/**
 * @Description: imp对象
 */
type Imp struct {
	ID       string  `json:"id"`
	Tagid    string  `json:"tagid"`
	Bidfloor float64 `json:"bidfloor"`
}

/**
 * @Description: 设备信息
 */
type Device struct {
	Ua    string `json:"ua"`
	IP    string `json:"ip"`
	Geo   *Geo   `json:"geo"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Os    string `json:"os"`
	Osv   string `json:"osv"`
}

/**
 * @Description: 地理位置信息
 */
type Geo struct {
	Lat     int    `json:"lat"`
	Lon     int    `json:"lon"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

/**
 * @Description: 利用gob进行深拷贝
 */
func DeepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

/**
 * @Description: 利用json进行深拷贝
 */
func DeepCopyByJson(src, dst *BidRequest) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}

/**
 * @Description: 通过自定义进行copy
 */
func DeepCopyByCustom(src, dst *BidRequest) {
	dst.ID = src.ID
	dst.Device = &Device{
		Ua: src.Device.Ua,
		IP: src.Device.IP,
		Geo: &Geo{
			Lat: src.Device.Geo.Lat,
			Lon: src.Device.Geo.Lon,
		},
		Make:  src.Device.Make,
		Model: src.Device.Model,
		Os:    src.Device.Os,
		Osv:   src.Device.Osv,
	}
	dst.Imp = make([]*Imp, len(src.Imp))
	for index, imp := range src.Imp {
		//注意此处因为imp对象里无指针对象,所以可以直接使用等于
		dst.Imp[index] = imp
	}
}

func initData() *BidRequest {
	str := "{\"id\":\"MM7dIXz4H05qtmViqnY5dW\",\"imp\":[{\"id\":\"1\",\"tagid\":\"3979722720\",\"bidfloor\":0.01}],\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 10; SM-G960N Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36 (Mobile; afma-sdk-a-v212621039.212621039.0)\",\"ip\":\"192.168.1.0\",\"geo\":{\"lat\":0,\"lon\":0,\"country\":\"KOR\",\"region\":\"KR-11\",\"city\":\"Seoul\"},\"make\":\"samsung\",\"model\":\"sm-g960n\",\"os\":\"android\",\"osv\":\"10\"}}"
	ans := new(BidRequest)
	json.Unmarshal([]byte(str), &ans)
	return ans
}
