package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigCard struct {
	CardId       int `json:"CardId"`
	Friendliness int `json:"Friendliness"`
	Check        int `json:"Check"`
}

var ConfigCardMap map[int]*ConfigCard
var ConfigCardMapByRoleId map[int]*ConfigCard

func init() {
	ConfigCardMap = make(map[int]*ConfigCard)
	utils.GetCsvUtilMgr().LoadCsv("Card", &ConfigCardMap)
	ConfigCardMapByRoleId = make(map[int]*ConfigCard)
	for _, v := range ConfigCardMap {
		ConfigCardMapByRoleId[v.Check] = v
	}

	fmt.Println("init csv_card")
}

func GetCardConfig(cardId int) *ConfigCard {
	return ConfigCardMap[cardId]
}

func GetCardConfigByRoleId(roleId int) *ConfigCard {
	return ConfigCardMapByRoleId[roleId]
}