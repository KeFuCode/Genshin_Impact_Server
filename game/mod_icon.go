package game

import (
	"fmt"
	"server/bin/csvs"
)

type Icon struct {
	IconId int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := self.IconInfo[iconId]
	
	return ok
}

func (self *ModIcon) AddItem(iconId int) {
	_, ok := self.IconInfo[iconId]
	if ok {
		fmt.Println("icon already exist: ", iconId)
		return
	}

	configIcon := csvs.GetIconConfig(iconId)
	if configIcon == nil {
		fmt.Println("illegal icon: ", iconId)
		return
	}

	self.IconInfo[iconId] = &Icon{IconId: iconId}
	fmt.Println("get icon: ", iconId)
}
