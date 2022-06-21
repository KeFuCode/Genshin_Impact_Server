package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

func NewTestPlayer() *Player {
	// mod init: if not, player_instance is nil ptr, then this goroutine will break down
	// server is normal, gamer need reopen game
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)

	//*******************************
	// 模块数据初始化
	player.ModPlayer.Icon = 0
	player.ModPlayer.Card = 0
	//*******************************

	return player
}

// external interface: receive client request
func (self *Player) RecvSetIcon(iconId int) {
	// debug site
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(cardId int) {
	// debug site
	self.ModPlayer.SetCard(cardId, self)
}

func (self *Player) RecvSetName(name string) {
	if GetManageBanWord().IsBanWord(name) {
		return
	}

	// debug site
	self.ModPlayer.RecvSetName(name, self)
}

func (self *Player) RecvSetSign(sign string) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}

	// debug site
	self.ModPlayer.RecvSetSign(sign, self)
}
