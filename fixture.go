package tester

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type (
	TableName  = string
	ColumnName = string
	DataValue  = interface{}

	RowData           = map[ColumnName]DataValue
	TableData         []RowData
	FixtureData       map[TableName]TableData
	FixtureCollection map[DbName]*FixtureElement
)

type check struct {
	column   string
	expected string
	code     errCode
}

func (c check) string() string {
	switch c.code {
	case checkErrorNotExist:
		{
			return "NE"
		}
	case checkErrorNotEqual:
		{
			return "!= " + c.expected
		}
	}

	return ""
}

func (t TableData) compare(expected TableData) (map[int][]check, error) {
	result := make(map[int][]check, len(t))

	for realIdx, realRow := range t {
		// If expected data smaller than real data
		if realIdx > len(expected)-1 {
			for realCol := range realRow {
				result[realIdx] = append(result[realIdx], check{
					column:   realCol,
					expected: "",
					code:     checkErrorNotExist,
				})
			}
			continue
		}

		expectRow := expected[realIdx]
		for expectCol, expectVal := range expectRow {
			realVal, e := realRow[expectCol]
			if !e {
				return nil, errors.New(fmt.Sprintf("Real Db Data Not Exists Column: %s", expectCol))
			}

			expectVal := interfaceToString(expectVal)
			if interfaceToString(realVal) != expectVal {
				if _, e := result[realIdx]; !e {
					result[realIdx] = make([]check, 0, 5)
				}
				result[realIdx] = append(result[realIdx], check{
					column:   expectCol,
					expected: expectVal,
					code:     checkErrorNotEqual,
				})
			}
		}
	}

	return result, nil
}

func (t TableData) getFormattedData(colNames []string) []row {
	rowsLen := len(t)
	colsNameLen := len(colNames)

	result := make([]row, 0, len(t))

	for i := 0; i < rowsLen; i++ {
		// row(map[string]interface{})
		newRow := make(row, 0, colsNameLen)

		for j := 0; j < colsNameLen; j++ {
			// colNames
			for k, v := range t[i] {

				// k -> column Name
				if k == colNames[j] {
					newRow = append(newRow, v)
					break
				}
			}
		}
		result = append(result, newRow)
	}

	return result
}

func (f FixtureData) getInsertQueries() []string {
	queryList := make([]string, 0, 5)

	for tableName, dataSetList := range f {
		for _, dataSet := range dataSetList {
			query := "INSERT IGNORE INTO "
			query += tableName

			count := 0
			dataLen := len(dataSet)
			var columnQuery, valueQuery string
			columnQuery += "("
			valueQuery += "("
			for columnName, value := range dataSet {

				columnQuery += columnName
				switch value.(type) {
				case int:
					valueQuery += strconv.Itoa(value.(int))
				case string:
					valueQuery += "\"" + value.(string) + "\""
				}
				if count < dataLen-1 {
					columnQuery += ", "
					valueQuery += ", "
				}
				count++
			}
			columnQuery += ") "
			valueQuery += ")"
			query += columnQuery + "VALUES" + valueQuery + ";"
			queryList = append(queryList, query)
		}
	}

	return queryList
}

func (f FixtureData) getDeleteAllQueries() []string {
	queryList := make([]string, 0, 5)

	for tableName := range f {
		query := "DELETE FROM "
		query += tableName
		queryList = append(queryList, query)
	}

	return queryList
}

type FixtureElement struct {
	Before func() FixtureData
	After  func() FixtureData
}

type fixtureInfo struct {
	db      *sql.DB
	element *FixtureElement
}
