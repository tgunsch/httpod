FROM golang:1.16 as build-env


WORKDIR /workdir
COPY go.mod go.mod
COPY go.sum go.sum
COPY . /workdir

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=readonly -a -installsuffix cgo -v -o ./build/bin/httpod github.com/tgunsch/httpod/cmd

FROM scratch as slim

USER 1334

COPY --from=build-env /workdir/build/bin/httpod /usr/bin/httpod

CMD ["httpod"]

FROM ubuntu as ubuntu

RUN apt-get update && apt-get install -y \
  curl \
  && rm -rf /var/lib/apt/lists/*

USER 1334

COPY --from=build-env /workdir/build/bin/httpod /usr/bin/httpod

CMD ["httpod"]