package main

import (
	"fmt"
	"server/game"
	"time"
)

func main() {
	//********************
	// init: load server config

	fmt.Printf("Data Test ---- start\n")

	// each 10s touch once
	go game.GetManageBanWord().Run()

	player := game.NewTestPlayer()
	player.RecvSetIcon(1)

	// each 3s touch once
	trikerIn := time.NewTicker(time.Second * 3)
	// each 5s touch once
	trikerOut := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-trikerIn.C:
			player.RecvSetIcon(int(time.Now().Unix()))
		case <-trikerOut.C:
			player.RecvSetName("test~test")
		}
	}

}
