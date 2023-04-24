
FROM golang:1.19

EXPOSE 3030

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/app/main.go 


CMD ["./main"]

CMD ["./main"]

