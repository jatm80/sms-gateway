package telegram

import (
	"bytes"
	"net/http"
	"testing"
)

type MockReadCloser struct {
	*bytes.Reader
}

func (m *MockReadCloser) Close() error {
	return nil
}

type mockRequest struct {
	body []byte
}

func (m *mockRequest) Read(p []byte) (n int, err error) {
	copy(p, m.body)
	return len(m.body), nil
}

func TestIsTelegramEnabled(t *testing.T) {
	botWithToken := &Bot{Token: "1234567890:abcdef", ChatId: "123456789"}
	botWithoutToken := &Bot{Token: "", ChatId: "123456789"}

	if !botWithToken.IsTelegramEnabled() {
		t.Errorf("Expected botWithToken to be enabled, got disabled")
	}

	if botWithoutToken.IsTelegramEnabled() {
		t.Errorf("Expected botWithoutToken to be disabled, got enabled")
	}
}

func TestSendToTelegram(t *testing.T) {
	bot := &Bot{Token: "1234567890:abcdef", ChatId: "123456789"}

	err := bot.SendToTelegram("Hello, Telegram!")
	if err != nil {
		t.Errorf("SendToTelegram returned error: %v", err)
	}
}

func TestParseTelegramRequest(t *testing.T) {
	updateJSON := `{"update_id":123,"message":{"text":"Hello","chat":{"id":456}}}`
	mockBody := []byte(updateJSON)

	mockReq := &http.Request{
		Body: &MockReadCloser{bytes.NewReader(mockBody)},
	}

	update, err := ParseTelegramRequest(mockReq)
	if err != nil {
		t.Errorf("ParseTelegramRequest returned error: %v", err)
	}

	if update.UpdateId != 123 {
		t.Errorf("Expected UpdateId to be 123, got %d", update.UpdateId)
	}

	if update.Message.Text != "Hello" {
		t.Errorf("Expected Message.Text to be 'Hello', got %s", update.Message.Text)
	}

	if update.Message.Chat.Id != 456 {
		t.Errorf("Expected Chat.Id to be 456, got %d", update.Message.Chat.Id)
	}
}

func TestExtractData(t *testing.T) {
	input := "command value1 text"
	cmd, val, txt, err := ExtractData(input)

	if err != nil {
		t.Errorf("ExtractData returned error: %v", err)
	}

	expectedCmd := "command"
	expectedVal := "value1"
	expectedTxt := "text"

	if cmd != expectedCmd || val != expectedVal || txt != expectedTxt {
		t.Errorf("Extracted data does not match expected values")
	}
}

func TestIsValidAustralianMobile(t *testing.T) {
	validNumber := "0412345678"
	invalidNumber := "1234567890"

	if !IsValidAustralianMobile(validNumber) {
		t.Errorf("Expected %s to be a valid Australian mobile number", validNumber)
	}

	if IsValidAustralianMobile(invalidNumber) {
		t.Errorf("Expected %s to be an invalid Australian mobile number", invalidNumber)
	}
}
