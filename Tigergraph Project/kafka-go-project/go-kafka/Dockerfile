# add base image
FROM golang:latest

WORKDIR /kafka-go-project

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
EXPOSE 8080

CMD [ "executable" , "/Users/ziruoke/Desktop/Tigergraph Project/kafka-go-project/go-kafka/main/main.go"]