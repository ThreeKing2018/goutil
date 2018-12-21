package timer

import "testing"

func TestStart(t *testing.T) {
	Start()
}
func TestStartTicker(t *testing.T) {
	go func() {
		StartTicker()
	}()
}
