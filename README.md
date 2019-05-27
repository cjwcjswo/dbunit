# Golang Database Testing tool
[[![GoDoc](https://godoc.org/github.com/cjwcjswo/dbunit?status.svg)](https://godoc.org/github.com/cjwcjswo/dbunit) [![Go Report Card](https://goreportcard.com/badge/github.com/cjwcjswo/dbunit)](https://goreportcard.com/report/github.com/cjwcjswo/dbunit)

## Install 

```bash
go get github.com/cjwcjswo/dbunit
```

## Usage

### 1. Init database config function

```go
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
				"fixtures/init_table.sql", // table scheme files
			},
		},
	}
}
```

### 2. Define 'Before Table Data' & "After Table Data"

- 'Before Table Data': Table data needed to run the handler

- 'After Table Data': Expected table Data after handler execution

```go
func BeforeFunc() dbunit.FixtureData {
	return dbunit.FixtureData{
		"TB_USER": {
			{
				"name":    "cjwoov",
				"age":     28,
				"country": "korea",
			},
			{
				"name":    "battlecook",
				"age":     34,
				"country": "korea",
			},
		},
	}
}
func AfterFunc() dbunit.FixtureData {
	return dbunit.FixtureData{
		"TB_USER": {
			{
				"name": "cjwoov",
				"age":  29,
			},
			{
				"name":    "battlecook",
				"age":     34,
				"country": "korea",
			},
		},
	}
}

```

### 3. Initialize the data needed to execute the logic

```go
	c := dbunit.FixtureCollection{
		DbName: &dbunit.FixtureElement{
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	}
	dbunit.Init(&c) // Table data initialized after this function is called.
```

### 4. Compare Real Table data to Expected Table data 

```go
dbunit.AssertTableData(t)
```

## Support drivers

- mysql

## License

MIT