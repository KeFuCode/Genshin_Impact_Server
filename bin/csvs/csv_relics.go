package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigRelics struct {
	RelicsId int `json:"RelicsId"`
	Type     int `json:"Type"`
	Pos      int `json:"Pos"`
	Star     int `json:"Star"`
}

var ConfigRelicsMap map[int]*ConfigRelics

func init() {
	ConfigRelicsMap = make(map[int]*ConfigRelics)
	utils.GetCsvUtilMgr().LoadCsv("Relics", &ConfigRelicsMap)

	fmt.Println("init csv_relics")
}

func GetRelicsConfig(RelicsId int) *ConfigRelics {
	return ConfigRelicsMap[RelicsId]
}
