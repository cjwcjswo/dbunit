package dbunit

const (
	showColumnsQuery = "SHOW COLUMNS FROM %s"
)

type errCode int16

const (
	checkErrorNone = iota
	checkErrorNotExist
	checkErrorNotEqual
)
