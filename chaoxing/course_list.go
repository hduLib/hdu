package chaoxing

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type CourseBrief struct {
	ClazzId  string
	CourseId string
	CoverURL string
	Title    string
	// Duration 意义不大，不同老师填的不同，不推荐使用
	Duration    string
	TeacherName string
	// CourseNum 可能是上课地点可能是课程号，感觉全看教师自己填了什么，不推荐使用
	CourseNum string
	url       string
	// point to cx for getting further information
	cx *Cx
	//本来应该有一个名为cpi的字段，但是仅仅出现在url内并且没有摸清他的意义，暂时不予理会
}

type CourseList struct {
	Courses []CourseBrief
}

func (cx *Cx) CourseList() *CourseList {
	resp, err := cx.get(courseListURL())
	if err != nil {
		return nil
	}
	var list CourseList
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	doc.Find(".learnCourse").Each(func(_ int, selection *goquery.Selection) {
		clazzId, exist := selection.Find(".clazzId").Attr("value")
		if !exist {
			log.Println("clazzId not existed")
		}
		courseId, exist := selection.Find(".courseId").Attr("value")
		if !exist {
			log.Println("courseId not existed")
		}
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
		list.Courses = append(list.Courses, CourseBrief{
			ClazzId:     clazzId,
			CourseId:    courseId,
			CoverURL:    imgUrl,
			url:         url,
			Title:       title,
			Duration:    dur,
			TeacherName: teacher,
			CourseNum:   CourseNum,
			cx:          cx,
		})
	})
	return &list
}

// todo
func (l *CourseList) FindByName(name string) *CourseBrief {
	return nil
}

func (l *CourseList) Each(f func(course *CourseBrief) bool) {
	for i := range l.Courses {
		if !f(&l.Courses[i]) {
			return
		}
	}
}
