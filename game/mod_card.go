package game

import (
	"fmt"
	"server/bin/csvs"
)

type Card struct {
	CardId int
}

type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) IsHasCard(cardId int) bool {
	_, ok := self.CardInfo[cardId]

	return ok
}

func (self *ModCard) AddItem(cardId int, friendliness int) {
	_, ok := self.CardInfo[cardId]
	if ok {
		fmt.Println("card already exist: ", cardId)
		return
	}

	configCard := csvs.GetCardConfig(cardId)
	if configCard == nil {
		fmt.Println("illegal card: ", cardId)
		return
	}
	if friendliness < configCard.Friendliness {
		fmt.Println(cardId, "friendliness not enough: ", friendliness)
		return
	}

	self.CardInfo[cardId] = &Card{CardId: cardId}
	fmt.Println("get card: ", cardId)
}

func (self *ModCard) CheckGetCard(roleId int, friendiness int) {
	config := csvs.GetCardConfigByRoleId(roleId)
	if config == nil {
		return
	}

	self.AddItem(config.CardId, friendiness)
}