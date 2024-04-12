package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var login Login
	var messages Messages

	BASE_URL := os.Getenv("BASE_URL")
	if len(BASE_URL) == 0 {
		BASE_URL = "http://192.168.8.1/api"
		}

	TELEGRAM_TOKEN := os.Getenv("TELEGRAM_TOKEN")
	if len(TELEGRAM_TOKEN) == 0 {
		fmt.Printf("Error TELEGRAM_TOKEN not defined")
		return
		}

	TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")
	if len(TELEGRAM_CHAT_ID) == 0 {
		fmt.Printf("Error TELEGRAM_CHAT_ID not defined")
		return
		}
	
	telegramUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",TELEGRAM_TOKEN)	
    telegramChatId, err := strconv.ParseInt(TELEGRAM_CHAT_ID, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}


	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL + "/webserver/SesTokInfo", nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	err = xml.Unmarshal(responseBody, &login)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		return
	}

	if len(login.SesInfo) > 0 {
	fmt.Println(login.SesInfo)
	fmt.Println(login.TokInfo)
	} else {
		fmt.Printf("Error obtaining session values\n")
		return
	}

	requestBody := `<?xml version='1.0' encoding='UTF-8'?><request><PageIndex>1</PageIndex><ReadCount>20</ReadCount><BoxType>1</BoxType><SortType>0</SortType><Ascending>0</Ascending><UnreadPreferred>0</UnreadPreferred></request>`


	req, err = http.NewRequest("POST", BASE_URL + "/sms/sms-list", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", login.SesInfo)
	req.Header.Set("__RequestVerificationToken", login.TokInfo)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	//fmt.Printf("Response Body:\n%s\n", string(responseBody))


	err = xml.Unmarshal(responseBody, &messages)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		return
	}

	fmt.Printf("Messages available: %d \n", messages.Count)
	for key, message := range messages.Messages {
		fmt.Printf("[%d] %s => %s \n", key, message.Date, message.Content)
		SendMessage(telegramUrl,&Telegram{ChatID: telegramChatId, Text: fmt.Sprintf("[%d] %s => %s \n", key, message.Date, message.Content)})
	}


}

func SendMessage(url string, message *Telegram) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	return nil
}