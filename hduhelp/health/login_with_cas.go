package health

import (
	"errors"
	"github.com/hduLib/hdu/cas"
	"net/http"
	"regexp"
)

const casLoginURL = "https://api.hduhelp.com/login/direct/cas?clientID=healthcheckin&redirect=https://healthcheckin.hduhelp.com/#/auth"

var tokenRegexp = regexp.MustCompile("https://healthcheckin.hduhelp.com/\\?auth=([0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12})")

func (h *Health) LoginWithCas(user, passwd string) error {

	req, err := cas.GenLoginReq(casLoginURL, user, passwd)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	tmp := tokenRegexp.FindStringSubmatch(res.Request.URL.String())
	if len(tmp) <= 1 {
		return errors.New("token is not included in redirect url")
	}
	h.token = tmp[1]
	return nil
}
