package api

import (
	"testing"

)

func TestGetRate(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}