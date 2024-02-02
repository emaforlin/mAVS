// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package mavs

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sort "sort"
	sync "sync"
)

var _ protoreflect.Map = (*_StoredVoting_3_map)(nil)

type _StoredVoting_3_map struct {
	m *map[string]uint64
}

func (x *_StoredVoting_3_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_StoredVoting_3_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfUint64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_StoredVoting_3_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_StoredVoting_3_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_StoredVoting_3_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfUint64(v)
}

func (x *_StoredVoting_3_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_StoredVoting_3_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_StoredVoting_3_map) NewValue() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_StoredVoting_3_map) IsValid() bool {
	return x.m != nil
}

var (
	md_StoredVoting            protoreflect.MessageDescriptor
	fd_StoredVoting_index      protoreflect.FieldDescriptor
	fd_StoredVoting_timewindow protoreflect.FieldDescriptor
	fd_StoredVoting_counting   protoreflect.FieldDescriptor
)

func init() {
	file_mavs_mavs_stored_voting_proto_init()
	md_StoredVoting = File_mavs_mavs_stored_voting_proto.Messages().ByName("StoredVoting")
	fd_StoredVoting_index = md_StoredVoting.Fields().ByName("index")
	fd_StoredVoting_timewindow = md_StoredVoting.Fields().ByName("timewindow")
	fd_StoredVoting_counting = md_StoredVoting.Fields().ByName("counting")
}

var _ protoreflect.Message = (*fastReflection_StoredVoting)(nil)

type fastReflection_StoredVoting StoredVoting

func (x *StoredVoting) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StoredVoting)(x)
}

func (x *StoredVoting) slowProtoReflect() protoreflect.Message {
	mi := &file_mavs_mavs_stored_voting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StoredVoting_messageType fastReflection_StoredVoting_messageType
var _ protoreflect.MessageType = fastReflection_StoredVoting_messageType{}

type fastReflection_StoredVoting_messageType struct{}

func (x fastReflection_StoredVoting_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StoredVoting)(nil)
}
func (x fastReflection_StoredVoting_messageType) New() protoreflect.Message {
	return new(fastReflection_StoredVoting)
}
func (x fastReflection_StoredVoting_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StoredVoting
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StoredVoting) Descriptor() protoreflect.MessageDescriptor {
	return md_StoredVoting
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StoredVoting) Type() protoreflect.MessageType {
	return _fastReflection_StoredVoting_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StoredVoting) New() protoreflect.Message {
	return new(fastReflection_StoredVoting)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StoredVoting) Interface() protoreflect.ProtoMessage {
	return (*StoredVoting)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StoredVoting) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != "" {
		value := protoreflect.ValueOfString(x.Index)
		if !f(fd_StoredVoting_index, value) {
			return
		}
	}
	if x.Timewindow != "" {
		value := protoreflect.ValueOfString(x.Timewindow)
		if !f(fd_StoredVoting_timewindow, value) {
			return
		}
	}
	if len(x.Counting) != 0 {
		value := protoreflect.ValueOfMap(&_StoredVoting_3_map{m: &x.Counting})
		if !f(fd_StoredVoting_counting, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_StoredVoting) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "mavs.mavs.StoredVoting.index":
		return x.Index != ""
	case "mavs.mavs.StoredVoting.timewindow":
		return x.Timewindow != ""
	case "mavs.mavs.StoredVoting.counting":
		return len(x.Counting) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredVoting) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "mavs.mavs.StoredVoting.index":
		x.Index = ""
	case "mavs.mavs.StoredVoting.timewindow":
		x.Timewindow = ""
	case "mavs.mavs.StoredVoting.counting":
		x.Counting = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StoredVoting) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "mavs.mavs.StoredVoting.index":
		value := x.Index
		return protoreflect.ValueOfString(value)
	case "mavs.mavs.StoredVoting.timewindow":
		value := x.Timewindow
		return protoreflect.ValueOfString(value)
	case "mavs.mavs.StoredVoting.counting":
		if len(x.Counting) == 0 {
			return protoreflect.ValueOfMap(&_StoredVoting_3_map{})
		}
		mapValue := &_StoredVoting_3_map{m: &x.Counting}
		return protoreflect.ValueOfMap(mapValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredVoting) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "mavs.mavs.StoredVoting.index":
		x.Index = value.Interface().(string)
	case "mavs.mavs.StoredVoting.timewindow":
		x.Timewindow = value.Interface().(string)
	case "mavs.mavs.StoredVoting.counting":
		mv := value.Map()
		cmv := mv.(*_StoredVoting_3_map)
		x.Counting = *cmv.m
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredVoting) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mavs.mavs.StoredVoting.counting":
		if x.Counting == nil {
			x.Counting = make(map[string]uint64)
		}
		value := &_StoredVoting_3_map{m: &x.Counting}
		return protoreflect.ValueOfMap(value)
	case "mavs.mavs.StoredVoting.index":
		panic(fmt.Errorf("field index of message mavs.mavs.StoredVoting is not mutable"))
	case "mavs.mavs.StoredVoting.timewindow":
		panic(fmt.Errorf("field timewindow of message mavs.mavs.StoredVoting is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StoredVoting) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mavs.mavs.StoredVoting.index":
		return protoreflect.ValueOfString("")
	case "mavs.mavs.StoredVoting.timewindow":
		return protoreflect.ValueOfString("")
	case "mavs.mavs.StoredVoting.counting":
		m := make(map[string]uint64)
		return protoreflect.ValueOfMap(&_StoredVoting_3_map{m: &m})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mavs.mavs.StoredVoting"))
		}
		panic(fmt.Errorf("message mavs.mavs.StoredVoting does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StoredVoting) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in mavs.mavs.StoredVoting", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StoredVoting) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredVoting) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_StoredVoting) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StoredVoting) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StoredVoting)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Index)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Timewindow)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Counting) > 0 {
			SiZeMaP := func(k string, v uint64) {
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 1 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]string, 0, len(x.Counting))
				for k := range x.Counting {
					sortme = append(sortme, k)
				}
				sort.Strings(sortme)
				for _, k := range sortme {
					v := x.Counting[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.Counting {
					SiZeMaP(k, v)
				}
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*StoredVoting)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Counting) > 0 {
			MaRsHaLmAp := func(k string, v uint64) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i -= len(k)
				copy(dAtA[i:], k)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(k)))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x1a
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForCounting := make([]string, 0, len(x.Counting))
				for k := range x.Counting {
					keysForCounting = append(keysForCounting, string(k))
				}
				sort.Slice(keysForCounting, func(i, j int) bool {
					return keysForCounting[i] < keysForCounting[j]
				})
				for iNdEx := len(keysForCounting) - 1; iNdEx >= 0; iNdEx-- {
					v := x.Counting[string(keysForCounting[iNdEx])]
					out, err := MaRsHaLmAp(keysForCounting[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.Counting {
					v := x.Counting[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.Timewindow) > 0 {
			i -= len(x.Timewindow)
			copy(dAtA[i:], x.Timewindow)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Timewindow)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Index) > 0 {
			i -= len(x.Index)
			copy(dAtA[i:], x.Index)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Index)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*StoredVoting)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoredVoting: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoredVoting: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Index = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timewindow", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Timewindow = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Counting", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Counting == nil {
					x.Counting = make(map[string]uint64)
				}
				var mapkey string
				var mapvalue uint64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						wire |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					fieldNum := int32(wire >> 3)
					if fieldNum == 1 {
						var stringLenmapkey uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							stringLenmapkey |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						intStringLenmapkey := int(stringLenmapkey)
						if intStringLenmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.Counting[mapkey] = mapvalue
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: mavs/mavs/stored_voting.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoredVoting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      string            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Timewindow string            `protobuf:"bytes,2,opt,name=timewindow,proto3" json:"timewindow,omitempty"`
	Counting   map[string]uint64 `protobuf:"bytes,3,rep,name=counting,proto3" json:"counting,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *StoredVoting) Reset() {
	*x = StoredVoting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mavs_mavs_stored_voting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredVoting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredVoting) ProtoMessage() {}

// Deprecated: Use StoredVoting.ProtoReflect.Descriptor instead.
func (*StoredVoting) Descriptor() ([]byte, []int) {
	return file_mavs_mavs_stored_voting_proto_rawDescGZIP(), []int{0}
}

func (x *StoredVoting) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *StoredVoting) GetTimewindow() string {
	if x != nil {
		return x.Timewindow
	}
	return ""
}

func (x *StoredVoting) GetCounting() map[string]uint64 {
	if x != nil {
		return x.Counting
	}
	return nil
}

var File_mavs_mavs_stored_voting_proto protoreflect.FileDescriptor

var file_mavs_mavs_stored_voting_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x61, 0x76, 0x73, 0x2f, 0x6d, 0x61, 0x76, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x5f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x61, 0x76, 0x73, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x0c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x83, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x42, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x76, 0x73, 0x2f,
	0x6d, 0x61, 0x76, 0x73, 0xa2, 0x02, 0x03, 0x4d, 0x4d, 0x58, 0xaa, 0x02, 0x09, 0x4d, 0x61, 0x76,
	0x73, 0x2e, 0x4d, 0x61, 0x76, 0x73, 0xca, 0x02, 0x09, 0x4d, 0x61, 0x76, 0x73, 0x5c, 0x4d, 0x61,
	0x76, 0x73, 0xe2, 0x02, 0x15, 0x4d, 0x61, 0x76, 0x73, 0x5c, 0x4d, 0x61, 0x76, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x61, 0x76,
	0x73, 0x3a, 0x3a, 0x4d, 0x61, 0x76, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mavs_mavs_stored_voting_proto_rawDescOnce sync.Once
	file_mavs_mavs_stored_voting_proto_rawDescData = file_mavs_mavs_stored_voting_proto_rawDesc
)

func file_mavs_mavs_stored_voting_proto_rawDescGZIP() []byte {
	file_mavs_mavs_stored_voting_proto_rawDescOnce.Do(func() {
		file_mavs_mavs_stored_voting_proto_rawDescData = protoimpl.X.CompressGZIP(file_mavs_mavs_stored_voting_proto_rawDescData)
	})
	return file_mavs_mavs_stored_voting_proto_rawDescData
}

var file_mavs_mavs_stored_voting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mavs_mavs_stored_voting_proto_goTypes = []interface{}{
	(*StoredVoting)(nil), // 0: mavs.mavs.StoredVoting
	nil,                  // 1: mavs.mavs.StoredVoting.CountingEntry
}
var file_mavs_mavs_stored_voting_proto_depIdxs = []int32{
	1, // 0: mavs.mavs.StoredVoting.counting:type_name -> mavs.mavs.StoredVoting.CountingEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mavs_mavs_stored_voting_proto_init() }
func file_mavs_mavs_stored_voting_proto_init() {
	if File_mavs_mavs_stored_voting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mavs_mavs_stored_voting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredVoting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mavs_mavs_stored_voting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mavs_mavs_stored_voting_proto_goTypes,
		DependencyIndexes: file_mavs_mavs_stored_voting_proto_depIdxs,
		MessageInfos:      file_mavs_mavs_stored_voting_proto_msgTypes,
	}.Build()
	File_mavs_mavs_stored_voting_proto = out.File
	file_mavs_mavs_stored_voting_proto_rawDesc = nil
	file_mavs_mavs_stored_voting_proto_goTypes = nil
	file_mavs_mavs_stored_voting_proto_depIdxs = nil
}
