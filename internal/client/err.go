package client

import "fmt"

type ErrNotOk struct {
	StatusCode int
	Body       string
}

// for further err, do type assertion
func (e *ErrNotOk) Error() string {
	return fmt.Sprintf("status code is %d", e.StatusCode)
}
