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
	playerTest.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerTest.ModBag.AddItemToBag(1000003, 500)
	playerTest.ModBag.AddItemToBag(1000003, 800)
	playerTest.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerTest.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerTest.ModBag.RemoveItemToBagGM(1000003, 1000)
	playerTest.ModBag.RemoveItemToBagGM(1000003, 1000)

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
