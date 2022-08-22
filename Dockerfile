FROM golang:1.18.4-alpine
RUN mkdir /apigateway
COPY . /apigateway
WORKDIR /apigateway/cmd
RUN go build -o main .
CMD ["/apigateway/cmd/main"]
