FROM scratch
COPY sms-gateway /
ENTRYPOINT ["/sms-gateway"]