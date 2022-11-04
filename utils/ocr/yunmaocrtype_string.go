// Code generated by "stringer -type=YunmaOCRType"; DO NOT EDIT.

package ocr

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Common-1]
	_ = x[Slide-2]
	_ = x[SinSlide-3]
	_ = x[TrafficSlide-4]
	_ = x[Click-5]
	_ = x[Rotate-6]
	_ = x[Google-7]
	_ = x[Hcaptcha-8]
	_ = x[FunCaptcha-9]
}

const _YunmaOCRType_name = "CommonSlideSinSlideTrafficSlideClickRotateGoogleHcaptchaFunCaptcha"

var _YunmaOCRType_index = [...]uint8{0, 6, 11, 19, 31, 36, 42, 48, 56, 66}

func (i YunmaOCRType) String() string {
	i -= 1
	if i < 0 || i >= YunmaOCRType(len(_YunmaOCRType_index)-1) {
		return "YunmaOCRType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _YunmaOCRType_name[_YunmaOCRType_index[i]:_YunmaOCRType_index[i+1]]
}
