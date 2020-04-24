package handler

import (
	"encoding/json"
	"fmt"
	cache_store "github-cn-search/service/cache-store"
	"github-cn-search/service/common"
	result_expoter "github-cn-search/service/result-expoter"
	"github.com/astaxie/beego/config"
	"net/http"
)

type ReturnData struct {
	Code int `json:code`
	Data MenuData `json:data`
}

type MenuData struct {
	UnitData []MenuUnitData `json:"unitData"`
}

type MenuUnitData struct {
	UnitEN string `json:"unitEN"`// like name description and so on
	UnitCN string `json:"unitCN"`// explain for zh
	UnitValue []string `json:"unitValue"`
}

func MenuIndex(w http.ResponseWriter, r *http.Request) (e error) {
	menu := menu()

	fmt.Println("MenuIndex get menu result:", menu)
	bytes, e := json.Marshal(menu)
	if e != nil {
		fmt.Println("MenuIndex parse json fail...err=",e)
		failResult,_ := json.Marshal(result_expoter.FailReturn("system error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	fmt.Fprintf(w, string(bytes))
	return nil
}

func menu() (r ReturnData){
	configer,e := config.NewConfig("json", "./config/menu.json")
	if e != nil {
		fmt.Println("MenuIndex read config menu.json fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	}

	menuArray, err := configer.DIY("rootArray")
	if err != nil {
		fmt.Println("MenuIndex read config DIY fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	}

	var m MenuData
	menuArrayCasted := menuArray.([]interface{})
	if menuArrayCasted == nil {
		fmt.Println("MenuIndex read config cast fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	} else {
		var MenuUnitData MenuUnitData
		for _,v := range menuArrayCasted  {
			if v != nil {
				element := v.(map[string]interface{})
				MenuUnitData.UnitCN = element["cn"].(string)
				MenuUnitData.UnitEN = element["en"].(string)
				MenuUnitData.UnitValue = nil
				menuCacheResult := cache_store.CACHE.Get(element["key"].(string))
				fmt.Println(MenuUnitData.UnitEN,":cache value:", menuCacheResult)
				if len(menuCacheResult) > 0 {
					var UnitValue []string
					e := json.Unmarshal([]byte(menuCacheResult), &UnitValue)
					if e == nil {
						MenuUnitData.UnitValue = UnitValue
					}
				}
				m.UnitData = append(m.UnitData, MenuUnitData)
			}
		}
	}

	var returnData ReturnData
	returnData.Code = common.Code.OK
	returnData.Data = m

	return returnData
}


