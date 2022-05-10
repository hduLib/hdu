//健康打卡自动打卡
package main

/*
qq机器人+自动健康打卡
需要自主抓包获取token，打卡需要微信或钉钉oauth得到的token，这一步很难自动化
*/

import (
	"fmt"
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	checkin "github.com/BaiMeow/hdu/hduhelp/health"
	"log"
	"regexp"
	"time"
)

var b *bot.Bot

var students = make(map[int64]*checkin.Health)

var regexpToken = regexp.MustCompile("/checkin token ([0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12})")
var regexpLoc = regexp.MustCompile("/checkin location ([0-9]{6})")

func main() {
	b = bot.New(driver.NewWsDriver("ws://localhost:6700", ""))
	b.Attach(&bot.PrivateMsgHandler{
		Priority: 1, F: func(MsgID int32, UserID int64, Msg message.Msg) bool {
			if Msg[0].GetType() != "text" {
				return false
			}
			msg := Msg[0].(message.Text).Text
			//添加token
			matches := regexpToken.FindStringSubmatch(msg)
			if len(matches) == 2 {
				c := checkin.New()
				if err := c.SetToken(matches[1]); err != nil {
					log.Println(err)
					sendMsg(UserID, fmt.Sprintf("%v", err))
					return true
				}
				students[UserID] = c
				sendMsg(UserID, "添加token成功")
				return true
			}
			//修改地区编码
			matches = regexpLoc.FindStringSubmatch(msg)
			if len(matches) == 2 {
				if students[UserID] == nil {
					sendMsg(UserID, "未绑定token")
					return true
				}
				if err := students[UserID].SetLocation(matches[1]); err != nil {
					sendMsg(UserID, fmt.Sprintf("%v", err))
					return true
				}
				sendMsg(UserID, "已将地区编码设置为"+matches[1])
				return true
			}
			//修改状态
			if msg == "/checkin at home" {
				if students[UserID] == nil {
					sendMsg(UserID, "未绑定token")
					return true
				}
				students[UserID].AtHome()
				sendMsg(UserID, "已修改为在家")
				return true
			}
			if msg == "/checkin at school" {
				if students[UserID] == nil {
					sendMsg(UserID, "未绑定token")
					return true
				}
				students[UserID].AtSchool()
				sendMsg(UserID, "已修改为在学校")
				return true
			}
			//人工打卡
			if msg == "/checkin checkin" {
				if students[UserID] != nil {
					checkinAndValidate(students[UserID], UserID)
				} else {
					sendMsg(UserID, "未添加Token")
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
		for qq, c := range students {
			checkinAndValidate(c, qq)
		}
	}
}

func checkinAndValidate(c *checkin.Health, qq int64) {
	if err := c.Checkin(); err != nil {
		sendMsg(qq, err.Error())
		return
	}
	validate, err := c.Validate()
	if err != nil {
		sendMsg(qq, err.Error())
		return
	}
	lastDays := int(time.Until(time.Unix(validate.ExpiredTime, 0)).Hours() / 24)
	sendMsg(qq, fmt.Sprintf("打卡成功,token还能打卡%d天", lastDays))
}

func sendMsg(qq int64, txt string) {
	if _, err := b.SendPrivateMsg(qq, message.New().Text(txt)); err != nil {
		fmt.Printf("发送消息时出错(qq:%d,msg:%s)：%v", qq, txt, err)
	}
}
