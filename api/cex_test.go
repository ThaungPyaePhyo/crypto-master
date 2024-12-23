package api_test

import (
	"testing"
	"cryptomasters.com/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
	
}