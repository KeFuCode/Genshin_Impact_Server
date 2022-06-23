package game

import (
	"fmt"
	"server/bin/csvs"
)

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}

type ModBag struct {
	BagInfo map[int]*ItemInfo
}

func (self *ModBag) IsHasItem(itemId int) bool {
	return true
}

func (self *ModBag) AddItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.AddItemToBag(itemId, 1)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("role: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 12)
	default: // like to normal_item
		// self.AddItemToBag(itemId, 1)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum += num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: num}
	}

	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("get item:", config.ItemName, "---- num: ", num, "----now num:", self.BagInfo[itemId].ItemNum)
	}
}

func (self *ModBag) RemoveItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemToBagGM(itemId, 1)
	default: // like to normal_item
		// self.RemoveItemToBag(itemId, 1)
	}
}

func (self *ModBag) RemoveItemToBagGM(itemId int, num int64) {
	_, ok := self.BagInfo[itemId]
	if ok {
		self.BagInfo[itemId].ItemNum -= num
	} else {
		self.BagInfo[itemId] = &ItemInfo{ItemId: itemId, ItemNum: 0 - num}
	}

	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("reduce item:", config.ItemName, "----num: ", num, "----now num:", self.BagInfo[itemId].ItemNum)
	}
}
