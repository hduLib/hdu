package chaoxing

import "testing"

func TestEncAes(t *testing.T) {
	t.Log(encryptByAES("123123123") == "GszvoF3bQseqgnnv/WhUZA==")
}
