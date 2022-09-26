package tcpprotocol

import (
	"fmt"
)

type PowInteface interface {
	GenerateUniqKey() string
	ValidateMessage(version, message string) error
	GetComplexity() int
	GetVersion() string
}

type ProtocolServer struct {
	pow PowInteface
}

func NewProtocolServer(pow PowInteface) *ProtocolServer {
	return &ProtocolServer{
		pow: pow,
	}
}

func (p ProtocolServer) Request(data []byte, key string) (map[string]string, error) {
	result, err := parseToMap(data, p.pow.GetVersion())
	if err != nil {
		return nil, err
	}
	if mesKey, ok := result["POW_KEY"]; !ok || mesKey != key {
		return nil, ErrInvalidRequestKey
	}
	if err := p.pow.ValidateMessage(p.pow.GetVersion(), string(data)); err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProtocolServer) Response(data map[string]string) ([]byte, string) {
	data["complexity"] = fmt.Sprintf("%d", p.pow.GetComplexity())
	key := p.pow.GenerateUniqKey()
	data["POW_KEY"] = key
	return []byte(parseToBytes(data, p.pow.GetVersion())), key
}
func (p ProtocolServer) ResponseError(err error) ([]byte, string) {
	data := map[string]string{
		"error": err.Error(),
	}
	return p.Response(data)
}
