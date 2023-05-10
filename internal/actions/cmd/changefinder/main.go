// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/internal/actions/logg"
)

var (
	dir       = flag.String("dir", "", "the root directory to evaluate")
	format    = flag.String("format", "plain", "output format, one of [plain|github], defaults to 'plain'")
	ghVarName = flag.String("gh-var", "submodules", "github format's variable name to set output for, defaults to 'submodules'.")
	base      = flag.String("base", "origin/main", "the base ref to compare to, defaults to 'origin/main'")
)

func main() {
	flag.BoolVar(&logg.Quiet, "q", false, "quiet mode, minimal logging")
	flag.Parse()
	rootDir, err := os.Getwd()
	if err != nil {
		logg.Fatal(err)
	}
	if *dir != "" {
		rootDir = *dir
	}
	logg.Printf("Root dir: %q", rootDir)

	submodules, err := mods(rootDir)
	if err != nil {
		logg.Fatal(err)
	}

	changes, err := gitFilesChanges(rootDir)
	if err != nil {
		logg.Fatalf("unable to get files changed: %v", err)
	}

	modulesSeen := map[string]bool{}
	updatedSubmodules := []string{}
	for _, change := range changes {
		if strings.HasPrefix(change, "internal") {
			continue
		}
		submod, ok := owner(change, submodules)
		if !ok {
			logg.Printf("no module for: %s", change)
			continue
		}
		if _, seen := modulesSeen[submod]; !seen {
			logg.Printf("changes in submodule: %s", submod)
			updatedSubmodules = append(updatedSubmodules, submod)
			modulesSeen[submod] = true
		}
	}

	output(updatedSubmodules)
}

func output(s []string) error {
	switch *format {
	case "github":
		b, err := json.Marshal(s)
		if err != nil {
			logg.Fatalf("unable to marshal submodules: %v", err)
		}
		fmt.Printf("::set-output name=%s::%s", *ghVarName, b)
	case "plain":
		fallthrough
	default:
		fmt.Println(strings.Join(s, "\n"))
	}
	return nil
}

func owner(file string, submodules []string) (string, bool) {
	submod := ""
	for _, mod := range submodules {
		if strings.HasPrefix(file, mod) && len(mod) > len(submod) {
			submod = mod
		}
	}

	return submod, submod != ""
}

func mods(dir string) (submodules []string, err error) {
	c := exec.Command("go", "list", "-m")
	c.Dir = dir
	b, err := c.Output()
	if err != nil {
		return submodules, err
	}
	list := strings.Split(strings.TrimSpace(string(b)), "\n")

	submodules = []string{}
	for _, mod := range list {
		// Skip non-submodule or internal submodules.
		if mod == "cloud.google.com/go" || strings.Contains(mod, "internal") {
			continue
		}
		logg.Printf("found module: %s", mod)
		mod = strings.TrimPrefix(mod, "cloud.google.com/go/")
		submodules = append(submodules, mod)
	}

	return submodules, nil
}

func gitFilesChanges(dir string) ([]string, error) {
	c := exec.Command("git", "diff", "--name-only", *base)
	c.Dir = dir
	b, err := c.Output()
	if err != nil {
		return nil, err
	}
	b = bytes.TrimSpace(b)
	logg.Printf("Files changed:\n%s", b)
	return strings.Split(string(b), "\n"), nil
}
