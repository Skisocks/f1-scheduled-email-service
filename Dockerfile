FROM golang:1.17-alpine as Builder

WORKDIR /email-service
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.15.0
WORKDIR /email-service-binary
CMD ["crond", "-f"]
RUN echo "*/1 * * * * /email-service-binary/main" | crontab -
COPY --from=builder /email-service ./