package game

import (
	"fmt"
	"server/bin/csvs"
)

type ModBag struct {
}

func (self *ModBag) IsHasItem(itemId int) bool {
	return true
}

func (self *ModBag) AddItem(itemId int) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "item isn't exist")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		fmt.Println("common item: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("role: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("icon: ", itemConfig.ItemName)
	case csvs.ITEMTYPE_CARD:
		fmt.Println("card: ", itemConfig.ItemName)
	}
}
