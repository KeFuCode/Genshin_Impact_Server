package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigHomeItem struct {
	HomeItemId int `json:"HomeItemId"`
	Type   int `json:"Type"`
}

var ConfigHomeItemMap map[int]*ConfigHomeItem

func init() {
	ConfigHomeItemMap = make(map[int]*ConfigHomeItem)
	utils.GetCsvUtilMgr().LoadCsv("Home", &ConfigHomeItemMap)

	fmt.Println("init csv_homeItem")
}

func GetHomeItemConfig(homeItemId int) *ConfigHomeItem {
	return ConfigHomeItemMap[homeItemId]
}
