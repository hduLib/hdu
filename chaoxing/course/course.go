package course

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hduLib/hdu/chaoxing/request"
	"github.com/hduLib/hdu/chaoxing/utils"
)

type Course struct {
	ClazzId           string
	CourseId          string
	CoverURL          string
	Title             string
	Duration          string
	TeacherName       string
	CourseNum         string
	cpi               string
	cfid              string
	bbsid             string
	heardUt           string
	fid               string
	opEnc             string
	enc               string
	oldEnc            string
	workEnc           string
	examEnc           string
	v                 string
	t                 string
	courseEvaluateUrl string
	req               *request.Request
}

func NewCourse(resp []byte, cb *Brief) (*Course, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, fmt.Errorf("fail to parse resp:%v", err)
	}
	return &Course{
		ClazzId:           cb.ClazzId,
		CourseId:          cb.CourseId,
		CoverURL:          cb.CoverURL,
		Title:             cb.Title,
		Duration:          cb.Duration,
		TeacherName:       cb.TeacherName,
		CourseNum:         cb.CourseNum,
		cpi:               utils.GetValueAttrBySelector(doc, "#cpi"),
		cfid:              utils.GetValueAttrBySelector(doc, "#cfid"),
		bbsid:             utils.GetValueAttrBySelector(doc, "#bbsid"),
		heardUt:           utils.GetValueAttrBySelector(doc, "#heardUt"),
		fid:               utils.GetValueAttrBySelector(doc, "#fid"),
		opEnc:             utils.GetValueAttrBySelector(doc, "#openc"),
		enc:               utils.GetValueAttrBySelector(doc, "#enc"),
		oldEnc:            utils.GetValueAttrBySelector(doc, "#oldenc"),
		workEnc:           utils.GetValueAttrBySelector(doc, "#workEnc"),
		examEnc:           utils.GetValueAttrBySelector(doc, "#examEnc"),
		v:                 utils.GetValueAttrBySelector(doc, "#v"),
		t:                 utils.GetValueAttrBySelector(doc, "#t"),
		courseEvaluateUrl: utils.GetValueAttrBySelector(doc, "#courseEvaluateUrl"),
		req:               cb.req,
	}, nil
}