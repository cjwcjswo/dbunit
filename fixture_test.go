package dbunit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckToString(t *testing.T) {
	// given
	check := check{
		column:   "",
		expected: "",
		code:     checkErrorNone,
	}

	// when
	result := check.string()

	// then
	assert.Equal(t, "", result)
}

func TestCheckToStringNotExist(t *testing.T) {
	// given
	check := check{
		column:   "",
		expected: "",
		code:     checkErrorNotExist,
	}

	// when
	result := check.string()

	// then
	assert.Equal(t, "NE", result)
}

func TestCheckToStringNotEqual(t *testing.T) {
	// given
	check := check{
		column:   "",
		expected: "test",
		code:     checkErrorNotEqual,
	}

	// when
	result := check.string()

	// then
	assert.Equal(t, "!= test", result)
}

func TestCompareTableData(t *testing.T) {
	// given
	data := TableData{
		{
			"seq":     1,
			"name":    "cjwoov",
			"age":     28,
			"country": "Korea",
		},
		{
			"seq":     2,
			"name":    "tester",
			"age":     35,
			"country": "Japan",
		},
	}

	// when
	result, err := data.compare(data)

	// then
	assert.Equal(t, 0, len(result))
	assert.Nil(t, err)
}

func TestCompareTableData2(t *testing.T) {
	// given
	data1 := TableData{
		{
			"seq":     1,
			"name":    "cjwoov",
			"age":     28,
			"country": "Korea",
		},
		{
			"seq":     2,
			"name":    "tester",
			"age":     35,
			"country": "Japan",
		},
	}
	data2 := TableData{
		{
			"seq":     1,
			"name":    "cjwoov",
			"age":     38,
			"country": "Korea",
		},
		{
			"seq":     2,
			"name":    "tester2",
			"age":     35,
			"country": "Japan",
		},
	}

	// when
	result, err := data1.compare(data2)

	// then
	assert.Equal(t, 2, len(result))
	assert.Nil(t, err)
	assert.Equal(t, "age", result[0][0].column)
	assert.Equal(t, "38", result[0][0].expected)
	assert.Equal(t, errCode(checkErrorNotEqual), result[0][0].code)
	assert.Equal(t, "name", result[1][0].column)
	assert.Equal(t, "tester2", result[1][0].expected)
	assert.Equal(t, errCode(checkErrorNotEqual), result[1][0].code)
}

func TestCompareTableData3(t *testing.T) {
	// given
	data1 := TableData{
		{
			"seq":     1,
			"name":    "cjwoov",
			"age":     28,
			"country": "Korea",
		},
		{
			"seq":     2,
			"name":    "tester",
			"age":     35,
			"country": "Japan",
		},
	}
	data2 := TableData{
		{
			"invalid_seq": 1,
			"name":        "cjwoov",
			"age":         38,
			"country":     "Korea",
		},
	}

	// when
	result, err := data1.compare(data2)

	// then
	assert.Equal(t, 0, len(result))
	assert.NotNil(t, err)
}

func TestGetFormattedTableData(t *testing.T) {
	// given
	data := TableData{
		{
			"name":    "cjwoov",
			"age":     28,
			"seq":     1,
			"country": "Korea",
		},
		{
			"age":     35,
			"country": "Japan",
			"name":    "tester",
			"seq":     2,
		},
	}
	expected := []row{
		{
			1, "cjwoov", 28, "Korea",
		},
		{
			2, "tester", 35, "Japan",
		},
	}
	// when
	result := data.getFormattedData([]string{"seq", "name", "age", "country"})

	// then
	assert.Equal(t, 2, len(result))
	for i := 0; i < 2; i++ {
		for j := 0; j < len(result[i]); j++ {
			assert.Equal(t, expected[i][j], result[i][j])
		}
	}
}

func TestGetInsertQueryFixture(t *testing.T) {
	// given
	data := TableData{
		{
			"name":    "cjwoov",
			"age":     28,
			"seq":     1,
			"country": "Korea",
		},
		{
			"age":     35,
			"country": "Japan",
			"name":    "tester",
			"seq":     2,
		},
	}
	fixture := FixtureData{
		"TB_PERSON": data,
	}

	// when
	result := fixture.getInsertQueries()

	// then
	assert.Equal(t, 2, len(result))
	for _, query := range result {
		println(query)
	}
}

func TestGetDeleteQueryFixture(t *testing.T) {
	// given
	data := TableData{
		{
			"name":    "cjwoov",
			"age":     28,
			"seq":     1,
			"country": "Korea",
		},
		{
			"age":     35,
			"country": "Japan",
			"name":    "tester",
			"seq":     2,
		},
	}
	fixture := FixtureData{
		"TB_PERSON": data,
	}

	// when
	result := fixture.getDeleteAllQueries()

	// then
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "DELETE FROM TB_PERSON", result[0])
}
