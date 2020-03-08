FROM golang:1.13

WORKDIR /hex
COPY . ./

WORKDIR /hex/cmd/server
RUN go build -o /hex/run-server

WORKDIR /hex
EXPOSE 8080
ENTRYPOINT /hex/run-server
