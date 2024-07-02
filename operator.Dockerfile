FROM golang:1.21 as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy && go mod verify

COPY . .

WORKDIR /usr/src/app/operator/cmd
RUN go build -v -o /usr/local/bin/operator ./...

WORKDIR /usr/src/app/cli
RUN go build -v -o /usr/local/bin/cli ./...

FROM debian:latest

COPY --from=build /usr/local/bin/operator /usr/local/bin/operator
COPY --from=build /usr/local/bin/cli /usr/local/bin/cli

RUN apt-get update && apt-get install -y git curl bash
RUN curl -L https://foundry.paradigm.xyz | bash
RUN bash -c 'source /root/.profile && foundryup'
RUN cp -r /root/.foundry/bin/* /usr/local/bin/

ENTRYPOINT ["/bin/sh", "-c"]