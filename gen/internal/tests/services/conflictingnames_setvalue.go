// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package services

import (
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// ConflictingNames_SetValue_Args represents the arguments for the ConflictingNames.setValue function.
//
// The arguments for setValue are sent and received over the wire as this struct.
type ConflictingNames_SetValue_Args struct {
	Request *ConflictingNamesSetValueArgs `json:"request,omitempty"`
}

// ToWire translates a ConflictingNames_SetValue_Args struct into a Thrift-level intermediate
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
func (v *ConflictingNames_SetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ConflictingNamesSetValueArgs_Read(w wire.Value) (*ConflictingNamesSetValueArgs, error) {
	var v ConflictingNamesSetValueArgs
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a ConflictingNames_SetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ConflictingNames_SetValue_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ConflictingNames_SetValue_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ConflictingNames_SetValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _ConflictingNamesSetValueArgs_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ConflictingNames_SetValue_Args
// struct.
func (v *ConflictingNames_SetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("ConflictingNames_SetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ConflictingNames_SetValue_Args match the
// provided ConflictingNames_SetValue_Args.
//
// This function performs a deep comparison.
func (v *ConflictingNames_SetValue_Args) Equals(rhs *ConflictingNames_SetValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ConflictingNames_SetValue_Args.
func (v *ConflictingNames_SetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *ConflictingNames_SetValue_Args) GetRequest() (o *ConflictingNamesSetValueArgs) {
	if v != nil && v.Request != nil {
		return v.Request
	}

	return
}

// IsSetRequest returns true if Request is not nil.
func (v *ConflictingNames_SetValue_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "setValue" for this struct.
func (v *ConflictingNames_SetValue_Args) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ConflictingNames_SetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ConflictingNames_SetValue_Helper provides functions that aid in handling the
// parameters and return values of the ConflictingNames.setValue
// function.
var ConflictingNames_SetValue_Helper = struct {
	// Args accepts the parameters of setValue in-order and returns
	// the arguments struct for the function.
	Args func(
		request *ConflictingNamesSetValueArgs,
	) *ConflictingNames_SetValue_Args

	// IsException returns true if the given error can be thrown
	// by setValue.
	//
	// An error can be thrown by setValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for setValue
	// given the error returned by it. The provided error may
	// be nil if setValue did not fail.
	//
	// This allows mapping errors returned by setValue into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// setValue
	//
	//   err := setValue(args)
	//   result, err := ConflictingNames_SetValue_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from setValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*ConflictingNames_SetValue_Result, error)

	// UnwrapResponse takes the result struct for setValue
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if setValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := ConflictingNames_SetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ConflictingNames_SetValue_Result) error
}{}

func init() {
	ConflictingNames_SetValue_Helper.Args = func(
		request *ConflictingNamesSetValueArgs,
	) *ConflictingNames_SetValue_Args {
		return &ConflictingNames_SetValue_Args{
			Request: request,
		}
	}

	ConflictingNames_SetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ConflictingNames_SetValue_Helper.WrapResponse = func(err error) (*ConflictingNames_SetValue_Result, error) {
		if err == nil {
			return &ConflictingNames_SetValue_Result{}, nil
		}

		return nil, err
	}
	ConflictingNames_SetValue_Helper.UnwrapResponse = func(result *ConflictingNames_SetValue_Result) (err error) {
		return
	}

}

// ConflictingNames_SetValue_Result represents the result of a ConflictingNames.setValue function call.
//
// The result of a setValue execution is sent and received over the wire as this struct.
type ConflictingNames_SetValue_Result struct {
}

// ToWire translates a ConflictingNames_SetValue_Result struct into a Thrift-level intermediate
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
func (v *ConflictingNames_SetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ConflictingNames_SetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ConflictingNames_SetValue_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ConflictingNames_SetValue_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ConflictingNames_SetValue_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ConflictingNames_SetValue_Result
// struct.
func (v *ConflictingNames_SetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ConflictingNames_SetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ConflictingNames_SetValue_Result match the
// provided ConflictingNames_SetValue_Result.
//
// This function performs a deep comparison.
func (v *ConflictingNames_SetValue_Result) Equals(rhs *ConflictingNames_SetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ConflictingNames_SetValue_Result.
func (v *ConflictingNames_SetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "setValue" for this struct.
func (v *ConflictingNames_SetValue_Result) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ConflictingNames_SetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
