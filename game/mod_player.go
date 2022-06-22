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
	ShowCard       int

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
