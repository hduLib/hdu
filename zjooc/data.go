package zjooc

const (
	loginUrl   = "http://service.zjooc.cn/service/centro/api/auth/app/authorize"
	coursesUrl = "http://service.zjooc.cn/service/jxxt/api/app/course/student/course?publishStatus=3&pageNo=1&pageSize=10"
	chapterUrl = "http://service.zjooc.cn/service/jxxt/api/app/course/chapter/getStudentCourseChapters"
	videoUrl   = "http://service.zjooc.cn/service/learningmonitor/api/learning/monitor/videoPlaying"
	textUrl    = "http://service.zjooc.cn/service/learningmonitor/api/learning/monitor/finishTextChapter"
)

type Resp struct {
	// 成功为"操作成功"
	Message string `json:"message"`
	// 成功为0
	ResultCode int `json:"resultCode"`
	// 成功为true
	Success bool `json:"success"`
}

type LoginReq struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
	Type      int    `json:"type"`
}

type LoginResp struct {
	Data struct {
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
	} `json:"data"`
	Resp
}
