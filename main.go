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

	player := game.NewTestPlayer()
	player.RecvSetIcon(1)

	// each 1s touch once
	triker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-triker.C:
			if time.Now().Unix()%3 == 0 {
				player.RecvSetName("专业代练")
			} else if time.Now().Unix()%5 == 0 {
				player.RecvSetName("正常名字")
			}
		}
	}

}
