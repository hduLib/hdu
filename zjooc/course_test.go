package zjooc

import (
	"fmt"
	"os"
	"testing"
)

func TestUser_GetCurrentCourses(t *testing.T) {
	user, err := Login(os.Getenv("zjooc_account"), os.Getenv("zjooc_passwd"))
	if err != nil {
		t.Log(err)
		return
	}
	courses, err := user.CurrentCourses(Published)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(courses)
}

func TestUser_PapersByCourse(t *testing.T) {
	user, err := Login(os.Getenv("zjooc_account"), os.Getenv("zjooc_passwd"))
	if err != nil {
		t.Log(err)
		return
	}
	papers, err := user.PapersByCourse("2c91808281b87da50181cd026be14a85", Assignment, "20221")
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(papers)
}
