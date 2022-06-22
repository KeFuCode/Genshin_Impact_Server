package game

import "sync"

const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1
	TASK_STATE_FINASH = 2
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
}

func NewTestPlayer() *Player {
	// mod init: if not, player_instance is nil ptr, then this goroutine will break down
	// server is normal, gamer need reopen game
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	player.ModUniqueTask.Locker = new(sync.RWMutex)

	//*******************************
	// 模块数据初始化
	player.ModPlayer.Icon = 0
	player.ModPlayer.Card = 0
	player.ModPlayer.PlayerLevel = 1 // init level is 1
	player.ModPlayer.WorldLevel = 6
	player.ModPlayer.WorldLevelNow = 6

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

// reduce 1 world level 
func(self *Player) ReduceWorldLevel()  {
	self.ModPlayer.ReduceWorldLevel(self)
}
// recover world level
func(self *Player) ReturnWorldLevel()  {
	self.ModPlayer.ReturnWorldLevel(self)
}
