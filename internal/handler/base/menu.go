package base

import (
	"encoding/json"
	"log"
)

type MenuItem struct {
	Icon       string      `json:"icon"`
	Name       string      `json:"name"`
	Controller string      `json:"controller"`
	Action     string      `json:"action"`
	Sonlist    []*MenuItem `json:"sonlist"`
}

var (
	menu     []*MenuItem
	menuJson = `
		[
			{
				"name": "玩家信息",
				"icon": "",
				"sonlist": [
					{
						"name": 		"基础信息",
						"icon": 		"",
						"Controller": 	"player",
						"Action": 		"basic"
					},
					{
						"name": 		"任务",
						"icon": 		"",
						"Controller": 	"player",
						"Action": 		"task"
					},
					{
						"name": 		"资源",
						"icon": 		"",
						"Controller": 	"player",
						"Action": 		"resource"
					}
				]
			}
		]
	`
)

func init() {
	menu = make([]*MenuItem, 0)
	err := json.Unmarshal([]byte(menuJson), &menu)
	if err != nil {
		log.Fatalf("menu:init: json.Unmarshal error, err=%v.", err)
	}

}
