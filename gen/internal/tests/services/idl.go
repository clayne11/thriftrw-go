// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package services

import (
	exceptions "go.uber.org/thriftrw/gen/internal/tests/exceptions"
	unions "go.uber.org/thriftrw/gen/internal/tests/unions"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "services",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/services",
	FilePath: "services.thrift",
	SHA1:     "22d2183c133fb327925d7288a6191b17e8f28175",
	Includes: []*thriftreflect.ThriftModule{
		exceptions.ThriftModule,
		unions.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./unions.thrift\"\ninclude \"./exceptions.thrift\"\n\ntypedef string Key\n\nexception InternalError {\n    1: optional string message\n}\n\nservice KeyValue {\n    // void and no exceptions\n    void setValue(1: Key key, 2: unions.ArbitraryValue value)\n\n    void setValueV2(\n        /** Key to change. */\n        1: required Key key,\n        /**\n         * New value for the key.\n         *\n         * If the key already has an existing value, it will be overwritten.\n         */\n        2: required unions.ArbitraryValue value,\n    )\n\n    // Return with exceptions\n    unions.ArbitraryValue getValue(1: Key key)\n        throws (1: exceptions.DoesNotExistException doesNotExist)\n\n    // void with exceptions\n    void deleteValue(1: Key key)\n        throws (\n            /**\n             * Raised if a value with the given key doesn't exist.\n             */\n            1: exceptions.DoesNotExistException doesNotExist,\n            2: InternalError internalError\n        )\n\n    list<unions.ArbitraryValue> getManyValues(\n        1: list<Key> range  // < reserved keyword as an argument\n    ) throws (\n        1: exceptions.DoesNotExistException doesNotExist,\n    )\n\n    i64 size()  // < primitve return value\n}\n\nservice Cache {\n    oneway void clear()\n    oneway void clearAfter(1: i64 durationMS)\n}\n\nstruct ConflictingNames_SetValue_Args {\n    1: required string key\n    2: required binary value\n}\n\nservice ConflictingNames {\n    void setValue(1: ConflictingNames_SetValue_Args request)\n}\n\nservice non_standard_service_name {\n    void non_standard_function_name()\n}\n"
