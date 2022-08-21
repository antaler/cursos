ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine as builder

RUN go env -w GOPROXY=direct

RUN apk update && apk add --no-cache git

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"  -o /platzi-curso-ws

FROM scratch as runner

ADD  ./.env /

COPY --from=builder /platzi-curso-ws /platzi-curso-ws

EXPOSE 5050


ENTRYPOINT [ "/platzi-curso-ws" ]