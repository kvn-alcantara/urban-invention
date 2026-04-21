FROM golang:1.20-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY . .
RUN go build -o /out/app ./cmd/app

FROM alpine:3.20

WORKDIR /app
COPY --from=build /out/app /app/app
EXPOSE 8080
ENV PORT=8080
CMD ["/app/app"]
