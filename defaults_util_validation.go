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

package defaults

import (
	"context"
	"fmt"
)

//nolint:unused
func fValidation(name string) func(func()) error {
	return func(f func()) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fR1Validation[R1 any](name string) func(func() R1) error {
	return func(f func() R1) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fA1Validation[A1 any](name string) func(func(A1)) error {
	return func(f func(A1)) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fA1R1Validation[A1, R1 any](name string) func(func(A1) R1) error {
	return func(f func(A1) R1) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fA1R2Validation[A1, R1, R2 any](name string) func(func(A1) (R1, R2)) error {
	return func(f func(A1) (R1, R2)) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fA2Validation[A1, A2 any](name string) func(func(A1, A2)) error {
	return func(f func(A1, A2)) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

//nolint:unused
func fA2R1Validation[A1, A2, R1 any](name string) func(func(A1, A2) R1) error {
	return func(f func(A1, A2) R1) error {
		if f == nil {
			return fmt.Errorf("%s function must not be nil", name)
		}
		return nil
	}
}

// ------------------------------------------------------------------------ //

//nolint:unused
func fActxAifaceR1[R1 any](name string) func(func(context.Context, any) R1) error {
	return fA2R1Validation[context.Context, any, R1](name)
}

//nolint:unused
func fActxAiface(name string) func(func(context.Context, any)) error {
	return fA2Validation[context.Context, any](name)
}
