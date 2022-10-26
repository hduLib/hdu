package course

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hduLib/hdu/chaoxing/request"
	"github.com/hduLib/hdu/chaoxing/utils"
	"log"
	"strings"
)

type List struct {
	Courses []Brief
}

// todo
func (l *List) FindByName(name string) *Brief {
	var res *Brief
	l.Each(func(course *Brief) bool {
		if course.Title == name {
			res = course
			return false
		}
		return true
	})
	return res
}

func (l *List) Each(f func(course *Brief) bool) {
	for i := range l.Courses {
		if !f(&l.Courses[i]) {
			return
		}
	}
}

func NewCourseList(resp []byte, req *request.Request) (*List, error) {
	var list List
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	doc.Find(".learnCourse").Each(func(_ int, selection *goquery.Selection) {
		url, exist := selection.Find("a.color1").Attr("href")
		if !exist {
			log.Println("course url not existed")
		}
		imgUrl, exist := selection.Find("img").Attr("src")
		if !exist {
			log.Println("cover url not existed")
		}
		title := strings.TrimSpace(selection.Find("span").Contents().Text())
		info := selection.Find(".color3")
		teacher := info.Contents().Text()
		info = info.Next()
		var dur string
		fmt.Sscanf(info.Contents().Text(), "开课时间：%s", &dur)
		var CourseNum string
		fmt.Sscanf(info.Next().Contents().Text(), "班级：%s", &CourseNum)
		list.Courses = append(list.Courses, Brief{
			ClazzId:     utils.GetValueAttrBySelector(selection, ".clazzId"),
			CourseId:    utils.GetValueAttrBySelector(selection, ".courseId"),
			CoverURL:    imgUrl,
			url:         url,
			Title:       title,
			Duration:    dur,
			TeacherName: teacher,
			CourseNum:   CourseNum,
			req:         req,
		})
	})
	return &list, nil
}
