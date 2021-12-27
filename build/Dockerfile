FROM golang:1.17-buster as BUILDER

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0
RUN go build -o /go/bin/app

FROM scratch

COPY --from=BUILDER /go/bin/app /go/bin/app

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]
