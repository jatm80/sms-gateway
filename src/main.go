package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	send "github.com/jatm80/sms-gateway/http"
	"github.com/jatm80/sms-gateway/huawei"
)

type Telegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func main () {

	var BIND_ADDRESS_PORT = ":3000"
    if value, ok := os.LookupEnv("BIND_ADDRESS_PORT"); ok  {
		BIND_ADDRESS_PORT = value
	}

	r := mux.NewRouter()
	r.StrictSlash(false)
    r.HandleFunc("/send-sms",handleOutbound).Methods(http.MethodPost)

    go func(){
		h := &huawei.E3372{
			BaseURL: huawei.DEFAULT_BASE_URL,
		}
		m, err := h.GetSMSList()
		if err != nil {
			log.Println(err)
		}
		if m.Count > 0 {
			for k,message := range m.Messages {
				if message.Smstat == 0 {
					err := sendToTelegram(fmt.Sprintf("[%d] %s => %s \n", k, message.Date, message.Content))
					if err != nil {
						log.Println(err)
						return
					}
					_, err = h.MarkAsRead(message)
					if err != nil {
						log.Println(err)
						return
					}	
				} else {
					_, err := h.DeleteSMS(message)
					if err != nil {
						log.Println(err)
						return
					}	
				}

			}
		}
	}()

	http.ListenAndServe(BIND_ADDRESS_PORT,r)

}


func handleOutbound(w http.ResponseWriter, r *http.Request) {
	h := &huawei.E3372{
		BaseURL: huawei.DEFAULT_BASE_URL,
	}
    h.SendSMS("123","test")
}

func sendToTelegram(message string) error {

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

	payload, err := json.Marshal(&Telegram{
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