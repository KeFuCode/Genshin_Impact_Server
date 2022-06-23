package main

import (
	"fmt"
	"server/bin/csvs"
	"server/game"
	_ "time"
)

func main() {
	//********************
	// init: load server config
	csvs.CheckLoadCsv()
	// each 10s touch once
	go game.GetManageBanWord().Run()

	fmt.Printf("Data Test ---- start\n")

	playerTest := game.NewTestPlayer()
	playerTest.ModPlayer.SetIcon(3000001, playerTest)
	playerTest.ModBag.AddItem(3000001, playerTest)
	playerTest.ModPlayer.SetIcon(3000001, playerTest)

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
