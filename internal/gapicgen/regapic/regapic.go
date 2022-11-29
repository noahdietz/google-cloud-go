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

package regapic

import "cloud.google.com/go/internal/gapicgen/generator"

type Status struct {
	Import       string
	Regapic      bool
	NumericEnums bool
	GRPCOnly     bool
	RegapicOnly  bool
}

func Summarize() (summary []Status) {
	for _, m := range generator.MicrogenGapicConfigs {
		s := Status{
			Import:       m.ImportPath,
			NumericEnums: m.NumericEnumsEnabled,
			// Explicitly setting Transports to 'grpc' indicates it should
			// never have REGAPIC for whatever reason.
			GRPCOnly: len(m.Transports) == 1 && m.Transports[0] == "grpc",
			// Explicitly setting Transports to 'rest' indicates it should
			// never have gRPC for whatever reason.
			RegapicOnly: len(m.Transports) == 1 && m.Transports[0] == "rest",
			Regapic:     len(m.Transports) >= 1 && (m.Transports[0] == "rest" || m.Transports[1] == "rest"),
		}
		summary = append(summary, s)
	}

	return summary
}
