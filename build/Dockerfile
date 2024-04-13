FROM golang:1.21.1

ENV APP_DIR /go/src/app
ENV GO111MODULE=on

RUN mkdir -p $APP_DIR

COPY . $APP_DIR

WORKDIR $APP_DIR

RUN go mod download
RUN go build -o main ./cmd

CMD ["./app", "--storage=$STORAGE_VALUE"]
