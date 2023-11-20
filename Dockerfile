FROM golang:1.21.4-alpine as builder
ENV GO111MODULE=on
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN mkdir /go/src/Dota-Quiz
WORKDIR /go/src/Dota-Quiz
COPY go.mod ./
RUN go mod download && go mod verify
COPY . ./
RUN go build -o DotaTrivia .
ENTRYPOINT [ "./DotaTrivia" ]