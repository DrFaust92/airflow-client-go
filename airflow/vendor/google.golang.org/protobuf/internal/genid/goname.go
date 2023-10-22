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

package genid

// Go names of implementation-specific struct fields in generated messages.
const (
	State_goname = "state"

	SizeCache_goname  = "sizeCache"
	SizeCacheA_goname = "XXX_sizecache"

	WeakFields_goname  = "weakFields"
	WeakFieldsA_goname = "XXX_weak"

	UnknownFields_goname  = "unknownFields"
	UnknownFieldsA_goname = "XXX_unrecognized"

	ExtensionFields_goname  = "extensionFields"
	ExtensionFieldsA_goname = "XXX_InternalExtensions"
	ExtensionFieldsB_goname = "XXX_extensions"

	WeakFieldPrefix_goname = "XXX_weak_"
)
