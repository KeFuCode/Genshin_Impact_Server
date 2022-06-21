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

	playerGM := game.NewTestPlayer()
	playerGM.ModPlayer.AddExp(10000000, playerGM)

	/* 	// each 1s touch once
	   	triker := time.NewTicker(time.Second * 3)
	   	for {
	   		select {
	   		case <-triker.C:
	   			playerGM.ModPlayer.AddExp(5000)
	   		}
	   	} */

	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)

	for {
		//
	}

	return
}

func playerLoadConfig(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		config := csvs.ConfigUniqueTaskMap[10001]
		if config != nil {
			fmt.Println(config.TaskId)
		}
	}
}
