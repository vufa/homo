package homo

import (
	_ "github.com/gogo/protobuf/gogoproto"
)

//go:generate protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. function.proto
//go:generate protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. api/kv.proto
