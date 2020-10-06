FROM golang
WORKDIR /app
COPY . .
RUN go get -d && \
    go test && \
    go build ./main.go
