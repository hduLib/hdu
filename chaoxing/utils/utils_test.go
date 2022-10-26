package utils

import "testing"

func TestEncAes(t *testing.T) {
	t.Log(EncryptByAES("123123123") == "GszvoF3bQseqgnnv/WhUZA==")
}
