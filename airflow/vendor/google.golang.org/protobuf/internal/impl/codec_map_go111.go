// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.12

package impl

import "reflect"

type mapIter struct {
	v    reflect.Value
	keys []reflect.Value
}

// mapRange provides a less-efficient equivalent to
// the Go 1.12 reflect.Value.MapRange method.
func mapRange(v reflect.Value) *mapIter {
	return &mapIter{v: v}
}

func (i *mapIter) Next() bool {
	if i.keys == nil {
		i.keys = i.v.MapKeys()
	} else {
		i.keys = i.keys[1:]
	}
	return len(i.keys) > 0
}

func (i *mapIter) Key() reflect.Value {
	return i.keys[0]
}

func (i *mapIter) Value() reflect.Value {
	return i.v.MapIndex(i.keys[0])
}
