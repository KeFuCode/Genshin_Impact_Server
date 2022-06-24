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

func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

	switch itemConfig.SortType {
	// case csvs.ITEMTYPE_NORMAL:
	// 	self.AddItemToBag(itemId, num)
	case csvs.ITEMTYPE_ROLE:
		player.ModRole.AddItem(itemId, num)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 12)
	default: // like to normal_item
		self.AddItemToBag(itemId, num)
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

func (self *ModBag) RemoveItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemToBagGM(itemId, num)
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

func (self *ModBag) RemoveItemToBag(itemId int, num int64) {
	if !self.HasEnoughItem(itemId, num) {
		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
			fmt.Println(config.ItemName, "not enough", "----now num: ", nowNum)
		}
		return
	}

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

func (self *ModBag) HasEnoughItem(itemId int, num int64) bool {
	_, ok := self.BagInfo[itemId]
	if !ok {
		return false
	} else if self.BagInfo[itemId].ItemNum < num {
		return false
	}

	return true
}
