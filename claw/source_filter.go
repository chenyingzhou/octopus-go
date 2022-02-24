package claw

type SourceFilter struct {
	Type   string              //字段之间关系(AND/OR)
	Values map[string][]string //字段的值，e.g. {"column1": ["value1", "value2"], "column2": ["value3"]}
}
