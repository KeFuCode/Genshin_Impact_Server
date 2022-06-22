package game

import (
	"fmt"
	"server/bin/csvs"
	"time"
)

type ModPlayer struct {
	// 可见字段
	UserId         int // unique_id
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelNow  int
	WorldLevelCool int64 // operate world_level cool time
	ShowTeam       []int
	ShowCard       []int
	Birth          int

	// 隐藏字段
	IsProhibit int // account status
	IsGm       int // GM account status
}

// external interface: gamer set ModPlayer inner value
func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端，操作非法
		return
	}

	player.ModPlayer.Icon = iconId
	fmt.Println("Now Icon: ", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModCard.IsHasCard(cardId) {
		// 通知客户端，操作非法
		return
	}

	player.ModPlayer.Card = cardId
	fmt.Println("Now Card: ", player.ModPlayer.Card)
}

func (self *ModPlayer) RecvSetName(name string, player *Player) {

	player.ModPlayer.Name = name
	fmt.Println("Now Name: ", player.ModPlayer.Name)
}

func (self *ModPlayer) RecvSetSign(sign string, player *Player) {

	player.ModPlayer.Sign = sign
	fmt.Println("Now Sign: ", player.ModPlayer.Sign)
}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("operate fail: --- now world_level: ", self.WorldLevel)
		return
	}

	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("operate fail: --- now world_level: ", self.WorldLevel, "--- real world level: ", self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("operate fail: --- cooling")
		return
	}

	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("operate success: --- now world_level: ", self.WorldLevel, "--- real world level: ", self.WorldLevelNow)
	return
}

func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevel == self.WorldLevelNow {
		fmt.Println("operate fail: --- now world_level: ", self.WorldLevel, "--- real world level: ", self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("operate fail: --- cooling")
		return
	}

	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME
	fmt.Println("operate success: --- now world_level: ", self.WorldLevel, "--- real world level: ", self.WorldLevelNow)
	return
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {
	if self.Birth > 0 {
		fmt.Println("your birthday already set")
		return
	}

	month := birth / 100
	day := birth % 100

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "month doesn't have", day, "day")
			return
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "month doesn't have", day, "day")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println(month, "month doesn't have", day, "day")
			return
		}
	default:
		fmt.Println(month, "month isn't exist")
		return
	}

	self.Birth = birth
	fmt.Println("set success, birthday is: ", month, "month", day, "day")

	if self.IsBirthDay() {
		fmt.Println("happy birthday!")
	} else {
		fmt.Println("your birthday is coming soon~")
	}
}

func (self *ModPlayer) IsBirthDay() bool {
	month := time.Now().Month()
	day := time.Now().Day()

	if int(month) == self.Birth/100 && day == self.Birth%100 {
		return true
	}

	return false
}

func (self *ModPlayer) SetShowCard(showCard []int, player *Player) {
	cardExist := make(map[int]int)
	newList := make([]int, 0)

	for _, cardId := range showCard {
		_, ok := cardExist[cardId]
		if ok {
			continue
		}
		if !player.ModCard.IsHasCard(cardId) {
			continue
		}
		newList = append(newList, cardId)
		cardExist[cardId] = 1
	}

	self.ShowCard = newList
	fmt.Println(self.ShowCard)
}

// internal interface: gamer do something, then server give exp to gamer's role.
func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp

	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		if config.PlayerExp == 0 { // level max to 60
			break
		}
		// finish task? if finished, continue run; else break
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			break
		}

		// if finish the task, then do
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}

	fmt.Println("now level:", self.PlayerLevel, "---now exp:", self.PlayerExp)
}
