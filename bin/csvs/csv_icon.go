package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigIcon struct {
	IconId int `json:"IconId"`
	Check  int `json:"Check"`
}

var ConfigIconMap map[int]*ConfigIcon
var ConfigIconMapByRoleId map[int]*ConfigIcon

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)
	ConfigIconMapByRoleId = make(map[int]*ConfigIcon)
	for _, v := range ConfigIconMap {
		ConfigIconMapByRoleId[v.Check] = v
	}

	fmt.Println("init csv_icon")
}

func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func GetIconConfigByRoleId(roleId int) *ConfigIcon {
	return ConfigIconMapByRoleId[roleId]
}