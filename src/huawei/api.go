package huawei

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/jatm80/sms-gateway/http"
)

type E3372 struct {
    BaseURL string
}

type Session struct {
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

type Response struct {
    XMLName xml.Name `xml:"response"`
    Content string   `xml:",innerxml"`
}

type ErrorResponse struct {
    XMLName xml.Name `xml:"error"`
    Code    int      `xml:"code"`
    Message string   `xml:"message"`
}

func (e *E3372) Login() (Session, error) {

	var session Session

    c := &http.Request{
		Path: e.BaseURL + "/webserver/SesTokInfo",
		Method: "GET",
	}

	r, err := c.Send()
    if err != nil {
      return Session{},err
	}
     
	err = xml.Unmarshal(r, &session)
	if err != nil {
		return Session{},err
	}

	return session, nil
}


func (e *E3372) GetSMSList() (Messages, error)  {
	
	s, err := e.Login()
	if err != nil {
		return Messages{}, err
	}

	c := &http.Request{
		Path: e.BaseURL + "/sms/sms-list",
		Method: "POST",
		Headers: map[string]string{
			"Cookie": s.SesInfo,
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			"__RequestVerificationToken": s.TokInfo,
		},
		Body: []byte("<?xml version='1.0' encoding='UTF-8'?><request><PageIndex>1</PageIndex><ReadCount>20</ReadCount><BoxType>1</BoxType><SortType>0</SortType><Ascending>0</Ascending><UnreadPreferred>0</UnreadPreferred></request>"),
	}

	b,err := c.Send()
	if err != nil {
		return Messages{}, err
	}

	var messages Messages

	err = xml.Unmarshal(b, &messages)
	if err != nil {
		return Messages{}, err
	}

	return messages, nil
}

func (e *E3372) MarkAsRead(m Message) (bool, error) {

	s,err := e.Login()
	if err != nil {
		return false, err
	}

	c := &http.Request{
      Path: e.BaseURL + "/sms/set-read",
	  Method: "POST",
	  Headers: map[string]string{
		"Cookie": s.SesInfo,
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"__RequestVerificationToken": s.TokInfo,
		},
	  Body: []byte(fmt.Sprintf("<?xml version='1.0' encoding='UTF-8'?><request><Index>%d</Index></request>",m.Index)),
    }

	b,err := c.Send()
	if err != nil {
		return false, err
	}

	resp, err := parseXMLResponse(b)
    if err != nil {
        return false, err
    } 

	return resp, nil

}

func (e *E3372) DeleteSMS(m Message) (bool, error) {

	s,err := e.Login()
	if err != nil {
		return false, err
	}

	c := &http.Request{
      Path: e.BaseURL + "/sms/delete-sms",
	  Method: "POST",
	  Headers: map[string]string{
		"Cookie": s.SesInfo,
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"__RequestVerificationToken": s.TokInfo,
		},
	  Body: []byte(fmt.Sprintf("<?xml version='1.0' encoding='UTF-8'?><request><Index>%d</Index></request>",m.Index)),
    }

	b,err := c.Send()
	if err != nil {
		return false, err
	}

	resp, err := parseXMLResponse(b)
    if err != nil {
        return false, err
    } 

	return resp, nil
}

func (e *E3372) SendSMS(number string, message string) error {

	s,err := e.Login()
	if err != nil {
		return err
	}

	c := &http.Request{
      Path: e.BaseURL + "/sms/send-sms",
	  Method: "POST",
	  Headers: map[string]string{
		"Cookie": s.SesInfo,
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"__RequestVerificationToken": s.TokInfo,
		},
	  Body: []byte(fmt.Sprintf("<?xml version='1.0' encoding='UTF-8'?><request><Index>-1</Index><Phones><Phone>%s</Phone></Phones><Sca></Sca><Content>%s</Content><Length>-1</Length><Reserved>1</Reserved><Date>-1</Date></request>",number,message)),
    }

	resp,err := c.Send()
	if err != nil {
		return err
	}

	_, err = parseXMLResponse(resp)
    if err != nil {
        return err
    } 

	return nil
}


func parseXMLResponse(xmlData []byte) (bool, error) {
    var response Response
    err := xml.Unmarshal(xmlData, &response)
    if err == nil {
        if response.XMLName.Local == "response" {
            return true, nil
        }
    }

    var errorResponse ErrorResponse
    err = xml.Unmarshal(xmlData, &errorResponse)
    if err == nil {
        if errorResponse.XMLName.Local == "error" {
            return false, fmt.Errorf("error %d: %s", errorResponse.Code, errorResponse.Message)
        }
    }

    return false, fmt.Errorf("unexpected XML response format")
}

func GetAPIBaseUrl() string {
	var baseUrl = "http://192.168.8.1/api"
	if b, ok := os.LookupEnv("DEFAULT_BASE_URL"); ok  {
		baseUrl = b
	}
	return baseUrl
}