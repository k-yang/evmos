// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package feemarketv1

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Params                             protoreflect.MessageDescriptor
	fd_Params_no_base_fee                 protoreflect.FieldDescriptor
	fd_Params_base_fee_change_denominator protoreflect.FieldDescriptor
	fd_Params_elasticity_multiplier       protoreflect.FieldDescriptor
	fd_Params_enable_height               protoreflect.FieldDescriptor
	fd_Params_base_fee                    protoreflect.FieldDescriptor
	fd_Params_min_gas_price               protoreflect.FieldDescriptor
	fd_Params_min_gas_multiplier          protoreflect.FieldDescriptor
)

func init() {
	file_ethermint_feemarket_v1_feemarket_proto_init()
	md_Params = File_ethermint_feemarket_v1_feemarket_proto.Messages().ByName("Params")
	fd_Params_no_base_fee = md_Params.Fields().ByName("no_base_fee")
	fd_Params_base_fee_change_denominator = md_Params.Fields().ByName("base_fee_change_denominator")
	fd_Params_elasticity_multiplier = md_Params.Fields().ByName("elasticity_multiplier")
	fd_Params_enable_height = md_Params.Fields().ByName("enable_height")
	fd_Params_base_fee = md_Params.Fields().ByName("base_fee")
	fd_Params_min_gas_price = md_Params.Fields().ByName("min_gas_price")
	fd_Params_min_gas_multiplier = md_Params.Fields().ByName("min_gas_multiplier")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.NoBaseFee != false {
		value := protoreflect.ValueOfBool(x.NoBaseFee)
		if !f(fd_Params_no_base_fee, value) {
			return
		}
	}
	if x.BaseFeeChangeDenominator != uint32(0) {
		value := protoreflect.ValueOfUint32(x.BaseFeeChangeDenominator)
		if !f(fd_Params_base_fee_change_denominator, value) {
			return
		}
	}
	if x.ElasticityMultiplier != uint32(0) {
		value := protoreflect.ValueOfUint32(x.ElasticityMultiplier)
		if !f(fd_Params_elasticity_multiplier, value) {
			return
		}
	}
	if x.EnableHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.EnableHeight)
		if !f(fd_Params_enable_height, value) {
			return
		}
	}
	if x.BaseFee != "" {
		value := protoreflect.ValueOfString(x.BaseFee)
		if !f(fd_Params_base_fee, value) {
			return
		}
	}
	if x.MinGasPrice != "" {
		value := protoreflect.ValueOfString(x.MinGasPrice)
		if !f(fd_Params_min_gas_price, value) {
			return
		}
	}
	if x.MinGasMultiplier != "" {
		value := protoreflect.ValueOfString(x.MinGasMultiplier)
		if !f(fd_Params_min_gas_multiplier, value) {
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
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		return x.NoBaseFee != false
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		return x.BaseFeeChangeDenominator != uint32(0)
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		return x.ElasticityMultiplier != uint32(0)
	case "ethermint.feemarket.v1.Params.enable_height":
		return x.EnableHeight != int64(0)
	case "ethermint.feemarket.v1.Params.base_fee":
		return x.BaseFee != ""
	case "ethermint.feemarket.v1.Params.min_gas_price":
		return x.MinGasPrice != ""
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		return x.MinGasMultiplier != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		x.NoBaseFee = false
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		x.BaseFeeChangeDenominator = uint32(0)
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		x.ElasticityMultiplier = uint32(0)
	case "ethermint.feemarket.v1.Params.enable_height":
		x.EnableHeight = int64(0)
	case "ethermint.feemarket.v1.Params.base_fee":
		x.BaseFee = ""
	case "ethermint.feemarket.v1.Params.min_gas_price":
		x.MinGasPrice = ""
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		x.MinGasMultiplier = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		value := x.NoBaseFee
		return protoreflect.ValueOfBool(value)
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		value := x.BaseFeeChangeDenominator
		return protoreflect.ValueOfUint32(value)
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		value := x.ElasticityMultiplier
		return protoreflect.ValueOfUint32(value)
	case "ethermint.feemarket.v1.Params.enable_height":
		value := x.EnableHeight
		return protoreflect.ValueOfInt64(value)
	case "ethermint.feemarket.v1.Params.base_fee":
		value := x.BaseFee
		return protoreflect.ValueOfString(value)
	case "ethermint.feemarket.v1.Params.min_gas_price":
		value := x.MinGasPrice
		return protoreflect.ValueOfString(value)
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		value := x.MinGasMultiplier
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		x.NoBaseFee = value.Bool()
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		x.BaseFeeChangeDenominator = uint32(value.Uint())
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		x.ElasticityMultiplier = uint32(value.Uint())
	case "ethermint.feemarket.v1.Params.enable_height":
		x.EnableHeight = value.Int()
	case "ethermint.feemarket.v1.Params.base_fee":
		x.BaseFee = value.Interface().(string)
	case "ethermint.feemarket.v1.Params.min_gas_price":
		x.MinGasPrice = value.Interface().(string)
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		x.MinGasMultiplier = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		panic(fmt.Errorf("field no_base_fee of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		panic(fmt.Errorf("field base_fee_change_denominator of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		panic(fmt.Errorf("field elasticity_multiplier of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.enable_height":
		panic(fmt.Errorf("field enable_height of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.base_fee":
		panic(fmt.Errorf("field base_fee of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.min_gas_price":
		panic(fmt.Errorf("field min_gas_price of message ethermint.feemarket.v1.Params is not mutable"))
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		panic(fmt.Errorf("field min_gas_multiplier of message ethermint.feemarket.v1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ethermint.feemarket.v1.Params.no_base_fee":
		return protoreflect.ValueOfBool(false)
	case "ethermint.feemarket.v1.Params.base_fee_change_denominator":
		return protoreflect.ValueOfUint32(uint32(0))
	case "ethermint.feemarket.v1.Params.elasticity_multiplier":
		return protoreflect.ValueOfUint32(uint32(0))
	case "ethermint.feemarket.v1.Params.enable_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "ethermint.feemarket.v1.Params.base_fee":
		return protoreflect.ValueOfString("")
	case "ethermint.feemarket.v1.Params.min_gas_price":
		return protoreflect.ValueOfString("")
	case "ethermint.feemarket.v1.Params.min_gas_multiplier":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.feemarket.v1.Params"))
		}
		panic(fmt.Errorf("message ethermint.feemarket.v1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in ethermint.feemarket.v1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
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
		if x.NoBaseFee {
			n += 2
		}
		if x.BaseFeeChangeDenominator != 0 {
			n += 1 + runtime.Sov(uint64(x.BaseFeeChangeDenominator))
		}
		if x.ElasticityMultiplier != 0 {
			n += 1 + runtime.Sov(uint64(x.ElasticityMultiplier))
		}
		if x.EnableHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.EnableHeight))
		}
		l = len(x.BaseFee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MinGasPrice)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MinGasMultiplier)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*Params)
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
		if len(x.MinGasMultiplier) > 0 {
			i -= len(x.MinGasMultiplier)
			copy(dAtA[i:], x.MinGasMultiplier)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinGasMultiplier)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.MinGasPrice) > 0 {
			i -= len(x.MinGasPrice)
			copy(dAtA[i:], x.MinGasPrice)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinGasPrice)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.BaseFee) > 0 {
			i -= len(x.BaseFee)
			copy(dAtA[i:], x.BaseFee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BaseFee)))
			i--
			dAtA[i] = 0x32
		}
		if x.EnableHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EnableHeight))
			i--
			dAtA[i] = 0x28
		}
		if x.ElasticityMultiplier != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ElasticityMultiplier))
			i--
			dAtA[i] = 0x18
		}
		if x.BaseFeeChangeDenominator != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BaseFeeChangeDenominator))
			i--
			dAtA[i] = 0x10
		}
		if x.NoBaseFee {
			i--
			if x.NoBaseFee {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*Params)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NoBaseFee", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.NoBaseFee = bool(v != 0)
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseFeeChangeDenominator", wireType)
				}
				x.BaseFeeChangeDenominator = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BaseFeeChangeDenominator |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ElasticityMultiplier", wireType)
				}
				x.ElasticityMultiplier = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ElasticityMultiplier |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EnableHeight", wireType)
				}
				x.EnableHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EnableHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
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
				x.BaseFee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
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
				x.MinGasPrice = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinGasMultiplier", wireType)
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
				x.MinGasMultiplier = string(dAtA[iNdEx:postIndex])
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

// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: ethermint/feemarket/v1/feemarket.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the feemarket module parameters
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// no_base_fee forces the EIP-1559 base fee to 0 (needed for 0 price calls)
	NoBaseFee bool `protobuf:"varint,1,opt,name=no_base_fee,json=noBaseFee,proto3" json:"no_base_fee,omitempty"`
	// base_fee_change_denominator bounds the amount the base fee can change
	// between blocks.
	BaseFeeChangeDenominator uint32 `protobuf:"varint,2,opt,name=base_fee_change_denominator,json=baseFeeChangeDenominator,proto3" json:"base_fee_change_denominator,omitempty"`
	// elasticity_multiplier bounds the maximum gas limit an EIP-1559 block may
	// have.
	ElasticityMultiplier uint32 `protobuf:"varint,3,opt,name=elasticity_multiplier,json=elasticityMultiplier,proto3" json:"elasticity_multiplier,omitempty"`
	// enable_height defines at which block height the base fee calculation is enabled.
	EnableHeight int64 `protobuf:"varint,5,opt,name=enable_height,json=enableHeight,proto3" json:"enable_height,omitempty"`
	// base_fee for EIP-1559 blocks.
	BaseFee string `protobuf:"bytes,6,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
	// min_gas_price defines the minimum gas price value for cosmos and eth transactions
	MinGasPrice string `protobuf:"bytes,7,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price,omitempty"`
	// min_gas_multiplier bounds the minimum gas used to be charged
	// to senders based on gas limit
	MinGasMultiplier string `protobuf:"bytes,8,opt,name=min_gas_multiplier,json=minGasMultiplier,proto3" json:"min_gas_multiplier,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_ethermint_feemarket_v1_feemarket_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetNoBaseFee() bool {
	if x != nil {
		return x.NoBaseFee
	}
	return false
}

func (x *Params) GetBaseFeeChangeDenominator() uint32 {
	if x != nil {
		return x.BaseFeeChangeDenominator
	}
	return 0
}

func (x *Params) GetElasticityMultiplier() uint32 {
	if x != nil {
		return x.ElasticityMultiplier
	}
	return 0
}

func (x *Params) GetEnableHeight() int64 {
	if x != nil {
		return x.EnableHeight
	}
	return 0
}

func (x *Params) GetBaseFee() string {
	if x != nil {
		return x.BaseFee
	}
	return ""
}

func (x *Params) GetMinGasPrice() string {
	if x != nil {
		return x.MinGasPrice
	}
	return ""
}

func (x *Params) GetMinGasMultiplier() string {
	if x != nil {
		return x.MinGasMultiplier
	}
	return ""
}

var File_ethermint_feemarket_v1_feemarket_proto protoreflect.FileDescriptor

var file_ethermint_feemarket_v1_feemarket_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x6f, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x6f, 0x42, 0x61, 0x73,
	0x65, 0x46, 0x65, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x62, 0x61, 0x73, 0x65, 0x46,
	0x65, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x15, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x43, 0x0a,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x28, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x44, 0x65, 0x63, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x46,
	0x65, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xc8, 0xde, 0x1f, 0x00, 0xda,
	0xde, 0x1f, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65, 0x63, 0xa8, 0xe7,
	0xb0, 0x2a, 0x01, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xc8, 0xde,
	0x1f, 0x00, 0xda, 0xde, 0x1f, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e,
	0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65,
	0x63, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x3a, 0x1d, 0x8a, 0xe7, 0xb0, 0x2a, 0x18, 0x65,
	0x76, 0x6d, 0x6f, 0x73, 0x2f, 0x78, 0x2f, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x10, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x42,
	0xdb, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x46, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x33, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x46, 0x58, 0xaa, 0x02, 0x16, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x5c, 0x46, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x46, 0x65, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x18, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a,
	0x46, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_feemarket_v1_feemarket_proto_rawDescOnce sync.Once
	file_ethermint_feemarket_v1_feemarket_proto_rawDescData = file_ethermint_feemarket_v1_feemarket_proto_rawDesc
)

func file_ethermint_feemarket_v1_feemarket_proto_rawDescGZIP() []byte {
	file_ethermint_feemarket_v1_feemarket_proto_rawDescOnce.Do(func() {
		file_ethermint_feemarket_v1_feemarket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_feemarket_v1_feemarket_proto_rawDescData)
	})
	return file_ethermint_feemarket_v1_feemarket_proto_rawDescData
}

var file_ethermint_feemarket_v1_feemarket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_feemarket_v1_feemarket_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: ethermint.feemarket.v1.Params
}
var file_ethermint_feemarket_v1_feemarket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_feemarket_v1_feemarket_proto_init() }
func file_ethermint_feemarket_v1_feemarket_proto_init() {
	if File_ethermint_feemarket_v1_feemarket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_ethermint_feemarket_v1_feemarket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_feemarket_v1_feemarket_proto_goTypes,
		DependencyIndexes: file_ethermint_feemarket_v1_feemarket_proto_depIdxs,
		MessageInfos:      file_ethermint_feemarket_v1_feemarket_proto_msgTypes,
	}.Build()
	File_ethermint_feemarket_v1_feemarket_proto = out.File
	file_ethermint_feemarket_v1_feemarket_proto_rawDesc = nil
	file_ethermint_feemarket_v1_feemarket_proto_goTypes = nil
	file_ethermint_feemarket_v1_feemarket_proto_depIdxs = nil
}