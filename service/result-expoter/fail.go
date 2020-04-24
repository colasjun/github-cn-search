package result_expoter

import "github-cn-search/service/common"

type FailData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func FailReturn(msg string) (FailData FailData) {
	FailData.Code = common.Code.OK
	FailData.Msg = msg
	return FailData
}