package dbunit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceToStringTypeInteger(t *testing.T) {
	// given

	// when
	result := interfaceToString(1)

	// then
	assert.Equal(t, "1", result)
}

func TestInterfaceToStringTypeString(t *testing.T) {
	// given

	// when
	result := interfaceToString("1")

	// then
	assert.Equal(t, "1", result)
}

func TestInterfaceToStringTypeBytes(t *testing.T) {
	// given

	// when
	result := interfaceToString([]byte("1"))

	// then
	assert.Equal(t, "1", result)
}
