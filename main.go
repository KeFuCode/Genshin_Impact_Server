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

	// each 10s touch once
	triker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-triker.C:
			playerTest := game.NewTestPlayer()
			go playerTest.Run()
		}
	}

	return
}
