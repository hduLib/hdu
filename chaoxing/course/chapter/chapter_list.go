package chapter

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/hduLib/hdu/chaoxing/request"
)

type List struct {
	Chapters []Brief
}

func NewList(resp []byte, req *request.Request) (*List, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	chapters := doc.Find(".chapter_item")
	chapters.Each(func(_ int, selection *goquery.Selection) {
		//todo: parse chapter
	})
	return nil, nil
}
