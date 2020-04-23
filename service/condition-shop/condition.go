package condition_shop

import "fmt"

type ConditionMsg struct {
	code int // 0 success 1 failure
	reason string // some msg
	condition string // success result
}

func Condition(msg string) (c ConditionMsg){
	c = ConditionMsg{0,"ok",""}

	fmt.Printf("condition_shop recieve msg success:%s", msg)

	if len(msg) < 1 {
		c.code = 1
		c.reason = "params empty"
		return c
	}

	return c
}