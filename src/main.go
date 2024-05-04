package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/forPelevin/gomoji"
	"github.com/gorilla/mux"
	"github.com/jatm80/sms-gateway/huawei"
	"github.com/jatm80/sms-gateway/telegram"
)

var telegramToken = os.Getenv("TELEGRAM_TOKEN")
var telegramChatID = os.Getenv("TELEGRAM_CHAT_ID")

func main() {
	certBase64 := os.Getenv("CERT_BASE64")
	keyBase64 := os.Getenv("KEY_BASE64")

	if certBase64 == "" || keyBase64 == "" {
		log.Fatal("CERT_BASE64 and KEY_BASE64 environment variables must be set")
	}

	var bindAddressPort = ":8443"
	if value, ok := os.LookupEnv("BIND_ADDRESS_PORT"); ok {
		bindAddressPort = value
	}

	t := &telegram.Bot{
		Token:  telegramToken,
		ChatID: telegramChatID,
	}

	r := mux.NewRouter()
	r.StrictSlash(false)
	if t.IsTelegramEnabled() {
		r.HandleFunc("/"+t.GetTelegramToken(), TelegramHandler).Methods(http.MethodPost)
	}
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not Found"))
	})

	go func() {
		h := &huawei.E3372{
			BaseURL: huawei.GetAPIBaseURL(),
		}

		for {
			m, err := h.GetSMSList()
			if err != nil {
				log.Println(err)
			}

			log.Printf("SMS received: %v", m.Messages)

			if m.Count > 0 {
				for k, message := range m.Messages {
					if message.Smstat == 0 {
						log.Println(message)
						if t.IsTelegramEnabled() {
							err := t.SendToTelegram(fmt.Sprintf("[%d] %s => %s \n", k, message.Phone, gomoji.RemoveEmojis(message.Content)))
							if err != nil {
								log.Println(err)
								return
							}
							_, err = h.MarkAsRead(message)
							if err != nil {
								log.Println(err)
								return
							}
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
			time.Sleep(30 * time.Second)
		}
	}()

	cert, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		log.Fatalf("Failed to decode certificate: %v", err)
	}

	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		log.Fatalf("Failed to decode key: %v", err)
	}

	err = http.ListenAndServeTLS(bindAddressPort, string(cert), string(key), r)
	if err != nil {
		log.Fatalf("Something went wrong starting web service: %v", err)
	}
}

func TelegramHandler(_ http.ResponseWriter, r *http.Request) {
	h := &huawei.E3372{
		BaseURL: huawei.GetAPIBaseURL(),
	}

	t := &telegram.Bot{
		Token:  telegramToken,
		ChatID: telegramChatID,
	}

	message, err := telegram.ParseTelegramRequest(r)
	if err != nil {
		return
	}

	cmd, num, text, err := telegram.ExtractData(message.Message.Text)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if cmd == "/sms" {
		if !telegram.IsValidAustralianMobile(num) {
			log.Println("not valid AU mobile number")
			return
		}

		ne := gomoji.RemoveEmojis(text)

		log.Printf("Received /sms - sending message %s to %s \n", ne, num)

		err = h.SendSMS(num, ne)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if t.IsTelegramEnabled() {
			err := t.SendToTelegram("SMS Sent")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
