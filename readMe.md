To generate pb.go files
protoc service.proto --go_out=.

For 50000 Requests (100 concurrent users * 500 requests)
Rest - 35 seconds <-> ~1 seconds
gRPC - 5 seconds <-> 10 seconds 