// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/device.proto

package entity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type Device_Status int32

const (
	Device_Unknown      Device_Status = 0
	Device_Connected    Device_Status = 1
	Device_Disconnected Device_Status = 2
	Device_Available    Device_Status = 3
	Device_Myself       Device_Status = 42
)

var Device_Status_name = map[int32]string{
	0:  "Unknown",
	1:  "Connected",
	2:  "Disconnected",
	3:  "Available",
	42: "Myself",
}
var Device_Status_value = map[string]int32{
	"Unknown":      0,
	"Connected":    1,
	"Disconnected": 2,
	"Available":    3,
	"Myself":       42,
}

func (x Device_Status) String() string {
	return proto.EnumName(Device_Status_name, int32(x))
}
func (Device_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorDevice, []int{0, 0} }

type Device struct {
	ID         string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt  time.Time     `protobuf:"bytes,2,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	UpdatedAt  time.Time     `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	DeletedAt  *time.Time    `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,stdtime" json:"deleted_at,omitempty"`
	Name       string        `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Status     Device_Status `protobuf:"varint,10,opt,name=status,proto3,enum=berty.entity.Device_Status" json:"status,omitempty"`
	ApiVersion uint32        `protobuf:"varint,11,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	ContactID  string        `protobuf:"bytes,12,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptorDevice, []int{0} }

func (m *Device) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Device) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Device) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *Device) GetDeletedAt() *time.Time {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetStatus() Device_Status {
	if m != nil {
		return m.Status
	}
	return Device_Unknown
}

func (m *Device) GetApiVersion() uint32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *Device) GetContactID() string {
	if m != nil {
		return m.ContactID
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "berty.entity.Device")
	proto.RegisterEnum("berty.entity.Device_Status", Device_Status_name, Device_Status_value)
}
func (m *Device) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Device) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDevice(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintDevice(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.DeletedAt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDevice(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.DeletedAt)))
		n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.DeletedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Status != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.Status))
	}
	if m.ApiVersion != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.ApiVersion))
	}
	if len(m.ContactID) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.ContactID)))
		i += copy(dAtA[i:], m.ContactID)
	}
	return i, nil
}

func encodeFixed64Device(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Device(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDevice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Device) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovDevice(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovDevice(uint64(l))
	if m.DeletedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.DeletedAt)
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDevice(uint64(m.Status))
	}
	if m.ApiVersion != 0 {
		n += 1 + sovDevice(uint64(m.ApiVersion))
	}
	l = len(m.ContactID)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	return n
}

func sovDevice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDevice(x uint64) (n int) {
	return sovDevice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Device) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
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
			return fmt.Errorf("proto: Device: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Device: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletedAt == nil {
				m.DeletedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.DeletedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Device_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiVersion", wireType)
			}
			m.ApiVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiVersion |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContactID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevice
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
func skipDevice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevice
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
					return 0, ErrIntOverflowDevice
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
					return 0, ErrIntOverflowDevice
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
				return 0, ErrInvalidLengthDevice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDevice
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
				next, err := skipDevice(dAtA[start:])
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
	ErrInvalidLengthDevice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevice   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("entity/device.proto", fileDescriptorDevice) }

var fileDescriptorDevice = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x4e, 0x30, 0x64, 0x92, 0x20, 0x6b, 0xa1, 0xb0, 0x02, 0x8a, 0x23, 0x57, 0x11,
	0x3a, 0xad, 0xa5, 0x3b, 0x2a, 0x1a, 0x94, 0x5c, 0x9a, 0x14, 0x14, 0x98, 0x3f, 0x05, 0x4d, 0xb4,
	0xf6, 0xce, 0x99, 0xd5, 0xd9, 0x5e, 0xcb, 0xde, 0x04, 0xf9, 0x2d, 0x28, 0x79, 0xa4, 0x2b, 0x79,
	0x82, 0x80, 0x4c, 0x4b, 0xc5, 0x13, 0x20, 0x7b, 0x1d, 0x44, 0x87, 0x68, 0x56, 0xb3, 0x33, 0xdf,
	0xfc, 0x3e, 0xcd, 0x07, 0x8f, 0x30, 0xd7, 0x52, 0xd7, 0x81, 0xc0, 0xa3, 0x8c, 0x91, 0x15, 0xa5,
	0xd2, 0x8a, 0x4e, 0x23, 0x2c, 0x75, 0xcd, 0xcc, 0x68, 0xfe, 0x38, 0x51, 0x89, 0xea, 0x06, 0x41,
	0x5b, 0x19, 0xcd, 0xdc, 0x4b, 0x94, 0x4a, 0x52, 0x0c, 0xba, 0x5f, 0x74, 0xb8, 0x09, 0xb4, 0xcc,
	0xb0, 0xd2, 0x3c, 0x2b, 0x8c, 0xc0, 0xff, 0x39, 0x04, 0x7b, 0xdb, 0x51, 0xe9, 0x05, 0x58, 0x52,
	0xb8, 0x64, 0x49, 0x56, 0xe3, 0xcd, 0xd3, 0xe6, 0xe4, 0x59, 0xbb, 0xed, 0xaf, 0x93, 0x47, 0x13,
	0x55, 0x66, 0x2f, 0xfc, 0xa2, 0x94, 0x19, 0x2f, 0xeb, 0xfd, 0x2d, 0xd6, 0x7e, 0x68, 0x49, 0x41,
	0xaf, 0x01, 0xe2, 0x12, 0xb9, 0x46, 0xb1, 0xe7, 0xda, 0xb5, 0x96, 0x64, 0x35, 0xb9, 0x9c, 0x33,
	0x63, 0xc7, 0xce, 0x76, 0xec, 0xed, 0xd9, 0x6e, 0xf3, 0xe0, 0xee, 0xe4, 0x0d, 0x3e, 0x7f, 0xf3,
	0x48, 0x38, 0xee, 0xf7, 0xd6, 0xba, 0x85, 0x1c, 0x0a, 0x71, 0x86, 0x0c, 0xff, 0x07, 0xd2, 0xef,
	0xad, 0x35, 0x7d, 0x09, 0x20, 0x30, 0xc5, 0x1e, 0x32, 0xfa, 0x27, 0x64, 0x64, 0x00, 0xfd, 0xce,
	0x5a, 0x53, 0x0a, 0xa3, 0x9c, 0x67, 0xe8, 0xde, 0x6b, 0x4f, 0x0f, 0xbb, 0x9a, 0x5e, 0x81, 0x5d,
	0x69, 0xae, 0x0f, 0x95, 0x0b, 0x4b, 0xb2, 0x7a, 0x78, 0xf9, 0x84, 0xfd, 0x9d, 0x36, 0x33, 0x91,
	0xb1, 0x37, 0x9d, 0x24, 0xec, 0xa5, 0xd4, 0x83, 0x09, 0x2f, 0xe4, 0xfe, 0x88, 0x65, 0x25, 0x55,
	0xee, 0x4e, 0x96, 0x64, 0x35, 0x0b, 0x81, 0x17, 0xf2, 0xbd, 0xe9, 0xd0, 0x0b, 0x80, 0x58, 0xe5,
	0x9a, 0xc7, 0x7a, 0x2f, 0x85, 0x3b, 0xed, 0xa2, 0x9e, 0x35, 0x27, 0x6f, 0x7c, 0x6d, 0xba, 0xbb,
	0x6d, 0x38, 0xee, 0x05, 0x3b, 0xe1, 0xbf, 0x06, 0xdb, 0x18, 0xd0, 0x09, 0xdc, 0x7f, 0x97, 0xdf,
	0xe6, 0xea, 0x53, 0xee, 0x0c, 0xe8, 0x0c, 0x5a, 0x79, 0x8e, 0xb1, 0x46, 0xe1, 0x10, 0xea, 0xc0,
	0x74, 0x2b, 0xab, 0xf8, 0x4f, 0xc7, 0x6a, 0x05, 0xeb, 0x23, 0x97, 0x29, 0x8f, 0x52, 0x74, 0x86,
	0x14, 0xc0, 0x7e, 0x55, 0x57, 0x98, 0xde, 0x38, 0xcf, 0x36, 0xcf, 0xef, 0x9a, 0x05, 0xf9, 0xda,
	0x2c, 0xc8, 0xf7, 0x66, 0x41, 0xbe, 0xfc, 0x58, 0x0c, 0x3e, 0xf8, 0x89, 0xd4, 0x1f, 0x0f, 0x11,
	0x8b, 0x55, 0x16, 0x74, 0x27, 0xf6, 0x6f, 0xac, 0x4a, 0x0c, 0xcc, 0xb5, 0x91, 0xdd, 0xa5, 0x78,
	0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x57, 0xc3, 0x49, 0xc3, 0x87, 0x02, 0x00, 0x00,
}
