Demo code to showcase basics of gRPC with Rest comparision part of a gRPC talk given by me. 
(slides can be found here https://docs.google.com/presentation/d/1w50uzYdM-Mxvt8dwGYlMEbyBQ4fCZq6gx8JKPjAFtZs/edit?usp=sharing)



To generate pb.go files
protoc service.proto --go_out=.


Performance - 
For 50000 Requests (100 concurrent users * 500 requests)
Rest - 35 seconds <-> ~1 seconds
gRPC - 5 seconds <-> 10 seconds 
