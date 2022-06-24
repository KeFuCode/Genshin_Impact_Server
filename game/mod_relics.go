package game

import (
	"fmt"
	"server/bin/csvs"
)

type Relics struct {
	RelicsId int
	KeyId    int // Relics unique_id
}

type ModRelics struct {
	RelicsInfo map[int]*Relics
	MaxKey     int
}

func (self *ModRelics) AddItem(relicsId int, num int64) {
	config := csvs.GetRelicsConfig(relicsId)
	if config == nil {
		fmt.Println("relics config isn't exist")
		return
	}

	if len(self.RelicsInfo)+int(num) > csvs.RELICS_MAX_COUNT {
		fmt.Println("relics overflow max count")
		return
	}

	for i := int64(0); i < num; i++ {
		relics := new(Relics)
		relics.RelicsId = relicsId
		self.MaxKey++
		relics.KeyId = self.MaxKey
		self.RelicsInfo[relics.KeyId] = relics
		fmt.Println("get Relics:", csvs.GetItemName(relicsId), "--- Relics id:", relics.KeyId)
	}
}
