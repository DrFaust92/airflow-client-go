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

// Package text implements the text format for protocol buffers.
// This package has no semantic understanding for protocol buffers and is only
// a parser and composer for the format.
//
// There is no formal specification for the protobuf text format, as such the
// C++ implementation (see google::protobuf::TextFormat) is the reference
// implementation of the text format.
//
// This package is neither a superset nor a subset of the C++ implementation.
// This implementation permits a more liberal grammar in some cases to be
// backwards compatible with the historical Go implementation.
// Future parsings unique to Go should not be added.
// Some grammars allowed by the C++ implementation are deliberately
// not implemented here because they are considered a bug by the protobuf team
// and should not be replicated.
//
// The Go implementation should implement a sufficient amount of the C++
// grammar such that the default text serialization by C++ can be parsed by Go.
// However, just because the C++ parser accepts some input does not mean that
// the Go implementation should as well.
//
// The text format is almost a superset of JSON except:
//	* message keys are not quoted strings, but identifiers
//	* the top-level value must be a message without the delimiters
package text
