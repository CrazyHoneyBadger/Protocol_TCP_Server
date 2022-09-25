package tcpprotocol

import "fmt"

const (
	ProtocolVersion = "0.1.0"
	POWVersion      = "0.1.0"
	OK              = iota
)

var (
	ErrWrongRequestFormat     = fmt.Errorf("wrong request format")
	ErrInvalidProtocolVersion = fmt.Errorf("invalid protocol version")
	ErrInvalidRequestKey      = fmt.Errorf("invalid request key")
	ErrInvalidPayloadFormat   = fmt.Errorf("the payload must be in key:value format")
)
