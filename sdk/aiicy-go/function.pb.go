// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: function.proto

package aiicy

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// FunctionMessage function message
type FunctionMessage struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	QOS                  uint32   `protobuf:"varint,2,opt,name=QOS,proto3" json:"QOS,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Timestamp            int64    `protobuf:"zigzag64,10,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	FunctionName         string   `protobuf:"bytes,11,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`
	FunctionInvokeID     string   `protobuf:"bytes,12,opt,name=FunctionInvokeID,proto3" json:"FunctionInvokeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionMessage) Reset()         { *m = FunctionMessage{} }
func (m *FunctionMessage) String() string { return proto.CompactTextString(m) }
func (*FunctionMessage) ProtoMessage()    {}
func (*FunctionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{0}
}
func (m *FunctionMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FunctionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FunctionMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FunctionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionMessage.Merge(m, src)
}
func (m *FunctionMessage) XXX_Size() int {
	return m.Size()
}
func (m *FunctionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FunctionMessage)(nil), "aiicy.FunctionMessage")
}

func init() { proto.RegisterFile("function.proto", fileDescriptor_8ac74addf543d91a) }

var fileDescriptor_8ac74addf543d91a = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2b, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xcc, 0xcc, 0x4c,
	0xae, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5,
	0x74, 0x91, 0x91, 0x8b, 0xdf, 0x0d, 0x6a, 0x90, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x10,
	0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x93, 0xa7, 0x8b, 0x90,
	0x00, 0x17, 0x73, 0xa0, 0x7f, 0xb0, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x88, 0x29, 0x24,
	0xc2, 0xc5, 0x1a, 0x92, 0x5f, 0x90, 0x99, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1,
	0x08, 0x49, 0x70, 0xb1, 0x07, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48, 0xb0, 0x28, 0x30, 0x6a,
	0xf0, 0x04, 0xc1, 0xb8, 0x42, 0x32, 0x5c, 0x9c, 0x21, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9,
	0x05, 0x12, 0x5c, 0x0a, 0x8c, 0x1a, 0x42, 0x41, 0x08, 0x01, 0x21, 0x25, 0x2e, 0x1e, 0x98, 0x13,
	0xfc, 0x12, 0x73, 0x53, 0x25, 0xb8, 0xc1, 0x86, 0xa2, 0x88, 0x09, 0x69, 0x71, 0x09, 0xc0, 0xf8,
	0x9e, 0x79, 0x65, 0xf9, 0xd9, 0xa9, 0x9e, 0x2e, 0x12, 0x3c, 0x60, 0x75, 0x18, 0xe2, 0x46, 0x2e,
	0x5c, 0x1c, 0x30, 0x31, 0x21, 0x0b, 0x2e, 0x16, 0xe7, 0xc4, 0x9c, 0x1c, 0x21, 0x31, 0x3d, 0x70,
	0xf0, 0xe8, 0xa1, 0xf9, 0x55, 0x0a, 0x87, 0xb8, 0x12, 0x83, 0x93, 0xc2, 0x89, 0x87, 0x72, 0x0c,
	0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x81, 0x47, 0x72,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x12, 0x1b,
	0x38, 0x08, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xfb, 0x9c, 0x5e, 0x8a, 0x01, 0x00,
	0x00,
}

func (this *FunctionMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionMessage)
	if !ok {
		that2, ok := that.(FunctionMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.QOS != that1.QOS {
		return false
	}
	if this.Topic != that1.Topic {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.FunctionInvokeID != that1.FunctionInvokeID {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FunctionClient is the client API for Function service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FunctionClient interface {
	Call(ctx context.Context, in *FunctionMessage, opts ...grpc.CallOption) (*FunctionMessage, error)
}

type functionClient struct {
	cc *grpc.ClientConn
}

func NewFunctionClient(cc *grpc.ClientConn) FunctionClient {
	return &functionClient{cc}
}

func (c *functionClient) Call(ctx context.Context, in *FunctionMessage, opts ...grpc.CallOption) (*FunctionMessage, error) {
	out := new(FunctionMessage)
	err := c.cc.Invoke(ctx, "/aiicy.Function/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunctionServer is the server API for Function service.
type FunctionServer interface {
	Call(context.Context, *FunctionMessage) (*FunctionMessage, error)
}

// UnimplementedFunctionServer can be embedded to have forward compatible implementations.
type UnimplementedFunctionServer struct {
}

func (*UnimplementedFunctionServer) Call(ctx context.Context, req *FunctionMessage) (*FunctionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterFunctionServer(s *grpc.Server, srv FunctionServer) {
	s.RegisterService(&_Function_serviceDesc, srv)
}

func _Function_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aiicy.Function/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Call(ctx, req.(*FunctionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Function_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aiicy.Function",
	HandlerType: (*FunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Function_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "function.proto",
}

func (m *FunctionMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FunctionMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FunctionInvokeID) > 0 {
		i -= len(m.FunctionInvokeID)
		copy(dAtA[i:], m.FunctionInvokeID)
		i = encodeVarintFunction(dAtA, i, uint64(len(m.FunctionInvokeID)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.FunctionName) > 0 {
		i -= len(m.FunctionName)
		copy(dAtA[i:], m.FunctionName)
		i = encodeVarintFunction(dAtA, i, uint64(len(m.FunctionName)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Timestamp != 0 {
		i = encodeVarintFunction(dAtA, i, uint64((uint64(m.Timestamp)<<1)^uint64((m.Timestamp>>63))))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x1a
	}
	if m.QOS != 0 {
		i = encodeVarintFunction(dAtA, i, uint64(m.QOS))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintFunction(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFunction(dAtA []byte, offset int, v uint64) int {
	offset -= sovFunction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedFunctionMessage(r randyFunction, easy bool) *FunctionMessage {
	this := &FunctionMessage{}
	this.ID = uint64(uint64(r.Uint32()))
	this.QOS = uint32(r.Uint32())
	this.Topic = string(randStringFunction(r))
	v1 := r.Intn(100)
	this.Payload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Payload[i] = byte(r.Intn(256))
	}
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	this.FunctionName = string(randStringFunction(r))
	this.FunctionInvokeID = string(randStringFunction(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFunction(r, 13)
	}
	return this
}

type randyFunction interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFunction(r randyFunction) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFunction(r randyFunction) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFunction(r)
	}
	return string(tmps)
}
func randUnrecognizedFunction(r randyFunction, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFunction(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFunction(dAtA []byte, r randyFunction, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFunction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFunction(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *FunctionMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovFunction(uint64(m.ID))
	}
	if m.QOS != 0 {
		n += 1 + sovFunction(uint64(m.QOS))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sozFunction(uint64(m.Timestamp))
	}
	l = len(m.FunctionName)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	l = len(m.FunctionInvokeID)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFunction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFunction(x uint64) (n int) {
	return sovFunction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FunctionMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FunctionMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QOS", wireType)
			}
			m.QOS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QOS |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFunction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFunction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.Timestamp = int64(v)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFunction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunctionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionInvokeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFunction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunctionInvokeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFunction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFunction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFunction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFunction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFunction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFunction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFunction = fmt.Errorf("proto: unexpected end of group")
)
