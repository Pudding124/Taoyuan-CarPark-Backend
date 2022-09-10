FROM golang:1.18.5-alpine as base

WORKDIR /app

ENV CGO_ENABLED=0

# Download necessary Go modules
COPY app/go.mod ./
COPY app/go.sum ./
RUN go mod download

COPY app ./

FROM base as building

RUN go build -o /bin/app .

FROM scratch as release

WORKDIR /opt/app

COPY --from=building /bin/app /opt/app/

CMD ["/opt/app/app"]