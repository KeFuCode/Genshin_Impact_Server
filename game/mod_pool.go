package game

import (
	"fmt"
	"server/bin/csvs"
)

type PoolInfo struct {
	PoolId        int
	FiveStarTimes int
	FourStarTimes int
	isMustUp      int
}

type ModPool struct {
	UpPoolInfo *PoolInfo
}

func (self *ModPool) AddTimes() {
	self.UpPoolInfo.FiveStarTimes++
	self.UpPoolInfo.FourStarTimes++
}

func (self *ModPool) DoUpPool() {
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	resultEach := make(map[int]int)
	resultEachTest := make(map[int]int)
	fiveTest := 0
	for i := 0; i < 100000000; i++ {
		self.AddTimes()
		if i%10 == 0 {
			fiveTest = 0
		}
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if self.UpPoolInfo.FiveStarTimes > csvs.FIVE_STAR_TIMES_LIMIT || self.UpPoolInfo.FourStarTimes > csvs.FOUR_STAR_TIMES_LIMIT {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (self.UpPoolInfo.FiveStarTimes - csvs.FIVE_STAR_TIMES_LIMIT) * csvs.FIVE_STAR_TIMES_LIMIT_EACH_VALUE
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (self.UpPoolInfo.FourStarTimes - csvs.FOUR_STAR_TIMES_LIMIT) * csvs.FOUR_STAR_TIMES_LIMIT_EACH_VALUE
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.DropId = config.DropId
				newConfig.Result = config.Result
				newConfig.IsEnd = config.IsEnd
				if config.Result == 10001 {
					newConfig.Weight = config.Weight + addFiveWeight
				} else if config.Result == 10002 {
					newConfig.Weight = config.Weight + addFourWeight
				} else if config.Result == 10003 {
					newConfig.Weight = config.Weight - addFiveWeight - addFourWeight
				}
				newDropGroup.DropConfigs = append(newDropGroup.DropConfigs, newConfig)
			}
			dropGroup = newDropGroup
		}

		roleIdConfig := csvs.GetRandDropNew(dropGroup)
		if roleIdConfig != nil {
			roleConfig := csvs.GetRoleConfig(roleIdConfig.Result)
			if roleConfig != nil {
				if roleConfig.Star == 5 {
					fiveTest++
					resultEach[self.UpPoolInfo.FiveStarTimes]++
					self.UpPoolInfo.FiveStarTimes = 0
					fiveNum++

					if self.UpPoolInfo.isMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("data error")
								return
							}
							self.UpPoolInfo.isMustUp = csvs.LOGIC_FALSE
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.isMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.isMustUp = csvs.LOGIC_TRUE
					}

				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
		}
		if i%10 == 9 {
			resultEachTest[fiveTest]++
		}
	}

	for k, v := range result {
		fmt.Printf("get %s times: %d\n", csvs.GetItemName(k), v)
	}
	fmt.Printf("get 4 star %d times\n", fourNum)
	fmt.Printf("get 5 star %d times\n", fiveNum)

	for k, v := range resultEach {
		fmt.Printf("no.%d times get 5 star times: %d\n", k, v)
	}

	for k, v := range resultEachTest {
		fmt.Printf("10_series get %d 5_star times: %d\n", k, v)
	}
}
