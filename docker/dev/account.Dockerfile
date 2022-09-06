FROM golang:alpine3.16 as build

WORKDIR /app

COPY ../../go.mod ./

RUN go mod download

COPY ../../ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o account ./account_service/cmd/main.go


FROM scratch

COPY --from=build /app/account /account
COPY ../../account_service/config /config
COPY ../../account_service/scripts /scripts
EXPOSE 8080

CMD ["./account"]


