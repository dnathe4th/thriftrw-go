// Code generated by thriftrw v1.15.0. DO NOT EDIT.
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/gen/internal/tests/exceptions"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// KeyValue_DeleteValue_Args represents the arguments for the KeyValue.deleteValue function.
//
// The arguments for deleteValue are sent and received over the wire as this struct.
type KeyValue_DeleteValue_Args struct {
	Key *Key `json:"key,omitempty"`
}

// ToWire translates a KeyValue_DeleteValue_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_DeleteValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = v.Key.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Key_Read(w wire.Value) (Key, error) {
	var x Key
	err := x.FromWire(w)
	return x, err
}

// FromWire deserializes a KeyValue_DeleteValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_DeleteValue_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_DeleteValue_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_DeleteValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_DeleteValue_Args
// struct.
func (v *KeyValue_DeleteValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}

	return fmt.Sprintf("KeyValue_DeleteValue_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Key_EqualsPtr(lhs, rhs *Key) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this KeyValue_DeleteValue_Args match the
// provided KeyValue_DeleteValue_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_DeleteValue_Args) Equals(rhs *KeyValue_DeleteValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Key_EqualsPtr(v.Key, rhs.Key) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_DeleteValue_Args.
func (v *KeyValue_DeleteValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", (string)(*v.Key))
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *KeyValue_DeleteValue_Args) GetKey() (o Key) {
	if v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *KeyValue_DeleteValue_Args) IsSetKey() bool {
	return v.Key != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "deleteValue" for this struct.
func (v *KeyValue_DeleteValue_Args) MethodName() string {
	return "deleteValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_DeleteValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_DeleteValue_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.deleteValue
// function.
var KeyValue_DeleteValue_Helper = struct {
	// Args accepts the parameters of deleteValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *Key,
	) *KeyValue_DeleteValue_Args

	// IsException returns true if the given error can be thrown
	// by deleteValue.
	//
	// An error can be thrown by deleteValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for deleteValue
	// given the error returned by it. The provided error may
	// be nil if deleteValue did not fail.
	//
	// This allows mapping errors returned by deleteValue into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// deleteValue
	//
	//   err := deleteValue(args)
	//   result, err := KeyValue_DeleteValue_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from deleteValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*KeyValue_DeleteValue_Result, error)

	// UnwrapResponse takes the result struct for deleteValue
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if deleteValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := KeyValue_DeleteValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_DeleteValue_Result) error
}{}

func init() {
	KeyValue_DeleteValue_Helper.Args = func(
		key *Key,
	) *KeyValue_DeleteValue_Args {
		return &KeyValue_DeleteValue_Args{
			Key: key,
		}
	}

	KeyValue_DeleteValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		case *InternalError:
			return true
		default:
			return false
		}
	}

	KeyValue_DeleteValue_Helper.WrapResponse = func(err error) (*KeyValue_DeleteValue_Result, error) {
		if err == nil {
			return &KeyValue_DeleteValue_Result{}, nil
		}

		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_DeleteValue_Result.DoesNotExist")
			}
			return &KeyValue_DeleteValue_Result{DoesNotExist: e}, nil
		case *InternalError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_DeleteValue_Result.InternalError")
			}
			return &KeyValue_DeleteValue_Result{InternalError: e}, nil
		}

		return nil, err
	}
	KeyValue_DeleteValue_Helper.UnwrapResponse = func(result *KeyValue_DeleteValue_Result) (err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.InternalError != nil {
			err = result.InternalError
			return
		}
		return
	}

}

// KeyValue_DeleteValue_Result represents the result of a KeyValue.deleteValue function call.
//
// The result of a deleteValue execution is sent and received over the wire as this struct.
type KeyValue_DeleteValue_Result struct {
	// Raised if a value with the given key doesn't exist.
	DoesNotExist  *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
	InternalError *InternalError                    `json:"internalError,omitempty"`
}

// ToWire translates a KeyValue_DeleteValue_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_DeleteValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalError != nil {
		w, err = v.InternalError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_DeleteValue_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DoesNotExistException_Read(w wire.Value) (*exceptions.DoesNotExistException, error) {
	var v exceptions.DoesNotExistException
	err := v.FromWire(w)
	return &v, err
}

func _InternalError_Read(w wire.Value) (*InternalError, error) {
	var v InternalError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a KeyValue_DeleteValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_DeleteValue_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_DeleteValue_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_DeleteValue_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalError, err = _InternalError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.DoesNotExist != nil {
		count++
	}
	if v.InternalError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("KeyValue_DeleteValue_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a KeyValue_DeleteValue_Result
// struct.
func (v *KeyValue_DeleteValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	if v.InternalError != nil {
		fields[i] = fmt.Sprintf("InternalError: %v", v.InternalError)
		i++
	}

	return fmt.Sprintf("KeyValue_DeleteValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_DeleteValue_Result match the
// provided KeyValue_DeleteValue_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_DeleteValue_Result) Equals(rhs *KeyValue_DeleteValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}
	if !((v.InternalError == nil && rhs.InternalError == nil) || (v.InternalError != nil && rhs.InternalError != nil && v.InternalError.Equals(rhs.InternalError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_DeleteValue_Result.
func (v *KeyValue_DeleteValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	if v.InternalError != nil {
		err = multierr.Append(err, enc.AddObject("internalError", v.InternalError))
	}
	return err
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *KeyValue_DeleteValue_Result) GetDoesNotExist() (o *exceptions.DoesNotExistException) {
	if v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// IsSetDoesNotExist returns true if DoesNotExist is not nil.
func (v *KeyValue_DeleteValue_Result) IsSetDoesNotExist() bool {
	return v.DoesNotExist != nil
}

// GetInternalError returns the value of InternalError if it is set or its
// zero value if it is unset.
func (v *KeyValue_DeleteValue_Result) GetInternalError() (o *InternalError) {
	if v.InternalError != nil {
		return v.InternalError
	}

	return
}

// IsSetInternalError returns true if InternalError is not nil.
func (v *KeyValue_DeleteValue_Result) IsSetInternalError() bool {
	return v.InternalError != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "deleteValue" for this struct.
func (v *KeyValue_DeleteValue_Result) MethodName() string {
	return "deleteValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_DeleteValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
