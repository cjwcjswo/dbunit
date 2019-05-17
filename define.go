package dbunit

const (
	showColumnsQuery = "SHOW COLUMNS FROM %s"
)

type errCode int16

const (
	checkErrorNone     = 0
	checkErrorNotExist = 1
	checkErrorNotEqual = 2
)
