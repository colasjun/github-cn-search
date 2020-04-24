package handler

import (
	"encoding/json"
	"fmt"
	cache_store "github-cn-search/service/cache-store"
	result_expoter "github-cn-search/service/result-expoter"
	"github.com/astaxie/beego/config"
	"net/http"
)

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

	fmt.Println("get menu result:", menu)
	bytes, e := json.Marshal(menu)
	if e != nil {
		fmt.Println("menu Engine parse json fail...err=",e)
		failResult,_ := json.Marshal(result_expoter.FailReturn("system error"))
		fmt.Fprintf(w, string(failResult))
		return nil
	}

	fmt.Fprintf(w, string(bytes))
	return nil
}

func menu() (m MenuData){
	configer,e := config.NewConfig("json", "./config/menu.json")
	if e != nil {
		fmt.Println("condition-shop menu read config init fail...err=",e)
		panic("condition-shop menu read config init fail...")
	}

	menuArray, err := configer.DIY("rootArray")
	if err != nil {
		fmt.Println("condition-shop menu read config DIY fail...err=",e)
		panic("condition-shop menu read config DIY fail...")
	}

	menuArrayCasted := menuArray.([]interface{})
	if menuArrayCasted == nil {
		fmt.Println("condition-shop menu read config cast fail...err=",e)
		panic("condition-shop menu read config cast fail...")
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

	return m
}


