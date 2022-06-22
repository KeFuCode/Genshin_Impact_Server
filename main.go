package main

import (
	"fmt"
	"server/bin/csvs"
	"server/game"
)

func main() {
	//********************
	// init: load server config
	csvs.CheckLoadCsv()
	// each 10s touch once
	go game.GetManageBanWord().Run()

	fmt.Printf("Data Test ---- start\n")

	playerGM := game.NewTestPlayer()
	playerGM.ModPlayer.SetBirth(3001, playerGM)
	playerGM.ModPlayer.SetBirth(1235, playerGM)
	playerGM.ModPlayer.SetBirth(10, playerGM)
	playerGM.ModPlayer.SetBirth(622, playerGM)
	playerGM.ModPlayer.SetBirth(520, playerGM)

	// each 1s touch once
	/* 	triker := time.NewTicker(time.Second * 1)
	   	for {
	   		select {
	   		case <-triker.C:
	   			if time.Now().Unix()%3 == 0 {
	   				playerGM.ReduceWorldLevel()
	   			} else if time.Now().Unix()%5 == 0 {
	   				playerGM.ReturnWorldLevel()
	   			}
	   		}
	   	} */

	for {
		
	}
	return
}
