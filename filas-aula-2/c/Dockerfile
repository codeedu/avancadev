FROM golang:1.15

WORKDIR /go/src/msc

COPY . .

RUN go mod download

RUN GOOS=linux go build

EXPOSE 9092

ENTRYPOINT "./msc"