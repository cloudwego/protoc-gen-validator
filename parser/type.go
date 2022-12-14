// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

type ValidationType int

const (
	NumericValidation ValidationType = iota
	BinaryValidation
	BoolValidation
	EnumValidation
	ListValidation
	MapValidation
	StructLikeFieldValidation
	StructLikeValidation
)

type ValueType int

const (
	FieldReferenceValue ValueType = iota
	DoubleValue
	IntValue
	BoolValue
	EnumValue
	BinaryValue
	FunctionValue
)

var ValueTypeName = [...]string{
	FieldReferenceValue: "field-reference-value",
	DoubleValue:         "double-value",
	IntValue:            "int-value",
	BoolValue:           "bool-value",
	EnumValue:           "enum-value",
	BinaryValue:         "binary-value",
	FunctionValue:       "function-value",
}

func (vt ValueType) String() string {
	return ValueTypeName[vt]
}

type Validation struct {
	ValidationType ValidationType
	Rules          []*Rule
}

type Rule struct {
	Key       Key
	Specified *ValidationValue
	Range     []*ValidationValue
	Inner     *Validation
}

type ToolFunction struct {
	Name      string
	Arguments []ValidationValue
}

type ValidationValue struct {
	ValueType  ValueType
	TypedValue TypedValidationValue
}

type TypedValidationValue struct {
	FieldReference *protogen.Field
	Double         float64
	Int            int64
	Bool           bool
	// Enum           *tp.EnumValue
	Binary   string
	Function *ToolFunction
}

func (t *TypedValidationValue) GetFieldReferenceName(ref string) string {
	fName := t.FieldReference.GoName
	reference := ref + fmt.Sprintf("Get%s()", fName)
	return reference
}

type RuleFactory struct {
	specifiedKeys []Key
	rangeKeys     []Key
	rangeRules    []*Rule
}

func NewRuleFactory(keys []Key) *RuleFactory {
	rangeKeys := PickRangeKeys(keys)
	specifiedKeys := PickSpecifiedKeys(keys)
	rangeRules := make([]*Rule, len(rangeKeys))
	return &RuleFactory{
		specifiedKeys: specifiedKeys,
		rangeKeys:     rangeKeys,
		rangeRules:    rangeRules,
	}
}

func (v *RuleFactory) NewRule(key Key, value *ValidationValue) (bool, *Rule) {
	for _, sk := range v.specifiedKeys {
		if sk == key {
			return true, &Rule{
				Key:       key,
				Specified: value,
			}
		}
	}
	for i, rk := range v.rangeKeys {
		if rk == key {
			if v.rangeRules[i] == nil {
				v.rangeRules[i] = &Rule{
					Key:   key,
					Range: []*ValidationValue{value},
				}
				return true, v.rangeRules[i]
			} else {
				v.rangeRules[i].Range = append(v.rangeRules[i].Range, value)
				return true, nil
			}
		}
	}
	return false, nil
}
