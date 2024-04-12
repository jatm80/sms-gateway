#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go build -C src -o ../bin/raspi2-sms-gateway
env GOOS=linux GOARCH=arm GOARM=7 go build -C src -o ../bin/raspi3-sms-gateway
