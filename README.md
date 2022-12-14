# protoc-gen-validator
`protoc-gen-validator` is a protoc plugin that can generate go structure validate functions.<br>
`protoc-gen-validator` will generate the `Validate() error` method for the corresponding message.<br>
## Example:
```
message Example {
  // The value of 'Int64Const' must be equal to '123'
  int64 Int64Const = 1 [(api.vt).const="123"];
  
  // The value of 'DoubleLe' must be less than '123.45'
  optional double DoubleLe = 2 [(api.vt).le="123.45"];
  
  // The value of 'BoolConst' must be 'true'
  optional bool BoolConst = 3 [(api.vt).const="true"];
  
  // The maximum length of 'StringMaxSize' is '12'
  optional string StringMaxSize = 4 [(api.vt).max_size="12"];
  
  // The prefix of 'bytesPrefix' must be 'validator'
  optional bytes bytesPrefix = 5 [(api.vt).prefix="validator"];
  
  // Each element of 'ListElem' must be a 'validator'
  repeated string ListElem = 6 [(api.vt).elem.const="validator"];
  
  // 'MapKeyValue' The value of all keys must be '123'
  // 'MapKeyValue' The value of all keys must be greater than '100'
  // 'MapKeyValue' All values of value must be 'validator'
  // 'MapKeyValue' All values must be prefixed with 'validator'
  map<int32, string> MapKeyValue = 7 [(api.vt).key.const="123", (api.vt).key.gt="100", (api.vt).value.const="validator", (api.vt).value.prefix="validator"];
  
  optional int64 Func1 = 8 [(api.vt).gt = "@add($Int64Const, 1000)"]; // Use the built-in function 'add' to get the constraint value
}
```
The method of generation is as follows:
```
func (m *Example) Validate() error {
	if m.GetInt64Const() != int64(123) {
		return fmt.Errorf("field Int64Const not match const value, current value: %v", m.GetInt64Const())
	}
	if m.GetDoubleLe() > float64(123.45) {
		return fmt.Errorf("field DoubleLe le rule failed, current value: %v", m.GetDoubleLe())
	}
	if m.GetBoolConst() != true {
		return fmt.Errorf("field BoolConst const rule failed, current value: %v", m.GetBoolConst())
	}
	if len(m.GetStringMaxSize()) > int(12) {
		return fmt.Errorf("field StringMaxSize max_len rule failed, current value: %d", len(m.GetStringMaxSize()))
	}
	_src := []byte("validator")
	if !bytes.HasPrefix(m.GetBytesPrefix(), _src) {
		return fmt.Errorf("field bytesPrefix prefix rule failed, current value: %v", m.GetBytesPrefix())
	}
	for i := 0; i < len(m.GetListElem()); i++ {
		_elem := m.GetListElem()[i]
		_src1 := "validator"
		if _elem != _src1 {
			return fmt.Errorf("field _elem not match const value, current value: %v", _elem)
		}
	}
	for k := range m.GetMapKeyValue() {
		if k != int32(123) {
			return fmt.Errorf("field k not match const value, current value: %v", k)
		}
		if k <= int32(100) {
			return fmt.Errorf("field k gt rule failed, current value: %v", k)
		}
	}
	for _, v := range m.GetMapKeyValue() {
		_src2 := "validator"
		if v != _src2 {
			return fmt.Errorf("field v not match const value, current value: %v", v)
		}
		_src3 := "validator"
		if !strings.HasPrefix(v, _src3) {
			return fmt.Errorf("field v prefix rule failed, current value: %v", v)
		}
	}
	_src4 := m.GetInt64Const() + int64(1000)
	if m.GetFunc1() <= int64(_src4) {
		return fmt.Errorf("field Func1 gt rule failed, current value: %v", m.GetFunc1())
	}
	return nil
}
```
# Usage
## Dependency
* [protoc](https://developers.google.com/protocol-buffers/docs/downloads) located under `$PATH`
* [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go) located under `$PATH`/`$GOPATH`
* `protoc-gen-validator` located under `$PATH`/`$GOPATH`
* Support `proto2`/`proto3` syntax, `proto3` syntax is recommended
## Installation
`go install github.com/cloudwego/protoc-gen-validator@latest`
## Parameters
* version: Print `protoc-gen-validator` version
* recurse: Recursively generate validate functions for dependent proto files
* func: Specify the path of the custom validation function
## Examples
The validate function(example_validate.pb.go) is generated at the same location as in [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go).
```
cd example

protoc \
  -I . \
  --go_out=. \
  --validator_out=. \
  example.proto
```
By default, `example_validate.pb.go` is generated in the same path as `option go_package`. 
You can generate it to the location you want with the `paths=source_relative:. ` parameter to generate it to the location you want
```
cd example

protoc \
  -I . \
  --go_out=. \
  --go_opt=source_relative \
  --validator_out=. \
  --validator_opt=source_relative \
  example.proto
```
If you want to specify `go module`, you need the prefix of `go module` to be consistent with `option go_package`, e.g.
```
// option go_package="example.com/validator/example";
cd example

protoc \
  -I . \
  --go_out=. \
  --go_opt=module=example.com/validator \
  --validator_out=. \
  --validator_opt=module=example.com/validator \
  example.proto
```
## Usage with hz
- Use adaptations for hz
```
// with no go module
hz new -I={$INCLUDEPATH} \ 
       -idl={$IDLPATH} \
       -protoc-plugins=validator:hz=true:.
// with go module
hz new -I={$INCLUDEPATH} \ 
       -idl={$IDLPATH} \
       -protoc-plugins=validator:hz=true,go_mod={$GOMODULE}:.
       -mod={$GOMODULE}
```
- Use a specific go_package
```
// option go_package={$GOMODULE}/{$MODELDIR}/x/y/z
// {$MODELDIR} defaults to "biz/model" 
hz new -I={$INCLUDEPATH} \ 
       -idl={$IDLPATH} \
       -protoc-plugins=validator:module={$GOMODULE},recurse=true:. \ 
       -mod={$GOMODULE}
```
# API Annotation
In order to use the constraint rules correctly, the file "[api.proto](https://github.com/cloudwego/protoc-gen-validator/blob/main/parser/api/api.proto)" needs to be introduced when writing the 'proto' file

# Constraint rules
> Currently, 'protoc-gen-validator' only supports the basic data types of protobuf, some [WKTs](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf) types, such as Any, Oneofs, etc., will be supported later.<br>
> The annotation "vt" is an abbreviation for "validate".
### Numeric
> All numeric types (`float`, `double`, `int32`, `int64`, `uint32`, `uint64`, `sint32`, `sint64`, `fixed32`, `fixed64`, `sfixed32`, `sfixed64`) share the same constraint rules.
* const: The value of the field must be a specific value
```
int32 Int32Const = 1 [(api.vt).const="123"];
```
* lt/le/gt/ge: denote respectively: <, <=, >, >=
```
optional double DoubleLe = 2 [(api.vt).le="123.54"];
```
* in/not_in: The value of the field must be/not be one of some specific values
```
// Since the 'in' constraint is a list, it is written slightly differently here
optional fixed32 Fix32In = 3 [(api.vt)={in: ["123","456","789"]}];
```
* not_nil: If the field is a pointer, then the pointer cannot be nil
```
optional int64 I64NotNil = 4 [(api.vt).not_nil="true"];
```

### Bool
* const: The value of the field must be a specific value (true/false)
```
optional bool BoolConst = 1 [(api.vt).const="true"];
```
* not_nil: If the field is a pointer, then the pointer cannot be nil
```
optional bool BoolNotNil = 2 [(api.vt).not_nil="true"];
```

### String/Bytes
* const: The value of the field must be a specific value
```
optional string StringConst = 1 [(api.vt).const="validator"];
optional bytes bytesConst = 2 [(api.vt).const="validator"];
```
* pattern: Regular Match
```
optional string StringPattern = 3 [(api.vt).pattern="[0-9A-Za-z]+"];
optional bytes bytesPattern = 4 [(api.vt).pattern="[0-9A-Za-z]+"];
```
* min_size/max_size: Minimum/maximum length
```
optional string StringMinSize = 5 [(api.vt).min_size="12"];
optional bytes bytesMaxSize = 6 [(api.vt).max_size="12"];
```
* prefix/suffix/contains/not_contains: prefix/suffix/contains/not_contains
```
optional string StringPrefix = 7 [(api.vt).prefix="validator"];
optional string StringSuffix = 8 [(api.vt).suffix="validator"];
optional bytes bytesContain = 9 [(api.vt).contains="validator"];
optional bytes bytesNotContain = 10 [(api.vt).not_contains="validator"];
```
* in/not_in: The value of the field must be/not be one of some specific values
```
// Since the 'in' constraint is a list, it is written slightly differently here
optional string StringIn = 11 [(api.vt)={in:["123","456","789"]}];
optional bytes bytesNotIn = 12 [(api.vt)={not_in:["123","456","789"]}];
```
* not_nil: If the field is a pointer, then the pointer cannot be nil
```
optional string StringNotNil = 13 [(api.vt).not_nil="true"];
```

### Enum
```
enum EnumType {
  TWEET = 0;
  RETWEET = 1;
}
```
* const: The value of the field must be a specific value
```
optional EnumType Enum1 = 1 [(api.vt).const="EnumType.TWEET"];
```
* defined_only: The value of the field must be the value defined by the enum
```
optional EnumType Enum2 = 2 [(api.vt).defined_only="true"];
```
* not_nil: If the field is a pointer, then the pointer cannot be nil
```
optional EnumType Enum3 = 3 [(api.vt).not_nil="true"];
```

### Repeated
* min_size/max_size: Minimum/maximum number of elements
```
repeated string ListMinSize = 1 [(api.vt).min_size="12"];
```
* elem: Constraints on elements within a list
```
repeated string ListBaseElem = 2 [(api.vt).elem.const="validator"];
```

### Map
* min_size/max_size: Minimum/maximum number of elements
```
map<int32, string> MapISMinSize = 1 [(api.vt).min_size="10", (api.vt).max_size="30"];
```
* key: For the constraints on the key in the map
```
map<int32, string> MapKey = 2 [(api.vt).key.const="123", (api.vt).key.gt="12"];
```
* value: For the constraints on the value in the map
```
map<int32, string> MapValue = 3 [(api.vt).value.const="validator", (api.vt).value.prefix="validator"];
```
* no_sparse: If the value of a map is a pointer, then the pointer cannot be nil
```
map<int32, MsgType> MapNoSparse = 4 [(api.vt).no_sparse="true"];
```

### Message Field
* skip: Skip validating for this structure
```
optional MapValidate MsgField = 1 [(api.vt).skip="true"];
```
### Message Level Rule
* msg_vt.assert: The result of the expression specified by 'assert' should be "true", in the perspective of the message to validate
```
message StructValidate {
  option (api.msg_vt).assert = "@equal($MsgValidate,1)";
  optional int64 MsgValidate = 1;
}
```

### Cross-field references
* Cross-field references: You can use the value of another field as the constraint value, with the scope of the current structure
```
optional double DoubleLe = 1;
optional double Reference = 2 [(api.vt).le="$DoubleLe"];
```

### Built-in functions
`protoc-gen-validator` provide a set of built-in functions for validating

| function name | arguments                                             | results                                                | remarks                                 |
| ------------- | ----------------------------------------------------- | ------------------------------------------------------ | --------------------------------------- |
| len           | 1: container filed                                    | 1: length of container (integer)                       | just like `len` of go                   |
| sprintf       | 1: format string <br /> 2+: arguments matching format | 1: formatted string (string)                           | just like `fmt.Sprintf` of go           |
| now_unix_nano | none                                                  | 1: nano seconds (int64)                                | just like `time.Now().UnixNano()` of go |
| equal         | 1, 2: comparable values                               | 1: whether two arguments is equal (bool)               | just like `==` of go                    |
| mod           | 1, 2: integer                                         | 1: remainder of $1 / $2 (integer)                      | just like `%` of go                     |
| add           | 1, 2: both are numeric or string                      | 1: sum of two arguments (integer or float64 or string) | just like `+` of go                     |

### Custom validation functions
`protoc-gen-validator` provides a way to expand the validate function
Now you can use parameter `func` to customize your validation function. Like below:
```
cd example

protoc \
  -I . \
  --go_out=. \
  --validator_out=. \
  --validator_opt=func=my_func=path_to_template.txt \
  example.proto
```
`my_func` is the function name, `path_to_template.txt` is the path to template file which should be a go template.
Available template variables:

| variable name | meaning                               | type                                                             |
| ------------- | ------------------------------------- | ---------------------------------------------------------------- |
| Source        | variable name that rule will refer to | string                                                           |
| Function      | data of current function              | *"github.com/cloudwego/protoc-gen-validator/parser".ToolFunction |
