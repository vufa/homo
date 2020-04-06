package aiicy

import (
	_ "github.com/gogo/protobuf/gogoproto"
)

//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogo_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:. api/kv.proto
