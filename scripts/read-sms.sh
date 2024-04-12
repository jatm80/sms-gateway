#!/bin/bash

DATA=`curl -s http://192.168.8.1/api/webserver/SesTokInfo`
SESSION_ID=`echo "$DATA" | grep "SessionID=" | cut -b 10-147`
TOKEN=`echo "$DATA" | grep "TokInfo" | cut -b 10-41`

echo "monitoring: "
curl http://192.168.8.1/api/monitoring/status -H "Cookie: $SESSION_ID"

echo "sms info: "
curl http://192.168.8.1/api/sms/sms-count -H "Cookie: $SESSION_ID"

echo "list SMS: "
curl -X POST http://192.168.8.1/api/sms/sms-list \
	-H "Cookie: $SESSION_ID" \
	-H "__RequestVerificationToken: $TOKEN" \
	-H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" \
	--data "<?xml version='1.0' encoding='UTF-8'?><request><PageIndex>1</PageIndex><ReadCount>20</ReadCount><BoxType>1</BoxType><SortType>0</SortType><Ascending>0</Ascending><UnreadPreferred>0</UnreadPreferred></request>"
