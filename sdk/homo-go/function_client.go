package homo

import "google.golang.org/grpc"

// FClient client of functions server
type FClient struct {
	cli  FunctionClient
	cfg  FunctionClientConfig
	conn *grpc.ClientConn
}
