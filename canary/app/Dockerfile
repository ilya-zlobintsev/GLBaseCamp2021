FROM golang:latest as build

WORKDIR /build

COPY . .

ARG VERSION

RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.VERSION=${VERSION}"

FROM scratch

WORKDIR /app

COPY --from=build /build/app .

EXPOSE 8080

CMD ["/app/app"]