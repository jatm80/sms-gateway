package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jatm80/sms-gateway/huawei"
	"github.com/jatm80/sms-gateway/telegram"
)



func main () {

	var BIND_ADDRESS_PORT = ":8443"
    if value, ok := os.LookupEnv("BIND_ADDRESS_PORT"); ok  {
		BIND_ADDRESS_PORT = value
	}

	r := mux.NewRouter()
	r.StrictSlash(false)
    if telegram.IsTelegramEnabled()  {
		r.HandleFunc("/" + telegram.GetTelegramToken(),TelegramHandler).Methods(http.MethodPost)
	}
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	})

    go func(){
		h := &huawei.E3372{
			BaseURL: huawei.GetAPIBaseUrl(),
		}

        for {
			m, err := h.GetSMSList()
			if err != nil {
				log.Println(err)
			}

			fmt.Println(m.Messages)

			if m.Count > 0 {
				for k,message := range m.Messages {
					if message.Smstat == 0 {
						log.Println(message)
						if telegram.IsTelegramEnabled() {
							err := telegram.SendToTelegram(fmt.Sprintf("[%d] %s => %s \n", k, message.Date, message.Content))
							if err != nil {
								log.Println(err)
								return
							} else {
								_, err = h.MarkAsRead(message)
								if err != nil {
									log.Println(err)
									return
								}
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

	http.ListenAndServeTLS(BIND_ADDRESS_PORT,"PUBLIC.pem","PRIVATE.key",r)

}


func TelegramHandler(w http.ResponseWriter, r *http.Request) {
	h := &huawei.E3372{
		BaseURL: huawei.GetAPIBaseUrl(),
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
		if ! telegram.IsValidAustralianMobile(num) {
			log.Println("not valid AU mobile number")
			return
		 }
	
		 fmt.Printf("Received /sms - sending message %s to %s \n",text, num)
	
		err = h.SendSMS(num,text)
		if err != nil {
			log.Println(err.Error())
			return
		 } else {
			if telegram.IsTelegramEnabled() {
				err := telegram.SendToTelegram("SMS Sent")
				if err != nil {
					log.Println(err)
					return
				} 
			}
		 }
	 }

}