FROM golang:alpine

WORKDIR /App

COPY go.mod .

RUN go mod tidy

COPY . .

RUN go build cmd/main.go

EXPOSE 8080

CMD [ "./main" ]