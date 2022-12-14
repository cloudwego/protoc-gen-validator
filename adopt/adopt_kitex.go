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

package adopt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/protoc-gen-validator/validator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const KitexGenPath = "kitex_gen"

func AdoptKitex(gen *protogen.Plugin) error {
	adopter := &KitexAdopter{
		Gen: gen,
	}
	err := adopter.Init()
	if err != nil {
		return err
	}
	newGen, err := adopter.newPluginGen()
	if err != nil {
		return err
	}
	newGen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, f := range newGen.Files {
		if strings.HasPrefix(f.Proto.GetPackage(), "google.protobuf") {
			continue
		}
		gopkg := f.Proto.GetOptions().GetGoPackage()
		if !strings.HasPrefix(gopkg, adopter.PackagePrefix) {
			fmt.Fprintf(os.Stderr, "[WARN] %q is skipped because its import path %q is not located in ./kitex_gen. Change the go_package option or use '--protobuf M%s=A-Import-Path-In-kitex_gen' to override it if you want this file to be generated under kitex_gen.\n",
				f.Proto.GetName(), gopkg, f.Proto.GetName())
			continue
		}
		f.GeneratedFilenamePrefix = strings.TrimPrefix(f.GeneratedFilenamePrefix, adopter.PackagePrefix)
		g, err := validator.NewGenerator(newGen, f)
		if err != nil {
			return err
		}
		err = g.Generate()
		if err != nil {
			return err
		}
	}

	*gen = *newGen
	return nil
}

type KitexAdopter struct {
	Gen           *protogen.Plugin
	ModuleName    string
	PackagePrefix string
}

func (ad *KitexAdopter) Init() error {
	params := strings.Split(ad.Gen.Request.GetParameter(), ",")
	// get module name
	for _, param := range params {
		if strings.Contains(param, "=") {
			ret := strings.Split(param, "=")
			if ret[0] == "GoMod" {
				ad.ModuleName = ret[1]
			}
		}
	}

	// get package prefix
	err := ad.checkPath()
	if err != nil {
		return err
	}

	return nil
}

func (ad *KitexAdopter) newPluginGen() (*protogen.Plugin, error) {
	oldRequest := ad.Gen.Request
	err := ad.ModifyGoPackage(oldRequest)
	if err != nil {
		return nil, err
	}
	newGen, err := protogen.Options{}.New(oldRequest)
	if err != nil {
		return nil, err
	}
	return newGen, nil
}

func (ad *KitexAdopter) ModifyGoPackage(req *pluginpb.CodeGeneratorRequest) error {
	for _, f := range req.ProtoFile {
		if strings.HasPrefix(f.GetPackage(), "google.protobuf") {
			continue
		}

		if f.Options == nil || f.Options.GoPackage == nil {
			return fmt.Errorf("ERROR: go_package is missing in proto file %q", f.GetName())
		}
		gopkg := f.GetOptions().GetGoPackage()
		if path, ok := ad.getImportPath(gopkg); ok {
			f.Options.GoPackage = &path
		}
	}

	return nil
}

func (ad *KitexAdopter) getImportPath(pkg string) (path string, ok bool) {
	parts := strings.Split(pkg, "/")
	if len(parts) == 0 {
		// malformed import path
		return "", false
	}

	if strings.HasPrefix(pkg, ad.PackagePrefix) {
		return pkg, true
	}
	if strings.Contains(parts[0], ".") || (ad.ModuleName != "" && strings.HasPrefix(pkg, ad.ModuleName)) {
		// already a complete import path, but outside the target path
		return "", false
	}
	// incomplete import path
	return filepath.Join(ad.PackagePrefix, pkg), true
}

func (ad *KitexAdopter) checkPath() error {
	goPath, err := GetGOPATH()
	if err != nil {
		return err
	}
	gosrc := filepath.Join(goPath, "src")
	gosrc, err = filepath.Abs(gosrc)
	if err != nil {
		return err
	}
	curpath, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	if strings.HasPrefix(curpath, gosrc) {
		if ad.PackagePrefix, err = filepath.Rel(gosrc, curpath); err != nil {
			return fmt.Errorf("get GOPATH/src relpath failed: %v", err.Error())
		}
		ad.PackagePrefix = filepath.Join(ad.PackagePrefix, KitexGenPath)
	} else {
		if ad.ModuleName == "" {
			return fmt.Errorf("outside of $GOPATH. Please specify a module name with the 'GoMod' flag for validator")
		}
	}

	if ad.ModuleName != "" {
		module, path, ok := SearchGoMod(curpath)
		if ok {
			// go.mod exists
			if module != ad.ModuleName {
				return fmt.Errorf("The module name given by the '-module' option ('%s') is not consist with the name defined in go.mod ('%s' from %s)\n",
					ad.ModuleName, module, path)
			}
			if ad.PackagePrefix, err = filepath.Rel(path, curpath); err != nil {
				return fmt.Errorf("get package prefix failed: %v", err.Error())
			}
			ad.PackagePrefix = filepath.Join(ad.ModuleName, ad.PackagePrefix, KitexGenPath)
		} else {
			ad.PackagePrefix = filepath.Join(ad.ModuleName, KitexGenPath)
		}
	}

	return nil
}
