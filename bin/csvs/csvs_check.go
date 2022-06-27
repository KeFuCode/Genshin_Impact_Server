package csvs

import (
	"fmt"
	"math/rand"
)

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
	ConfigDropGroupMap = make(map[int]*DropGroup)
	for _, v := range ConfigDropSlice {
		dropGroup, ok := ConfigDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			ConfigDropGroupMap[v.DropId] = dropGroup
		}
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}

	// RandDropTest()
}

func RandDropTest() {
	dropGroup := ConfigDropGroupMap[1000]
	if dropGroup == nil {
		return
	}

	num := 0
	for {
		config := GetRandDropNew(dropGroup)
		if config.IsEnd == LOGIC_TRUE {
			fmt.Println(GetItemName(config.Result))
			num++
			dropGroup = ConfigDropGroupMap[1000]
			if num >= 100 {
				break
			} else {
				continue
			}
		}
		dropGroup = ConfigDropGroupMap[config.Result]
		if dropGroup == nil {
			break
		}
	}
}

func GetRandDrop(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			return v
		}
	}

	return nil
}

func GetRandDropNew(dropGroup *DropGroup) *ConfigDrop {
	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew(dropGroup)
		}
	}

	return nil
}

func GetRandDropNew1(dropGroup *DropGroup, fiveInfo map[int]int, fourInfo map[int]int) *ConfigDrop {
	for _, v := range dropGroup.DropConfigs {
		_, ok := fiveInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fiveInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fiveInfo[config.Result] {
					maxGetTime = fiveInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}

		_, ok = fourInfo[v.Result]
		if ok {
			index := 0
			maxGetTime := 0
			for k, config := range dropGroup.DropConfigs {
				_, nowOK := fourInfo[config.Result]
				if !nowOK {
					continue
				}
				if maxGetTime < fourInfo[config.Result] {
					maxGetTime = fourInfo[config.Result]
					index = k
				}
			}
			return dropGroup.DropConfigs[index]
		}
	}

	randNum := rand.Intn(dropGroup.WeightAll)
	randNow := 0
	for _, v := range dropGroup.DropConfigs {
		randNow += v.Weight
		if randNum < randNow {
			if v.IsEnd == LOGIC_TRUE {
				return v
			}
			dropGroup := ConfigDropGroupMap[v.Result]
			if dropGroup == nil {
				return nil
			}
			return GetRandDropNew1(dropGroup, fiveInfo, fourInfo)
		}
	}
	return nil
}