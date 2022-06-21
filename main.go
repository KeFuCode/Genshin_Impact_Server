package main

import (
	"fmt"
	"server/bin/csvs"
	"server/game"
	"time"
)

func main() {
	//********************
	// init: load server config
	csvs.CheckLoadCsv()

	fmt.Printf("Data Test ---- start\n")

	// each 10s touch once
	go game.GetManageBanWord().Run()

	playerGM := game.NewTestPlayer()

	// each 1s touch once
	triker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-triker.C:
			playerGM.ModPlayer.AddExp(5000)
		}
	}

}
