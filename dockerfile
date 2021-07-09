# работа с модулями
FROM golang:1.15 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# сборка
FROM golang:1.15 as builder
COPY --from=modules /go/pkg /go/pkg
RUN mkdir -p /src
ADD . /src
WORKDIR /src
# добавляется пользователь без прав root
RUN useradd -u 10001 cropurluser
# сборка бинарного файла
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -o /cropurl ./cmd/cropurl

# запуск приложения
FROM scratch
# копирование данных пользователя
COPY --from=builder /etc/passwd /etc/passwd
USER cropurluser
COPY --from=builder /cropurl /cropurl
CMD ["/cropurl"]

