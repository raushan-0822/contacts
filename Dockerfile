FROM golang:1.12.6 as builder
# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*
# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx
# create a working directory
WORKDIR /starterkit

COPY . /starterkit/
# install packages
RUN go mod vendor

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go
# strip and compress the binary
RUN strip --strip-unneeded main
RUN upx main

# use scratch (base for a docker image)
FROM ubuntu
# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /starterkit/main .
COPY --from=builder /starterkit/config.json .
# set env variable
ENV MYSQL_URL="root:heimdall@tcp(172.18.0.2:3306)/contact?charset=utf8&parseTime=True"
# run the binary
CMD ["./main"]
