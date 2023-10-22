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

package impl

import (
	"fmt"

	pref "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// weakFields adds methods to the exported WeakFields type for internal use.
//
// The exported type is an alias to an unnamed type, so methods can't be
// defined directly on it.
type weakFields WeakFields

func (w weakFields) get(num pref.FieldNumber) (pref.ProtoMessage, bool) {
	m, ok := w[int32(num)]
	return m, ok
}

func (w *weakFields) set(num pref.FieldNumber, m pref.ProtoMessage) {
	if *w == nil {
		*w = make(weakFields)
	}
	(*w)[int32(num)] = m
}

func (w *weakFields) clear(num pref.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) HasWeak(w WeakFields, num pref.FieldNumber) bool {
	_, ok := w[int32(num)]
	return ok
}

func (Export) ClearWeak(w *WeakFields, num pref.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) GetWeak(w WeakFields, num pref.FieldNumber, name pref.FullName) pref.ProtoMessage {
	if m, ok := w[int32(num)]; ok {
		return m
	}
	mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
	if mt == nil {
		panic(fmt.Sprintf("message %v for weak field is not linked in", name))
	}
	return mt.Zero().Interface()
}

func (Export) SetWeak(w *WeakFields, num pref.FieldNumber, name pref.FullName, m pref.ProtoMessage) {
	if m != nil {
		mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
		if mt == nil {
			panic(fmt.Sprintf("message %v for weak field is not linked in", name))
		}
		if mt != m.ProtoReflect().Type() {
			panic(fmt.Sprintf("invalid message type for weak field: got %T, want %T", m, mt.Zero().Interface()))
		}
	}
	if m == nil || !m.ProtoReflect().IsValid() {
		delete(*w, int32(num))
		return
	}
	if *w == nil {
		*w = make(weakFields)
	}
	(*w)[int32(num)] = m
}
