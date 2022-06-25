package csvs

import "fmt"

type DropGroup struct {
	DropId      int
	WeightAll   int
	DropConfigs []*ConfigDrop
}

var ConfigDropGroupMap map[int]*DropGroup

func CheckLoadCsv() {
	MakeDropGroupMap()

	fmt.Println("csv init end")
}

func MakeDropGroupMap() {
	configDropGroupMap := make(map[int]*DropGroup)
	for _, v := range ConfigDropSlice {
		dropGroup, ok := configDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			configDropGroupMap[v.DropId] = dropGroup
		}
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
}