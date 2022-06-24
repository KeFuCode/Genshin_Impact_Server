package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigRole struct {
	RoleId   int    `json:"RoleId"`
	ItemName string    `json:"ItemName"`
	Star int `json:"Star"`
	Stuff int `json:"Stuff"`
	StuffNum int `json:"StuffNum"`
	StuffItem int `json:"StuffItem"`
	StuffItemNum int `json:"StuffItemNum"`
	MaxStuffItem int `json:"MaxStuffItem"`
	MaxStuffItemNum int `json:"MaxStuffItemNum"`
}

var ConfigRoleMap map[int]*ConfigRole

func init() {
	ConfigRoleMap = make(map[int]*ConfigRole)

	utils.GetCsvUtilMgr().LoadCsv("Role", &ConfigRoleMap)

	fmt.Println("init csv_role")
}

func GetRoleConfig(roleId int) *ConfigRole {
	return ConfigRoleMap[roleId]
}
