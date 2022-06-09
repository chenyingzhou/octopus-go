package claw

import (
	"fmt"
	"github.com/chenyingzhou/octopus-go/consts"
	"github.com/chenyingzhou/octopus-go/datasource"
	"log"
	"math"
	"strconv"
	"strings"
)

func (st *SourceTree) GetKey() string {
	return fmt.Sprintf("%s_%s_%d", st.DataSource, st.DataSet, st.Priority)
}

func (st *SourceTree) SearchNode(dataSource string, dataSet string, priority int32) *SourceTree {
	if st.GetKey() == fmt.Sprintf("%s_%s_%d", dataSource, dataSet, priority) {
		return st
	}
	for _, relation := range st.Relations {
		node := relation.SourceTree.SearchNode(dataSource, dataSet, priority)
		if node != nil {
			return node
		}
	}
	return nil
}

func (st *SourceTree) SearchParent(root *SourceTree) *SourceTree {
	for _, relation := range root.Relations {
		if relation.SourceTree.GetKey() == st.GetKey() {
			return root
		}
	}
	for _, relation := range root.Relations {
		parent := st.SearchParent(&relation.SourceTree)
		if parent != nil {
			return parent
		}
	}
	return nil
}

func (st *SourceTree) Fetch(sf SourceFilter, withRelation bool) map[string][]map[string]string {
	result := make(map[string][]map[string]string)
	emptyList := make([]map[string]string, 0)
	result[st.GetKey()] = emptyList
	conditions := make([]string, 0)
	for col, values := range sf.Values {
		if len(values) == 0 {
			continue
		}
		conditions = append(conditions, fmt.Sprintf("`%s` IN (%s)", col, "'"+strings.Join(values, "','")+"'"))
	}
	if len(conditions) == 0 {
		return result
	}
	where := strings.Join(conditions, string(sf.Type))
	wheres := make([]string, 0)
	wheres = append(wheres, where)
	if st.ExtraCondition != "" {
		wheres = append(wheres, st.ExtraCondition)
	}
	for _, condition := range sf.ExtraConditions {
		wheres = append(wheres, condition)
	}
	where = "(" + strings.Join(wheres, ") AND (") + ")"

	db, err := datasource.GetMysqlClient(st.DataSource)
	if err != nil {
		log.Printf("Connect to database fail, database: " + st.DataSource)
		return result
	}

	columns := make([]string, 0)
	if st.IdColumn != "" {
		columns = append(columns, st.IdColumn)
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
	result[st.GetKey()] = list

	// 处理子节点的数据
	if withRelation {
		relationSfMap := st.matchChildSourceFilters(list)
		for _, relation := range st.Relations {
			relationSf, ok := relationSfMap[relation.SourceTree.GetKey()]
			if ok {
				childResult := relation.SourceTree.Fetch(relationSf, true)
				for k, list := range childResult {
					result[k] = list
				}
			}
		}
	}
	return result
}

func (st *SourceTree) grouping(data map[string][]map[string]string, stRow map[string]string) map[string][]map[string]string {
	result := make(map[string][]map[string]string)
	singleList := make([]map[string]string, 0)
	singleList = append(singleList, stRow)
	result[st.GetKey()] = singleList
	for _, relation := range st.Relations {
		childKey := relation.SourceTree.GetKey()
		childRows, ok := data[childKey]
		if !ok {
			continue
		}
		for _, childRow := range childRows {
			valid := true
			for _, field := range relation.Fields {
				columnVal, ok1 := stRow[field.Column]
				targetVal, ok2 := childRow[field.Target]
				valid = valid && ok1 && ok2 && (columnVal == targetVal)
			}
			if valid {
				childResult := relation.SourceTree.grouping(data, childRow)
				for key, rows := range childResult {
					_, ok3 := result[key]
					if ok3 {
						result[key] = append(result[key], rows...)
					} else {
						result[key] = rows
					}
				}
			}
		}
	}
	return result
}

func (st *SourceTree) toDocument(dataGroup map[string][]map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	stRows, ok := dataGroup[st.GetKey()]
	if !ok {
		return result
	}
	for _, field := range st.Fields {
		colValues := make([]string, 0)
		for _, row := range stRows {
			value, ok1 := row[field.Column]
			if ok1 {
				colValues = append(colValues, value)
			}
		}
		var tarValue interface{}
		var tarValues = make([]interface{}, 0)

		switch field.SourceFieldTargetType {
		case consts.SourceFieldTargetTypeBoolExists:
			tarValue = len(colValues) > 0
		case consts.SourceFieldTargetTypeBoolMatchAny:
			tarValue = len(colValues) > 0
			for _, colValue := range colValues {
				if colValue != "" && colValue != "0" {
					tarValue = true
					break
				}
			}
		case consts.SourceFieldTargetTypeBoolMatchAll:
			tarValue = len(colValues) > 0
			for _, colValue := range colValues {
				if colValue == "" || colValue == "0" {
					tarValue = false
					break
				}
			}
		case consts.SourceFieldTargetTypeNumberCount:
			tarValue = len(colValues)
		case consts.SourceFieldTargetTypeNumberSum:
			var sum float64 = 0
			for _, colValue := range colValues {
				num, err := strconv.ParseFloat(colValue, 64)
				if err == nil {
					continue
				}
				sum = sum + num
				tarValue = sum
			}
		case consts.SourceFieldTargetTypeNumberMax:
			max := -math.MaxFloat64
			for _, colValue := range colValues {
				num, err := strconv.ParseFloat(colValue, 64)
				if err == nil {
					continue
				}
				max = math.Max(max, num)
				tarValue = max
			}
		case consts.SourceFieldTargetTypeNumberMin:
			min := math.MaxFloat64
			for _, colValue := range colValues {
				num, err := strconv.ParseFloat(colValue, 64)
				if err == nil {
					continue
				}
				min = math.Max(min, num)
				tarValue = min
			}
		default:
			for _, colValue := range colValues {
				tarValues = append(tarValues, colValue)
			}
		}
		if len(tarValues) == 0 {
			result[field.Target] = tarValue
		} else if len(tarValues) == 1 {
			result[field.Target] = tarValues[0]
		} else {
			result[field.Target] = tarValues
		}
	}
	for _, relation := range st.Relations {
		for k, v := range relation.SourceTree.toDocument(dataGroup) {
			result[k] = v
		}
	}
	return result
}

func (st *SourceTree) matchChildSourceFilters(rows []map[string]string) map[string]SourceFilter {
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

func (st *SourceTree) matchParentSourceFilters(rows []map[string]string, parent *SourceTree) SourceFilter {
	sf := SourceFilter{Type: consts.SourceRelationTypeAnd}
	for _, relation := range parent.Relations {
		if relation.SourceTree.GetKey() != st.GetKey() {
			continue
		}
		valuesMap := make(map[string][]string)
		extraColumns := make([]string, 0)
		for _, field := range relation.Fields {
			target := field.Target
			column := field.Column
			values := make([]string, 0)
			for _, row := range rows {
				value, ok := row[target]
				if ok {
					values = append(values, value)
				}
			}
			valuesMap[column] = values
			extraColumns = append(extraColumns, column)
		}
		sf.Values = valuesMap
		sf.ExtraColumns = extraColumns
	}
	return sf
}
