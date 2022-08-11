package skl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const pushBody = "{\"currentLocation\":\"浙江省杭州市钱塘区\",\"city\":\"杭州市\",\"districtAdcode\":\"330114\",\"province\":\"浙江省\",\"district\":\"钱塘区\",\"healthCode\":0,\"healthReport\":0,\"currentLiving\":0,\"last14days\":0}"
const pushURL = "https://skl.hdu.edu.cn/api/punch"

func (skl *skl) Push() error {
	req, err := http.NewRequest(http.MethodPost, pushURL, bytes.NewBufferString(pushBody))
	if err != nil {
		return err
	}
	req.Header.Add("X-Auth-Token", skl.XAuthToken)
	req.Header.Add("skl-Ticket", GenTicket())
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		e, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("fail to push:%s", string(e))
	}
	return nil

}
