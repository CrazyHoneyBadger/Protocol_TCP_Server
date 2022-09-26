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
func (mp MockPow) GetComplexity() int {
	return 1
}
func (mp MockPow) GetVersion() string {
	return "0.1.0"
}
func (mp MockPow) SignMessage(version, message string, complexity int) (string, error) {
	return "OK", nil
}
