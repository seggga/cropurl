# работа с модулями
FROM golang:1.15 as modules
ADD go.mod go.sum /modules/
RUN cd /modules && go mod download

# сборка
FROM golang:1.15 as builder
COPY --from=modules /go/pkg /go/pkg
RUN mkdir -p /src
ADD . /src
WORKDIR /src
#   добавляется пользователь без прав root
RUN useradd -u 10001 cropurluser
#   сборка бинарного файла
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -o cropurl /src/cmd/cropurl/main.go

# запуск приложения
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /src/cropurl /cropurl
USER cropurluser

CMD ["/cropurl"]

