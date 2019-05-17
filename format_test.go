package dbunit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettableLineWord(t *testing.T) {
	// given
	result := make([]string, 5)

	// when
	for i := 1; i <= 5; i++ {
		result[i-1] = gettableLineWord(i)
	}

	// then
	for i := 1; i <= 5; i++ {
		assert.Equal(t, tableLineWordNumPerColumn*i, len(result[i-1]))
		println(result[i-1])
	}
}

func TestGetTableColumnWord(t *testing.T) {
	// given

	// when
	println(getTableColumnWord(11, false))
	println(getTableColumnWord(11, true))
	println(getTableColumnWord("odd", false))
	println(getTableColumnWord("odd", true))
	println(getTableColumnWord([]byte("even"), false))
	println(getTableColumnWord([]byte("even"), true))

	// then
}

func TestAddCheckErrorToValue(t *testing.T) {
	// given

	// when
	value := addCheckErrorToValue("cjwoov", "name", nil)

	// then
	assert.Equal(t, "cjwoov", value)
}

func TestAddCheckErrorToValueNotExist(t *testing.T) {
	// given

	// when
	value := addCheckErrorToValue("cjwoov", "name", []check{{
		column:   "name",
		expected: "",
		code:     checkErrorNotExist,
	}})

	// then
	assert.Equal(t, "cjwoov(NE)", value)
}

func TestAddCheckErrorToValueNotEqual(t *testing.T) {
	// given

	// when
	value := addCheckErrorToValue("cjwoov", "name", []check{{
		column:   "name",
		expected: "battlecook",
		code:     checkErrorNotEqual,
	}})

	// then
	assert.Equal(t, "cjwoov(!= battlecook)", value)
}

func TestGetCheckResult(t *testing.T) {
	// given
	colNames := []string{"seq", "name", "age", "country"}
	rowDataList := []row{
		{
			"1", "cjwoov", 28, "Korea",
		},
		{
			"2", "tester", 33, "Japan",
		},
	}
	checkMap := map[int][]check{
		0: {{
			column:   "age",
			expected: "30",
			code:     checkErrorNotEqual,
		}},
		1: {{
			column:   "age",
			expected: "35",
			code:     checkErrorNotEqual,
		}, {
			column:   "country",
			expected: "USA",
			code:     checkErrorNotEqual,
		},
		},
	}

	// when
	println(getCheckResult("TB_PERSON", colNames, rowDataList, checkMap))

	// then
}
