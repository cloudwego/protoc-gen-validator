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
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/protoc-gen-validator/validator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func AdoptHz(gen *protogen.Plugin) error {
	adopter := &HzAdopter{
		Gen: gen,
	}
	adopter.Init()
	newGen, err := adopter.newPluginGen()
	if err != nil {
		return err
	}
	newGen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, f := range newGen.Files {
		if strings.HasPrefix(f.Proto.GetPackage(), "google.protobuf") {
			continue
		}
		g, err := validator.NewGenerator(newGen, f)
		if err != nil {
			return err
		}
		impt := string(f.GoImportPath)
		if strings.HasPrefix(impt, adopter.GoModule) {
			impt = impt[len(adopter.GoModule):]
		}
		f.GeneratedFilenamePrefix = filepath.Join(ImportToPath(impt, ""), BaseName(f.Proto.GetName(), ".proto"))
		err = g.Generate()
		if err != nil {
			return err
		}
	}

	*gen = *newGen
	return nil
}

type HzAdopter struct {
	Gen      *protogen.Plugin
	OutDir   string
	GoModule string
	ModelDir string
}

func (ad *HzAdopter) Init() {
	params := strings.Split(ad.Gen.Request.GetParameter(), ",")
	// init parameters
	ad.ModelDir = "biz/model"
	ad.SetGoModule()
	for _, param := range params {
		if strings.Contains(param, "=") {
			ret := strings.Split(param, "=")
			if ret[0] == "out_dir" {
				ad.OutDir = ret[1]
			}
			if ret[0] == "go_mod" {
				ad.GoModule = ret[1]
			}
			if ret[0] == "model_dir" {
				ad.ModelDir = ret[1]
			}
		}
	}
}

func (ad *HzAdopter) newPluginGen() (*protogen.Plugin, error) {
	oldRequest := ad.Gen.Request
	ad.ModifyGoPackage(oldRequest)
	newGen, err := protogen.Options{}.New(oldRequest)
	if err != nil {
		return nil, err
	}
	return newGen, nil
}

func (ad *HzAdopter) ModifyGoPackage(req *pluginpb.CodeGeneratorRequest) {
	goMod := ad.GoModule
	for _, f := range req.ProtoFile {
		if strings.HasPrefix(f.GetPackage(), "google.protobuf") {
			continue
		}
		goPkg := getGoPackage(f, nil)
		if !strings.Contains(goPkg, goMod) {
			if strings.HasPrefix(goPkg, "/") {
				goPkg = goMod + goPkg
			} else {
				goPkg = goMod + "/" + goPkg
			}
		}
		impt, _ := ad.fixModelPathAndPackage(goPkg)
		*f.Options.GoPackage = impt
	}
}

func (ad *HzAdopter) fixModelPathAndPackage(pkg string) (impt, path string) {
	if strings.HasPrefix(pkg, ad.GoModule) {
		impt = ImportToPathAndConcat(pkg[len(ad.GoModule):], "")
		if !strings.HasPrefix(impt, "/") {
			impt = "/" + impt
		}
	}
	if ad.ModelDir != "" && ad.ModelDir != "." {
		modelImpt := PathToImport(string(filepath.Separator)+ad.ModelDir, "")
		// trim model dir for go package
		if strings.HasPrefix(impt, modelImpt) {
			impt = impt[len(modelImpt):]
		}
		impt = PathToImport(ad.ModelDir, "") + impt
	}
	path = ImportToPath(impt, "")
	impt = ad.GoModule + "/" + impt
	if IsWindows() {
		impt = PathToImport(impt, "")
	}
	return
}

func (ad *HzAdopter) SetGoModule() {
	goPath, err := GetGOPATH()
	if err != nil {
		return
	}
	if goPath == "" {
		return
	}
	goSrc := filepath.Join(goPath, "src")
	cwd, _ := os.Getwd()

	// Generate the project under gopath, use the relative path as the package name
	if strings.HasPrefix(cwd, goSrc) {
		if goMod, _ := filepath.Rel(goSrc, cwd); err != nil {
			return
		} else {
			ad.GoModule = goMod
		}
	}
}

// getGoPackage  get option go_package
func getGoPackage(f *descriptorpb.FileDescriptorProto, pkgMap map[string]string) string {
	if f.Options == nil {
		f.Options = new(descriptorpb.FileOptions)
	}
	if f.Options.GoPackage == nil {
		f.Options.GoPackage = new(string)
	}
	goPkg := f.Options.GetGoPackage()

	// if go_package has ";", for example go_package="/a/b/c;d", we will use "/a/b/c" as go_package
	if strings.Contains(goPkg, ";") {
		pkg := strings.Split(goPkg, ";")
		if len(pkg) == 2 {
			goPkg = pkg[0]
		}

	}

	if goPkg == "" {
		goPkg = f.GetPackage()
	}
	if opt, ok := pkgMap[f.GetName()]; ok {
		return opt
	}
	return goPkg
}
