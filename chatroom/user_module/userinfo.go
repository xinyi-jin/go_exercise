package usermodule

import (
	"math/rand"
	"time"
)

type UserInfo struct {
	username string
}

func (u *UserInfo) GetUserName() string {
	return u.username
}

func (u *UserInfo) SetUserName(username string) {
	u.username = username
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func GenUserName() string {
	rand.Seed(time.Now().UnixNano())
	return userNameArr[rand.Intn(len(userNameArr))]

}

var userNameArr = []string{
	"毛毛",
	"猪猪",
	"偷偷",
	"狗狗",
	"啦啦",
	"小飞侠",
	"无敌金身人",
	"金刚不坏身",
	"斗罗大陆神",
	"远古异能兽",
	"上天入地猴",
	"古灵精怪鸟",
}
