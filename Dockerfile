FROM alpine
ADD wechatApp-service /wechatApp-service
ENTRYPOINT [ "/wechatApp-service" ]
