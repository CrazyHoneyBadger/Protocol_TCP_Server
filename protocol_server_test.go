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
