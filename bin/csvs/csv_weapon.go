package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigWeapon struct {
	WeaponId int `json:"WeaponId"`
	Type     int `json:"Type"`
	Star     int `json:"Star"`
}

var ConfigWeaponMap map[int]*ConfigWeapon

func init() {
	ConfigWeaponMap = make(map[int]*ConfigWeapon)
	utils.GetCsvUtilMgr().LoadCsv("Weapon", &ConfigWeaponMap)

	fmt.Println("init csv_weapon")
}

func GetWeaponConfig(weaponId int) *ConfigWeapon {
	return ConfigWeaponMap[weaponId]
}