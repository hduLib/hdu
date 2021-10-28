package health

import (
	"testing"
	"time"
)

func TestCheckin(t *testing.T) {
	h := New()
	if err := h.SetToken("b478bccf-5c2f-478b-8410-30cd3b40facf"); err != nil {
		t.Error(err)
	}
	if err := h.Checkin(); err != nil {
		t.Error(err)
	}
}

func TestValidate(t *testing.T) {
	h := New()
	if err := h.SetToken("b478bccf-5c2f-478b-8410-30cd3b40facf"); err != nil {
		t.Error(err)
	}
	validate, err := h.Validate()
	if err != nil {
		t.Error(err)
	}
	t.Logf("还剩下%d天", int(time.Until(time.Unix(validate.ExpiredTime, 0)).Hours())/24)
}

func TestHDUHealth_Info(t *testing.T) {
	h := New()
	if err := h.SetToken("b478bccf-5c2f-478b-8410-30cd3b40facf"); err != nil {
		t.Error(err)
	}
	t.Log(h.Info())
}
