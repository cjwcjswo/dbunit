package dbunit

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	tableLineWord             = "-"
	tableLineWordNumPerColumn = 30
	tableSpaceWord            = "|"
	blank                     = " "

	lineSeparator = "\n"
)

type row = []interface{}

func getCheckResult(tableName string, colNames []string, rowDataList []row, checkMap map[int][]check) string {
	var buf bytes.Buffer
	colNamesLen := len(colNames)
	rowDataListLen := len(rowDataList)

	// Write Table name
	buf.WriteString(tableName + lineSeparator)
	buf.WriteString(gettableLineWord(colNamesLen) + lineSeparator)

	// Write Columns Name
	for idx, colName := range colNames {
		addWord := false
		if idx == colNamesLen-1 {
			addWord = true
		}
		buf.WriteString(getTableColumnWord(colName, addWord))
	}
	buf.WriteString(lineSeparator)

	for rowIdx, rowData := range rowDataList {
		// Write row data
		rowDataLen := len(rowData)
		buf.WriteString(gettableLineWord(rowDataLen) + lineSeparator)
		idx := 0
		for colIdx, value := range rowData {
			addWord := false
			idx++
			if idx == len(rowData) {
				addWord = true
			}

			if checkList, e := checkMap[rowIdx]; e {
				value = addCheckErrorToValue(value, colNames[colIdx], checkList)
			}
			buf.WriteString(getTableColumnWord(value, addWord))
		}
		buf.WriteString(lineSeparator)

		// If row data is end, write table line
		if rowIdx == rowDataListLen-1 {
			buf.WriteString(gettableLineWord(rowDataLen) + lineSeparator)
		}
	}

	// Return String Buffer
	return buf.String() + lineSeparator
}

func addCheckErrorToValue(value interface{}, columnName string, checkList []check) string {
	result := interfaceToString(value)
	for _, check := range checkList {
		if columnName == check.column {
			result += fmt.Sprintf("(%s)", check.string())
			break
		}
	}
	return result
}

func getTableColumnWord(value interface{}, isAddEndWord bool) string {
	colName := interfaceToString(value)

	// Calculate text center align
	emptyNum := tableLineWordNumPerColumn - 2
	emptyNum -= len(colName)
	endWord := tableSpaceWord
	if !isAddEndWord {
		endWord = blank
	}
	addWord := ""
	if emptyNum%2 == 1 {
		addWord = blank
	}
	return fmt.Sprintf("%s%s%s%s%s", tableSpaceWord, strings.Repeat(blank, emptyNum/2), colName, strings.Repeat(blank, emptyNum/2)+addWord, endWord)
}

func gettableLineWord(columnNum int) string {
	return strings.Repeat(strings.Repeat(tableLineWord, tableLineWordNumPerColumn), columnNum)
}
