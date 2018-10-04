// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/metrics/v2/metrics_service.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import _ "github.com/lyft/protoc-gen-validate/validate"
import prometheus "istio.io/gogo-genproto/prometheus"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StreamMetricsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsResponse) Reset()         { *m = StreamMetricsResponse{} }
func (m *StreamMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsResponse) ProtoMessage()    {}
func (*StreamMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_metrics_service_99fb3b9d736e913c, []int{0}
}
func (m *StreamMetricsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamMetricsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsResponse.Merge(dst, src)
}
func (m *StreamMetricsResponse) XXX_Size() int {
	return m.Size()
}
func (m *StreamMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsResponse proto.InternalMessageInfo

type StreamMetricsMessage struct {
	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamMetricsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// A list of metric entries
	EnvoyMetrics         []*prometheus.MetricFamily `protobuf:"bytes,2,rep,name=envoy_metrics,json=envoyMetrics" json:"envoy_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StreamMetricsMessage) Reset()         { *m = StreamMetricsMessage{} }
func (m *StreamMetricsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage) ProtoMessage()    {}
func (*StreamMetricsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_metrics_service_99fb3b9d736e913c, []int{1}
}
func (m *StreamMetricsMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamMetricsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamMetricsMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamMetricsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage.Merge(dst, src)
}
func (m *StreamMetricsMessage) XXX_Size() int {
	return m.Size()
}
func (m *StreamMetricsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage proto.InternalMessageInfo

func (m *StreamMetricsMessage) GetIdentifier() *StreamMetricsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamMetricsMessage) GetEnvoyMetrics() []*prometheus.MetricFamily {
	if m != nil {
		return m.EnvoyMetrics
	}
	return nil
}

type StreamMetricsMessage_Identifier struct {
	// The node sending metrics over the stream.
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamMetricsMessage_Identifier) Reset()         { *m = StreamMetricsMessage_Identifier{} }
func (m *StreamMetricsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsMessage_Identifier) ProtoMessage()    {}
func (*StreamMetricsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_metrics_service_99fb3b9d736e913c, []int{1, 0}
}
func (m *StreamMetricsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamMetricsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamMetricsMessage_Identifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamMetricsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsMessage_Identifier.Merge(dst, src)
}
func (m *StreamMetricsMessage_Identifier) XXX_Size() int {
	return m.Size()
}
func (m *StreamMetricsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsMessage_Identifier proto.InternalMessageInfo

func (m *StreamMetricsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsResponse)(nil), "envoy.service.metrics.v2.StreamMetricsResponse")
	proto.RegisterType((*StreamMetricsMessage)(nil), "envoy.service.metrics.v2.StreamMetricsMessage")
	proto.RegisterType((*StreamMetricsMessage_Identifier)(nil), "envoy.service.metrics.v2.StreamMetricsMessage.Identifier")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MetricsService service

type MetricsServiceClient interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/envoy.service.metrics.v2.MetricsService/StreamMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceStreamMetricsClient{stream}
	return x, nil
}

type MetricsService_StreamMetricsClient interface {
	Send(*StreamMetricsMessage) error
	CloseAndRecv() (*StreamMetricsResponse, error)
	grpc.ClientStream
}

type metricsServiceStreamMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceStreamMetricsClient) Send(m *StreamMetricsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsClient) CloseAndRecv() (*StreamMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MetricsService service

type MetricsServiceServer interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(MetricsService_StreamMetricsServer) error
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).StreamMetrics(&metricsServiceStreamMetricsServer{stream})
}

type MetricsService_StreamMetricsServer interface {
	SendAndClose(*StreamMetricsResponse) error
	Recv() (*StreamMetricsMessage, error)
	grpc.ServerStream
}

type metricsServiceStreamMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceStreamMetricsServer) SendAndClose(m *StreamMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsServer) Recv() (*StreamMetricsMessage, error) {
	m := new(StreamMetricsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.metrics.v2.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/metrics/v2/metrics_service.proto",
}

func (m *StreamMetricsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamMetricsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamMetricsMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamMetricsMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identifier != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetricsService(dAtA, i, uint64(m.Identifier.Size()))
		n1, err := m.Identifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.EnvoyMetrics) > 0 {
		for _, msg := range m.EnvoyMetrics {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetricsService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamMetricsMessage_Identifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamMetricsMessage_Identifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetricsService(dAtA, i, uint64(m.Node.Size()))
		n2, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMetricsService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StreamMetricsResponse) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamMetricsMessage) Size() (n int) {
	var l int
	_ = l
	if m.Identifier != nil {
		l = m.Identifier.Size()
		n += 1 + l + sovMetricsService(uint64(l))
	}
	if len(m.EnvoyMetrics) > 0 {
		for _, e := range m.EnvoyMetrics {
			l = e.Size()
			n += 1 + l + sovMetricsService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamMetricsMessage_Identifier) Size() (n int) {
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovMetricsService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetricsService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetricsService(x uint64) (n int) {
	return sovMetricsService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamMetricsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetricsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamMetricsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamMetricsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetricsService
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
func (m *StreamMetricsMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetricsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamMetricsMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamMetricsMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetricsService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Identifier == nil {
				m.Identifier = &StreamMetricsMessage_Identifier{}
			}
			if err := m.Identifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvoyMetrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetricsService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvoyMetrics = append(m.EnvoyMetrics, &prometheus.MetricFamily{})
			if err := m.EnvoyMetrics[len(m.EnvoyMetrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetricsService
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
func (m *StreamMetricsMessage_Identifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetricsService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Identifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetricsService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetricsService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &core.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetricsService
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
func skipMetricsService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetricsService
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
					return 0, ErrIntOverflowMetricsService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetricsService
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetricsService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetricsService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetricsService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetricsService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetricsService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/metrics/v2/metrics_service.proto", fileDescriptor_metrics_service_99fb3b9d736e913c)
}

var fileDescriptor_metrics_service_99fb3b9d736e913c = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0xbd, 0x54, 0x1d, 0x5e, 0xad, 0x48, 0x54, 0x5a, 0x82, 0x94, 0xd2, 0xa9, 0xd3, 0x0b,
	0xc4, 0x41, 0x5c, 0x0b, 0x2a, 0x0e, 0x75, 0x48, 0x27, 0x5d, 0xca, 0x35, 0x79, 0xea, 0x41, 0x93,
	0x0b, 0xb9, 0x33, 0xd0, 0xd1, 0x45, 0xc4, 0x8f, 0xe4, 0xe4, 0xe8, 0xe8, 0x47, 0x90, 0x6e, 0x7e,
	0x0b, 0x49, 0xee, 0x52, 0x2d, 0x28, 0xe8, 0xf6, 0x92, 0xf7, 0xff, 0xff, 0xde, 0xff, 0xdd, 0x03,
	0xa4, 0xb4, 0x90, 0x73, 0x5f, 0x51, 0x5e, 0x88, 0x88, 0xfc, 0x84, 0x74, 0x2e, 0x22, 0xe5, 0x17,
	0x41, 0x5d, 0x4e, 0x6c, 0x0b, 0xb3, 0x5c, 0x6a, 0xe9, 0x76, 0x2a, 0x3d, 0xd6, 0x3f, 0xad, 0x08,
	0x8b, 0xc0, 0x3b, 0x30, 0x24, 0x9e, 0x89, 0xd2, 0x1d, 0xc9, 0x9c, 0xfc, 0x29, 0x57, 0xd6, 0xe7,
	0xb5, 0x6a, 0xa5, 0xf9, 0x6c, 0x17, 0x7c, 0x26, 0x62, 0xae, 0xc9, 0xaf, 0x0b, 0xd3, 0xe8, 0xb7,
	0x61, 0x7f, 0xac, 0x73, 0xe2, 0xc9, 0xc8, 0xe8, 0x43, 0x52, 0x99, 0x4c, 0x15, 0xf5, 0xef, 0x1d,
	0xd8, 0x5b, 0xe9, 0x8c, 0x48, 0x29, 0x7e, 0x43, 0xee, 0x25, 0x80, 0x88, 0x29, 0xd5, 0xe2, 0x5a,
	0x50, 0xde, 0x61, 0x3d, 0x36, 0x68, 0x06, 0xc7, 0xf8, 0x5b, 0x4c, 0xfc, 0x89, 0x81, 0xe7, 0x4b,
	0x40, 0xf8, 0x0d, 0xe6, 0x9e, 0x41, 0xab, 0xe2, 0x4c, 0xac, 0xbf, 0xe3, 0xf4, 0x1a, 0x83, 0x66,
	0xd0, 0x47, 0x21, 0xcb, 0xb8, 0x09, 0xe9, 0x5b, 0xba, 0x53, 0x18, 0xcd, 0x04, 0xa5, 0x1a, 0x0d,
	0xf3, 0x94, 0x27, 0x62, 0x36, 0x0f, 0xb7, 0x2a, 0xa3, 0x1d, 0xe3, 0x9d, 0x00, 0x7c, 0x8d, 0x70,
	0x8f, 0x60, 0x3d, 0x95, 0x31, 0xd9, 0xac, 0x6d, 0x9b, 0x95, 0x67, 0xa2, 0xcc, 0x57, 0x3e, 0x1c,
	0x5e, 0xc8, 0x98, 0x86, 0xf0, 0xfc, 0xf1, 0xd2, 0xd8, 0x78, 0x62, 0xce, 0x0e, 0x0b, 0x2b, 0x43,
	0xf0, 0xc0, 0x60, 0xdb, 0x22, 0xc7, 0x66, 0x33, 0x57, 0x43, 0x6b, 0x65, 0x23, 0x17, 0xff, 0xb7,
	0xba, 0xe7, 0xff, 0x51, 0xbf, 0x3c, 0xc4, 0xda, 0x80, 0x0d, 0x77, 0x5f, 0x17, 0x5d, 0xf6, 0xb6,
	0xe8, 0xb2, 0xf7, 0x45, 0x97, 0x5d, 0x39, 0x45, 0xf0, 0xc8, 0xd8, 0x74, 0xb3, 0xba, 0xe0, 0xe1,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x33, 0xa2, 0xd2, 0x53, 0x02, 0x00, 0x00,
}
