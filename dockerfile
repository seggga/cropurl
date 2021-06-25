FROM golang:1.15 as builder
RUN mkdir -p /cropurl
ADD . /cropurl
WORKDIR /cropurl
# Собираем бинарный файл
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -o /cropurl ./cmd/cropurl

FROM scratch
COPY --from=builder /cropurl /cropurl
CMD ["/cropurl"]
