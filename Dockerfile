FROM golang:1.22

WORKDIR /app

COPY . .

ARG FILE_NAME=main.go
ARG APP_NAME=myapp

RUN go build -o $APP_NAME "$FILE_NAME"

CMD ["./myapp"]
