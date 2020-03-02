// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protostructure.proto

package protostructure

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Struct represents a struct type.
//
// This has the following limitations:
//
//   * Circular references are not allowed between any struct types
//   * Embedded structs are not supported
//   * Methods are not preserved
//
type Struct struct {
	// fields is the list of fields in the struct
	Fields               []*Struct_Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Struct) Reset()         { *m = Struct{} }
func (m *Struct) String() string { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()    {}
func (*Struct) Descriptor() ([]byte, []int) {
	return fileDescriptor_protostructure_6518de6f1f7e1710, []int{0}
}
func (m *Struct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Struct.Unmarshal(m, b)
}
func (m *Struct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Struct.Marshal(b, m, deterministic)
}
func (dst *Struct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Struct.Merge(dst, src)
}
func (m *Struct) XXX_Size() int {
	return xxx_messageInfo_Struct.Size(m)
}
func (m *Struct) XXX_DiscardUnknown() {
	xxx_messageInfo_Struct.DiscardUnknown(m)
}

var xxx_messageInfo_Struct proto.InternalMessageInfo

func (m *Struct) GetFields() []*Struct_Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Field is a field type. See reflect.StructField in the Go stdlib
// since the fields in this message match that almost exactly.
type Struct_Field struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	PkgPath              string   `protobuf:"bytes,2,opt,name=PkgPath,proto3" json:"PkgPath,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
	Type                 *Type    `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Struct_Field) Reset()         { *m = Struct_Field{} }
func (m *Struct_Field) String() string { return proto.CompactTextString(m) }
func (*Struct_Field) ProtoMessage()    {}
func (*Struct_Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_protostructure_6518de6f1f7e1710, []int{0, 0}
}
func (m *Struct_Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Struct_Field.Unmarshal(m, b)
}
func (m *Struct_Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Struct_Field.Marshal(b, m, deterministic)
}
func (dst *Struct_Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Struct_Field.Merge(dst, src)
}
func (m *Struct_Field) XXX_Size() int {
	return xxx_messageInfo_Struct_Field.Size(m)
}
func (m *Struct_Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Struct_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Struct_Field proto.InternalMessageInfo

func (m *Struct_Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Struct_Field) GetPkgPath() string {
	if m != nil {
		return m.PkgPath
	}
	return ""
}

func (m *Struct_Field) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Struct_Field) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// Type represents a Go type.
type Type struct {
	// Types that are valid to be assigned to Type:
	//	*Type_Primitive
	//	*Type_Container
	//	*Type_Struct
	Type                 isType_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_protostructure_6518de6f1f7e1710, []int{1}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (dst *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(dst, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

type isType_Type interface {
	isType_Type()
}

type Type_Primitive struct {
	Primitive *Primitive `protobuf:"bytes,1,opt,name=primitive,proto3,oneof"`
}

type Type_Container struct {
	Container *Container `protobuf:"bytes,2,opt,name=container,proto3,oneof"`
}

type Type_Struct struct {
	Struct *Struct `protobuf:"bytes,3,opt,name=struct,proto3,oneof"`
}

func (*Type_Primitive) isType_Type() {}

func (*Type_Container) isType_Type() {}

func (*Type_Struct) isType_Type() {}

func (m *Type) GetType() isType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type) GetPrimitive() *Primitive {
	if x, ok := m.GetType().(*Type_Primitive); ok {
		return x.Primitive
	}
	return nil
}

func (m *Type) GetContainer() *Container {
	if x, ok := m.GetType().(*Type_Container); ok {
		return x.Container
	}
	return nil
}

func (m *Type) GetStruct() *Struct {
	if x, ok := m.GetType().(*Type_Struct); ok {
		return x.Struct
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Type) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Type_OneofMarshaler, _Type_OneofUnmarshaler, _Type_OneofSizer, []interface{}{
		(*Type_Primitive)(nil),
		(*Type_Container)(nil),
		(*Type_Struct)(nil),
	}
}

func _Type_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Type)
	// type
	switch x := m.Type.(type) {
	case *Type_Primitive:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Primitive); err != nil {
			return err
		}
	case *Type_Container:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case *Type_Struct:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Struct); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Type.Type has unexpected type %T", x)
	}
	return nil
}

func _Type_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Type)
	switch tag {
	case 1: // type.primitive
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Primitive)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Primitive{msg}
		return true, err
	case 2: // type.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Container)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Container{msg}
		return true, err
	case 3: // type.struct
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Struct)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Struct{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Type_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Type)
	// type
	switch x := m.Type.(type) {
	case *Type_Primitive:
		s := proto.Size(x.Primitive)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Container:
		s := proto.Size(x.Container)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Struct:
		s := proto.Size(x.Struct)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Primitive is a primitive type such as int, bool, etc.
type Primitive struct {
	// kind is the reflect.Kind value for this primitive. This MUST be
	// a primitive value. For example, reflect.Ptr would be invalid here.
	Kind                 uint32   `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Primitive) Reset()         { *m = Primitive{} }
func (m *Primitive) String() string { return proto.CompactTextString(m) }
func (*Primitive) ProtoMessage()    {}
func (*Primitive) Descriptor() ([]byte, []int) {
	return fileDescriptor_protostructure_6518de6f1f7e1710, []int{2}
}
func (m *Primitive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Primitive.Unmarshal(m, b)
}
func (m *Primitive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Primitive.Marshal(b, m, deterministic)
}
func (dst *Primitive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Primitive.Merge(dst, src)
}
func (m *Primitive) XXX_Size() int {
	return xxx_messageInfo_Primitive.Size(m)
}
func (m *Primitive) XXX_DiscardUnknown() {
	xxx_messageInfo_Primitive.DiscardUnknown(m)
}

var xxx_messageInfo_Primitive proto.InternalMessageInfo

func (m *Primitive) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

// Container represents any "container" type such as a sliec, array, map, etc.
type Container struct {
	// kind must be one of: array, map, ptr, slice
	Kind uint32 `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// elem is the type of the element of this container
	Elem *Type `protobuf:"bytes,2,opt,name=elem,proto3" json:"elem,omitempty"`
	// key is the type of the key, only if kind == map
	Key *Type `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// count is the number of elements, only if kind == array
	Count                int32    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_protostructure_6518de6f1f7e1710, []int{3}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *Container) GetElem() *Type {
	if m != nil {
		return m.Elem
	}
	return nil
}

func (m *Container) GetKey() *Type {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Container) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Struct)(nil), "protostructure.Struct")
	proto.RegisterType((*Struct_Field)(nil), "protostructure.Struct.Field")
	proto.RegisterType((*Type)(nil), "protostructure.Type")
	proto.RegisterType((*Primitive)(nil), "protostructure.Primitive")
	proto.RegisterType((*Container)(nil), "protostructure.Container")
}

func init() {
	proto.RegisterFile("protostructure.proto", fileDescriptor_protostructure_6518de6f1f7e1710)
}

var fileDescriptor_protostructure_6518de6f1f7e1710 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x27, 0xb6, 0x53, 0xe9, 0x2d, 0x8a, 0x84, 0x22, 0x51, 0x04, 0x4b, 0x17, 0xd2, 0x55,
	0x91, 0xea, 0xc6, 0xad, 0x82, 0xcc, 0x4a, 0x4a, 0x9c, 0x17, 0xa8, 0x9d, 0x38, 0x86, 0x4e, 0x7f,
	0xec, 0xa4, 0x42, 0x9f, 0xc0, 0xe7, 0x71, 0xe3, 0xf3, 0x49, 0x6e, 0x7f, 0xc4, 0x52, 0x77, 0xf7,
	0x5c, 0xbe, 0x93, 0x7b, 0x72, 0xc0, 0xad, 0xea, 0x52, 0x95, 0x7b, 0x55, 0x37, 0xa9, 0x6a, 0x6a,
	0x11, 0xa2, 0xa4, 0xc7, 0x7f, 0xb7, 0xfe, 0x17, 0x01, 0xeb, 0x19, 0x15, 0xbd, 0x05, 0xeb, 0x55,
	0x8a, 0xdd, 0x66, 0xcf, 0x88, 0x67, 0x04, 0x4e, 0x74, 0x11, 0x4e, 0x5e, 0xe8, 0xb8, 0xf0, 0x51,
	0x43, 0xbc, 0x67, 0xcf, 0xdf, 0x61, 0x89, 0x0b, 0x4a, 0xc1, 0x7c, 0x4a, 0x72, 0xc1, 0x88, 0x47,
	0x02, 0x9b, 0xe3, 0x4c, 0x19, 0x1c, 0xc6, 0xd9, 0x36, 0x4e, 0xd4, 0x1b, 0x3b, 0xc0, 0xf5, 0x20,
	0xe9, 0x09, 0x18, 0xeb, 0x64, 0xcb, 0x0c, 0xdc, 0xea, 0x91, 0x06, 0x60, 0xaa, 0xb6, 0x12, 0xcc,
	0xf4, 0x48, 0xe0, 0x44, 0xee, 0xf4, 0xf8, 0xba, 0xad, 0x04, 0x47, 0xc2, 0xff, 0x26, 0x60, 0x6a,
	0x49, 0xef, 0xc0, 0xae, 0x6a, 0x99, 0x4b, 0x25, 0x3f, 0xba, 0xbb, 0x4e, 0x74, 0x36, 0xf5, 0xc5,
	0x03, 0xb0, 0x5a, 0xf0, 0x5f, 0x5a, 0x5b, 0xd3, 0xb2, 0x50, 0x89, 0x2c, 0x44, 0x8d, 0xd9, 0x66,
	0xac, 0x0f, 0x03, 0xa0, 0xad, 0x23, 0x4d, 0xaf, 0xc1, 0xea, 0x18, 0x4c, 0xef, 0x44, 0xa7, 0xf3,
	0x3d, 0xad, 0x16, 0xbc, 0xe7, 0xee, 0xad, 0xee, 0x6b, 0xfe, 0x25, 0xd8, 0x63, 0x1c, 0xdd, 0x57,
	0x26, 0x8b, 0x0d, 0xe6, 0x3e, 0xe2, 0x38, 0xfb, 0x9f, 0x04, 0xec, 0xf1, 0xea, 0x1c, 0xa1, 0x5b,
	0x12, 0x3b, 0x91, 0xf7, 0x91, 0xff, 0x69, 0x49, 0x13, 0xf4, 0x0a, 0x8c, 0x4c, 0xb4, 0x7d, 0xc6,
	0x79, 0x50, 0x03, 0xd4, 0x85, 0x65, 0x5a, 0x36, 0x85, 0xc2, 0xe2, 0x97, 0xbc, 0x13, 0x2f, 0x16,
	0xf2, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x7e, 0x7c, 0x3a, 0x46, 0x02, 0x00, 0x00,
}