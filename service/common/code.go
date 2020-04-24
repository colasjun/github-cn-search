package common

type CodeData struct {
	FAIL int
	OK int
}

type FailMsgData struct {
	PanicMsg string
}

var Code *CodeData
var FailMsg *FailMsgData

func init() {
	Code = &CodeData{
		FAIL:500,
		OK:200,
	}

	FailMsg = &FailMsgData{
		PanicMsg:"fail...",
	}
}