FROM golang:1.11-stretch as builder

RUN mkdir -p /go/src/github.com/raushan-0822/contacts

COPY . /go/src/github.com/raushan-0822/contacts/

WORKDIR /go/src/github.com/raushan-0822/contacts/

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

RUN adduser -S -D -H -h /go/src/github.com/raushan-0822/contacts/ user
USER user

#service
FROM alpine:3.8
RUN apk --no-cache add ca-certificates supervisor

#Configure
ADD etc/supervisord.conf /etc/
RUN supervisord --nodaemon --configuration /etc/supervisord.conf

WORKDIR /app/

COPY --from=builder /go/src/github.com/raushan-0822/contacts/app .
COPY --from=builder /go/src/github.com/raushan-0822/contacts/config.json .

RUN app /usr/local/bin/

CMD ["./usr/local/bin/app"]
