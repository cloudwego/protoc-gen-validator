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

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/protoc-gen-validator/adopt"
	"github.com/cloudwego/protoc-gen-validator/validator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), validator.Version)
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		recurse = flags.Bool("recurse", false, "recurse generate")
		_       = flags.String("func", "", "customize function")
		isHz    = flags.Bool("hz", false, "adopt hz")
		isKitex = flags.Bool("kitex", false, "adopt kitex")
		_       = flags.String("out_dir", ".", "output dir")
		_       = flags.String("go_mod", "", "go module")
		_       = flags.String("model_dir", "biz/model", "model dir")
	)

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *isHz {
			err := adopt.AdoptHz(gen)
			if err != nil {
				return err
			}
			return nil
		}

		if *isKitex {
			err := adopt.AdoptKitex(gen)
			if err != nil {
				return err
			}
			return nil
		}

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if strings.HasPrefix(f.Proto.GetPackage(), "google.protobuf") {
				continue
			}
			g, err := validator.NewGenerator(gen, f)
			if err != nil {
				return err
			}
			if *recurse {
				err = g.Generate()
				if err != nil {
					return err
				}
			} else {
				if f.Generate {
					err = g.Generate()
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
}
