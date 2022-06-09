package main

import (
	"github.com/chenyingzhou/octopus-go/claw"
	"github.com/chenyingzhou/octopus-go/plate"
)

func main() {
	food := plate.Food{
		SourceType: "MYSQL",
		DataSource: "octopus_b",
		DataSet:    "x",
		Ids:        []int32{1},
		Rows:       nil,
		Conditions: nil,
	}
	c := claw.Container[1]
	c.Handle(food)
}
