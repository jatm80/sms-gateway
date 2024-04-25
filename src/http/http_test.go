package http

import (
	"strings"
	"testing"
)


func TestSend_Success(t *testing.T) {

	r := &Request{
		Path:    "http://example.com/api",
		Method:  "GET",
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    []byte{},
	}

	responseBody, err := r.Send()
	if err != nil {
		t.Errorf("Send failed unexpectedly: %v", err)
	}

	expectedResponseBody := "Example Domain"
	if ! strings.Contains(string(responseBody), expectedResponseBody) {
		t.Errorf("Unexpected response body. Expected: %s, Got: %s", expectedResponseBody, responseBody)
	}
}

func TestSend_RequestError(t *testing.T) {

	r := &Request{
		Path:    "invalid-url",
		Method:  "GET",
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    []byte{},
	}

	_, err := r.Send()
	if err == nil {
		t.Error("Expected an error but got none")
	}

}
