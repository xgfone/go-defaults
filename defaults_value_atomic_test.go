// Copyright 2023 xgfone
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

//go:build atomic

package defaults

import "testing"

func TestValueIsAtomic(t *testing.T) {
	v := NewValue(123)
	if _v := v.value.Load().(valuer[int]).Value; _v != 123 { // If not atomic, fail to compile.
		t.Errorf("expect %d, but got %d", 123, _v)
	}
}
