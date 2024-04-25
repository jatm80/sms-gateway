package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

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

var telegramToken = os.Getenv("TELEGRAM_TOKEN")
var telegramChatId = os.Getenv("TELEGRAM_CHAT_ID")


func SendToTelegram(message string) error {

	if ! IsTelegramEnabled() {
		return errors.New("error TELEGRAM_TOKEN or TELEGRAM_CHAT_ID not defined")
	}
	
	telegramUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",telegramToken)	
    telegramChatId, err := strconv.ParseInt(telegramChatId, 10, 64)
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

func ExtractData(m string) (string, string, string, error) {
 
	s := strings.SplitN(m," ",3)

	if len(s) < 3 {
		return "", "", "", fmt.Errorf("invalid data, received: %s",m)
	}

    return s[0],s[1],s[2],nil
}

func IsValidAustralianMobile(phoneNumber string) bool {
    regex := `^(?:\+?61)?(?:04|\(04\))[0-9]{8}$`
    pattern := regexp.MustCompile(regex)
    return pattern.MatchString(phoneNumber)
}

func IsTelegramEnabled() bool {
	
	if len(telegramToken) == 0 {
		return false
		}
	
	if len(telegramChatId) == 0 {
		return false
		}

	return true
}

func GetTelegramToken () string {
	return telegramToken
}