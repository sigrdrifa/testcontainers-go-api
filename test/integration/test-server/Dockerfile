FROM docker.io/golang:latest

COPY . ./

WORKDIR test/integration/test-server
RUN ls

CMD go run server.go
