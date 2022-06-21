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

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端，操作非法
		return
	}

	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", player.ModPlayer.Icon)
}
