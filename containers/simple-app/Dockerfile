FROM golang:latest as build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w"

FROM scratch

WORKDIR /app

COPY --from=build /build/simple-app .

EXPOSE 8000

CMD ["/app/simple-app"]