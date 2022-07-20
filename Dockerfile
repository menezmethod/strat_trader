FROM golang:1.18-alpine

RUN apk add --no-cache git && go install github.com/cespare/reflex@latest

WORKDIR /strat_trader
COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY configs/reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]