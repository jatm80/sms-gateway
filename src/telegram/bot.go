package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"net/http"

	send "github.com/jatm80/sms-gateway/http"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text     string   `json:"text"`
	Chat     Chat     `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

type OutboundMsg struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func SendToTelegram(message string) error {

	TELEGRAM_TOKEN := os.Getenv("TELEGRAM_TOKEN")
	if len(TELEGRAM_TOKEN) == 0 {
		return errors.New("error TELEGRAM_TOKEN not defined")
		}

	TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")
	if len(TELEGRAM_CHAT_ID) == 0 {
		return errors.New("error TELEGRAM_CHAT_ID not defined")
		}
	
	telegramUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",TELEGRAM_TOKEN)	
    telegramChatId, err := strconv.ParseInt(TELEGRAM_CHAT_ID, 10, 64)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(&OutboundMsg{
		ChatID: telegramChatId,
		Text: message,
		
	})
	if err != nil {
		return err
	}

	c := &send.Request{
		Path: telegramUrl,
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

	 err = json.Unmarshal(body,&update)
	 if err != nil {
		return nil, err
	 }

	return &update, nil
}