package tester

import (
	"database/sql"
	"strconv"
)

// Parse interface type to string type
func interfaceToString(input interface{}) string {
	var result string

	switch r := input.(type) {
	case int:
		{
			result = strconv.Itoa(r)
		}
	case []byte:
		{
			result = string(r)
		}
	case string:
		result = r
	}

	return result
}

// Return rows formatted map[string]interface{} data
func getFormattedTableDataFromQuery(rows *sql.Rows) (TableData, error) {
	cols, _ := rows.Columns()
	resultRows := make([]RowData, 0, 5)

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		resultRows = append(resultRows, m)
	}

	return resultRows, nil
}
