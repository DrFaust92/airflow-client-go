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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protoreflect

import (
	"google.golang.org/protobuf/internal/pragma"
)

// The following types are used by the fast-path Message.ProtoMethods method.
//
// To avoid polluting the public protoreflect API with types used only by
// low-level implementations, the canonical definitions of these types are
// in the runtime/protoiface package. The definitions here and in protoiface
// must be kept in sync.
type (
	methods = struct {
		pragma.NoUnkeyedLiterals
		Flags            supportFlags
		Size             func(sizeInput) sizeOutput
		Marshal          func(marshalInput) (marshalOutput, error)
		Unmarshal        func(unmarshalInput) (unmarshalOutput, error)
		Merge            func(mergeInput) mergeOutput
		CheckInitialized func(checkInitializedInput) (checkInitializedOutput, error)
	}
	supportFlags = uint64
	sizeInput    = struct {
		pragma.NoUnkeyedLiterals
		Message Message
		Flags   uint8
	}
	sizeOutput = struct {
		pragma.NoUnkeyedLiterals
		Size int
	}
	marshalInput = struct {
		pragma.NoUnkeyedLiterals
		Message Message
		Buf     []byte
		Flags   uint8
	}
	marshalOutput = struct {
		pragma.NoUnkeyedLiterals
		Buf []byte
	}
	unmarshalInput = struct {
		pragma.NoUnkeyedLiterals
		Message  Message
		Buf      []byte
		Flags    uint8
		Resolver interface {
			FindExtensionByName(field FullName) (ExtensionType, error)
			FindExtensionByNumber(message FullName, field FieldNumber) (ExtensionType, error)
		}
	}
	unmarshalOutput = struct {
		pragma.NoUnkeyedLiterals
		Flags uint8
	}
	mergeInput = struct {
		pragma.NoUnkeyedLiterals
		Source      Message
		Destination Message
	}
	mergeOutput = struct {
		pragma.NoUnkeyedLiterals
		Flags uint8
	}
	checkInitializedInput = struct {
		pragma.NoUnkeyedLiterals
		Message Message
	}
	checkInitializedOutput = struct {
		pragma.NoUnkeyedLiterals
	}
)
