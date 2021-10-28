//健康打卡自动打卡
package main

/*
q群机器人+自动健康打卡
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

func main() {
	b = bot.New(driver.NewWsDriver("ws://localhost:6700", ""))
	b.Run()
	b.Attach(&bot.PrivateMsgHandler{
		Priority: 1, F: func(MsgID int32, UserID int64, Msg message.Msg) bool {
			if Msg[0].GetType() != "text" {
				return false
			}
			msg := Msg[0].(message.Text).Text
			matches := regexpToken.FindStringSubmatch(msg)
			if len(matches) != 2 {
				return false
			}
			c := checkin.New()
			if err := c.SetToken(matches[1]); err != nil {
				log.Fatal(err)
			}
			students[UserID] = c
			sendMsg(UserID, "添加token成功")
			return true
		}})
	log.Println("开始自动打卡")

	for {
		now := time.Now()
		t := time.NewTimer(time.Until(time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, now.Location())))
		<-t.C
		for qq, c := range students {
			if err := c.Checkin(); err != nil {
				sendMsg(qq, err.Error())
				continue
			}
			validate, err := c.Validate()
			if err != nil {
				sendMsg(qq, err.Error())
				continue
			}
			lastDays := int(time.Until(time.Unix(validate.ExpiredTime, 0)).Hours() / 24)
			sendMsg(qq, fmt.Sprintf("打卡成功,token还能打卡%d天", lastDays))
		}
	}
}

func sendMsg(qq int64, txt string) error {
	_, err := b.SendPrivateMsg(qq, message.New().Text(txt))
	return err
}
