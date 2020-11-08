To generate pb.go files
protoc service.proto --go_out=.

For 50000 Requests (100 concurrent users * 500 requests)
Rest - 35 seconds <-> 50 seconds
gRPC - 4 seconds <-> 7 seconds 