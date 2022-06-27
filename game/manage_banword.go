package game

import (
	"fmt"
	"regexp"
	"server/bin/csvs"
	"time"
)

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
	Test         map[int]int
	MsgChannel   chan int
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"原神"}
		manageBanWord.Test = make(map[int]int)
		manageBanWord.MsgChannel = make(chan int)
	}

	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}

	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}

	return false
}

func (self *ManageBanWord) Run() {
	GetServer().AddGo()
	// load base word library
	self.BanWordBase = csvs.GetBanWordBase()

	triker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-triker.C: // triker will touch every 1s
			if time.Now().Unix()%10 == 0 {
				fmt.Println("update word library")
				GetServer().UpdateBanWord(self.BanWordBase)
			} else {

			}
		case _, ok := <-self.MsgChannel:
			if !ok {
				GetServer().GoDone()
				return
			}
		}
	}
}

func (self *ManageBanWord) Close() {
	close(self.MsgChannel)
}
