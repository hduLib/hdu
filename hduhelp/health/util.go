package health

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func b64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func s1(data string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
