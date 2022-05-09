package health

import (
	"testing"
	"time"
)

func TestCheckin(t *testing.T) {
	h := New()
	if err := h.SetToken("b1585681-0f01-4ded-be80-e162fd070b82"); err != nil {
		t.Error(err)
	}
	if err := h.Checkin(); err != nil {
		t.Error(err)
	}
}

func TestValidate(t *testing.T) {
	h := New()
	if err := h.SetToken("b1585681-0f01-4ded-be80-e162fd070b82"); err != nil {
		t.Error(err)
	}
	validate, err := h.Validate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("还剩下%d天", int(time.Until(time.Unix(validate.ExpiredTime, 0)).Hours())/24)
}

func TestHDUHealth_Info(t *testing.T) {
	h := New()
	if err := h.SetToken("b1585681-0f01-4ded-be80-e162fd070b82"); err != nil {
		t.Error(err)
	}
	t.Log(h.Info())
}

func TestLoginWithCas(t *testing.T) {
	h := New()
	if err := h.LoginWithCas("21111111", "666666"); err != nil {
		t.Error(err)
	}
}
