package tcpprotocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToMap(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.0;KEY:OK")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.Nil(t, err)
	assert.Equal(t, "0.1.0", result["PROT_VER"])
	assert.Equal(t, "0.1.0", result["POW_VER"])
	assert.Equal(t, "OK", result["KEY"])
}
func TestParseToMap_Empty(t *testing.T) {
	data := []byte("")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrWrongRequestFormat)
	assert.Nil(t, result)
}
func TestParseToMap_RandomString(t *testing.T) {
	data := []byte("sadasdasdsadsadasad")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrWrongRequestFormat)
	assert.Nil(t, result)
}
func TestParseToMap_NotValue(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.0;KEY")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidPayloadFormat)
	assert.Nil(t, result)
}
func TestParseToMap_NotKey(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.0;:OK")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidPayloadFormat)
	assert.Nil(t, result)
}
func TestParseToMap_NotKeyAndValue(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.0;:")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidPayloadFormat)
	assert.Nil(t, result)
}
func TestParseToMap_LastChar(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.0;KEY:OK;")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.Nil(t, err)
	assert.Equal(t, "0.1.0", result["PROT_VER"])
	assert.Equal(t, "0.1.0", result["POW_VER"])
	assert.Equal(t, "OK", result["KEY"])
}
func TestParseToMap_InvalidProtocolVersion(t *testing.T) {
	data := []byte("PROT_VER:0.1.1;POW_VER:0.1.0;KEY:OK")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidProtocolVersion)
	assert.Nil(t, result)
}
func TestParseToMap_InvalidPowVersion(t *testing.T) {
	data := []byte("PROT_VER:0.1.0;POW_VER:0.1.1;KEY:OK")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidPowVersion)
	assert.Nil(t, result)
}
func TestParseToMap_InvalidPowVersionAndProtocolVersion(t *testing.T) {
	data := []byte("PROT_VER:0.1.1;POW_VER:0.1.1;KEY:OK")
	powVersion := "0.1.0"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidProtocolVersion)
	assert.Nil(t, result)
}
func TestParseToMap_NotProtocolVersion(t *testing.T) {
	data := []byte("POW_VER:0.1.0;KEY:OK")
	powVersion := "0.1.1"
	result, err := parseToMap(data, powVersion)
	assert.ErrorIs(t, err, ErrInvalidProtocolVersion)
	assert.Nil(t, result)
}

func TestParseToBytes(t *testing.T) {
	data := map[string]string{
		"PROT_VER": "0.1.0",
		"POW_VER":  "0.1.0",
		"KEY":      "OK",
	}
	powVersion := "0.1.0"
	result := parseToBytes(data, powVersion)
	resMap, err := parseToMap([]byte(result), powVersion)
	assert.Nil(t, err)
	assert.Equal(t, data, resMap)
}

func TestParseToBytes_Empty(t *testing.T) {
	data := map[string]string{}
	powVersion := "0.1.0"
	result := parseToBytes(data, powVersion)
	dataFinal := map[string]string{
		"PROT_VER": "0.1.0",
		"POW_VER":  "0.1.0",
	}
	resMap, err := parseToMap([]byte(result), powVersion)
	assert.Nil(t, err)
	assert.Equal(t, dataFinal, resMap)
}
