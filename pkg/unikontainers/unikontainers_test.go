// Copyright (c) 2023-2026, Nubificus LTD
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

package unikontainers

import (
	"testing"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func TestGetNetworkType(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    string
	}{
		{
			name: "user-container",
			annotations: map[string]string{
				"io.kubernetes.cri.container-name": "user-container",
			},
			expected: "static",
		},
		{
			name: "other-container",
			annotations: map[string]string{
				"io.kubernetes.cri.container-name": "other-container",
			},
			expected: "dynamic",
		},
		{
			name:        "missing annotation",
			annotations: map[string]string{},
			expected:    "dynamic",
		},
		{
			name:        "nil annotations",
			annotations: nil,
			expected:    "dynamic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unikontainer{
				Spec: &specs.Spec{
					Annotations: tt.annotations,
				},
			}
			got := u.getNetworkType()
			if got != tt.expected {
				t.Errorf("getNetworkType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
