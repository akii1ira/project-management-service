FROM golang:1.25.2

WORKDIR /app

COPY . .

# Собираем под ARM (чтобы совпадало с архитектурой контейнера)
RUN GOOS=linux GOARCH=amd64 go build -o main .


RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
