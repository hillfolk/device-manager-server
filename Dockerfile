FROM golang:1.11 as goimage
ENV SRC=/go/src/
ENV GO111MODULE=on
RUN mkdir -p /go/src/
WORKDIR /go/src/github.com/hillfolk/device-manager-server
RUN git clone -b master --single-branch https://github.com/hillfolk/device-manager-server.git /go/src/github.com/hillfolk/device-manager-server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o bin/device-manager-server


FROM alpine:3.9 as baseimagealp
ENV WORK_DIR=/docker/bin
WORKDIR $WORK_DIR
RUN mkdir -p ./data/
COPY --from=goimage /go/src/github.com/hillfolk/device-manager-server/bin/ ./
ENTRYPOINT /docker/bin/device-manager-server run --port=8080 --db=mongo:27017
EXPOSE 8080