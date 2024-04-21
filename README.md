# sms-gateway

## Error Codes:
```
ERROR_UNKNOWN	100001
ERROR_NOT_SUPPORT	100002
ERROR_NO_RIGHT	100003
ERROR_BUSY	100004
ERROR_FORMAT_ERROR	100005
ERROR_PARAMETER_ERROR	100006
ERROR_SAVE_CONFIG_FILE_ERROR	100007
ERROR_GET_CONFIG_FILE_ERROR	100008
ERROR_NO_SIM_CARD_OR_INVALID_SIM_CARD	101001
ERROR_CHECK_SIM_CARD_PIN_LOCK	101002
ERROR_CHECK_SIM_CARD_PUN_LOCK	101003
ERROR_CHECK_SIM_CARD_CAN_UNUSEABLE	101004
ERROR_ENABLE_PIN_FAILED	101005
ERROR_DISABLE_PIN_FAILED	101006
ERROR_UNLOCK_PIN_FAILED	101007
ERROR_DISABLE_AUTO_PIN_FAILED	101008
ERROR_ENABLE_AUTO_PIN_FAILED	101009
ERROR_GET_NET_TYPE_FAILED	102001
ERROR_GET_SERVICE_STATUS_FAILED	102002
ERROR_GET_ROAM_STATUS_FAILED	102003
ERROR_GET_CONNECT_STATUS_FAILED	102004
ERROR_DEVICE_AT_EXECUTE_FAILED	103001
ERROR_DEVICE_PIN_VALIDATE_FAILED	103002
ERROR_DEVICE_PIN_MODIFFY_FAILED	103003
ERROR_DEVICE_PUK_MODIFFY_FAILED	103004
ERROR_DEVICE_GET_AUTORUN_VERSION_FAILED	103005
ERROR_DEVICE_GET_API_VERSION_FAILED	103006
ERROR_DEVICE_GET_PRODUCT_INFORMATON_FAILED	103007
ERROR_DEVICE_SIM_CARD_BUSY	103008
ERROR_DEVICE_SIM_LOCK_INPUT_ERROR	103009
ERROR_DEVICE_NOT_SUPPORT_REMOTE_OPERATE	103010
ERROR_DEVICE_PUK_DEAD_LOCK	103011
ERROR_DEVICE_GET_PC_AISSST_INFORMATION_FAILED	103012
ERROR_DEVICE_SET_LOG_INFORMATON_LEVEL_FAILED	103013
ERROR_DEVICE_GET_LOG_INFORMATON_LEVEL_FAILED	103014
ERROR_DEVICE_COMPRESS_LOG_FILE_FAILED	103015
ERROR_DEVICE_RESTORE_FILE_DECRYPT_FAILED	103016
ERROR_DEVICE_RESTORE_FILE_VERSION_MATCH_FAILED	103017
ERROR_DEVICE_RESTORE_FILE_FAILED	103018
ERROR_DEVICE_SET_TIME_FAILED	103101
ERROR_COMPRESS_LOG_FILE_FAILED	103102
ERROR_DHCP_ERROR	104001
ERROR_SAFE_ERROR	106001
ERROR_DIALUP_GET_CONNECT_FILE_ERROR	107720
ERROR_DIALUP_SET_CONNECT_FILE_ERROR	107721
ERROR_DIALUP_DIALUP_MANAGMENT_PARSE_ERROR	107722
ERROR_DIALUP_ADD_PRORILE_ERROR	107724
ERROR_DIALUP_MODIFY_PRORILE_ERROR	107725
ERROR_DIALUP_SET_DEFAULT_PRORILE_ERROR	107726
ERROR_DIALUP_GET_PRORILE_LIST_ERROR	107727
ERROR_DIALUP_GET_AUTO_APN_MATCH_ERROR	107728
ERROR_DIALUP_SET_AUTO_APN_MATCH_ERROR	107729
ERROR_LOGIN_NO_EXIST_USER	108001
ERROR_LOGIN_PASSWORD_ERROR	108002
ERROR_LOGIN_ALREADY_LOGINED	108003
ERROR_LOGIN_MODIFY_PASSWORD_FAILED	108004
ERROR_LOGIN_TOO_MANY_USERS_LOGINED	108005
ERROR_LOGIN_USERNAME_OR_PASSWORD_ERROR	108006
ERROR_LOGIN_TOO_MANY_TIMES	108007
ERROR_LANGUAGE_GET_FAILED	109001
ERROR_LANGUAGE_SET_FAILED	109002
ERROR_ONLINE_UPDATE_SERVER_NOT_ACCESSED	110001
ERROR_ONLINE_UPDATE_ALREADY_BOOTED	110002
ERROR_ONLINE_UPDATE_GET_DEVICE_INFORMATION_FAILED	110003
ERROR_ONLINE_UPDATE_GET_LOCAL_GROUP_COMMPONENT_INFORMATION_FAILED	110004
ERROR_ONLINE_UPDATE_NOT_FIND_FILE_ON_SERVER	110005
ERROR_ONLINE_UPDATE_NEED_RECONNECT_SERVER	110006
ERROR_ONLINE_UPDATE_CANCEL_DOWNLODING	110007
ERROR_ONLINE_UPDATE_SAME_FILE_LIST	110008
ERROR_ONLINE_UPDATE_CONNECT_ERROR	110009
ERROR_ONLINE_UPDATE_INVALID_URL_LIST	110021
ERROR_ONLINE_UPDATE_NOT_SUPPORT_URL_LIST	110022
ERROR_ONLINE_UPDATE_NOT_BOOT	110023
ERROR_ONLINE_UPDATE_LOW_BATTERY	110024
ERROR_USSD_ERROR	111001
ERROR_USSD_FUCNTION_RETURN_ERROR	111012
ERROR_USSD_IN_USSD_SESSION	111013
ERROR_USSD_TOO_LONG_CONTENT	111014
ERROR_USSD_EMPTY_COMMAND	111016
ERROR_USSD_CODING_ERROR	111017
ERROR_USSD_AT_SEND_FAILED	111018
ERROR_USSD_NET_NO_RETURN	111019
ERROR_USSD_NET_OVERTIME	111020
ERROR_USSD_XML_SPECIAL_CHARACTER_TRANSFER_FAILED	111021
ERROR_USSD_NET_NOT_SUPPORT_USSD	111022
ERROR_SET_NET_MODE_AND_BAND_WHEN_DAILUP_FAILED	112001
ERROR_SET_NET_SEARCH_MODE_WHEN_DAILUP_FAILED	112002
ERROR_SET_NET_MODE_AND_BAND_FAILED	112003
ERROR_SET_NET_SEARCH_MODE_FAILED	112004
ERROR_NET_REGISTER_NET_FAILED	112005
ERROR_NET_NET_CONNECTED_ORDER_NOT_MATCH	112006
ERROR_NET_CURRENT_NET_MODE_NOT_SUPPORT	112007
ERROR_NET_SIM_CARD_NOT_READY_STATUS	112008
ERROR_NET_MEMORY_ALLOC_FAILED	112009
ERROR_SMS_NULL_ARGUMENT_OR_ILLEGAL_ARGUMENT	113017
ERROR_SMS_OVERTIME	113018
ERROR_SMS_QUERY_SMS_INDEX_LIST_ERROR	113020
ERROR_SMS_SET_SMS_CENTER_NUMBER_FAILED	113031
ERROR_SMS_DELETE_SMS_FAILED	113036
ERROR_SMS_SAVE_CONFIG_FILE_FAILED	113047
ERROR_SMS_LOCAL_SPACE_NOT_ENOUGH	113053
ERROR_SMS_TELEPHONE_NUMBER_TOO_LONG	113054
ERROR_SD_FILE_EXIST	114001
ERROR_SD_DIRECTORY_EXIST	114002
ERROR_SD_FILE_OR_DIRECTORY_NOT_EXIST	114004
ERROR_SD_IS_OPERTED_BY_OTHER_USER	114004
ERROR_SD_FILE_NAME_TOO_LONG	114005
ERROR_SD_NO_RIGHT	114006
ERROR_SD_FILE_IS_UPLOADING	114007
ERROR_PB_NULL_ARGUMENT_OR_ILLEGAL_ARGUMENT	115001
ERROR_PB_OVERTIME	115002
ERROR_PB_CALL_SYSTEM_FUCNTION_ERROR	115003
ERROR_PB_WRITE_FILE_ERROR	115004
ERROR_PB_READ_FILE_ERROR	115005
ERROR_PB_LOCAL_TELEPHONE_FULL_ERROR	115199
ERROR_STK_NULL_ARGUMENT_OR_ILLEGAL_ARGUMENT	116001
ERROR_STK_OVERTIME	116002
ERROR_STK_CALL_SYSTEM_FUCNTION_ERROR	116003
ERROR_STK_WRITE_FILE_ERROR	116004
ERROR_STK_READ_FILE_ERROR	116005
ERROR_WIFI_STATION_CONNECT_AP_PASSWORD_ERROR	117001
ERROR_WIFI_WEB_PASSWORD_OR_DHCP_OVERTIME_ERROR	117002
ERROR_WIFI_PBC_CONNECT_FAILED	117003
ERROR_WIFI_STATION_CONNECT_AP_WISPR_PASSWORD_ERROR	117004
ERROR_CRADLE_GET_CRURRENT_CONNECTED_USER_IP_FAILED	118001
ERROR_CRADLE_GET_CRURRENT_CONNECTED_USER_MAC_FAILED	118002
ERROR_CRADLE_SET_MAC_FAILED	118003
ERROR_CRADLE_GET_WAN_INFORMATION_FAILED	118004
ERROR_CRADLE_CODING_FAILED	118005
ERROR_CRADLE_UPDATE_PROFILE_FAILED	118006
```
## References:
https://www.0xf8.org/2017/01/flashing-a-huawei-e3372h-4g-lte-stick-from-hilink-to-stick-mode/
https://stephenmonro.wordpress.com/2019/02/13/getting-sms-messages-from-the-huawei-e3372-lte-modem/
https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
https://core.telegram.org/bots/webhooks
https://core.telegram.org/bots/webhooks#how-do-i-set-a-webhook-for-either-type
