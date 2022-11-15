package zjooc

import (
	"fmt"
)

const (
	loginUrl   = "https://service.zjooc.cn/service/centro/api/auth/app/authorize"
	chapterUrl = "https://service.zjooc.cn/service/jxxt/api/app/course/chapter/getStudentCourseChapters"
	videoUrl   = "https://service.zjooc.cn/service/learningmonitor/api/learning/monitor/videoPlaying"
	textUrl    = "https://service.zjooc.cn/service/learningmonitor/api/learning/monitor/finishTextChapter"
)

type publishStatus = int

const (
	Published publishStatus = iota + 3
	Finished
	// 未发布课程的publishStatus未知
)

type paperType = int

const (
	// Exam 考试
	Exam paperType = iota
	// Test 测验
	Test
	// Assignment 作业
	Assignment
)

func coursesUrl(status publishStatus, pageNo, pageSize int) string {
	return fmt.Sprintf("https://service.zjooc.cn/service/jxxt/api/app/course/student/course?publishStatus=%d&pageNo=%d&pageSize=%d", status, pageNo, pageSize)
}

func paperUrl(Type paperType, courseId string, batchKey string, pageNo, pageSize int) string {
	return fmt.Sprintf("https://service.zjooc.cn/service/tkksxt/api/admin/paper/student/page?paperType=%d&courseId=%s&batchKey=%s&pageNo=%d&pageSize=%d", Type, courseId, batchKey, pageNo, pageSize)
}

type Resp[T any] struct {
	// 成功为"操作成功"
	Message string `json:"message"`
	// 成功为0
	ResultCode int `json:"resultCode"`
	// 成功为true
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type LoginReq struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
	Type      int    `json:"type"`
}

type LoginResp struct {
	LoginResult struct {
		AccessToken       string `json:"access_token"`
		AuthorizationCode string `json:"authorization_code"`
		// 后续仅需要openid
		Openid           string `json:"openid"`
		RefreshToken     string `json:"refresh_token"`
		UserCenterOpenId string `json:"userCenterOpenId"`
	} `json:"loginResult"`
	User struct {
		// 证件号
		Certificate string `json:"certificate"`
		// 证件类型，"1"为身份证
		CertificateType string `json:"certificateType"`
		CompleteInfo    struct {
			// "杭州电子科技大学"
			CorpName          string `json:"corpName"`
			InputCertificate  int    `json:"inputCertificate"`
			InputEmail        int    `json:"inputEmail"`
			InputLoginName    int    `json:"inputLoginName"`
			InputName         int    `json:"inputName"`
			InputPhone        int    `json:"inputPhone"`
			NeedComplete      int    `json:"needComplete"`
			RepeatCertificate int    `json:"repeatCertificate"`
			RepeatPhone       int    `json:"repeatPhone"`
		} `json:"completeInfo"`
		// 邮箱
		Email string `json:"email"`
		// 疑似与学期相关，第几个学期就是几
		Gender             int    `json:"gender"`
		Id                 string `json:"id"`
		IsEmailVerified    int    `json:"isEmailVerified"`
		IsPhoneVerified    int    `json:"isPhoneVerified"`
		IsUserNameModified int    `json:"isUserNameModified"`
		// "hdu_"+学号
		LoginName string `json:"loginName"`
		// 姓名
		Name string `json:"name"`
		// 学号
		No string `json:"no"`
		// 电话
		Phone   string `json:"phone"`
		PsdAuth string `json:"psdAuth"`
	} `json:"user"`
}

type Course struct {
	Id               string      `json:"id"`
	TeacherName      string      `json:"teacherName"`
	TeacherId        string      `json:"teacherId"`
	CourseName       string      `json:"courseName"`
	CourseImgUrl     string      `json:"courseImgUrl"`
	CourseProgress   float64     `json:"courseProgress"`
	PersistentPeriod int         `json:"persistentPeriod"`
	PublishStatus    int         `json:"publishStatus"`
	StartDate        string      `json:"startDate"`
	EndDate          string      `json:"endDate"`
	CurrentCycle     int         `json:"currentCycle"`
	CorpId           string      `json:"corpId"`
	TemplateType     interface{} `json:"templateType"`
	Source           int         `json:"source"`
	Qxfbzt           interface{} `json:"qxfbzt"`
	Profile          string      `json:"profile"`
	Current          int         `json:"current"`
	Duration         int         `json:"duration"`
	BatchId          string      `json:"batchId"`
}

type Paper struct {
	PublishTime string  `json:"publishTime"`
	EndTime     string  `json:"endTime"`
	ClassId     string  `json:"classId"`
	PaperId     string  `json:"paperId"`
	PaperName   string  `json:"paperName"`
	FinalScore  float64 `json:"finalScore"`
	TotalScore  float64 `json:"totalScore"`
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	// 0 沒做, 1 做了
	ReviewStatus int `json:"reviewStatus"`
	// 0 开放, 2 截止, 4 应该是截止了但是没交
	ProcessStatus int    `json:"processStatus"`
	ScorePropor   string `json:"scorePropor"`
	PaperStyle    int    `json:"paperStyle"`
	PaperArchive  int    `json:"paperArchive"`
}
