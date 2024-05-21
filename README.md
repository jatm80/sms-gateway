# sms-gateway
[![CI](https://github.com/jatm80/sms-gateway/actions/workflows/ci.yaml/badge.svg)](https://github.com/jatm80/sms-gateway/actions/workflows/ci.yaml)

![SMS Gateway](./docs/img/sms-gateway.jpg)

## Requirements
- Huawei E3372 USB Modem
- Some compute (raspberry pi, docker)

## Using sms-gateway

* Generate self signed certificate:
```
openssl req -newkey rsa:4096 -sha256 -nodes -keyout PRIVATE.key -x509 -days 365 -out PUBLIC.pem
```
* Run the service:

```
export TELEGRAM_TOKEN=123:abc 
export TELEGRAM_CHAT_ID=567
export CERT_PATH=PUBLIC.pem
export KEY_PATH=PRIVATE.key
./sms-gateway
```
Enviroment variables:
- TELEGRAM_TOKEN: Required
- TELEGRAM_CHAT_ID: Required
- CERT_PATH: Required either self-signed or signed
- KEY_PATH: Required either self-signed or signed
- BIND_ADDRESS_PORT: defaults to 8443
- DEFAULT_BASE_URL:  defaults to `http://192.168.8.1`

## Deploy using ansible in `deploy/ansible-playbook`


* `inventory.ini` update the remote IP and ansible_user   
```
[raspberry_pi]
192.168.1.100 ansible_user=pi ansible_ssh_private_key_file=~/.ssh/id_rsa
```

* `sms-gateway.yaml`  update the variables with the right information:
```
    country: "US"
    state: "New Jersey"
    locality: "Gotham City"
    organization: "jondoe"
    organizational_unit: "None"
    common_name: "sms-bot.example.com"
    email: "webmaster@example.com"
    env_vars:
      TELEGRAM_TOKEN: "your telegram bot token"
      TELEGRAM_CHAT_ID: "your telegram bot chat id"
```

* Run ansible-playbook:

```
cd deploy/ansible-playbook/
ansible-playbook -i inventory.ini sms-gateway.yaml
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
* https://www.0xf8.org/2017/01/flashing-a-huawei-e3372h-4g-lte-stick-from-hilink-to-stick-mode/
* https://stephenmonro.wordpress.com/2019/02/13/getting-sms-messages-from-the-huawei-e3372-lte-modem/
* https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
* https://core.telegram.org/bots/webhooks
* https://core.telegram.org/bots/webhooks#how-do-i-set-a-webhook-for-either-type
