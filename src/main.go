package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
    r.HandleFunc("/send-sms",UpdateHandler).Methods(http.MethodPost)

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
					err := telegram.SendToTelegram(fmt.Sprintf("[%d] %s => %s \n", k, message.Date, message.Content))
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


func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// h := &huawei.E3372{
	// 	BaseURL: huawei.DEFAULT_BASE_URL,
	// }

	message, err := telegram.ParseTelegramRequest(r)
    if err != nil {
		return 
	}

	log.Println(message.Message.Text)
}