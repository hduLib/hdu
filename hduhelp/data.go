package hduhelp

const timeURL = "https://api.hduhelp.com/time"

type TimeResp struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		SchoolYear             string `json:"schoolYear"`
		Semester               string `json:"semester"`
		SemesterStartTimestamp int    `json:"semester_start_timestamp"`
		WeekNow                int    `json:"weekNow"`
		WeekDayNow             int    `json:"weekDayNow"`
		TimeStamp              int    `json:"timeStamp"`
		Section                int    `json:"section"`
	} `json:"data"`
	Cache bool `json:"cache"`
}
