# How-to:

## Test and build
A docker image will be built if all the go tests pass.
```
docker build -t kpmg2 .
```

### You can also run the 
With specific key
```
go run ./main.go -key compute.name
```
or
```
go run ./main.go -key network.interface.0.ipv4.ipAddress.0.privateIpAddress
```
