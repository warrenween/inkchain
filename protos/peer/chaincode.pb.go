// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{5, 0}
}

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ChaincodeID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeSpec) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type SenderSpec struct {
	Sender   []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Counter  uint64 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	InkLimit []byte `protobuf:"bytes,3,opt,name=ink_limit,json=inkLimit,proto3" json:"ink_limit,omitempty"`
	Msg      []byte `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *SenderSpec) Reset()                    { *m = SenderSpec{} }
func (m *SenderSpec) String() string            { return proto.CompactTextString(m) }
func (*SenderSpec) ProtoMessage()               {}
func (*SenderSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SenderSpec) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *SenderSpec) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *SenderSpec) GetInkLimit() []byte {
	if m != nil {
		return m.InkLimit
	}
	return nil
}

func (m *SenderSpec) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type SignContent struct {
	ChaincodeSpec   *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	IdGenerationAlg string         `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
	SenderSpec      *SenderSpec    `protobuf:"bytes,3,opt,name=sender_spec,json=senderSpec" json:"sender_spec,omitempty"`
}

func (m *SignContent) Reset()                    { *m = SignContent{} }
func (m *SignContent) String() string            { return proto.CompactTextString(m) }
func (*SignContent) ProtoMessage()               {}
func (*SignContent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SignContent) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *SignContent) GetIdGenerationAlg() string {
	if m != nil {
		return m.IdGenerationAlg
	}
	return ""
}

func (m *SignContent) GetSenderSpec() *SenderSpec {
	if m != nil {
		return m.SenderSpec
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effective_date,json=effectiveDate" json:"effective_date,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if m != nil {
		return m.CodePackage
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetExecEnv() ChaincodeDeploymentSpec_ExecutionEnvironment {
	if m != nil {
		return m.ExecEnv
	}
	return ChaincodeDeploymentSpec_DOCKER
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string      `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
	SenderSpec      *SenderSpec `protobuf:"bytes,3,opt,name=sender_spec,json=senderSpec" json:"sender_spec,omitempty"`
	Sig             []byte      `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeInvocationSpec) GetIdGenerationAlg() string {
	if m != nil {
		return m.IdGenerationAlg
	}
	return ""
}

func (m *ChaincodeInvocationSpec) GetSenderSpec() *SenderSpec {
	if m != nil {
		return m.SenderSpec
	}
	return nil
}

func (m *ChaincodeInvocationSpec) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*SenderSpec)(nil), "protos.SenderSpec")
	proto.RegisterType((*SignContent)(nil), "protos.SignContent")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6e, 0xeb, 0x44,
	0x10, 0x3e, 0x6e, 0x72, 0xfa, 0x33, 0x4e, 0x82, 0x59, 0xca, 0x21, 0x2a, 0x17, 0xf4, 0x58, 0x5c,
	0x54, 0x15, 0x72, 0x44, 0x5a, 0x71, 0x85, 0x90, 0xd2, 0xd8, 0xad, 0x0c, 0x21, 0xa9, 0x36, 0x69,
	0x25, 0xb8, 0x89, 0x1c, 0x7b, 0xe2, 0xae, 0x62, 0xef, 0x5a, 0xf6, 0x26, 0x6a, 0xde, 0x89, 0x67,
	0xe0, 0x01, 0x78, 0x29, 0xd0, 0xae, 0x63, 0x87, 0xaa, 0xbd, 0x83, 0x73, 0xe5, 0x99, 0xf1, 0xcc,
	0xce, 0xf7, 0x7d, 0x33, 0xbb, 0x70, 0x9a, 0x21, 0xe6, 0xbd, 0xf0, 0x29, 0x60, 0x3c, 0x14, 0x11,
	0x3a, 0x59, 0x2e, 0xa4, 0x20, 0x87, 0xfa, 0x53, 0x9c, 0x7d, 0x13, 0x0b, 0x11, 0x27, 0xd8, 0xd3,
	0xee, 0x62, 0xbd, 0xec, 0x49, 0x96, 0x62, 0x21, 0x83, 0x34, 0x2b, 0x13, 0xed, 0x09, 0x98, 0xc3,
	0xaa, 0xd6, 0x77, 0x09, 0x81, 0x66, 0x16, 0xc8, 0xa7, 0xae, 0x71, 0x6e, 0x5c, 0x9c, 0x50, 0x6d,
	0xab, 0x18, 0x0f, 0x52, 0xec, 0x1e, 0x94, 0x31, 0x65, 0x93, 0x2e, 0x1c, 0x6d, 0x30, 0x2f, 0x98,
	0xe0, 0xdd, 0x86, 0x0e, 0x57, 0xae, 0xfd, 0x2d, 0x74, 0xf6, 0x07, 0xf2, 0x6c, 0x2d, 0x55, 0x7d,
	0x90, 0xc7, 0x45, 0xd7, 0x38, 0x6f, 0x5c, 0xb4, 0xa8, 0xb6, 0xed, 0xbf, 0x0d, 0x68, 0xd7, 0x69,
	0xd3, 0x0c, 0x43, 0xe2, 0x40, 0x53, 0x6e, 0x33, 0xd4, 0x9d, 0x3b, 0xfd, 0xb3, 0x12, 0x5e, 0xe1,
	0xbc, 0x48, 0x72, 0x66, 0xdb, 0x0c, 0xa9, 0xce, 0x23, 0x3f, 0x40, 0xab, 0x26, 0x3d, 0x67, 0x91,
	0x46, 0x67, 0xf6, 0xbf, 0x78, 0x55, 0xe7, 0xbb, 0xd4, 0xac, 0x13, 0xfd, 0x88, 0x7c, 0x07, 0xef,
	0x99, 0x82, 0xa5, 0x71, 0x9b, 0xfd, 0x0f, 0xaf, 0x0b, 0xd4, 0x5f, 0x5a, 0x26, 0x29, 0x9e, 0x4a,
	0x31, 0xb1, 0x96, 0xdd, 0xe6, 0xb9, 0x71, 0xf1, 0x9e, 0x56, 0xae, 0xfd, 0x13, 0x34, 0x15, 0x1a,
	0xd2, 0x86, 0x93, 0x87, 0xb1, 0xeb, 0xdd, 0xfa, 0x63, 0xcf, 0xb5, 0xde, 0x11, 0x80, 0xc3, 0xbb,
	0xc9, 0x68, 0x30, 0xbe, 0xb3, 0x0c, 0x72, 0x0c, 0xcd, 0xf1, 0xc4, 0xf5, 0xac, 0x03, 0x72, 0x04,
	0x8d, 0xe1, 0x80, 0x5a, 0x0d, 0x15, 0xfa, 0x79, 0xf0, 0x38, 0xb0, 0x9a, 0x76, 0x0a, 0x30, 0x45,
	0x1e, 0x61, 0xae, 0xd9, 0x7f, 0x80, 0xc3, 0x42, 0x7b, 0x9a, 0x7f, 0x8b, 0xee, 0x3c, 0xd5, 0x3f,
	0x14, 0x6b, 0x2e, 0x31, 0xd7, 0x04, 0x9b, 0xb4, 0x72, 0xc9, 0xd7, 0x70, 0xc2, 0xf8, 0x6a, 0x9e,
	0xb0, 0x94, 0x95, 0x5c, 0x5a, 0xf4, 0x98, 0xf1, 0xd5, 0x48, 0xf9, 0xc4, 0x82, 0x46, 0x5a, 0xc4,
	0x1a, 0x72, 0x8b, 0x2a, 0xd3, 0xfe, 0xc3, 0x00, 0x73, 0xca, 0x62, 0x3e, 0x14, 0x5c, 0x22, 0x97,
	0xe4, 0x47, 0xe8, 0xec, 0xe5, 0x2b, 0x32, 0x0c, 0x75, 0x63, 0xb3, 0xff, 0xe5, 0x9b, 0xc2, 0xd3,
	0x76, 0xf8, 0x62, 0x58, 0x97, 0xf0, 0x39, 0x8b, 0xe6, 0x31, 0x72, 0xcc, 0x03, 0xc9, 0x04, 0x9f,
	0x07, 0x49, 0xbc, 0xdb, 0x8f, 0xcf, 0x58, 0x74, 0x57, 0xc7, 0x07, 0x49, 0x4c, 0xae, 0xc0, 0x2c,
	0xc9, 0x94, 0x6d, 0x4a, 0xd9, 0x49, 0xd5, 0x66, 0xaf, 0x01, 0x85, 0xa2, 0xb6, 0xed, 0x3f, 0x0f,
	0xe0, 0xab, 0x1a, 0x81, 0x8b, 0x59, 0x22, 0xb6, 0x29, 0x72, 0xa9, 0x9b, 0xff, 0x37, 0xe8, 0x03,
	0xe8, 0xe0, 0x72, 0x89, 0xa1, 0x64, 0x1b, 0x9c, 0x47, 0x81, 0xc4, 0xdd, 0xe6, 0x9c, 0x39, 0xe5,
	0x55, 0x71, 0xaa, 0xab, 0xe2, 0xcc, 0xaa, 0xab, 0x42, 0xdb, 0x75, 0x85, 0x1b, 0x48, 0x24, 0x1f,
	0xa1, 0xa5, 0x7b, 0x67, 0x41, 0xb8, 0x0a, 0x62, 0xdc, 0xa9, 0x6f, 0xaa, 0xd8, 0x7d, 0x19, 0x22,
	0x13, 0x38, 0xc6, 0x67, 0x0c, 0xe7, 0xc8, 0x37, 0x7a, 0x0a, 0x9d, 0xfe, 0xf5, 0x2b, 0x74, 0x2f,
	0x69, 0x39, 0xde, 0x33, 0x86, 0x6b, 0x25, 0x9a, 0xc7, 0x37, 0x2c, 0x17, 0x5c, 0xfd, 0xa0, 0x47,
	0xea, 0x14, 0x8f, 0x6f, 0x6c, 0x07, 0x4e, 0xdf, 0x4a, 0x50, 0xfb, 0xe6, 0x4e, 0x86, 0xbf, 0x78,
	0xb4, 0xdc, 0xbd, 0xe9, 0x6f, 0xd3, 0x99, 0xf7, 0xab, 0x65, 0xd8, 0x7f, 0x19, 0xff, 0x12, 0xd0,
	0xe7, 0x1b, 0x11, 0xea, 0x81, 0xfc, 0x0f, 0x02, 0x7e, 0xea, 0xd9, 0xab, 0xe5, 0x2d, 0x58, 0xbd,
	0xbc, 0x05, 0x8b, 0x2f, 0xaf, 0xe1, 0x74, 0x28, 0xf8, 0x92, 0x45, 0xc8, 0x25, 0x0b, 0x12, 0x26,
	0xb7, 0x23, 0xdc, 0x60, 0xa2, 0x08, 0xdf, 0x3f, 0xdc, 0x8c, 0xfc, 0xa1, 0xf5, 0x8e, 0x58, 0xd0,
	0x1a, 0x4e, 0xc6, 0xb7, 0xbe, 0xeb, 0x8d, 0x67, 0xfe, 0x60, 0x64, 0x19, 0x37, 0x8f, 0xf0, 0x51,
	0xe4, 0xb1, 0xc3, 0xf8, 0x4a, 0x13, 0xd8, 0x1b, 0xbb, 0xf6, 0xea, 0xe1, 0xfc, 0xfd, 0xfb, 0x98,
	0xc9, 0xa7, 0xf5, 0xc2, 0x09, 0x45, 0xda, 0x63, 0x7c, 0x95, 0x04, 0x8b, 0x62, 0x29, 0xd6, 0x3c,
	0xd2, 0x14, 0x7a, 0x55, 0x49, 0xf9, 0x7e, 0x16, 0x3d, 0x55, 0xb2, 0x28, 0xdf, 0xd6, 0xab, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x64, 0xf6, 0x0a, 0x7a, 0x05, 0x00, 0x00,
}
