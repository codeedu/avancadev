FROM golang:1.15.3

WORKDIR /go/src/fullcycle

COPY . .

RUN GOOS=linux go build b.go

ENTRYPOINT [ "./b" ]