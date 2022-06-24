package game

import (
	"fmt"
	_ "sync"
)

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
	ModRole       *ModRole
	ModBag        *ModBag
}

func NewTestPlayer() *Player {
	// mod init: if not, player_instance is nil ptr, then this goroutine will break down
	// server is normal, gamer need reopen game
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)
	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	// player.ModUniqueTask.Locker = new(sync.RWMutex)
	player.ModRole = new(ModRole)
	player.ModRole.RoleInfo = make(map[int]*RoleInfo)
	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)

	//*******************************
	// 模块数据初始化
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
func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}

// recover world level
func (self *Player) ReturnWorldLevel() {
	self.ModPlayer.ReturnWorldLevel(self)
}

// set birthday: month * 100 + day
func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self)
}

// set show card
func (self *Player) SetShowCard(showCard []int) {
	self.ModPlayer.SetShowCard(showCard, self)
}

// set show team
func (self *Player) SetShowTeam(showRole []int) {
	self.ModPlayer.SetShowTeam(showRole, self)
}

// hide show team
func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}

// internal function
func (self *Player) Run() {

	for {
		fmt.Println("please input operate: 1 --- add item")
		var action int
		fmt.Scan(&action)
		fmt.Println(action)
		switch action {
		case 1:
			self.HandleAddItem()
		}
	}

	// triker := time.NewTicker(time.Second * 1)

	// for {
	// 	select {
	// 	case <-triker.C:
	// 		if time.Now().Unix()%5 == 0 {
	// 			self.ModRole.AddItem(2000017, 7)
	// 		}
	// 	}
	// }
}

func (self *Player) HandleAddItem() {
	itemId, itemNum := 0, 0
	fmt.Println("input itemId")
	fmt.Scan(&itemId)
	fmt.Println("input itemNum")
	fmt.Scan(&itemNum)
	self.ModBag.AddItem(itemId, int64(itemNum), self)
}
