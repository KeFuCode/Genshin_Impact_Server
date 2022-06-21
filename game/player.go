package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
}

func NewTestPlayer() *Player {
	// 模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)

	//*******************************
	// 模块数据初始化
	player.ModPlayer.Icon = 0
	//*******************************

	return player
}
