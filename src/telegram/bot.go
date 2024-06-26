package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"net/http"

	send "github.com/jatm80/sms-gateway/http"
)

type Bot struct {
	Token  string
	ChatID string
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

type OutboundMsg struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func (b *Bot) GetTelegramToken() string {
	return b.Token
}

func (b *Bot) GetTelegramChatID() string {
	return b.ChatID
}

func (b *Bot) IsTelegramEnabled() bool {

	if len(b.Token) == 0 {
		return false
	}

	if len(b.ChatID) == 0 {
		return false
	}

	t := regexp.MustCompile(`^\d+:(.*)$`)
	c := regexp.MustCompile(`^[0-9]+$`)
	if t.MatchString(b.Token) && c.MatchString(b.ChatID) {
		return true
	}

	return false
}

func (b *Bot) SendToTelegram(message string) error {

	if !b.IsTelegramEnabled() {
		return errors.New("error TELEGRAM_TOKEN or TELEGRAM_CHAT_ID not defined or valid")
	}

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
	telegramChatID, err := strconv.ParseInt(b.ChatID, 10, 64)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(&OutboundMsg{
		ChatID: telegramChatID,
		Text:   message,
	})
	if err != nil {
		return err
	}

	c := &send.Request{
		Path:   telegramURL,
		Method: "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: payload,
	}
	_, err = c.Send()
	if err != nil {
		return err
	}
	return nil
}

func ParseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &update)
	if err != nil {
		return nil, err
	}

	return &update, nil
}

func ExtractData(m string) (string, string, string, error) {

	s := strings.SplitN(m, " ", 3)

	if len(s) < 3 {
		return "", "", "", fmt.Errorf("invalid data, received: %s", m)
	}

	return s[0], s[1], s[2], nil
}

func IsValidAustralianMobile(phoneNumber string) bool {
	regex := `^(?:\+?61)?(?:04|\(04\))[0-9]{8}$`
	pattern := regexp.MustCompile(regex)
	return pattern.MatchString(phoneNumber)
}
