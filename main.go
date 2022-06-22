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
	playerGM.ModPlayer.SetShowTeam([]int{1001, 1001, 1001, 1002, 1001, 1005, 1001, 1001, 1001, 1002, 1001, 1005}, playerGM)

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
