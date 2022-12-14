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

package validator

import (
	"fmt"

	"github.com/cloudwego/protoc-gen-validator/parser"
	"google.golang.org/protobuf/compiler/protogen"
)

type ValidateContext struct {
	*parser.Validation
	RawField     *protogen.Field
	PbFile       *protogen.File
	Msg          *protogen.Message
	FieldName    string // Go name for field
	RawFieldName string // raw field name in idl
	GetNameFunc  string // Get***() func for getting the generated value
	IsOptional   bool
	ids          map[string]int
}

func mkMsgValidateContext(message *protogen.Message, file *protogen.File) ([]*ValidateContext, error) {
	var ret []*ValidateContext
	p := parser.NewParser()
	msgValidation, fieldValidations, err := p.Parse(message)
	if err != nil {
		return nil, err
	}
	ids := map[string]int{}
	for _, field := range message.Fields {
		ret = append(ret, &ValidateContext{
			PbFile:       file,
			FieldName:    field.GoName,
			RawFieldName: string(field.Desc.Name()),
			GetNameFunc:  "m.Get" + field.GoName + "()",
			IsOptional:   field.Desc.HasOptionalKeyword(),
			ids:          ids,
			Validation:   fieldValidations[field.Desc.Number()],
			RawField:     field,
			Msg:          message,
		})
	}
	ret = append(ret, &ValidateContext{
		PbFile:     file,
		Msg:        message,
		Validation: msgValidation,
		ids:        ids,
	})

	return ret, nil
}

func (v *ValidateContext) GenID(prefix string) (name string) {
	name = prefix
	if id := v.ids[prefix]; id > 0 {
		name += fmt.Sprint(id)
	}
	v.ids[prefix]++
	return
}
