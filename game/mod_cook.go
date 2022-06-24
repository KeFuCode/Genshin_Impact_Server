package game

import (
	"fmt"
	"server/bin/csvs"
)

type Cook struct {
	CookId int
}

type ModCook struct {
	CookInfo map[int]*Cook
}

func (self *ModCook) AddItem(cookId int) {
	_, ok := self.CookInfo[cookId]
	if ok {
		fmt.Println("already learned cook_skill:", csvs.GetItemName(cookId))
		return
	}

	configCook := csvs.GetCookConfig(cookId)
	if configCook == nil {
		fmt.Println("illegal cook_skill: ", csvs.GetItemName(cookId))
		return
	}

	self.CookInfo[cookId] = &Cook{CookId: cookId}
	fmt.Println("get cook_skill: ", cookId)
}