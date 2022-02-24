package claw

import (
	"database/sql"
	"fmt"
	"github.com/chenyingzhou/octopus-go/datasource"
	"log"
	"strings"
)

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
	where := strings.Join(conditions, sf.Type)
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

	query := fmt.Sprintf("SELECT %s FROM `%s` WHERE %s", "`"+strings.Join(columns, "`,`")+"`", st.DataSet, where)
	rows, err := db.Query(query)
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	if err != nil {
		log.Printf("Queray fail, query: %s, #%v", query, err)
		return
	}

	for rows.Next() {
		row := make(map[string]string)
		for _, column := range columns {
			row[column] = ""
		}
		err = rows.Scan(&row)
		if err != nil {
			continue
		}
		(*data)[st.DataSet] = append((*data)[st.DataSet], row)
	}

}
