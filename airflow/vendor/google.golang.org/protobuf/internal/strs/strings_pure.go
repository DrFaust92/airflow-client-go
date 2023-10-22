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

package strs

import pref "google.golang.org/protobuf/reflect/protoreflect"

func UnsafeString(b []byte) string {
	return string(b)
}

func UnsafeBytes(s string) []byte {
	return []byte(s)
}

type Builder struct{}

func (*Builder) AppendFullName(prefix pref.FullName, name pref.Name) pref.FullName {
	return prefix.Append(name)
}

func (*Builder) MakeString(b []byte) string {
	return string(b)
}
