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

package proto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Size returns the size in bytes of the wire-format encoding of m.
func Size(m Message) int {
	return MarshalOptions{}.Size(m)
}

// Size returns the size in bytes of the wire-format encoding of m.
func (o MarshalOptions) Size(m Message) int {
	// Treat a nil message interface as an empty message; nothing to output.
	if m == nil {
		return 0
	}

	return o.size(m.ProtoReflect())
}

// size is a centralized function that all size operations go through.
// For profiling purposes, avoid changing the name of this function or
// introducing other code paths for size that do not go through this.
func (o MarshalOptions) size(m protoreflect.Message) (size int) {
	methods := protoMethods(m)
	if methods != nil && methods.Size != nil {
		out := methods.Size(protoiface.SizeInput{
			Message: m,
		})
		return out.Size
	}
	if methods != nil && methods.Marshal != nil {
		// This is not efficient, but we don't have any choice.
		// This case is mainly used for legacy types with a Marshal method.
		out, _ := methods.Marshal(protoiface.MarshalInput{
			Message: m,
		})
		return len(out.Buf)
	}
	return o.sizeMessageSlow(m)
}

func (o MarshalOptions) sizeMessageSlow(m protoreflect.Message) (size int) {
	if messageset.IsMessageSet(m.Descriptor()) {
		return o.sizeMessageSet(m)
	}
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += o.sizeField(fd, v)
		return true
	})
	size += len(m.GetUnknown())
	return size
}

func (o MarshalOptions) sizeField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (size int) {
	num := fd.Number()
	switch {
	case fd.IsList():
		return o.sizeList(num, fd, value.List())
	case fd.IsMap():
		return o.sizeMap(num, fd, value.Map())
	default:
		return protowire.SizeTag(num) + o.sizeSingular(num, fd.Kind(), value)
	}
}

func (o MarshalOptions) sizeList(num protowire.Number, fd protoreflect.FieldDescriptor, list protoreflect.List) (size int) {
	if fd.IsPacked() && list.Len() > 0 {
		content := 0
		for i, llen := 0, list.Len(); i < llen; i++ {
			content += o.sizeSingular(num, fd.Kind(), list.Get(i))
		}
		return protowire.SizeTag(num) + protowire.SizeBytes(content)
	}

	for i, llen := 0, list.Len(); i < llen; i++ {
		size += protowire.SizeTag(num) + o.sizeSingular(num, fd.Kind(), list.Get(i))
	}
	return size
}

func (o MarshalOptions) sizeMap(num protowire.Number, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) (size int) {
	mapv.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		size += protowire.SizeTag(num)
		size += protowire.SizeBytes(o.sizeField(fd.MapKey(), key.Value()) + o.sizeField(fd.MapValue(), value))
		return true
	})
	return size
}
