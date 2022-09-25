package tcpprotocol

type MockPow struct {
}

func (mp MockPow) GenerateUniqKey() string {
	return "OK"
}
func (mp MockPow) ValidateMessage(version, message string) error {
	if message[len(message)-1] == 's' {
		return ErrInvalidProtocolVersion
	}
	return nil
}
