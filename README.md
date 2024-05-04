# sms-gateway
[![CI](https://github.com/jatm80/sms-gateway/actions/workflows/ci.yaml/badge.svg)](https://github.com/jatm80/sms-gateway/actions/workflows/ci.yaml)

![SMS Gateway](./docs/img/sms-gateway.jpg)

## Requirements
- Huawei E3372 USB Modem
- Some compute (raspberry pi, docker)

## Using sms-gateway

```
export CERT_BASE64=$(base64 -w 0 /path/to/cert.pem)
export KEY_BASE64=$(base64 -w 0 /path/to/key.pem)
export TELEGRAM_TOKEN=123:abc 
export TELEGRAM_CHAT_ID=567 
./sms-gateway
```
Enviroment variables:
- TELEGRAM_TOKEN: Required
- TELEGRAM_CHAT_ID: Required
- CERT_BASE64: Required either self-signed or signed
- KEY_BASE64: Required either self-signed or signed
- BIND_ADDRESS_PORT: defaults to 8443
- DEFAULT_BASE_URL:  defaults to `http://192.168.8.1`


## Generate self signed certificate
```
openssl req -newkey rsa:4096 -sha256 -nodes -keyout PRIVATE.key -x509 -days 365 -out PUBLIC.pem
```

## Configure Telegram Bot with Self signed certificate
* Modify the file in `scripts/self-signed-cert.html` and defined the following:
```
        token: 'xxx',
        port: 8443,
        host: 'your-url-host',

  // token: is your telegram bot token.  
  // port: is the exposed port of the service sms-gateway. Telegram only allows for 443 or 8443.
  // host: is the url to reach to service sms-gateway once deployed.         
```  

* Click in `Set Webhook`, Telegram will respond with the URL of the webhook


## References:
https://www.0xf8.org/2017/01/flashing-a-huawei-e3372h-4g-lte-stick-from-hilink-to-stick-mode/
https://stephenmonro.wordpress.com/2019/02/13/getting-sms-messages-from-the-huawei-e3372-lte-modem/
https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
https://core.telegram.org/bots/webhooks
https://core.telegram.org/bots/webhooks#how-do-i-set-a-webhook-for-either-type
