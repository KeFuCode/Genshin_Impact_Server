package game

import (
	"fmt"
	"math/rand"
	"server/bin/csvs"
	"sync"
	"time"
)

type Server struct {
	Wait sync.WaitGroup
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
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
