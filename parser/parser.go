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
	"log"
	"strconv"

	"github.com/cloudwego/protoc-gen-validator/parser/api"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(msg *protogen.Message) (*Validation, map[protoreflect.FieldNumber]*Validation, error) {
	ret := make(map[protoreflect.FieldNumber]*Validation)
	for _, f := range msg.Fields {
		fieldAnnos := proto.GetExtension(f.Desc.Options(), api.E_Vt)
		validAnnotations, err := RulesToAnnotations(fieldAnnos.(*api.FieldRules))
		if err != nil {
			return nil, nil, err
		}
		v, err := p.parseField(msg, f.Desc, validAnnotations, f.Desc.IsList(), f.Desc.IsMap())
		if err != nil {
			return nil, nil, err
		}
		f.Desc.Number()
		ret[f.Desc.Number()] = v
	}
	msgAnno := proto.GetExtension(msg.Desc.Options(), api.E_MsgVt)
	msgRule, err := RulesToAnnotations(msgAnno.(*api.FieldRules))
	if err != nil {
		return nil, nil, err
	}
	v, err := p.parseStruct(msg, msgRule)
	if err != nil {
		return nil, nil, fmt.Errorf("[annotation parser] parse %s's annotations failed: %w", msg.Desc.Name(), err)
	}

	return v, ret, nil
}

func (p *Parser) parseStruct(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: StructLikeValidation}
	rf := NewRuleFactory(StructLikeKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFunctionValidation(msg, annoVal)
			if err != nil {
				log.Printf("%s parse as a function failed: %v", annoVal, err)
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized struct-like annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}

func (p *Parser) parseField(msg *protogen.Message, field protoreflect.FieldDescriptor, annotations []*Annotation, isList, isMap bool) (*Validation, error) {
	if isList {
		return p.parseList(msg, field, annotations)
	}

	if isMap {
		return p.parseMap(msg, field, annotations)
	}

	switch field.Kind() {
	case protoreflect.BoolKind:
		return p.parseBool(msg, annotations)
	case protoreflect.EnumKind:
		return p.parseEnum(msg, annotations)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind,
		protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind:
		return p.parseInt(msg, annotations)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return p.parseDouble(msg, annotations)
	case protoreflect.StringKind, protoreflect.BytesKind:
		return p.parseBytes(msg, annotations)
	case protoreflect.MessageKind:
		return p.parseMessageField(msg, annotations)
	default:
		return nil, fmt.Errorf("type %s not recognized", field.Kind())
	}
}

func (p *Parser) parseInt(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: NumericValidation}
	rf := NewRuleFactory(NumericKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case Const,
					LessThan,
					LessEqual,
					GreatThan,
					GreatEqual,
					In,
					NotIn:
					val, err := strconv.ParseInt(annoVal, 0, 64)
					if err != nil {
						return nil, fmt.Errorf("parse int value failed: %w", err)
					}
					value = &ValidationValue{
						ValueType:  IntValue,
						TypedValue: TypedValidationValue{Int: val},
					}
				case NotNil:
					val, err := strconv.ParseBool(annoVal)
					if err != nil {
						return nil, fmt.Errorf("parse int value failed: %w", err)
					}
					value = &ValidationValue{
						ValueType:  BoolValue,
						TypedValue: TypedValidationValue{Bool: val},
					}
				default:
					return nil, fmt.Errorf("unrecognized numeric annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized numeric annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}

	return validation, nil
}

func (p *Parser) parseBool(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: BoolValidation}
	rf := NewRuleFactory(BoolKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				val, err := strconv.ParseBool(annoVal)
				if err != nil {
					return nil, fmt.Errorf("parse bool value failed: %w", err)
				}
				value = &ValidationValue{
					ValueType:  BoolValue,
					TypedValue: TypedValidationValue{Bool: val},
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized bool annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}

func (p *Parser) parseDouble(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: NumericValidation}
	rf := NewRuleFactory(NumericKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case Const,
					LessThan,
					LessEqual,
					GreatThan,
					GreatEqual,
					In,
					NotIn:
					val, err := strconv.ParseFloat(annoVal, 64)
					if err != nil {
						return nil, fmt.Errorf("parse double value failed: %w", err)
					}
					value = &ValidationValue{
						ValueType:  DoubleValue,
						TypedValue: TypedValidationValue{Double: val},
					}
				case NotNil:
					val, err := strconv.ParseBool(annoVal)
					if err != nil {
						return nil, fmt.Errorf("parse int value failed: %w", err)
					}
					value = &ValidationValue{
						ValueType:  BoolValue,
						TypedValue: TypedValidationValue{Bool: val},
					}
				default:
					return nil, fmt.Errorf("unrecognized numeric annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized numeric annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}

func (p *Parser) parseBytes(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: BinaryValidation}
	rf := NewRuleFactory(BinaryKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case Const,
					Pattern,
					Prefix,
					Suffix,
					Contains,
					NotContains,
					In,
					NotIn:
					value = &ValidationValue{
						ValueType:  BinaryValue,
						TypedValue: TypedValidationValue{Binary: annoVal},
					}
				case MinSize,
					MaxSize:
					len, err := strconv.ParseInt(annoVal, 0, 64)
					if err != nil {
						return nil, err
					}
					value = &ValidationValue{
						ValueType:  IntValue,
						TypedValue: TypedValidationValue{Int: len},
					}
				case NotNil:
					val, err := strconv.ParseBool(annoVal)
					if err != nil {
						return nil, err
					}
					value = &ValidationValue{
						ValueType:  BoolValue,
						TypedValue: TypedValidationValue{Bool: val},
					}
				default:
					return nil, fmt.Errorf("unrecognized binary annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized binary annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}

func (p *Parser) parseMessageField(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: StructLikeFieldValidation}
	rf := NewRuleFactory(StructLikeFieldKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				val, err := strconv.ParseBool(annoVal)
				if err != nil {
					return nil, fmt.Errorf("parse struct-like value failed: %w", err)
				}
				value = &ValidationValue{
					ValueType:  BoolValue,
					TypedValue: TypedValidationValue{Bool: val},
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized struct-like annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}

func (p *Parser) parseList(msg *protogen.Message, field protoreflect.FieldDescriptor, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: ListValidation}
	rf := NewRuleFactory(ListKeys)
	var elemAnnotations []*Annotation
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		if nodeKey == Elem {
			elemKey := kp.toElemKey()
			elemAnnotations = append(elemAnnotations, &Annotation{Key: elemKey, Values: annoVals})
			continue
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case MinSize,
					MaxSize:
					len, err := strconv.ParseInt(annoVal, 0, 64)
					if err != nil {
						return nil, err
					}
					value = &ValidationValue{
						ValueType:  IntValue,
						TypedValue: TypedValidationValue{Int: len},
					}
				default:
					return nil, fmt.Errorf("unrecognized list annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized list annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}

	if len(elemAnnotations) > 0 {
		elemValidation, err := p.parseField(msg, field, elemAnnotations, false, false)
		if err != nil {
			return nil, fmt.Errorf("parse element annotation failed: %w", err)
		}
		validation.Rules = append(validation.Rules, &Rule{
			Key:   Elem,
			Inner: elemValidation,
		})
	}

	return validation, nil
}

func (p *Parser) parseMap(msg *protogen.Message, field protoreflect.FieldDescriptor, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: MapValidation}
	rf := NewRuleFactory(MapKeys)
	var keyAnnotations []*Annotation
	var valAnnotations []*Annotation
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		if nodeKey == MapKey {
			elemKey := kp.toElemKey()
			keyAnnotations = append(keyAnnotations, &Annotation{Key: elemKey, Values: annoVals})
			continue
		} else if nodeKey == MapValue {
			elemKey := kp.toElemKey()
			valAnnotations = append(valAnnotations, &Annotation{Key: elemKey, Values: annoVals})
			continue
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case MinSize, MaxSize:
					len, err := strconv.ParseInt(annoVal, 0, 64)
					if err != nil {
						return nil, err
					}
					value = &ValidationValue{
						ValueType:  IntValue,
						TypedValue: TypedValidationValue{Int: len},
					}
				case NoSparse:
					val, err := strconv.ParseBool(annoVal)
					if err != nil {
						return nil, err
					}
					value = &ValidationValue{
						ValueType:  BoolValue,
						TypedValue: TypedValidationValue{Bool: val},
					}
				default:
					return nil, fmt.Errorf("unrecognized map annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized map annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	if len(keyAnnotations) > 0 {
		keyValidation, err := p.parseField(msg, field.MapKey(), keyAnnotations, false, false)
		if err != nil {
			return nil, fmt.Errorf("parse key annotation failed: %w", err)
		}
		validation.Rules = append(validation.Rules, &Rule{
			Key:   MapKey,
			Inner: keyValidation,
		})
	}
	if len(valAnnotations) > 0 {
		valValidation, err := p.parseField(msg, field.MapValue(), valAnnotations, false, false)
		if err != nil {
			return nil, fmt.Errorf("parse value annotation failed: %w", err)
		}
		validation.Rules = append(validation.Rules, &Rule{
			Key:   MapValue,
			Inner: valValidation,
		})
	}
	return validation, nil
}

func (p *Parser) parseEnum(msg *protogen.Message, annotations []*Annotation) (*Validation, error) {
	validation := &Validation{ValidationType: EnumValidation}
	rf := NewRuleFactory(EnumKeys)
	for _, anno := range annotations {
		annoKey, annoVals := anno.Key, anno.Values
		kp, err := newKeyParser(annoKey)
		if err != nil {
			return nil, err
		}
		nodeStr := kp.next()
		nodeKey, ok := KeyFromString(nodeStr)
		if !ok {
			return nil, fmt.Errorf("invalid key %s", nodeStr)
		}
		for _, annoVal := range annoVals {
			value, err := getFieldReferenceValidation(msg, annoVal)
			if err != nil {
				return nil, err
			}
			if value == nil {
				value, err = getFunctionValidation(msg, annoVal)
				if err != nil {
					log.Printf("%s parse as a function failed: %v\n", annoVal, err)
				}
			}
			if value == nil {
				switch nodeKey {
				case Const:
					value = &ValidationValue{
						ValueType:  BinaryValue,
						TypedValue: TypedValidationValue{Binary: annoVal},
					}
				case DefinedOnly, NotNil:
					val, err := strconv.ParseBool(annoVal)
					if err != nil {
						return nil, fmt.Errorf("parse bool value failed: %w", err)
					}
					value = &ValidationValue{
						ValueType:  BoolValue,
						TypedValue: TypedValidationValue{Bool: val},
					}
				default:
					return nil, fmt.Errorf("unrecognized enum annotation key %s", annoKey)
				}
			}
			exist, rule := rf.NewRule(nodeKey, value)
			if !exist {
				return nil, fmt.Errorf("unrecognized enum annotation key %s", annoKey)
			}
			if rule != nil {
				validation.Rules = append(validation.Rules, rule)
			}
		}
	}
	return validation, nil
}
