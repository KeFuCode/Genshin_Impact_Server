package main

import (
	"fmt"
	"server/game"
)

func main() {
	fmt.Printf("Data Test ---- start\n")

	player := game.NewTestPlayer()

	player.RecvSetIcon(1) // HuTao
	player.RecvSetIcon(2) // WenDi
	player.RecvSetIcon(3) // ZhongLi

}
