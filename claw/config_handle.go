package claw

import (
	"encoding/json"
	"github.com/chenyingzhou/octopus-go/consts"
	"github.com/chenyingzhou/octopus-go/plate"
	"log"
	"strconv"
)

func (cfg *Config) Handle(food plate.Food) {
	rootSourceTree := cfg.SourceTree
	sourceTree := rootSourceTree.SearchNode(food.DataSource, food.DataSet, 0)
	// 简化处理，仅处理主键
	ids := make([]string, 0)
	for _, id := range food.Ids {
		ids = append(ids, strconv.Itoa(int(id)))
	}
	for _, row := range food.Rows {
		id, ok := row[sourceTree.IdColumn]
		if ok {
			ids = append(ids, id)
		}
	}
	sourceFilterValues := make(map[string][]string)
	sourceFilterValues["id"] = ids
	// 向上查询直到根结点
	sourceFilter := SourceFilter{
		Type:            consts.SourceRelationTypeAnd,
		Values:          sourceFilterValues,
		ExtraConditions: food.Conditions,
		ExtraColumns:    make([]string, 0),
	}
	for rootSourceTree.GetKey() != sourceTree.GetKey() {
		parent := sourceTree.SearchParent(rootSourceTree)
		for _, relation := range parent.Relations {
			if relation.SourceTree.GetKey() == sourceTree.GetKey() {
				for _, field := range relation.Fields {
					sourceFilter.ExtraColumns = append(sourceFilter.ExtraColumns, field.Target)
				}
			}
		}
		list := sourceTree.Fetch(sourceFilter, false)[sourceTree.GetKey()]
		sourceFilter = sourceTree.matchParentSourceFilters(*list, parent)
		sourceTree = parent
	}

	data := rootSourceTree.Fetch(sourceFilter, true)
	text, _ := json.Marshal(data)
	log.Fatalln(string(text))
}
