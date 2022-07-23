package usermodule

import (
	"go_exercise/chatroom/log"
	"math/rand"
	"time"
)

var logger *log.Logger

type User struct {
	username string
}

func (u *User) GetUserName() string {
	return u.username
}

func (u *User) ModifyUserName(name string) {
	logger.Infof("(u *User) ModifyUserName %s-->%s\n", u.username, name)
	u.username = name
}

func (u *User) SetUserName(username string) {
	u.username = username
}

func GenUserName() string {
	return userNameArr[rand.Intn(len(userNameArr))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
	logger = log.NewLogger("user_log", log.DEBUG)
}

func NewUser(name string) *User {
	if name == "" {
		name = GenUserName()
	}
	return &User{
		username: name,
	}
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
