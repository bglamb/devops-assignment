FROM golang:1.16-alpine
WORKDIR /mockapi
COPY go.mod ./
RUN go mod download
COPY mockapi.go ./
RUN go build -o /mockapi
EXPOSE 8080
COPY entrypoint.sh ./
CMD [ "./entrypoint.sh" ]
