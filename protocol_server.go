package tcpprotocol

import (
	"fmt"
	"strings"
)

type PowInteface interface {
	GenerateUniqKey() string
	ValidateMessage(version, message string) error
}

type ProtocolServer struct {
	pow PowInteface
}

func NewProtocolServer(pow PowInteface) *ProtocolServer {
	return &ProtocolServer{
		pow: pow,
	}
}

func (p ProtocolServer) Request(data []byte, key string) (*map[string]string, error) {
	payload := strings.Split(string(data), ";")
	if len(payload) <= 1 {
		return nil, ErrWrongRequestFormat
	}
	if payload[0] != ProtocolVersion {
		return nil, ErrInvalidProtocolVersion
	}
	pkey := strings.Split(payload[len(payload)-1], ":")
	if pkey[0] != key {
		return nil, ErrInvalidRequestKey
	}
	if err := p.pow.ValidateMessage(POWVersion, string(data)); err != nil {
		return nil, err
	}
	payload = payload[1 : len(payload)-1]
	result := make(map[string]string)
	for _, item := range payload {
		kv := strings.Split(item, ":")
		if len(kv) != 2 {
			return nil, ErrInvalidPayloadFormat
		}
		result[kv[0]] = kv[1]
	}
	return &result, nil
}

func (p ProtocolServer) Response(data map[string]string) ([]byte, string) {
	var result strings.Builder
	result.WriteString(ProtocolVersion + ";")
	for k, v := range data {
		result.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	key := p.pow.GenerateUniqKey()
	result.WriteString(key)
	return []byte(result.String()), key
}
func (p ProtocolServer) ResponseError(err string) ([]byte, string) {
	key := p.pow.GenerateUniqKey()
	return []byte(fmt.Sprintf("%s;ERROR:%s;%s", ProtocolVersion, err, key)), key
}
