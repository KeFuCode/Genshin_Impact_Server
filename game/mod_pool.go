package game

import (
	"fmt"
	"server/bin/csvs"
)

type PoolInfo struct {
	PoolId        int
	FiveStarTimes int
}

type ModPool struct {
	UpPoolInfo *PoolInfo
}

func (self *ModPool) DoUpPool() {
	result := make(map[int]int)
	for i := 0; i < 10000000; i++ {
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if self.UpPoolInfo.FiveStarTimes > csvs.FIVE_STAR_TIMES_LIMIT {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (self.UpPoolInfo.FiveStarTimes - csvs.FIVE_STAR_TIMES_LIMIT) * csvs.FIVE_STAR_TIMES_LIMIT_EACH_VALUE

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.DropId = config.DropId
				newConfig.Result = config.Result
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight
				} else {
					newConfig.Weight = config.Weight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}

		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil && roleConfig.Star == 5 {
				self.UpPoolInfo.FiveStarTimes = 0
			} else {
				self.UpPoolInfo.FiveStarTimes++
			}
			result[roleIdConfig.Result]++
		}
	}

	for k, v := range result {
		fmt.Printf("get %s times: %d\n", csvs.GetItemName(k), v)
	}
}
