// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// IngressPolicy is the name of the policy which incoming Broadcast messages are filtered against
type IngressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IngressPolicy) Reset()                    { *m = IngressPolicy{} }
func (m *IngressPolicy) String() string            { return proto.CompactTextString(m) }
func (*IngressPolicy) ProtoMessage()               {}
func (*IngressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// IngressPolicy is the name of the policy which incoming Deliver messages are filtered against
type EgressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *EgressPolicy) Reset()                    { *m = EgressPolicy{} }
func (m *EgressPolicy) String() string            { return proto.CompactTextString(m) }
func (*EgressPolicy) ProtoMessage()               {}
func (*EgressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation
	// policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*IngressPolicy)(nil), "orderer.IngressPolicy")
	proto.RegisterType((*EgressPolicy)(nil), "orderer.EgressPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x96, 0xd6, 0x0e, 0xad, 0xc2, 0x0a, 0x12, 0xea, 0xa5, 0xc4, 0x4b, 0x40, 0x69,
	0x0e, 0xe2, 0x5d, 0x1a, 0x3c, 0x88, 0x08, 0x12, 0x7b, 0xf2, 0xb6, 0x49, 0xa7, 0xc9, 0xd2, 0x66,
	0x27, 0xcc, 0x6e, 0xc0, 0xf8, 0xeb, 0x25, 0x9b, 0x6d, 0x0f, 0x5e, 0x3c, 0xe5, 0x7d, 0x33, 0x2f,
	0xb3, 0xf3, 0x18, 0xb8, 0x25, 0xde, 0x22, 0x23, 0x27, 0x05, 0xe9, 0x9d, 0x2a, 0x5b, 0x96, 0x56,
	0x91, 0x5e, 0x35, 0x4c, 0x96, 0xc4, 0xc4, 0x37, 0x17, 0xd7, 0x05, 0xd5, 0x35, 0xe9, 0x64, 0xf8,
	0x0c, 0xdd, 0xe8, 0x0e, 0xe6, 0x29, 0x69, 0x83, 0xda, 0xb4, 0x66, 0xd3, 0x35, 0x28, 0x04, 0x8c,
	0x6c, 0xd7, 0x60, 0x18, 0x2c, 0x83, 0x78, 0x9a, 0x39, 0x1d, 0x3d, 0xc1, 0x74, 0x2d, 0x6d, 0x51,
	0x7d, 0xaa, 0x1f, 0x14, 0x31, 0x5c, 0xd5, 0xf2, 0xfb, 0x1d, 0x8d, 0x91, 0x25, 0xa6, 0xd4, 0x6a,
	0xeb, 0xbc, 0xf3, 0xec, 0x6f, 0x39, 0x8a, 0x61, 0xe6, 0x7e, 0xdb, 0xa8, 0x1a, 0xa9, 0xb5, 0x22,
	0x84, 0x89, 0x1d, 0xa4, 0x9f, 0x7e, 0xc4, 0xe8, 0x19, 0x2e, 0x53, 0x46, 0xb7, 0xf5, 0x07, 0x1d,
	0x54, 0xd1, 0x89, 0x1b, 0x18, 0x37, 0x4e, 0x79, 0xab, 0xa7, 0xbe, 0xbe, 0x55, 0x25, 0x1a, 0x1b,
	0x9e, 0x2d, 0x83, 0x78, 0x96, 0x79, 0xea, 0x73, 0xbc, 0xea, 0x92, 0xd1, 0x18, 0x3f, 0x40, 0xc0,
	0x48, 0xcb, 0xfa, 0x94, 0xa3, 0xd7, 0x51, 0x04, 0xb3, 0x97, 0xff, 0x3c, 0xf7, 0x30, 0x4f, 0x2b,
	0xa9, 0xb4, 0xdb, 0x87, 0xd8, 0x88, 0x05, 0x5c, 0xb8, 0xb7, 0x15, 0x9a, 0x30, 0x58, 0x9e, 0xc7,
	0xd3, 0xec, 0xc4, 0x7d, 0xc2, 0x37, 0xb9, 0xdb, 0xcb, 0x35, 0xd3, 0x1e, 0xd9, 0xf4, 0x09, 0xf3,
	0x41, 0x7a, 0xeb, 0x11, 0xd7, 0xab, 0xaf, 0x87, 0x52, 0xd9, 0xaa, 0xcd, 0x57, 0x05, 0xd5, 0x49,
	0xd5, 0x35, 0xc8, 0x07, 0xdc, 0x96, 0xc8, 0xc9, 0x4e, 0xe6, 0xac, 0x8a, 0xc4, 0x9d, 0xc3, 0x24,
	0xfe, 0x58, 0xf9, 0xd8, 0xf1, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x31, 0x29, 0x3f,
	0xdb, 0x01, 0x00, 0x00,
}
