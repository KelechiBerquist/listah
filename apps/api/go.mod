module cornucopia/listah/apps/api

go 1.22.4

replace cornucopia/listah/internal => ../../internal

require (
	cornucopia/listah/internal v1.0.0
	google.golang.org/grpc v1.65.0
)

require (
	connectrpc.com/connect v1.16.2 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)