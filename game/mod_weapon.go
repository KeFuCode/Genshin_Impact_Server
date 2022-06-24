package game

import (
	"fmt"
	"server/bin/csvs"
)

type Weapon struct {
	WeaponId int
	KeyId    int // weapon unique_id
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey     int
}

func (self *ModWeapon) AddItem(weaponId int, num int64) {
	config := csvs.GetWeaponConfig(weaponId)
	if config == nil {
		fmt.Println("weapon config isn't exist")
		return
	}

	if len(self.WeaponInfo)+int(num) > csvs.WEAPON_MAX_COUNT {
		fmt.Println("weapon overflow max count")
		return
	}

	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = weaponId
		self.MaxKey++
		weapon.KeyId = self.MaxKey
		self.WeaponInfo[weapon.KeyId] = weapon
		fmt.Println("get weapon:", csvs.GetItemName(weaponId), "--- weapon id:", weapon.KeyId)
	}
}
