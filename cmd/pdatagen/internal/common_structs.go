// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

var commonFile = &File{
	Name: "common",
	imports: []string{
		`otlpcommon "go.opentelemetry.io/collector/model/internal/data/protogen/common/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`otlpcommon "go.opentelemetry.io/collector/model/internal/data/protogen/common/v1"`,
	},
	structs: []baseStruct{
		instrumentationLibrary,
		anyValueArray,
	},
}

var instrumentationLibrary = &messageValueStruct{
	structName:     "InstrumentationLibrary",
	description:    "// InstrumentationLibrary is a message representing the instrumentation library information.",
	originFullName: "otlpcommon.InstrumentationLibrary",
	fields: []baseField{
		nameField,
		&primitiveField{
			fieldName:       "Version",
			originFieldName: "Version",
			returnType:      "string",
			defaultVal:      `""`,
			testVal:         `"test_version"`,
		},
	},
}

// This will not be generated by this class.
// Defined here just to be available as returned message for the fields.
var stringMap = &sliceOfPtrs{
	structName: "StringMap",
	element:    stringKeyValue,
}

var stringKeyValue = &messageValueStruct{}

// This will not be generated by this class.
// Defined here just to be available as returned message for the fields.
var attributeMap = &sliceOfPtrs{
	structName: "AttributeMap",
	element:    attributeKeyValue,
}

var attributeKeyValue = &messageValueStruct{}

var instrumentationLibraryField = &messageValueField{
	fieldName:       "InstrumentationLibrary",
	originFieldName: "InstrumentationLibrary",
	returnMessage:   instrumentationLibrary,
}

var startTimeField = &primitiveTypedField{
	fieldName:       "StartTimestamp",
	originFieldName: "StartTimeUnixNano",
	returnType:      "Timestamp",
	rawType:         "uint64",
	defaultVal:      "Timestamp(0)",
	testVal:         "Timestamp(1234567890)",
}

var timeField = &primitiveTypedField{
	fieldName:       "Timestamp",
	originFieldName: "TimeUnixNano",
	returnType:      "Timestamp",
	rawType:         "uint64",
	defaultVal:      "Timestamp(0)",
	testVal:         "Timestamp(1234567890)",
}

var endTimeField = &primitiveTypedField{
	fieldName:       "EndTimestamp",
	originFieldName: "EndTimeUnixNano",
	returnType:      "Timestamp",
	rawType:         "uint64",
	defaultVal:      "Timestamp(0)",
	testVal:         "Timestamp(1234567890)",
}

var attributes = &sliceField{
	fieldName:       "Attributes",
	originFieldName: "Attributes",
	returnSlice:     attributeMap,
}

var nameField = &primitiveField{
	fieldName:       "Name",
	originFieldName: "Name",
	returnType:      "string",
	defaultVal:      `""`,
	testVal:         `"test_name"`,
}

var anyValue = &messageValueStruct{
	structName:     "AttributeValue",
	originFullName: "otlpcommon.AnyValue",
}

var anyValueArray = &sliceOfValues{
	structName: "AnyValueArray",
	element:    anyValue,
}

var schemaURLField = &primitiveField{
	fieldName:       "SchemaUrl",
	originFieldName: "SchemaUrl",
	returnType:      "string",
	defaultVal:      `""`,
	testVal:         `"https://opentelemetry.io/schemas/1.5.0"`,
}
