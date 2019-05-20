package dbunit

import (
	"database/sql"
	"fmt"
	"strings"
)

type (
	DbName    = string
	ConfigMap map[DbName]DbConfig
)

type DbConfig struct {
	DriverName string
	Host       string
	Port       int
	UserName   string
	Password   string
	Name       DbName

	SqlBytes []byte
}

var (
	debug   bool
	setup   func() ConfigMap
	cfgMap  ConfigMap
	connMap map[DbName]fixtureInfo
)

// If debug mode is true, print the result log.
func DebugMode(mode bool) {
	debug = mode
}

// Initialized 'setup' function.
// loadConfigFunc is called by 'Init' function
func InitSetupFunc(loadConfigFunc func() ConfigMap) {
	setup = loadConfigFunc
}

// 1. Load config
//
// Config function is initialized by 'InitSetupFunc'.
// If you want to set load config function, call 'InitSetupFunc'.
//
// 2. Connect to Database
//
// It is set based on the configuration file.
// If the 'SqlBytes' field of 'Config' struct is set, Database executes 'SqlBytes'.
// 'SqlBytes' means byte stream of '.sql' file.
//
// 3. Insert Database row data, that defined in the 'Before function' field of 'FixtureElement'.
func Init(collection *FixtureCollection) {
	if setup == nil {
		panic("database setup func not initialized. please define load config function by call InitSetupFunc")
	}

	cfgMap = setup()
	if len(cfgMap) < 1 {
		panic("config file load fail")
	}

	connect(collection)
	initCollection(collection)
}

// Compare After & Before Database Table Data
//
// If debug mode on, print the log regardless of the result value
//
// If return value is true, After & Before data same
func CompareAfterBefore() bool {
	defer func() {
		deleteAllData()
		for _, conn := range connMap {
			_ = conn.db.Close()
		}
		connMap = nil
	}()

	for _, info := range connMap {
		if info.element == nil {
			continue
		}

		afterData := info.element.After()
		for tbName, afterTbData := range afterData {
			tbColsInfoRaw, err := info.db.Query(fmt.Sprintf(showColumnsQuery, tbName))
			if err != nil {
				panic(err.Error())
			}

			tbColsInfo, err := getFormattedTableDataFromQuery(tbColsInfoRaw)
			if err != nil {
				panic(err.Error())
			}
			colNames := make([]string, 0, 5)
			for _, col := range tbColsInfo {
				colNames = append(colNames, string(col["Field"].([]byte)))
			}

			realTbDataRaw, err := info.db.Query(fmt.Sprintf("SELECT * FROM %s", tbName))
			if err != nil {
				panic(err.Error())
			}
			realTbData, err := getFormattedTableDataFromQuery(realTbDataRaw)
			if err != nil {
				panic(err.Error())
			}
			checkMap, err := realTbData.compare(afterTbData)
			if err != nil {
				panic(err.Error())
			}

			// Print Check Result
			if checkMap != nil && len(checkMap) > 0 {
				println(getCheckResult(tbName, colNames, realTbData.getFormattedData(colNames), checkMap))
				return false
			}
			if debug {
				println(getCheckResult(tbName, colNames, realTbData.getFormattedData(colNames), checkMap))
			}
		}
	}

	return true
}

func deleteAllData() {
	for _, info := range connMap {
		if info.element == nil {
			continue
		}

		queries := info.element.Before().getDeleteAllQueries()
		for _, query := range queries {
			_, _ = info.db.Exec(query)
		}
	}
}

func connect(collection *FixtureCollection) {
	connMap = make(map[DbName]fixtureInfo, len(cfgMap))
	for dbName, config := range cfgMap {
		openDb, err := sql.Open(config.DriverName, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.UserName, config.Password, config.Host, config.Port, config.Name))
		connMap[dbName] = fixtureInfo{
			db:      openDb,
			element: (*collection)[dbName],
		}
		if err != nil {
			panic(err.Error())
		}

		// if byte stream set, init database
		if len(config.SqlBytes) > 0 {
			initSql(connMap[dbName].db, config.SqlBytes)
		}
	}
}

func initSql(db *sql.DB, sqlBytes []byte) {
	queries := strings.Split(string(sqlBytes), ";")
	for _, query := range queries {
		// if query is empty
		if len(strings.TrimSpace(query)) < 1 {
			continue
		}

		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}
}

func initCollection(collection *FixtureCollection) {
	for dbName, data := range *collection {
		db := connMap[dbName].db
		queries := data.Before().getInsertQueries()
		for _, query := range queries {
			_, err := db.Exec(query)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
