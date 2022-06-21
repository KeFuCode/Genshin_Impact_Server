package main

import (
	"fmt"
	"server/game"
)

func main() {
	fmt.Printf("Data Test ---- start\n")

	game.GetManageBanWord().Run()

	// player := game.NewTestPlayer()

}
