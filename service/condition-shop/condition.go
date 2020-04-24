package condition_shop

import (
	"fmt"
	"github-cn-search/service/common"
)

type ConditionMsg struct {
	code int // 0 success 1 failure
	reason string // some msg
	condition string // success result
}

func Condition(msg string) (s string,code int){
	fmt.Printf("condition_shop recieve msg success:%s", msg)

	if len(msg) < 1 {
		return "params empty",common.Code.FAIL
	}

	return "",common.Code.OK
}