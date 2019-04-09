// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"strings"
)

// Errors is an error that contains multiple errors
type Errors struct {
	Errs []error
}

func (e *Errors) IsErr() bool {
	return len(e.Errs) > 0
}

func (e *Errors) Append(err error) {
	e.Errs = append(e.Errs, err)
}

func (e *Errors) Error() string {
	errs := []string{}
	for _, err := range e.Errs {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, ", ")
}

func (e *Errors) Equal(e2 error) bool {
	errs1 := []string{}
	errs2 := []string{}
	for _, err := range e.Errs {
		errs1 = append(errs1, err.Error())
	}
	if es2, ok := e2.(*Errors); ok {
		for _, err := range es2.Errs {
			errs2 = append(errs2, err.Error())
		}
	} else {
		errs2 = append(errs2, e2.Error())
	}

	return CompareStringSliceNoOrder(errs1, errs2)
}

// ErrBadRequest represent an error caused by a bad command request
// it's used to differentiate an internal error from an user error
type ErrBadRequest struct {
	Err error
}

func (e *ErrBadRequest) Error() string {
	return e.Err.Error()
}

func NewErrBadRequest(err error) *ErrBadRequest {
	return &ErrBadRequest{Err: err}
}

func IsErrBadRequest(err error) bool {
	_, ok := err.(*ErrBadRequest)
	return ok
}

// ErrNotFound represent a not found error
// it's used to differentiate an internal error from an user error
type ErrNotFound struct {
	Err error
}

func (e *ErrNotFound) Error() string {
	return e.Err.Error()
}

func NewErrNotFound(err error) *ErrNotFound {
	return &ErrNotFound{Err: err}
}

func IsErrNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	return ok
}
