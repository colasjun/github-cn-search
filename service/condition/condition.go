package condition

import (
	"fmt"
	"github-cn-search/service/common"
	"github.com/astaxie/beego/config"
)

type conditionMsg struct {
	Msg string
}

func Condition(params map[string]string) (s string,code int){
	fmt.Println("Condition receive msg success", params)

	if len(params) < 1 {
		return "params empty",common.Code.FAIL
	}

	// assemble condition
	config, e := config.NewConfig("ini", "./condition/condition.ini")
	if e != nil {
		return "Condition parse config ini error",common.Code.FAIL
	}

	var conditionMsg conditionMsg
	for k,v := range params {
		i := config.String(k)
		if len(i) < 1 {
			fmt.Printf("Condition config not exist for param %s\n", k)
			continue
		}

		conditionMsg.Msg = conditionMsg.Msg + " " + switchCondition(i,k,v)
	}

	if conditionMsg.Msg == "" {
		return "no param to search...",common.Code.FAIL
	}

	return conditionMsg.Msg,common.Code.OK
}

func switchCondition(conType string,k string,v string) (s string) {
	switch conType {
		case "in":
			return conditionIn(k, v)
		case "gt":
			return conditionGT(k, v)
		default:
			return ""
	}
}