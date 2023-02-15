FROM golang:1.20-alpine as base

WORKDIR /go/src/dev

# Download necessary Go modules into docker container
COPY go.mod .
COPY go.sum .
RUN go mod download


FROM base as builder

COPY . .
RUN go build -o ./app-bin
# Outputs the compiled file as 'app-bin' in '/go/src/dev/app-bin'


FROM alpine:3.10 as runtime

WORKDIR /

COPY --from=builder /go/src/dev/.env .
COPY --from=builder /go/src/dev/app-bin .

EXPOSE 8080

CMD ["./app-bin"]
