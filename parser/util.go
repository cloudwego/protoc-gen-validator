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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudwego/protoc-gen-validator/parser/api"
	"google.golang.org/protobuf/compiler/protogen"
)

// RulesToAnnotations convert the rule struct to []*Annotation
func RulesToAnnotations(fieldRules *api.FieldRules) ([]*Annotation, error) {
	var validAnno []*Annotation
	ruleBytes, err := json.Marshal(fieldRules)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(ruleBytes, &m)
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		if k == KeyString[MapKey] || k == KeyString[MapValue] || k == KeyString[Elem] {
			rule := v.(map[string]interface{})
			ret, err := getElemRule(rule, k)
			if err != nil {
				return nil, err
			}
			for k, v := range ret {
				var value []string
				for _, val := range v {
					value = append(value, val)
				}
				anno := &Annotation{
					Key:    "vt." + k,
					Values: value,
				}
				validAnno = append(validAnno, anno)
			}
		} else if k == KeyString[In] || k == KeyString[NotIn] {
			var values []string
			for _, val := range v.([]interface{}) {
				values = append(values, val.(string))
			}
			anno := &Annotation{
				Key:    "vt." + k,
				Values: values,
			}
			validAnno = append(validAnno, anno)
		} else {
			var values []string
			values = append(values, v.(string))
			anno := &Annotation{
				Key:    "vt." + k,
				Values: values,
			}
			validAnno = append(validAnno, anno)
		}
	}

	return validAnno, nil
}

func getElemRule(rules map[string]interface{}, originkey string) (map[string][]string, error) {
	ret := make(map[string][]string, len(rules))
	// elem rule don't nest elem rule in protobuf, so no need to process "MapKey"、"MapValue"、"Elem"
	for ruleKey, ruleContent := range rules {
		if ruleKey == KeyString[In] || ruleKey == KeyString[NotIn] {
			var value []string
			for _, val := range ruleContent.([]interface{}) {
				value = append(value, val.(string))
			}
			ret[originkey+"."+ruleKey] = value
		} else {
			var value []string
			value = append(value, ruleContent.(string))
			ret[originkey+"."+ruleKey] = value
		}
	}
	return ret, nil
}

func getFieldReferenceValidation(msg *protogen.Message, anno string) (*ValidationValue, error) {
	if !strings.HasPrefix(anno, "$") {
		return nil, nil
	}
	f := getFieldReference(anno, msg)
	if f == nil {
		return nil, fmt.Errorf("filed reference %s not found", anno)
	}
	return &ValidationValue{
		ValueType:  FieldReferenceValue,
		TypedValue: TypedValidationValue{FieldReference: f},
	}, nil
}

func getFieldReference(anno string, msg *protogen.Message) *protogen.Field {
	name := strings.TrimPrefix(anno, "$")
	for _, f := range msg.Fields {
		rawName := string(f.Desc.Name())
		if rawName == name {
			return f
		}
	}

	return nil
}

func (p *Function) pegText(node *node32) string {
	for n := node; n != nil; n = n.next {
		if s := p.pegText(n.up); s != "" {
			return s
		}
		if n.pegRule != rulePegText {
			continue
		}
		text := string(p.buffer[int(n.begin):int(n.end)])
		if text != "" {
			return text
		}
	}
	return ""
}

func (p *Function) pegAllText(node *node32) (s string) {
	for n := node; n != nil; n = n.next {
		text := string(p.buffer[int(n.begin):int(n.end)])
		if text != "" {
			return text
		}
	}
	return ""
}

func getFunctionValidation(st *protogen.Message, anno string) (*ValidationValue, error) {
	if !strings.HasPrefix(anno, "@") {
		return nil, nil
	}
	f, err := parseFunction(anno, st)
	if err != nil {
		return nil, err
	}
	node := f.AST().up
	// '@' Identifier LPAR Arguments RPAR
	name := f.pegText(node)
	node = node.next
	node = node.next // skip LPAR
	node = node.up   // Arguments
	arguments, err := parseFunctionArguments(st, f, node)
	if err != nil {
		return nil, err
	}
	return &ValidationValue{
		ValueType: FunctionValue,
		TypedValue: TypedValidationValue{Function: &ToolFunction{
			Name:      name,
			Arguments: arguments,
		}},
	}, nil
}

func parseFunction(anno string, st *protogen.Message) (*Function, error) {
	f := &Function{
		Buffer: anno,
	}
	if err := f.Init(); err != nil {
		return nil, err
	}
	if err := f.Parse(); err != nil {
		return nil, err
	}
	return f, nil
}

func parseFunctionArguments(st *protogen.Message, p *Function, node *node32) ([]ValidationValue, error) {
	// (ConstValue ListSeparator?)*
	var ret []ValidationValue
	for ; node != nil; node = node.next {
		switch node.pegRule {
		case ruleListSeparator:
			continue
		case ruleConstValue:
			node := node.up
			switch node.pegRule {
			case ruleDoubleConstant:
				value, err := strconv.ParseFloat(p.pegText(node), 64)
				if err != nil {
					return nil, err
				}
				ret = append(ret, ValidationValue{
					ValueType: DoubleValue,
					TypedValue: TypedValidationValue{
						Double: value,
					},
				})
			case ruleIntConstant:
				value, err := strconv.ParseInt(p.pegText(node), 0, 64)
				if err != nil {
					return nil, err
				}
				ret = append(ret, ValidationValue{
					ValueType: IntValue,
					TypedValue: TypedValidationValue{
						Int: value,
					},
				})
			case ruleLiteral:
				ret = append(ret, ValidationValue{
					ValueType: BinaryValue,
					TypedValue: TypedValidationValue{
						Binary: p.pegText(node),
					},
				})
			case ruleFieldReference:
				val, err := getFieldReferenceValidation(st, "$"+p.pegText(node))
				if err != nil {
					return nil, err
				}
				if val != nil {
					ret = append(ret, *val)
				}
			default:
				return nil, fmt.Errorf("unsupported const value %s for function arguments", rul3s[node.pegRule])
			}
		case ruleFunction:
			fv, err := getFunctionValidation(st, p.pegAllText(node))
			if err != nil {
				return nil, fmt.Errorf("unsupported const value %s for function arguments", rul3s[node.pegRule])
			}
			ret = append(ret, ValidationValue{
				ValueType: FunctionValue,
				TypedValue: TypedValidationValue{
					Function: fv.TypedValue.Function,
				},
			})
		default:
			return nil, fmt.Errorf("unsupported rule %s for function arguments", rul3s[node.pegRule])
		}
	}
	return ret, nil
}
