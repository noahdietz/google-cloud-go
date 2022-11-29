// Copyright 2022 Google LLC
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
	"flag"
	"fmt"

	"cloud.google.com/go/internal/gapicgen/regapic"
)

var printReady = flag.Bool("print_ready", false, "")
var printDone = flag.Bool("print_done", false, "")

func main() {
	flag.Parse()
	ready := 0
	done := 0
	for _, s := range regapic.Summarize() {
		if (s.RegapicOnly || s.GRPCOnly) || s.Regapic {
			if *printDone {
				fmt.Printf("Done: %s\n", s.Import)
			}
			done++
			continue
		}
		if *printReady {
			fmt.Printf("Ready: %s\n", s.Import)
		}
		ready++
	}
	fmt.Println("Ready:", ready, "Done:", done)
}
