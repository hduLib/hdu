package zjooc

const pageSize = 10

func (u *User) CurrentCourses(status publishStatus) ([]Course, error) {
	var no = 0
	return allPages[Course](u, func() string {
		no++
		return coursesUrl(status, no, pageSize)
	})
}

// PapersByCourse batchKey 同 Course 的 batchId
func (u *User) PapersByCourse(courseId string, Type paperType, batchKey string) ([]Paper, error) {
	var no = 0
	return allPages[Paper](u, func() string {
		no++
		return paperUrl(Type, courseId, batchKey, no, pageSize)
	})
}
