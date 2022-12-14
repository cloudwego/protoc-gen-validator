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
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func ImportToPathAndConcat(path, subFix string) string {
	path = strings.TrimSuffix(path, subFix)
	path = strings.ReplaceAll(path, "/", string(filepath.Separator))
	if i := strings.LastIndex(path, string(filepath.Separator)); i >= 0 && i < len(path)-1 && strings.Contains(path[i+1:], ".") {
		base := strings.ReplaceAll(path[i+1:], ".", "_")
		dir := path[:i]
		return dir + string(filepath.Separator) + base
	}
	return path
}

func PathToImport(path, subFix string) string {
	path = strings.TrimSuffix(path, subFix)
	// path = RelativePath(path)
	return strings.ReplaceAll(path, string(filepath.Separator), "/")
}

func ImportToPath(path, subFix string) string {
	// path = RelativePath(path)
	return strings.ReplaceAll(path, "/", string(filepath.Separator)) + subFix
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// BaseName get base name for path. ex: "github.com/p.s.m" => "p.s.m"
func BaseName(include, subFixToTrim string) string {
	include = unifyPath(include)
	subFixToTrim = unifyPath(subFixToTrim)
	last := include
	if id := strings.LastIndex(last, "/"); id >= 0 && id < len(last)-1 {
		last = last[id+1:]
	}
	if !strings.HasSuffix(last, subFixToTrim) {
		return last
	}
	return last[:len(last)-len(subFixToTrim)]
}

// unifyPath will convert "\" to "/" in path if the os is windows
func unifyPath(path string) string {
	if IsWindows() {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}

// GetGOPATH retrieves the GOPATH from environment variables or the `go env` command.
func GetGOPATH() (string, error) {
	goPath := os.Getenv("GOPATH")
	// If there are many path in GOPATH, pick up the first one.
	if GoPaths := strings.Split(goPath, ":"); len(GoPaths) > 1 {
		return GoPaths[0], nil
	}
	// GOPATH not set through environment variables, try to get one by executing "go env GOPATH"
	output, err := exec.Command("go", "env", "GOPATH").Output()
	if err != nil {
		return "", err
	}

	goPath = strings.TrimSpace(string(output))
	if len(goPath) == 0 {
		buildContext := build.Default
		goPath = buildContext.GOPATH
	}

	if len(goPath) == 0 {
		return "", fmt.Errorf("GOPATH not found")
	}
	return goPath, nil
}

// SearchGoMod searches go.mod from the given directory (which must be an absolute path) to
// the root directory. When the go.mod is found, its module name and path will be returned.
func SearchGoMod(cwd string) (moduleName, path string, found bool) {
	for {
		path = filepath.Join(cwd, "go.mod")
		data, err := os.ReadFile(path)
		if err == nil {
			re := regexp.MustCompile(`^\s*module\s+(\S+)\s*`)
			for _, line := range strings.Split(string(data), "\n") {
				m := re.FindStringSubmatch(line)
				if m != nil {
					return m[1], cwd, true
				}
			}
			return fmt.Sprintf("<module name not found in '%s'>", path), path, true
		}

		if !os.IsNotExist(err) {
			return
		}
		if cwd == "/" {
			break
		}
		cwd = filepath.Dir(cwd)
	}
	return
}
