package tcpprotocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtocolServer_Request(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("0.1.0;key1:value1;key2:value2;OK")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, err)
	assert.Equal(t, "value1", (*result)["key1"])
	assert.Equal(t, "value2", (*result)["key2"])
}
func TestProtocolServer_Request_Payload_error(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("0.1.0;key1:value1;key2;OK")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrInvalidPayloadFormat)

}
func TestProtocolServer_Request_Pow_error(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("0.1.0;key1:value1;key2;OK:OKs")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}
func TestProtocolServer_Request_Key_error(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("0.1.0;key1:value1;key2;OKs:OKs")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrInvalidRequestKey)

}
func TestProtocolServer_Request_Version_error(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("1.1.0;key1:value1;key2;OKs:OKs")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrInvalidProtocolVersion)

}
func TestProtocolServer_Request_String_error(t *testing.T) {
	pow := &MockPow{}
	ps := NewProtocolServer(pow)
	data := []byte("asdasdadasdsad")
	result, err := ps.Request(data, "OK")
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrWrongRequestFormat)

}
