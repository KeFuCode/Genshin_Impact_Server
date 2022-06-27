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

func (self *ModRole) AddItem(roleId int, num int64, player *Player) {
	config := csvs.GetRoleConfig(roleId)
	if config == nil {
		fmt.Println("role isn't exist:", roleId)
		return
	}

	for i := 0; i < int(num); i++ {
		_, ok := self.RoleInfo[roleId]
		if !ok {
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			self.RoleInfo[roleId] = data
		} else {
			// judge the real item
			// fmt.Println("get real item ...")
			self.RoleInfo[roleId].GetTimes++
			if self.RoleInfo[roleId].GetTimes >= csvs.ADD_ROLE_TIME_NORMAL_MIN &&
				self.RoleInfo[roleId].GetTimes <= csvs.ADD_ROLE_TIME_NORMAL_MAX {
				player.ModBag.AddItemToBag(config.Stuff, int64(config.StuffNum))
				player.ModBag.AddItemToBag(config.StuffItem, int64(config.StuffItemNum))
			} else {
				player.ModBag.AddItemToBag(config.MaxStuffItem, int64(config.MaxStuffItemNum))
			}
		}
	}

	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("get role: ", itemConfig.ItemName, "times: ", roleId, "----", self.RoleInfo[roleId].GetTimes, "times")
	}

	// check icon when role geted
	player.ModIcon.CheckGetIcon(roleId)

	// check card when role geted
	player.ModCard.CheckGetCard(roleId, 10) // friendliness: 10

}

func (self *ModRole) HandleSendRoleInfo() {
	fmt.Println(fmt.Sprintf("当前拥有角色信息如下:"))
	for _, v := range self.RoleInfo {
		v.SendRoleInfo()
	}
}

func (self *RoleInfo) SendRoleInfo() {
	fmt.Println(fmt.Sprintf("%s:,累计获得次数:%d", csvs.GetItemName(self.RoleId), self.GetTimes))
}

func (self *ModRole) GetRoleInfoForPoolCheck() (map[int]int, map[int]int) {
	fiveInfo := make(map[int]int)
	fourInfo := make(map[int]int)

	for _, v := range self.RoleInfo {
		roleConfig := csvs.GetRoleConfig(v.RoleId)
		if roleConfig == nil {
			continue
		}
		if roleConfig.Star == 5 {
			fiveInfo[roleConfig.RoleId] = v.GetTimes
		} else if roleConfig.Star == 4 {
			fourInfo[roleConfig.RoleId] = v.GetTimes
		}
	}
	return fiveInfo, fourInfo
}