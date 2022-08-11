package skl

import (
	"math/rand"
	"time"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")

func GenTicket() string {
	var res [21]byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 21; i++ {
		res[i] = letters[rand.Intn(64)]
	}
	return string(res[:])
}
