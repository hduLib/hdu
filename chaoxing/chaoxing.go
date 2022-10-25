package chaoxing

import (
	"net/http"
)

type Cx struct {
	cookie []*http.Cookie
}

func newUser(ck []*http.Cookie) *Cx {
	user := new(Cx)
	user.cookie = ck
	return user
}

func (cx *Cx) addCookieAndHeader2Req(req *http.Request) {
	for _, v := range cx.cookie {
		req.AddCookie(v)
	}
}
