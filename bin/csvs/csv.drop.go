package csvs

import (
	"fmt"
	"server/utils"
)

type ConfigDrop struct {
	DropId int `json:"DropId"`
	Weight int `json:"Weight"`
	Result int `json:"Result"`
	IsEnd  int `json:"IsEnd"`
}

var ConfigDropSlice []*ConfigDrop

func init() {
	ConfigDropSlice = make([]*ConfigDrop, 0)
	utils.GetCsvUtilMgr().LoadCsv("Drop", &ConfigDropSlice)

	fmt.Println("init csv_drop")
}
