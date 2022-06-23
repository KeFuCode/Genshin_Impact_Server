package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigCard struct {
	CardId       int `json:"CardId"`
	Friendliness int `json:"Friendliness"`
}

var ConfigCardMap map[int]*ConfigCard

func init() {
	ConfigCardMap = make(map[int]*ConfigCard)

	utils.GetCsvUtilMgr().LoadCsv("Card", &ConfigCardMap)

	fmt.Println("init csv_card")
}

func GetCardConfig(cardId int) *ConfigCard  {
	return ConfigCardMap[cardId]
}