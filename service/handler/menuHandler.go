package handler

import (
	"encoding/json"
	"fmt"
	cache_store "github-cn-search/service/cache-store"
	"github-cn-search/service/common"
	"github-cn-search/service/result"
	"github.com/astaxie/beego/config"
	"net/http"
)

type ReturnData struct {
	Code int `json:"code"`
	Data MenuData `json:"data"`
}

type MenuData struct {
	UnitData []MenuUnitData `json:"unitData"`
}

type MenuUnitData struct {
	UnitEN string `json:"unitEN"`// like name description and so on
	UnitCN string `json:"unitCN"`// explain for zh
	SearchType string `json:"type"` // search type
	UnitValue []string `json:"unitValue"`
}

func MenuIndex(w http.ResponseWriter, r *http.Request) (e error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

	menu := menu()

	fmt.Println("MenuIndex get menu result:", menu)
	bytes, e := json.Marshal(menu)
	if e != nil {
		fmt.Println("MenuIndex parse json fail...err=",e)
		failResult,_ := json.Marshal(result.FailReturn("system error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	fmt.Fprintf(w, string(bytes))
	return nil
}

func menu() (r ReturnData){
	config,e := config.NewConfig("json", "./config/menu.json")
	if e != nil {
		fmt.Println("MenuIndex read config menu.json fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	}

	menuArray, err := config.DIY("rootArray")
	if err != nil {
		fmt.Println("MenuIndex read config DIY fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	}

	var m MenuData
	menuArrayCasted := menuArray.([]interface{})
	if menuArrayCasted == nil {
		fmt.Println("MenuIndex read config cast fail...err=",e)
		panic(common.FailMsg.PanicMsg)
	}

	var MenuUnitData MenuUnitData
	for _,v := range menuArrayCasted  {
		if v != nil {
			element := v.(map[string]interface{})
			MenuUnitData.UnitCN = element["cn"].(string)
			MenuUnitData.UnitEN = element["en"].(string)
			MenuUnitData.SearchType = element["type"].(string)
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

	var returnData ReturnData
	returnData.Code = common.Code.OK
	returnData.Data = m

	return returnData
}


