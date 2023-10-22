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

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build purego appengine

package protoreflect

import "google.golang.org/protobuf/internal/pragma"

type valueType int

const (
	nilType valueType = iota
	boolType
	int32Type
	int64Type
	uint32Type
	uint64Type
	float32Type
	float64Type
	stringType
	bytesType
	enumType
	ifaceType
)

// value is a union where only one type can be represented at a time.
// This uses a distinct field for each type. This is type safe in Go, but
// occupies more memory than necessary (72B).
type value struct {
	pragma.DoNotCompare // 0B

	typ   valueType   // 8B
	num   uint64      // 8B
	str   string      // 16B
	bin   []byte      // 24B
	iface interface{} // 16B
}

func valueOfString(v string) Value {
	return Value{typ: stringType, str: v}
}
func valueOfBytes(v []byte) Value {
	return Value{typ: bytesType, bin: v}
}
func valueOfIface(v interface{}) Value {
	return Value{typ: ifaceType, iface: v}
}

func (v Value) getString() string {
	return v.str
}
func (v Value) getBytes() []byte {
	return v.bin
}
func (v Value) getIface() interface{} {
	return v.iface
}
