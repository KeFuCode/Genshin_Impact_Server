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
		player.ModRole.AddItem(itemId, num, player)
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 12)
	case csvs.ITEMTYPE_WEAPON:
		player.ModWeapon.AddItem(itemId, num)
	case csvs.ITEMTYPE_RELICS:
		player.ModRelics.AddItem(itemId, num)
	case csvs.ITEMTYPE_COOk:
		player.ModCook.AddItem(itemId)
	case csvs.ITEMTYPE_HOME_ITEM:
		player.ModHome.AddItem(itemId, num, player)
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

func (self *ModBag) RemoveItemToBag(itemId int, num int64, player *Player) {
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

func (self *ModBag) UseItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

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

	switch itemConfig.SortType {
	// case csvs.ITEMTYPE_NORMAL:
	// 	self.AddItemToBag(itemId, num)
	case csvs.ITEMTYPE_COOKBOOK:
		self.UseCookBook(itemId, num, player)
	case csvs.ITEMTYPE_FOOD:
		// add temp attribute to role
	default: // like to normal_item
		fmt.Println(itemId, "the item can't be used")
		return
	}
}

func (self *ModBag) UseCookBook(itemId int, num int64, player *Player) {
	cookBookConfig := csvs.GetCookBookConfig(itemId)
	if cookBookConfig == nil {
		fmt.Println(itemId, "cookbook isn't exist")
		return
	}

	self.RemoveItem(itemId, num, player)
	self.AddItem(cookBookConfig.Reward, num, player)
}
