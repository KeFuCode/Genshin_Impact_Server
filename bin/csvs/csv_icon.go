package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigIcon struct {
	IconId   int    `json:"IconId"`
}

var ConfigIconMap map[int]*ConfigIcon

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)

	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)

	fmt.Println("init csv_icon")
}

func GetIconConfig(iconId int) *ConfigIcon  {
	return ConfigIconMap[iconId]
}