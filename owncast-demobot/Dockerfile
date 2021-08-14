FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app/owncast-demobot

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/owncast-demobot .

EXPOSE 8100
CMD ["./out/owncast-demobot"]