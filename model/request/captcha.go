package request

type CaptchaRequest struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
