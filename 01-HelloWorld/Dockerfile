FROM golang:1.17-alpine3.15 as builder

WORKDIR /app

ENV CGO_BUILD = 0

COPY . . 

RUN go build main.go
#alpine / scratch 
FROM alpine

COPY --from=builder app/main app/main 

CMD ["./app/main"]
