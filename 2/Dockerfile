FROM golang as build
COPY . /app
WORKDIR /app
RUN go get -d && \
    go test && \
    go build