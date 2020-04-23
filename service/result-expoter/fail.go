package result_expoter

type FailData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func FailReturn(msg string) (FailData FailData) {
	FailData.Code = 500
	FailData.Msg = msg
	return FailData
}