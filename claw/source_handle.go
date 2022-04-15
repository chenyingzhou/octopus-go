package claw

import (
	"fmt"
	"github.com/chenyingzhou/octopus-go/consts"
	"github.com/chenyingzhou/octopus-go/datasource"
	"log"
	"strings"
)

func (st *SourceTree) GetKey() string {
	return fmt.Sprintf("%s_%s_%d", st.DataSource, st.DataSet, st.Priority)
}

func (st *SourceTree) Fetch(sf SourceFilter, data *map[string][]map[string]string) {
	conditions := make([]string, 0)
	for col, vals := range sf.Values {
		if len(vals) == 0 {
			continue
		}
		conditions = append(conditions, fmt.Sprintf("`%s` IN (%s)", col, "'"+strings.Join(vals, "','")+"'"))
	}
	if len(conditions) == 0 {
		return
	}
	where := strings.Join(conditions, string(sf.Type))
	if st.ExtraCondition != "" {
		where = "(" + st.ExtraCondition + ")" + " AND " + "(" + where + ")"
	}

	db, err := datasource.GetMysqlClient(st.DataSource)
	if err != nil {
		log.Printf("Connect to database fail, database: " + st.DataSource)
		return
	}

	columns := make([]string, 0)
	if st.IdColumn != "" {
		columns = append(columns, st.IdColumn)
	}
	if st.TimeColumn != "" {
		columns = append(columns, st.TimeColumn)
	}
	if st.DeleteColumn != "" {
		columns = append(columns, st.DeleteColumn)
	}
	for _, field := range st.Fields {
		columns = append(columns, field.Column)
	}
	for _, relation := range st.Relations {
		for _, field := range relation.Fields {
			columns = append(columns, field.Column)
		}
	}
	if sf.ExtraColumns != nil {
		columns = append(columns, sf.ExtraColumns...)
	}

	query := fmt.Sprintf("SELECT %s FROM `%s` WHERE %s", "`"+strings.Join(columns, "`,`")+"`", st.DataSet, where)
	list := datasource.QueryForMapSlice(db, query)
	(*data)[st.GetKey()] = list

	// 处理子节点的数据
	relationSfMap := st.matchSourceFilters(list)
	for _, relation := range st.Relations {
		relationSf, ok := relationSfMap[relation.SourceTree.GetKey()]
		if ok {
			relation.SourceTree.Fetch(relationSf, data)
		}
	}
}

func (st *SourceTree) matchSourceFilters(rows []map[string]string) map[string]SourceFilter {
	sfMap := make(map[string]SourceFilter)
	for _, relation := range st.Relations {
		valuesMap := make(map[string][]string)
		extraColumns := make([]string, 0)
		for _, field := range relation.Fields {
			target := field.Target
			column := field.Column
			values := make([]string, 0)
			for _, row := range rows {
				value, ok := row[column]
				if ok {
					values = append(values, value)
				}
			}
			valuesMap[target] = values
			extraColumns = append(extraColumns, target)
		}
		sfMap[relation.SourceTree.GetKey()] = SourceFilter{
			Type:         consts.SourceRelationTypeAnd,
			Values:       valuesMap,
			ExtraColumns: extraColumns,
		}
	}
	return sfMap
}
