//健康打卡自动打卡
package main

/*
qq机器人+自动健康打卡
已经适配新打卡系统
*/

import (
	"fmt"
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	"github.com/BaiMeow/hdu/skl"
	"log"
	"regexp"
	"time"
)

type profile struct {
	username, password string
	UserID             int64
}

var b *bot.Bot

var students = make(map[int64]*profile)

var regexpLogin = regexp.MustCompile("/checkin login (\\d{8,9}) (.*)")

func main() {
	b = bot.New(driver.NewWsDriver("ws://localhost:6700", ""))
	b.Attach(&bot.PrivateMsgHandler{
		Priority: 1, F: func(MsgID int32, UserID int64, Msg message.Msg) bool {
			if Msg[0].GetType() != "text" {
				return false
			}
			msg := Msg[0].(message.Text).Text
			//login
			matches := regexpLogin.FindStringSubmatch(msg)
			if len(matches) == 3 {
				c := profile{username: matches[1], password: matches[2], UserID: UserID}
				if _, err := skl.Login(c.username, c.password); err != nil {
					sendMsg(UserID, "登录失败")
				} else {
					sendMsg(UserID, "登录成功")
					students[UserID] = &c
				}
				return true
			}
			//人工打卡
			if msg == "/checkin checkin" {
				if students[UserID] != nil {
					students[UserID].checkin()
				} else {
					sendMsg(UserID, "未登录")
				}
				return true
			}
			return false
		},
	})
	err := b.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("开始自动打卡")

	for {
		now := time.Now()
		t := time.NewTimer(time.Until(time.Date(now.Year(), now.Month(), now.Day()+1, 7, 0, 0, 0, now.Location())))
		<-t.C
		for _, c := range students {
			c.checkin()
		}
	}
}

func (p *profile) checkin() {
	user, err := skl.Login(p.username, p.password)
	if err != nil {
		sendMsg(p.UserID, err.Error())
		return
	}
	err = user.Push()
	if err != nil {
		sendMsg(p.UserID, err.Error())
		return
	}
	sendMsg(p.UserID, fmt.Sprintf("打卡完成:%s", time.Now().Format("Jan 2 15:04:05")))
}

func sendMsg(qq int64, txt string) {
	if _, err := b.SendPrivateMsg(qq, message.New().Text(txt)); err != nil {
		fmt.Printf("发送消息时出错(qq:%d,msg:%s)：%v", qq, txt, err)
	}
}
