package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"server/bin/csvs"
	"sync"
	"syscall"
	"time"
)

type DBConfig struct {
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
}

type ServerConfig struct {
	ServerId int       `json:"serverid"`
	Host     string    `json:"host"`
	DBConfig *DBConfig `json:"database"`
}

type Server struct {
	Wait        sync.WaitGroup
	BanWordBase []string
	Lock        *sync.RWMutex

	Config *ServerConfig
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
		server.Lock = new(sync.RWMutex)
	}
	return server
}

func (self *Server) Start() {
	//********************
	rand.Seed(time.Now().Unix())
	// init: load server config
	csvs.CheckLoadCsv()
	// each 10s touch once
	go GetManageBanWord().Run()

	fmt.Printf("Data Test ---- start\n")

	playerTest := NewTestPlayer()
	go playerTest.Run()
	go self.SignalHandle()
	self.LoadConfig()

	// each 10s touch once
	// triker := time.NewTicker(time.Second * 10)
	// for {
	// 	select {
	// 	case <-triker.C:
	// 		playerTest := game.NewTestPlayer()
	// 		go playerTest.Run()
	// 	}
	// }

	self.Wait.Wait()
	fmt.Println("server closed success!")
}

func (self *Server) Close() {
	GetManageBanWord().Close()
}

func (self *Server) AddGo() {
	self.Wait.Add(1)
}

func (self *Server) GoDone() {
	self.Wait.Done()
}

func (self *Server) IsBanWord(txt string) bool {
	self.Lock.RLock()
	defer self.Lock.RUnlock()
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}

	return false
}

func (self *Server) UpdateBanWord(banWord []string) {
	self.Lock.Lock()
	defer self.Lock.Unlock()
	self.BanWordBase = banWord
}

func (self *Server) SignalHandle() {
	channelSignal := make(chan os.Signal)
	signal.Notify(channelSignal, syscall.SIGINT)

	for {
		select {
		case <-channelSignal:
			fmt.Println("get syscall.SIGINT")
			self.Close()
		}
	}
}

func (self *Server) LoadConfig() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(configFile, &self.Config)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}