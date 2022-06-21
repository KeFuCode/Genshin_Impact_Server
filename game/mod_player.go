package game

import "fmt"

type ModPlayer struct {
	// 可见字段
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelCool int64
	ShowTeam       []int
	ShowCard       int

	// 隐藏字段
	IsProhibit int
	IsGm       int
}

// internal interface: set ModPlayer inner value
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
	fmt.Println("Now Name: ", player.ModPlayer.Sign)
}
