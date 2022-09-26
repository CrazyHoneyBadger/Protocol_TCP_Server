package tcpprotocol

import (
	"fmt"
	"strings"
)

const (
	ProtocolVersion = "0.1.0"
)

var (
	ErrWrongRequestFormat     = fmt.Errorf("wrong request format")
	ErrInvalidProtocolVersion = fmt.Errorf("invalid protocol version")
	ErrInvalidRequestKey      = fmt.Errorf("invalid request key")
	ErrInvalidPayloadFormat   = fmt.Errorf("the payload must be in key:value format")
	ErrInvalidPowVersion      = fmt.Errorf("invalid pow version")
)

func parseToMap(data []byte, powVersion string) (map[string]string, error) {
	payload := strings.Split(string(data), ";")
	payload = payload[:len(payload)-1]
	if len(payload) <= 1 {
		return nil, ErrWrongRequestFormat
	}
	result := make(map[string]string)
	for _, item := range payload {
		kv := strings.Split(item, ":")
		if len(kv) < 2 || kv[0] == "" || kv[1] == "" {
			return nil, ErrInvalidPayloadFormat
		}
		result[kv[0]] = kv[1]
	}
	if prtVersion, ok := result["PROT_VER"]; !ok || prtVersion != ProtocolVersion {
		return nil, ErrInvalidProtocolVersion
	}
	if powVer, ok := result["POW_VER"]; !ok || powVer != powVersion {
		return nil, ErrInvalidPowVersion
	}
	return result, nil
}

func parseToBytes(data map[string]string, powVersion string) string {
	data["PROT_VER"] = ProtocolVersion
	data["POW_VER"] = powVersion
	var result strings.Builder
	for k, v := range data {
		result.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	return result.String()
}
