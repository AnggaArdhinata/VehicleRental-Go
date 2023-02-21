FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/backend-go

EXPOSE 4040

ENTRYPOINT [ "/app/backend-go" ]
CMD [ "server" ]