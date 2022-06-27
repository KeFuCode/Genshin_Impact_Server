package game

import (
	"fmt"
	"server/bin/csvs"
)

type PoolInfo struct {
	PoolId        int
	FiveStarTimes int
	FourStarTimes int
	IsMustUp      int
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

					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("data error")
								return
							}
							self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
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

func (self *ModPool) HandleUpPoolTen(player *Player) {
	for i := 0; i < 10; i++ {
		self.AddTimes()
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
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
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
					self.UpPoolInfo.FiveStarTimes = 0
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
					}
				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
				}
			}
			// fmt.Println(fmt.Sprintf("第%d抽抽中:%s", i+1, csvs.GetItemName(roleIdConfig.Result)))
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)
		}
	}
	if self.UpPoolInfo.IsMustUp == csvs.LOGIC_FALSE {
		fmt.Println(fmt.Sprintf("当前处于小保底区间！"))
	} else {
		fmt.Println(fmt.Sprintf("当前处于大保底区间！"))
	}
	fmt.Println(fmt.Sprintf("当前累计未出5星次数：%d", self.UpPoolInfo.FiveStarTimes))
	fmt.Println(fmt.Sprintf("当前累计未出4星次数：%d", self.UpPoolInfo.FourStarTimes))

}

func (self *ModPool) HandleUpPoolSingle(times int, player *Player) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	result := make(map[int]int)
	fourNum := 0
	fiveNum := 0
	for i := 0; i < times; i++ {
		self.AddTimes()
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
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
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
					self.UpPoolInfo.FiveStarTimes = 0
					fiveNum++
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
					}
				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
					fourNum++
				}
			}
			result[roleIdConfig.Result]++
			player.ModBag.AddItem(roleIdConfig.Result, 1, player)

		}
	}

	for k, v := range result {
		fmt.Println(fmt.Sprintf("抽中%s次数：%d", csvs.GetItemName(k), v))
	}
	fmt.Println(fmt.Sprintf("抽中4星：%d", fourNum))
	fmt.Println(fmt.Sprintf("抽中5星：%d", fiveNum))
}

func (self *ModPool) HandleUpPoolTimesTest(times int) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	resultEach := make(map[int]int)
	for i := 0; i < times; i++ {
		self.AddTimes()
		dropGroup := csvs.ConfigDropGroupMap[1000]
		if dropGroup == nil {
			return
		}

		if self.UpPoolInfo.FiveStarTimes > csvs.FIVE_STAR_TIMES_LIMIT || self.UpPoolInfo.FourStarTimes > csvs.FOUR_STAR_TIMES_LIMIT {
			newDropGroup := new(csvs.DropGroup)
			newDropGroup.DropId = dropGroup.DropId
			newDropGroup.WeightAll = dropGroup.WeightAll
			addFiveWeight := (self.UpPoolInfo.FiveStarTimes + 1 - csvs.FIVE_STAR_TIMES_LIMIT) * csvs.FIVE_STAR_TIMES_LIMIT_EACH_VALUE
			if addFiveWeight < 0 {
				addFiveWeight = 0
			}
			addFourWeight := (self.UpPoolInfo.FourStarTimes + 1 - csvs.FOUR_STAR_TIMES_LIMIT) * csvs.FOUR_STAR_TIMES_LIMIT_EACH_VALUE
			if addFourWeight < 0 {
				addFourWeight = 0
			}

			for _, config := range dropGroup.DropConfigs {
				newConfig := new(csvs.ConfigDrop)
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
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
					resultEach[self.UpPoolInfo.FiveStarTimes]++
					self.UpPoolInfo.FiveStarTimes = 0
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
					}
				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
				}
			}
		}
	}

	for k, v := range resultEach {
		fmt.Println(fmt.Sprintf("第%d抽抽出5星的次数：%d", k, v))
	}
}

func (self *ModPool) HandleUpPoolFiveTest(times int) {
	if times <= 0 || times > 100000000 {
		fmt.Println("请输入正确的数值(1~100000000)")
		return
	} else {
		fmt.Println(fmt.Sprintf("累计抽取%d次,结果如下:", times))
	}
	resultEachTest := make(map[int]int)
	fiveTest := 0
	for i := 0; i < times; i++ {
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
				newConfig.Result = config.Result
				newConfig.DropId = config.DropId
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
					self.UpPoolInfo.FiveStarTimes = 0
					if self.UpPoolInfo.IsMustUp == csvs.LOGIC_TRUE {
						dropGroup := csvs.ConfigDropGroupMap[100012]
						if dropGroup != nil {
							roleIdConfig = csvs.GetRandDropNew(dropGroup)
							if roleIdConfig == nil {
								fmt.Println("数据异常")
								return
							}
						}
					}
					if roleIdConfig.DropId == 100012 {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_FALSE
					} else {
						self.UpPoolInfo.IsMustUp = csvs.LOGIC_TRUE
					}
				} else if roleConfig.Star == 4 {
					self.UpPoolInfo.FourStarTimes = 0
				}
			}
		}
		if i%10 == 9 {
			resultEachTest[fiveTest]++
		}
	}

	for k, v := range resultEachTest {
		fmt.Println(fmt.Sprintf("10连%d黄次数：%d", k, v))
	}
}