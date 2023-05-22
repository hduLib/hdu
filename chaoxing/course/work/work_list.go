package work

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/hduLib/hdu/chaoxing/request"
	"github.com/hduLib/hdu/chaoxing/utils"
	"log"
	"strings"
)

type List struct {
	Works []Brief
}

func NewList(resp []byte, req *request.Request) (*List, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	var list List
	doc.Find("li").Each(func(i int, selection *goquery.Selection) {
		url, exist := selection.Attr("data")
		if !exist {
			log.Println("url not exist")
		}
		list.Works = append(list.Works, Brief{
			url:    url,
			Title:  selection.Find(".overHidden2").Contents().Text(),
			Time:   utils.ParseLeftTime2Deadline(strings.TrimSpace(selection.Find(".time").Contents().Text())),
			Status: selection.Find(".status").Contents().Text(),
			req:    req,
		})
	})
	return &list, nil
}

func NewListPhoneAPI(resp []byte, req *request.Request) (*List, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	var list List
	doc.Find("li").Each(func(i int, selection *goquery.Selection) {
		url, exist := selection.Attr("data")
		if !exist {
			log.Println("url not exist")
		}
		pos := strings.Index(url, "clazzId=")

		list.Works = append(list.Works, Brief{
			url:     url,
			ClazzId: url[pos+8 : pos+16],
			Title:   selection.Find("p").Contents().Text(),
			Time:    utils.ParseLeftTime2Deadline(strings.TrimSpace(selection.Find(".fr").Contents().Text())),
			Status:  selection.Find(".status").Contents().Text(),
			req:     req,
		})
	})
	return &list, nil
}
