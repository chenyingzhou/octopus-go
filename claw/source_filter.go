package claw

import "github.com/chenyingzhou/octopus-go/consts"

type SourceFilter struct {
	Type            consts.SourceRelationType //字段之间关系(AND/OR)
	Values          map[string][]string       //字段的值，e.g. {"column1": ["value1", "value2"], "column2": ["value3"]}
	ExtraConditions []string                  // 额外条件
	ExtraColumns    []string                  //需要查询的额外列，一般用于分组
}
