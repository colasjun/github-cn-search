package condition_shop

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type MenuData struct {
	UnitData []MenuUnitData `json:"unitData"`
}

type MenuUnitData struct {
	UnitEN string `json:"unitEN"`// like name description and so on
	UnitCN string `json:"unitCN"`// explain for zh
	UnitValue []string `json:"unitValue"`
}

func Menu() (m MenuData){
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
				m.UnitData = append(m.UnitData, MenuUnitData)
			}
		}
	}

	return m
}