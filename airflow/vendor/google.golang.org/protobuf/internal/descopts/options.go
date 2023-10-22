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

// Package descopts contains the nil pointers to concrete descriptor options.
//
// This package exists as a form of reverse dependency injection so that certain
// packages (e.g., internal/filedesc and internal/filetype can avoid a direct
// dependency on the descriptor proto package).
package descopts

import pref "google.golang.org/protobuf/reflect/protoreflect"

// These variables are set by the init function in descriptor.pb.go via logic
// in internal/filetype. In other words, so long as the descriptor proto package
// is linked in, these variables will be populated.
//
// Each variable is populated with a nil pointer to the options struct.
var (
	File           pref.ProtoMessage
	Enum           pref.ProtoMessage
	EnumValue      pref.ProtoMessage
	Message        pref.ProtoMessage
	Field          pref.ProtoMessage
	Oneof          pref.ProtoMessage
	ExtensionRange pref.ProtoMessage
	Service        pref.ProtoMessage
	Method         pref.ProtoMessage
)
