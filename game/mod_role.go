package game

import (
	"fmt"
	"server/bin/csvs"
)

type RoleInfo struct {
	RoleId   int
	GetTimes int
	// level experience relic
}

type ModRole struct {
	RoleInfo map[int]*RoleInfo
}

func (self *ModRole) IsHasRole(roleId int) bool {
	return true
}

func (self *ModRole) GetRoleLevel(roleId int) int {
	return 80
}

func (self *ModRole) AddItem(roleId int, num int64) {
	for i := 0; i < int(num); i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			self.RoleInfo[roleId] = data
		} else {
			// judge the real item
			fmt.Println("get real item ...")
			self.RoleInfo[roleId].GetTimes++
		}
	}

	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("get role: ", itemConfig.ItemName, "times: ", roleId, "----", self.RoleInfo[roleId].GetTimes, "times")
	}
}
