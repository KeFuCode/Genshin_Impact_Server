package game

import (
	"fmt"
	"server/bin/csvs"
)

type PoolInfo struct {
	PoolId int
}

type ModPool struct {
	UpPoolInfo map[int]*PoolInfo
}

func (self *ModPool) DoPool() {
	result := make(map[int]int)
	for i := 0; i < 10000000; i++ {
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}
		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			result[roleIdConfig.Result]++
		}
	}

	for k, v := range result {
		fmt.Printf("get %s times: %d\n", csvs.GetItemName(k), v)
	}
}
