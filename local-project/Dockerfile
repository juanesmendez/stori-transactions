FROM golang:1.22 AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest

ENV PORT=8080
WORKDIR /app
COPY --from=build /app .
CMD ["/app/main"]
