package tcpprotocol

import (
	"fmt"
	"strconv"
)

type (
	POWClient interface {
		SignMessage(version, message string, complexity int) (string, error)
		GetVersion() string
	}
	ProtocolClient struct {
		pow POWClient
	}
)

var ErrInvalidComplexity = fmt.Errorf("complexity not found")

func NewProtocolClient(pow POWClient) *ProtocolClient {
	return &ProtocolClient{
		pow: pow,
	}
}

func (p ProtocolClient) Request(data []byte) (map[string]string, error) {
	result, err := parseToMap(data, p.pow.GetVersion())
	if err != nil {
		return nil, err
	}
	complexity, ok := result["complexity"]
	if !ok {
		return nil, ErrInvalidComplexity
	}
	if _, err := strconv.Atoi(complexity); err != nil {
		return nil, ErrInvalidComplexity
	}
	if _, ok := result["POW_KEY"]; !ok {
		return nil, ErrInvalidRequestKey
	}
	return result, nil
}

func (p ProtocolClient) Response(data map[string]string, key string, complexity int) ([]byte, error) {
	data["POW_KEY"] = key
	str := parseToBytes(data, p.pow.GetVersion())
	signstring, err := p.pow.SignMessage(p.pow.GetVersion(), str, complexity)
	if err != nil {
		return nil, err
	}

	return []byte(signstring), nil
}
