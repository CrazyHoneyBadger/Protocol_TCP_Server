package tcpprotocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtocolClient_create(t *testing.T) {
	pow := &MockPow{}
	pc := NewProtocolClient(pow)
	assert.NotNil(t, pc)
}
func TestProtocolClient_Request_wrongRequest(t *testing.T) {
	pow := &MockPow{}
	pc := NewProtocolClient(pow)
	data := []byte("asdasdadasdsad")
	result, err := pc.Request(data)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrWrongRequestFormat)
}
func TestProtocolClient_Request_wrongVersion(t *testing.T) {
	pow := &MockPow{}
	pc := NewProtocolClient(pow)
	data := []byte("1.1.0;key1:value1;key2:value2;OK")
	result, err := pc.Request(data)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrInvalidProtocolVersion)
}
func TestProtocolClient_Request_NotComplexity(t *testing.T) {
	pow := &MockPow{}
	pc := NewProtocolClient(pow)
	data := []byte("0.1.0;key1:value1;key2:value2;OK")
	result, err := pc.Request(data)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrInvalidComplexity)
}
