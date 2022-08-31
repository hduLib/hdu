package skl

const (
	pushURL     = "https://skl.hdu.edu.cn/api/punch"
	casLogin    = "https://skl.hdu.edu.cn/api/userinfo?type=&index=passcard.html"
	myURL       = "https://skl.hdu.edu.cn/api/passcard/my"
	userInfoURL = "https://skl.hdu.edu.cn/api/userinfo?type="
	leaveURL    = "https://skl.hdu.edu.cn/api/pass-leave/add"
	courseURL   = "https://skl.hdu.edu.cn/api/course"
)

var PushReqHDU = PushReq{
	CurrentLocation: "浙江省杭州市钱塘区",
	City:            "杭州市",
	DistrictAdcode:  "330114",
	Province:        "浙江省",
	District:        "钱塘区",
	HealthCode:      0,
	HealthReport:    0,
	CurrentLiving:   0,
	Last14Days:      0,
}

type PushReq struct {
	// 定位地址精确到县/区一级，如"浙江省杭州市钱塘区"
	CurrentLocation string `json:"currentLocation"`
	// 定位地级市，如"杭州市"
	City string `json:"city"`
	// 中国行政区划代码，精确到县/区一级，如钱塘区为 330114
	DistrictAdcode string `json:"districtAdcode"`
	// 省份，如"浙江省"
	Province string `json:"province"`
	// 县/区一级，如"钱塘区"
	District string `json:"district"`
	// 健康码状态，0绿码，1红码，2橙码，3未领取
	HealthCode int `json:"healthCode"`
	// 健康状况
	// 0 健康
	// 1 发烧
	// 2 咳嗽腹泻
	// 3 确诊病例
	// 4 疑似病例
	HealthReport int `json:"healthReport"`
	// 生活状况
	// 0 正常
	// 1 发热送检
	// 2 集中隔离
	// 3 社区要求居家隔离
	// 4 学校要求居家隔离
	// 5 其他
	CurrentLiving int `json:"currentLiving"`
	// 14天内密接触情况
	// 0 无
	// 1 密接
	// 2 次密接
	Last14Days int `json:"last14days"`
}

type MyResp struct {
	// 学号
	Id string `json:"id"`
	// 未知
	UnitId string `json:"unitId"`
	// 打卡状态
	HeathCheckStatus int `json:"heathCheckStatus"`
	// 健康码状态
	HeathCodeStatus int `json:"heathCodeStatus"`
	// 上次核酸的报告日期当天的0点的unix时间（ms）
	HeathCheckStartDate int64 `json:"heathCheckStartDate"`
	// 核酸状态，0为有有效的核酸报告，其他暂时未知
	HsjcStatus int `json:"hsjcStatus"`
	// 核酸检测有效期截止时间
	HsjcValidTime int64 `json:"hsjcValidTime"`
	// 最后一次核酸检测的报告时间
	HsjcLastTime int64 `json:"hsjcLastTime"`
	// 未知
	EntryStatus int `json:"entryStatus"`
	// 疑似为离校开始时间
	OutStartTime int64 `json:"outStartTime"`
	// 最后一次返校的时间
	InStartTime int64 `json:"inStartTime"`
	// 未知
	OutValidTime int64 `json:"outValidTime"`
	// 未知
	OutStatus int `json:"outStatus"`
	// 疑似为在寝室状态
	DormitoryStatus int `json:"dormitoryStatus"`
	// 疑似为最新寝室闸机刷脸时间
	DormitoryArrivalTime int64 `json:"dormitoryArrivalTime"`
	// 未知
	UpdateTime int64 `json:"updateTime"`
	// 未知
	Status int `json:"status"`
	// 未知
	Reason string `json:"reason"`
}

type UserInfoResp struct {
	// 姓名
	UserName string `json:"userName"`
	// 学生为1
	UserType int `json:"userType"`
	// 学院
	UnitId string `json:"unitId"`
	// 学院
	UnitCode string `json:"unitCode"`
	// 学院名称
	UnitName string `json:"unitName"`
	// 年级（入学年份）
	Grade string `json:"grade"`
	// 班号
	ClassNo string `json:"classNo"`
	// 性别 1为男
	Sex string `json:"sex"`
	// 专业
	Major string `json:"major"`
	// 手机号
	Phone string `json:"phone"`
	// 学号
	Id string `json:"id"`
	// 生日时间戳(ms)
	Birthday int64 `json:"birthday"`
	// 未知
	SchoolDay           interface{}   `json:"schoolDay"`
	Degree              interface{}   `json:"degree"`
	AcademicCredentials interface{}   `json:"academicCredentials"`
	RoleList            []interface{} `json:"roleList"`
	RoleIdList          interface{}   `json:"roleIdList"`
	TeacherName         interface{}   `json:"teacherName"`
}

type LeaveReq struct {
	// 格式yyyy-mm-dd
	StartDate string `json:"startDate"`
	// 留空
	EndDate string `json:"endDate"`
	// 原因
	Reason string `json:"reason"`
	// 未知
	AuditType int `json:"auditType"`
	// 离校时间 ms时间戳
	OutTime string `json:"outTime"`
	// 返校时间 ms时间戳
	InTime string `json:"inTime"`
	// 前往的地区的行政区划代码
	AreaCode string `json:"areaCode"`
	// 目的地，格式如"浙江省-杭州市-上城区"
	Destination string `json:"destination"`
	// 附件列表，疑似先上传到指定oss
	FileList []OSSFile `json:"fileList"`
}

type CourseResp struct {
	// 周次，请求参数中的startTime所对应的周次
	Week int `json:"week"`
	// 学年，如"2022-2023"
	Xn string `json:"xn"`
	// 学期，"1"或"2"
	Xq string `json:"xq"`
	// 请求参数的startTime,默认为当前学期第一天
	StartTime int64 `json:"startTime"`
	List      []struct {
		// 教师学院编号，如计算机为05
		TeacherUnitNo string `json:"teacherUnitNo"`
		// 教师学院名称
		TeacherUnitName string `json:"teacherUnitName"`
		// 未知，可能是教师号
		TeacherNo string `json:"teacherNo"`
		// 开课学年
		SchoolYear string `json:"schoolYear"`
		// 开课学期
		Semester string `json:"semester"`
		// 教师职称
		TeacherMajor string `json:"teacherMajor"`
		// 未知
		CourseSchemaId string `json:"courseSchemaId"`
		// 课程Id（不是课程代码，是UUID）
		CourseId string `json:"courseId"`
		// 课程名称
		CourseName string `json:"courseName"`
		// 上课节次
		StartSection int `json:"startSection"`
		// 下课节次
		EndSection int `json:"endSection"`
		// 开始上课周次
		StartWeek int `json:"startWeek"`
		// 结束上课周次
		EndWeek int `json:"endWeek"`
		// 单双周，可能为"单","双",""
		Period string `json:"period"`
		// 上课地址（教室）
		ClassRoom string `json:"classRoom"`
		// 上课weekday(1-6),周日未知
		WeekDay int `json:"weekDay"`
		// 教室名称
		TeacherName string `json:"teacherName"`
		// 课程代码（长，如："(2022-2023-1)-C5692034-2"）
		CourseCode string `json:"courseCode"`
		// 课程代码 (如:"C5692034")
		CourseNo string `json:"courseNo"`
		// 课程归属 (如：“艺术创作与审美体验”)
		CourseType string `json:"courseType"`
		// 学分
		Mark float64 `json:"mark"`
		// 未知
		ListenTime int `json:"listenTime"`
		// 未知
		ListenStatus interface{} `json:"listenStatus"`
		// 教学班组成
		CourseClass string `json:"courseClass"`
		// 未知
		TotalTime int `json:"totalTime"`
		// 学生数量
		StudentCount int `json:"studentCount"`
		// 开课学院名称
		UnitName string `json:"unitName"`
		// 开课学院编号
		UnitCode string `json:"unitCode"`
		// 开课时间（如："星期三第1-2节{1-17周}"）
		CourseSchema string `json:"courseSchema"`
		// 未知，貌似均为"1",怀疑研究室不是"1"
		StudentType string `json:"studentType"`
	} `json:"list"`
}
