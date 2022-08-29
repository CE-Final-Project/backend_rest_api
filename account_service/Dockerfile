FROM golang:alpine3.16 as build

WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o account_service

FROM scratch as final
COPY --from=build /app/account_service .
ENV TZ=Asia/Bangkok
CMD ["account_service"]