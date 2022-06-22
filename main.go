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
	// each 10s touch once
	go game.GetManageBanWord().Run()

	fmt.Printf("Data Test ---- start\n")

	playerGM := game.NewTestPlayer()

	// each 1s touch once
	triker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-triker.C:
			if time.Now().Unix()%3 == 0 {
				playerGM.ReduceWorldLevel()
			} else if time.Now().Unix()%5 == 0 {
				playerGM.ReturnWorldLevel()
			}
		}
	}

	return
}
