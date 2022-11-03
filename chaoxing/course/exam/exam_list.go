package exam

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hduLib/hdu/chaoxing/request"
	"github.com/hduLib/hdu/chaoxing/utils"
	"strings"
)

type List struct {
	Exams []Brief
}

func NewList(resp []byte, req *request.Request, cpi, classId string) (*List, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	var list List
	doc.Find("li").Each(func(i int, selection *goquery.Selection) {
		status := selection.Find(".status").Contents().Text()
		onclick := selection.Contents().Get(1).Attr[1].Val
		var courseId, tId, paperId, lookpaperEnc, url string
		if len(onclick) > 10 {
			subs := strings.Split(onclick[8:len(onclick)-2], ",")
			if len(subs) == 7 {
				courseId = strings.Trim(subs[0], "'")
				tId = subs[1]
				//id := subs[2]
				//endTime := strings.Trim(subs[3], "'")
				paperId = subs[4]
				//isRetest := false
				//if subs[5] == "true" {
				//	isRetest = true
				//}
				lookpaperEnc = strings.Trim(subs[6], "'")
				url = fmt.Sprintf("https://mooc2-ans.chaoxing.com/exam-ans/exam/lookPaper?courseId=%s&classId=%s&paperId=%s&position=test&examRelationId=%s&cpi=%s&enc=%s&newMooc=true", courseId, classId, paperId, tId, cpi, lookpaperEnc)
			}
		}
		list.Exams = append(list.Exams, Brief{
			url:    url,
			Title:  selection.Find(".overHidden2").Contents().Text(),
			Time:   utils.ParseLeftTime2Deadline(strings.TrimSpace(selection.Find(".time").Contents().Text())),
			Status: status,
			req:    req,
		})
	})
	return &list, nil
}
