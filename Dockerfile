FROM golang:1.15 as build-env


WORKDIR /workdir
COPY go.mod go.mod
COPY go.sum go.sum
COPY . /workdir

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=readonly -a -installsuffix cgo -v -o ./build/bin/httpod github.com/tgunsch/httpod/cmd

FROM scratch

USER 1334

COPY --from=build-env /workdir/build/bin/httpod /usr/bin/httpod

ENTRYPOINT ["httpod"]
