package three_service

import (
	"fmt"
	"testing"
)

func TestValidateOnlineBankCode(t *testing.T) {
	code := "6228481101100634315"
	b := ValidateOnlineBankCode(code)
	if b {
		fmt.Println("卡正确")
	} else {
		fmt.Println("卡不正确")
	}
}
