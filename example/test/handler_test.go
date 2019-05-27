package test

import (
	"github.com/cjwcjswo/dbunit"
	"github.com/cjwcjswo/dbunit/example/logic"
	"github.com/cjwcjswo/dbunit/example/test/fixtures/fail"
	"github.com/cjwcjswo/dbunit/example/test/fixtures/success"
	"github.com/stretchr/testify/assert"
	"testing"
)

const DbName = "DB_TEST"

func init() {
	dbunit.InitSetupFunc(loadConfig)
}

func loadConfig() dbunit.ConfigMap {
	return dbunit.ConfigMap{
		DbName: {
			DriverName: "mysql",
			Host:       "127.0.0.1",
			Port:       3306,
			UserName:   "root",
			Password:   "",
			Name:       DbName,
			SqlPaths: []string{
				"fixtures/init_table.sql",
			},
		},
	}
}

func TestSuccess(t *testing.T) {
	// given
	c := dbunit.FixtureCollection{
		DbName: &dbunit.FixtureElement{
			Before: success.Before,
			After:  success.After,
		},
	}
	dbunit.Init(&c)

	// when
	err := logic.UpdateAgeHandler("cjwoov", 29)

	// then
	assert.Nil(t, err)
	dbunit.AssertTableData(t)
}

func TestFail(t *testing.T) {
	// given
	c := dbunit.FixtureCollection{
		DbName: &dbunit.FixtureElement{
			Before: fail.Before,
			After:  fail.After,
		},
	}
	dbunit.Init(&c)

	// when
	err := logic.UpdateAgeHandler("cjwoov", 29)

	// then
	assert.Nil(t, err)
	dbunit.AssertTableData(t)
}
