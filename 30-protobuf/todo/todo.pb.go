// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	Task
*/
package todo

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

type Task struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Done bool   `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4f, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xf4, 0xb8, 0x58, 0x42, 0x12,
	0x8b, 0xb3, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x90, 0x58, 0x4a, 0x7e, 0x5e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10,
	0x98, 0x9d, 0xc4, 0x06, 0xd6, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x28, 0x3c, 0x18,
	0x4a, 0x00, 0x00, 0x00,
}