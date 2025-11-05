#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go build -C src -o ../bin/sms-gateway_dev_linux_armv6
env GOOS=linux GOARCH=arm GOARM=7 go build -C src -o ../bin/sms-gateway_dev_linux_armv7
