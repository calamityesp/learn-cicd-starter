package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(T *testing.T) {
	testheader := http.Header{}
	testheader.Add("Authorization", "ApiKey apikey")

	apiKeyExpected := "apikey"
	testApiKey, _ := GetAPIKey(testheader)
	if !reflect.DeepEqual(apiKeyExpected, testApiKey) {
		T.Fatalf("Expected %s but got %s ", apiKeyExpected, testApiKey)
	}
}
