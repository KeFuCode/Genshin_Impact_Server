package main

import (
	"fmt"
	"server/game"
)

func main() {
	fmt.Printf("Data Test ---- start\n")

	player := game.NewTestPlayer()

	player.RecvSetName("好人")
	player.RecvSetName("坏人")
	player.RecvSetName("求外挂带")
	player.RecvSetName("好玩")
	player.RecvSetName("感觉不如原神...画质")
}
