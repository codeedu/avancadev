FROM golang:1.15

WORKDIR /go/src/fullcycle

COPY templates .

# Install go dependencies
RUN go get github.com/joho/godotenv && \
  go get github.com/wesleywillians/go-rabbitmq/queue

RUN GOOS=linux go build

ENTRYPOINT [ "./fullcycle" ]