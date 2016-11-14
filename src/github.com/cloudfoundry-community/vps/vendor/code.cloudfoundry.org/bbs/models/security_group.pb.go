// Code generated by protoc-gen-gogo.
// source: security_group.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PortRange struct {
	Start uint32 `protobuf:"varint,1,opt,name=start" json:"start"`
	End   uint32 `protobuf:"varint,2,opt,name=end" json:"end"`
}

func (m *PortRange) Reset()                    { *m = PortRange{} }
func (*PortRange) ProtoMessage()               {}
func (*PortRange) Descriptor() ([]byte, []int) { return fileDescriptorSecurityGroup, []int{0} }

func (m *PortRange) GetStart() uint32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *PortRange) GetEnd() uint32 {
	if m != nil {
		return m.End
	}
	return 0
}

type ICMPInfo struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type"`
	Code int32 `protobuf:"varint,2,opt,name=code" json:"code"`
}

func (m *ICMPInfo) Reset()                    { *m = ICMPInfo{} }
func (*ICMPInfo) ProtoMessage()               {}
func (*ICMPInfo) Descriptor() ([]byte, []int) { return fileDescriptorSecurityGroup, []int{1} }

func (m *ICMPInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ICMPInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type SecurityGroupRule struct {
	Protocol     string     `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	Destinations []string   `protobuf:"bytes,2,rep,name=destinations" json:"destinations,omitempty"`
	Ports        []uint32   `protobuf:"varint,3,rep,name=ports" json:"ports,omitempty"`
	PortRange    *PortRange `protobuf:"bytes,4,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	IcmpInfo     *ICMPInfo  `protobuf:"bytes,5,opt,name=icmp_info,json=icmpInfo" json:"icmp_info,omitempty"`
	Log          bool       `protobuf:"varint,6,opt,name=log" json:"log"`
}

func (m *SecurityGroupRule) Reset()                    { *m = SecurityGroupRule{} }
func (*SecurityGroupRule) ProtoMessage()               {}
func (*SecurityGroupRule) Descriptor() ([]byte, []int) { return fileDescriptorSecurityGroup, []int{2} }

func (m *SecurityGroupRule) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *SecurityGroupRule) GetDestinations() []string {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *SecurityGroupRule) GetPorts() []uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *SecurityGroupRule) GetPortRange() *PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *SecurityGroupRule) GetIcmpInfo() *ICMPInfo {
	if m != nil {
		return m.IcmpInfo
	}
	return nil
}

func (m *SecurityGroupRule) GetLog() bool {
	if m != nil {
		return m.Log
	}
	return false
}

func init() {
	proto.RegisterType((*PortRange)(nil), "models.PortRange")
	proto.RegisterType((*ICMPInfo)(nil), "models.ICMPInfo")
	proto.RegisterType((*SecurityGroupRule)(nil), "models.SecurityGroupRule")
}
func (this *PortRange) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PortRange)
	if !ok {
		that2, ok := that.(PortRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	return true
}
func (this *ICMPInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ICMPInfo)
	if !ok {
		that2, ok := that.(ICMPInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *SecurityGroupRule) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SecurityGroupRule)
	if !ok {
		that2, ok := that.(SecurityGroupRule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if len(this.Destinations) != len(that1.Destinations) {
		return false
	}
	for i := range this.Destinations {
		if this.Destinations[i] != that1.Destinations[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if this.Ports[i] != that1.Ports[i] {
			return false
		}
	}
	if !this.PortRange.Equal(that1.PortRange) {
		return false
	}
	if !this.IcmpInfo.Equal(that1.IcmpInfo) {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	return true
}
func (this *PortRange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.PortRange{")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ICMPInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.ICMPInfo{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SecurityGroupRule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&models.SecurityGroupRule{")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	if this.Destinations != nil {
		s = append(s, "Destinations: "+fmt.Sprintf("%#v", this.Destinations)+",\n")
	}
	if this.Ports != nil {
		s = append(s, "Ports: "+fmt.Sprintf("%#v", this.Ports)+",\n")
	}
	if this.PortRange != nil {
		s = append(s, "PortRange: "+fmt.Sprintf("%#v", this.PortRange)+",\n")
	}
	if this.IcmpInfo != nil {
		s = append(s, "IcmpInfo: "+fmt.Sprintf("%#v", this.IcmpInfo)+",\n")
	}
	s = append(s, "Log: "+fmt.Sprintf("%#v", this.Log)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSecurityGroup(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringSecurityGroup(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *PortRange) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PortRange) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSecurityGroup(data, i, uint64(m.Start))
	data[i] = 0x10
	i++
	i = encodeVarintSecurityGroup(data, i, uint64(m.End))
	return i, nil
}

func (m *ICMPInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ICMPInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSecurityGroup(data, i, uint64(m.Type))
	data[i] = 0x10
	i++
	i = encodeVarintSecurityGroup(data, i, uint64(m.Code))
	return i, nil
}

func (m *SecurityGroupRule) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SecurityGroupRule) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintSecurityGroup(data, i, uint64(len(m.Protocol)))
	i += copy(data[i:], m.Protocol)
	if len(m.Destinations) > 0 {
		for _, s := range m.Destinations {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Ports) > 0 {
		for _, num := range m.Ports {
			data[i] = 0x18
			i++
			i = encodeVarintSecurityGroup(data, i, uint64(num))
		}
	}
	if m.PortRange != nil {
		data[i] = 0x22
		i++
		i = encodeVarintSecurityGroup(data, i, uint64(m.PortRange.Size()))
		n1, err := m.PortRange.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.IcmpInfo != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintSecurityGroup(data, i, uint64(m.IcmpInfo.Size()))
		n2, err := m.IcmpInfo.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	data[i] = 0x30
	i++
	if m.Log {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func encodeFixed64SecurityGroup(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32SecurityGroup(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSecurityGroup(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *PortRange) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSecurityGroup(uint64(m.Start))
	n += 1 + sovSecurityGroup(uint64(m.End))
	return n
}

func (m *ICMPInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSecurityGroup(uint64(m.Type))
	n += 1 + sovSecurityGroup(uint64(m.Code))
	return n
}

func (m *SecurityGroupRule) Size() (n int) {
	var l int
	_ = l
	l = len(m.Protocol)
	n += 1 + l + sovSecurityGroup(uint64(l))
	if len(m.Destinations) > 0 {
		for _, s := range m.Destinations {
			l = len(s)
			n += 1 + l + sovSecurityGroup(uint64(l))
		}
	}
	if len(m.Ports) > 0 {
		for _, e := range m.Ports {
			n += 1 + sovSecurityGroup(uint64(e))
		}
	}
	if m.PortRange != nil {
		l = m.PortRange.Size()
		n += 1 + l + sovSecurityGroup(uint64(l))
	}
	if m.IcmpInfo != nil {
		l = m.IcmpInfo.Size()
		n += 1 + l + sovSecurityGroup(uint64(l))
	}
	n += 2
	return n
}

func sovSecurityGroup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSecurityGroup(x uint64) (n int) {
	return sovSecurityGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PortRange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PortRange{`,
		`Start:` + fmt.Sprintf("%v", this.Start) + `,`,
		`End:` + fmt.Sprintf("%v", this.End) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ICMPInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ICMPInfo{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SecurityGroupRule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SecurityGroupRule{`,
		`Protocol:` + fmt.Sprintf("%v", this.Protocol) + `,`,
		`Destinations:` + fmt.Sprintf("%v", this.Destinations) + `,`,
		`Ports:` + fmt.Sprintf("%v", this.Ports) + `,`,
		`PortRange:` + strings.Replace(fmt.Sprintf("%v", this.PortRange), "PortRange", "PortRange", 1) + `,`,
		`IcmpInfo:` + strings.Replace(fmt.Sprintf("%v", this.IcmpInfo), "ICMPInfo", "ICMPInfo", 1) + `,`,
		`Log:` + fmt.Sprintf("%v", this.Log) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSecurityGroup(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PortRange) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PortRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Start |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.End |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func (m *ICMPInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ICMPInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICMPInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func (m *SecurityGroupRule) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecurityGroupRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecurityGroupRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ports = append(m.Ports, v)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PortRange == nil {
				m.PortRange = &PortRange{}
			}
			if err := m.PortRange.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcmpInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcmpInfo == nil {
				m.IcmpInfo = &ICMPInfo{}
			}
			if err := m.IcmpInfo.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Log = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func skipSecurityGroup(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSecurityGroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSecurityGroup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipSecurityGroup(data[start:])
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
	ErrInvalidLengthSecurityGroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSecurityGroup   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("security_group.proto", fileDescriptorSecurityGroup) }

var fileDescriptorSecurityGroup = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x3f, 0xee, 0xda, 0x30,
	0x14, 0xc7, 0xe3, 0x5f, 0x08, 0x4a, 0xdc, 0x22, 0x15, 0x0b, 0x21, 0x8b, 0xc1, 0x44, 0x99, 0x32,
	0x40, 0xa8, 0x3a, 0x74, 0x6c, 0x25, 0x3a, 0x54, 0x0c, 0x95, 0x90, 0x7b, 0x00, 0x04, 0x89, 0x49,
	0x23, 0x25, 0x79, 0x51, 0xe2, 0x0c, 0x6c, 0x3d, 0x42, 0x8f, 0xd1, 0xa3, 0x30, 0x32, 0x76, 0x42,
	0x25, 0x5d, 0xaa, 0x4e, 0x9c, 0xa0, 0xaa, 0xec, 0x90, 0xa8, 0xbf, 0xed, 0xbd, 0xef, 0x1f, 0xcb,
	0xef, 0x83, 0x27, 0x95, 0x08, 0xeb, 0x32, 0x91, 0xa7, 0x5d, 0x5c, 0x42, 0x5d, 0x04, 0x45, 0x09,
	0x12, 0xc8, 0x30, 0x83, 0x48, 0xa4, 0xd5, 0x6c, 0x19, 0x27, 0xf2, 0x4b, 0x7d, 0x08, 0x42, 0xc8,
	0x56, 0x31, 0xc4, 0xb0, 0xd2, 0xf6, 0xa1, 0x3e, 0xea, 0x4d, 0x2f, 0x7a, 0x6a, 0x6b, 0xde, 0x7b,
	0xec, 0x6c, 0xa1, 0x94, 0x7c, 0x9f, 0xc7, 0x82, 0xcc, 0xb0, 0x55, 0xc9, 0x7d, 0x29, 0x29, 0x72,
	0x91, 0x3f, 0x5a, 0x0f, 0xce, 0xd7, 0xb9, 0xc1, 0x5b, 0x89, 0x4c, 0xb1, 0x29, 0xf2, 0x88, 0x3e,
	0xfd, 0xe7, 0x28, 0xc1, 0x7b, 0x87, 0xed, 0xcd, 0x87, 0x4f, 0xdb, 0x4d, 0x7e, 0x04, 0x42, 0xf1,
	0x40, 0x9e, 0x0a, 0xa1, 0xeb, 0xd6, 0x23, 0xa4, 0x15, 0xe5, 0x84, 0x10, 0x09, 0x5d, 0xef, 0x1d,
	0xa5, 0x78, 0x7f, 0x11, 0x1e, 0x7f, 0x7e, 0x1c, 0xf4, 0x51, 0xdd, 0xc3, 0xeb, 0x54, 0x90, 0xb7,
	0xd8, 0xd6, 0xff, 0x0b, 0x21, 0xd5, 0xaf, 0x39, 0xeb, 0x99, 0xea, 0xfc, 0xb9, 0xce, 0x49, 0xa7,
	0x2f, 0x20, 0x4b, 0xa4, 0xc8, 0x0a, 0x79, 0xe2, 0x7d, 0x96, 0x78, 0xf8, 0x65, 0x24, 0x2a, 0x99,
	0xe4, 0x7b, 0x99, 0x40, 0x5e, 0xd1, 0x27, 0xd7, 0xf4, 0x1d, 0xfe, 0x4c, 0x23, 0x13, 0x6c, 0x15,
	0x50, 0xca, 0x8a, 0x9a, 0xae, 0xe9, 0x8f, 0x78, 0xbb, 0x90, 0xd7, 0x18, 0xab, 0x61, 0x57, 0x2a,
	0x12, 0x74, 0xe0, 0x22, 0xff, 0xc5, 0x9b, 0x71, 0xd0, 0x42, 0x0d, 0x7a, 0x44, 0xdc, 0x29, 0x7a,
	0x5a, 0x4b, 0xec, 0x24, 0x61, 0x56, 0xec, 0x92, 0xfc, 0x08, 0xd4, 0xd2, 0x85, 0x57, 0x5d, 0xa1,
	0x43, 0xc2, 0x6d, 0x15, 0xd1, 0x70, 0xa6, 0xd8, 0x4c, 0x21, 0xa6, 0x43, 0x17, 0xf9, 0x76, 0x07,
	0x30, 0x85, 0x78, 0xbd, 0xb8, 0xdc, 0x98, 0xf1, 0xe3, 0xc6, 0x8c, 0xfb, 0x8d, 0xa1, 0xaf, 0x0d,
	0x43, 0xdf, 0x1b, 0x86, 0xce, 0x0d, 0x43, 0x97, 0x86, 0xa1, 0x9f, 0x0d, 0x43, 0xbf, 0x1b, 0x66,
	0xdc, 0x1b, 0x86, 0xbe, 0xfd, 0x62, 0xc6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xf6, 0xb1,
	0x16, 0xfd, 0x01, 0x00, 0x00,
}
