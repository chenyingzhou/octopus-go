package datasource

import (
	"database/sql"
	"log"
)

func QueryForMapSlice(db *sql.DB, query string) []map[string]string {
	list := make([]map[string]string, 0)
	rows, err := db.Query(query)
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	if err != nil {
		log.Printf("Queray fail, query: %s, #%v", query, err)
		return list
	}

	columns, _ := rows.Columns()
	length := len(columns)
	values := make([]sql.RawBytes, length)
	pointers := make([]interface{}, length)
	for i := 0; i < length; i++ {
		pointers[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			continue
		}
		row := make(map[string]string)
		for i := 0; i < length; i++ {
			row[columns[i]] = string(values[i])
		}
		list = append(list, row)
	}
	return list
}
