package main

import (
	"fmt"
	"github.com/hduLib/hdu/chaoxing"
	"github.com/hduLib/hdu/chaoxing/course"
	"log"
	"os"
)

var phone = os.Getenv("phone")
var passwd = os.Getenv("passwd")

func main() {
	user, err := chaoxing.LoginByPhoneAndPwd(phone, passwd)
	if err != nil {
		log.Fatalln(err)
		return
	}
	list, err := user.CourseList()
	if err != nil {
		log.Fatalln(err)
		return
	}
	list.Each(func(course *course.Brief) bool {
		c, err := course.Detail()
		if err != nil {
			log.Fatalln(err)
		}
		workList, err := c.WorkList()
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range workList.Works {
			if v.Status == "未交" {
				fmt.Printf("[%s]%s---%s\n", course.Title, v.Title, v.Time)
			}
		}
		return true
	})
}
