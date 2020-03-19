FROM alpine:3.11

# ssl certs
RUN apk add --update ca-certificates

WORKDIR /usr/app
COPY ./bin .

RUN ls /usr/app

ENTRYPOINT [ "/usr/app/lifehub" ]
