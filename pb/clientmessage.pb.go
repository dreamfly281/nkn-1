// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/clientmessage.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

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

type OutboundMessage struct {
	Dest              string   `protobuf:"bytes,1,opt,name=dest,proto3" json:"dest,omitempty"`
	Payload           []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Dests             []string `protobuf:"bytes,3,rep,name=dests" json:"dests,omitempty"`
	MaxHoldingSeconds uint32   `protobuf:"varint,4,opt,name=max_holding_seconds,json=maxHoldingSeconds,proto3" json:"max_holding_seconds,omitempty"`
}

func (m *OutboundMessage) Reset()      { *m = OutboundMessage{} }
func (*OutboundMessage) ProtoMessage() {}
func (*OutboundMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessage_30a1957f28ca9b83, []int{0}
}
func (m *OutboundMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutboundMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutboundMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *OutboundMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutboundMessage.Merge(dst, src)
}
func (m *OutboundMessage) XXX_Size() int {
	return m.Size()
}
func (m *OutboundMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutboundMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OutboundMessage proto.InternalMessageInfo

func (m *OutboundMessage) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

func (m *OutboundMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *OutboundMessage) GetDests() []string {
	if m != nil {
		return m.Dests
	}
	return nil
}

func (m *OutboundMessage) GetMaxHoldingSeconds() uint32 {
	if m != nil {
		return m.MaxHoldingSeconds
	}
	return 0
}

type InboundMessage struct {
	Src     string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *InboundMessage) Reset()      { *m = InboundMessage{} }
func (*InboundMessage) ProtoMessage() {}
func (*InboundMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessage_30a1957f28ca9b83, []int{1}
}
func (m *InboundMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InboundMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InboundMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *InboundMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboundMessage.Merge(dst, src)
}
func (m *InboundMessage) XXX_Size() int {
	return m.Size()
}
func (m *InboundMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InboundMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InboundMessage proto.InternalMessageInfo

func (m *InboundMessage) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *InboundMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*OutboundMessage)(nil), "pb.OutboundMessage")
	proto.RegisterType((*InboundMessage)(nil), "pb.InboundMessage")
}
func (this *OutboundMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutboundMessage)
	if !ok {
		that2, ok := that.(OutboundMessage)
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
	if this.Dest != that1.Dest {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if len(this.Dests) != len(that1.Dests) {
		return false
	}
	for i := range this.Dests {
		if this.Dests[i] != that1.Dests[i] {
			return false
		}
	}
	if this.MaxHoldingSeconds != that1.MaxHoldingSeconds {
		return false
	}
	return true
}
func (this *InboundMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InboundMessage)
	if !ok {
		that2, ok := that.(InboundMessage)
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
	if this.Src != that1.Src {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	return true
}
func (this *OutboundMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.OutboundMessage{")
	s = append(s, "Dest: "+fmt.Sprintf("%#v", this.Dest)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Dests: "+fmt.Sprintf("%#v", this.Dests)+",\n")
	s = append(s, "MaxHoldingSeconds: "+fmt.Sprintf("%#v", this.MaxHoldingSeconds)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InboundMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.InboundMessage{")
	s = append(s, "Src: "+fmt.Sprintf("%#v", this.Src)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClientmessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *OutboundMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutboundMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dest) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClientmessage(dAtA, i, uint64(len(m.Dest)))
		i += copy(dAtA[i:], m.Dest)
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClientmessage(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if len(m.Dests) > 0 {
		for _, s := range m.Dests {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.MaxHoldingSeconds != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintClientmessage(dAtA, i, uint64(m.MaxHoldingSeconds))
	}
	return i, nil
}

func (m *InboundMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InboundMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Src) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClientmessage(dAtA, i, uint64(len(m.Src)))
		i += copy(dAtA[i:], m.Src)
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClientmessage(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	return i, nil
}

func encodeVarintClientmessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedOutboundMessage(r randyClientmessage, easy bool) *OutboundMessage {
	this := &OutboundMessage{}
	this.Dest = string(randStringClientmessage(r))
	v1 := r.Intn(100)
	this.Payload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Payload[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(10)
	this.Dests = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Dests[i] = string(randStringClientmessage(r))
	}
	this.MaxHoldingSeconds = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedInboundMessage(r randyClientmessage, easy bool) *InboundMessage {
	this := &InboundMessage{}
	this.Src = string(randStringClientmessage(r))
	v3 := r.Intn(100)
	this.Payload = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Payload[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyClientmessage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneClientmessage(r randyClientmessage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringClientmessage(r randyClientmessage) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneClientmessage(r)
	}
	return string(tmps)
}
func randUnrecognizedClientmessage(r randyClientmessage, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldClientmessage(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldClientmessage(dAtA []byte, r randyClientmessage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateClientmessage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateClientmessage(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *OutboundMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dest)
	if l > 0 {
		n += 1 + l + sovClientmessage(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovClientmessage(uint64(l))
	}
	if len(m.Dests) > 0 {
		for _, s := range m.Dests {
			l = len(s)
			n += 1 + l + sovClientmessage(uint64(l))
		}
	}
	if m.MaxHoldingSeconds != 0 {
		n += 1 + sovClientmessage(uint64(m.MaxHoldingSeconds))
	}
	return n
}

func (m *InboundMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Src)
	if l > 0 {
		n += 1 + l + sovClientmessage(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovClientmessage(uint64(l))
	}
	return n
}

func sovClientmessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClientmessage(x uint64) (n int) {
	return sovClientmessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *OutboundMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutboundMessage{`,
		`Dest:` + fmt.Sprintf("%v", this.Dest) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Dests:` + fmt.Sprintf("%v", this.Dests) + `,`,
		`MaxHoldingSeconds:` + fmt.Sprintf("%v", this.MaxHoldingSeconds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InboundMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InboundMessage{`,
		`Src:` + fmt.Sprintf("%v", this.Src) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClientmessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *OutboundMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientmessage
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
			return fmt.Errorf("proto: OutboundMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutboundMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClientmessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClientmessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dests", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClientmessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dests = append(m.Dests, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHoldingSeconds", wireType)
			}
			m.MaxHoldingSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxHoldingSeconds |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClientmessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientmessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InboundMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientmessage
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
			return fmt.Errorf("proto: InboundMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InboundMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Src", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClientmessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Src = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClientmessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientmessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientmessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClientmessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientmessage
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
					return 0, ErrIntOverflowClientmessage
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
					return 0, ErrIntOverflowClientmessage
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
				return 0, ErrInvalidLengthClientmessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClientmessage
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
				next, err := skipClientmessage(dAtA[start:])
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
	ErrInvalidLengthClientmessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientmessage   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pb/clientmessage.proto", fileDescriptor_clientmessage_30a1957f28ca9b83)
}

var fileDescriptor_clientmessage_30a1957f28ca9b83 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x4f,
	0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xc9, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x25, 0x95, 0xa6, 0x81, 0x79, 0x60,
	0x0e, 0x98, 0x05, 0xd1, 0xa2, 0xd4, 0xca, 0xc8, 0xc5, 0xef, 0x5f, 0x5a, 0x92, 0x94, 0x5f, 0x9a,
	0x97, 0xe2, 0x0b, 0x31, 0x4c, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb5, 0xb8, 0x44, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0x41, 0x2a, 0x8a,
	0x25, 0x98, 0x15, 0x98, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x3d, 0x2e, 0xe1, 0xdc, 0xc4, 0x8a,
	0xf8, 0x8c, 0xfc, 0x9c, 0x94, 0xcc, 0xbc, 0xf4, 0xf8, 0xe2, 0xd4, 0xe4, 0xfc, 0xbc, 0x94, 0x62,
	0x09, 0x16, 0x05, 0x46, 0x0d, 0xde, 0x20, 0xc1, 0xdc, 0xc4, 0x0a, 0x0f, 0x88, 0x4c, 0x30, 0x44,
	0x42, 0xc9, 0x86, 0x8b, 0xcf, 0x33, 0x0f, 0xc5, 0x15, 0x02, 0x5c, 0xcc, 0xc5, 0x45, 0xc9, 0x50,
	0x47, 0x80, 0x98, 0xb8, 0xdd, 0xe0, 0x64, 0x73, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c,
	0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39,
	0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x41, 0x61, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x62, 0x5f, 0x54, 0x40, 0x57, 0x01, 0x00, 0x00,
}
