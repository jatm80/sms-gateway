package main

import (
	"encoding/xml"
)

type Login struct {
	XMLName  xml.Name `xml:"response"`
	SesInfo string `xml:"SesInfo"`
	TokInfo string `xml:"TokInfo"`
}

type Messages struct {
	XMLName  xml.Name `xml:"response"`
	Count    int      `xml:"Count"`
	Messages []Message `xml:"Messages>Message"`
}

type Message struct {
	Smstat int `xml:"Smstat"`
	Index int `xml:"Index"`
	Phone string `xml:"Phone"`
	Content string `xml:"Content"`
	Date string `xml:"Date"`
	Sca string `xml:"Sca"`
	SaveType int `xml:"SaveType"`
	Priority int `xml:"Priority"`
	SmsType int `xml:"SmsType"`
}


type Telegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}