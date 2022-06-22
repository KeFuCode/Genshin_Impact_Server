package main

import (
	"fmt"
	"server/bin/csvs"
	"server/game"
	_"time"
)

func main() {
	//********************
	// init: load server config
	csvs.CheckLoadCsv()
	// each 10s touch once
	go game.GetManageBanWord().Run()

	fmt.Printf("Data Test ---- start\n")

	playerTest := game.NewTestPlayer()
	playerTest.ModBag.AddItem(1000003)
	playerTest.ModBag.AddItem(1000006)
	playerTest.ModBag.AddItem(1000008)
	playerTest.ModBag.AddItem(2000002)
	playerTest.ModBag.AddItem(2000021)
	playerTest.ModBag.AddItem(2000088)
	playerTest.ModBag.AddItem(3000004)
	playerTest.ModBag.AddItem(4000025)

	// each 10s touch once
	// triker := time.NewTicker(time.Second * 10)
	// for {
	// 	select {
	// 	case <-triker.C:
	// 		playerTest := game.NewTestPlayer()
	// 		go playerTest.Run()
	// 	}
	// }

	return
}
