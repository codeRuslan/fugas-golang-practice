# Base image, golang 1.18

FROM golang:1.18.3
WORKDIR /workspace
COPY . .
RUN go mod download

EXPOSE 8000

ENTRYPOINT ["go","run","main.go"]