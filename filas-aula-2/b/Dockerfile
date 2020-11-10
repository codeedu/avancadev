FROM golang:1.15

WORKDIR /go/src/msb

COPY . .

RUN go mod download

RUN GOOS=linux go build

ENTRYPOINT "./msb"