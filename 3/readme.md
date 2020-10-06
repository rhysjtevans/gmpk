# How-to:

## Test and build using docker
A docker image will be built if all the go tests pass.
```
docker build -t kpmg3 .
```
Run normally
```
echo '{"a":{"b":{"c":"d"}}}' | docker run -i kpmg3 /app/main -key a.b.c
```
Run with specific key
```
echo '{"a":{"b":{"c":"d"}}}' | docker run -i kpmg3 /app/main -key a.b.c
```
## Run natively with go

```
go get -d
echo '{"a":{"b":{"c":"d"}}}' | go run ./main.go
```
or
```
go get -d
echo '{"a":{"b":{"c":"d"}}}' | go run ./main.go -key a.b.c
```
