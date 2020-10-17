FROM golang:1.15.2 AS builder
RUN apt-get update && apt-get -y install ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /bin/twitter-action
RUN chmod +x /bin/twitter-action
CMD ["/bin/twitter-action"]

FROM ubuntu:bionic
LABEL "com.github.actions.name"="Simple Tweet Action"
LABEL "com.github.actions.description"="Tweets from the cmdline for a user"
LABEL "com.github.actions.icon"="cloud"
LABEL "com.github.actions.color"="blue"

RUN apt-get update && apt-get -y install ca-certificates

COPY --from=builder /bin/twitter-action /usr/bin/twitter-action
ENTRYPOINT ["/usr/bin/twitter-action"]
