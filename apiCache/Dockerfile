FROM golang:alpine

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /app/cmd/beer-server
RUN go build main.go


CMD ["./main"]
