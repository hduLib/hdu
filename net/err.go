package net

import "fmt"

type ErrNotOk struct {
	StatusCode int
	Body       string
}

func (e *ErrNotOk) Error() string {
	return fmt.Sprintf("status code is %d:%s", e.StatusCode, e.Body)
}
