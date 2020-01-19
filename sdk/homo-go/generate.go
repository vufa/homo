package homo

import (
	_ "github.com/gogo/protobuf/gogoproto"
)

//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=plugins=grpc:. function.proto
//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogo_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:. api/kv.proto

//go:generate ./templates/gen.sh hub
//go:generate ./templates/gen.sh function
