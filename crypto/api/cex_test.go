package api_test

import (
	"testing"

	"fem.com/go/crypto/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("error was not found")
	}
}
