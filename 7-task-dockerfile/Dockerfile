# Base image, golang 1.19

FROM golang:1.19
WORKDIR /workspace
COPY . .
RUN go mod download

EXPOSE 8000

ENTRYPOINT ["go","run","main.go"]