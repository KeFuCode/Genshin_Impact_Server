package game

import (
	"fmt"
	"server/bin/csvs"
)

type HomeItem struct {
	HomeItemId  int
	HomeItemNum int64
	keyId       int
}

type ModHome struct {
	HomeItemInfo map[int]*HomeItem
	// wait to do ...
	// UseHomeItemInfo map[int]*UseHomeItem
	// MapInfo map[int]*Map
}

func (self *ModHome) AddItem(itemId int, num int64, player *Player) {
	_, ok := self.HomeItemInfo[itemId]
	if ok {
		self.HomeItemInfo[itemId].HomeItemNum += num
	} else {
		self.HomeItemInfo[itemId] = &HomeItem{HomeItemId: itemId, HomeItemNum: num}
	}

	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("get home_item:", config.ItemName, "---- num: ", num, "----now num:", self.HomeItemInfo[itemId].HomeItemNum)
	}
}
