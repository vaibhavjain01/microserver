FROM alpine
ADD hello /hello
ENTRYPOINT [ "/hello" ]
