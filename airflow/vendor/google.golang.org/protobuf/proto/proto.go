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

package proto

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Message is the top-level interface that all messages must implement.
// It provides access to a reflective view of a message.
// Any implementation of this interface may be used with all functions in the
// protobuf module that accept a Message, except where otherwise specified.
//
// This is the v2 interface definition for protobuf messages.
// The v1 interface definition is "github.com/golang/protobuf/proto".Message.
//
// To convert a v1 message to a v2 message,
// use "github.com/golang/protobuf/proto".MessageV2.
// To convert a v2 message to a v1 message,
// use "github.com/golang/protobuf/proto".MessageV1.
type Message = protoreflect.ProtoMessage

// Error matches all errors produced by packages in the protobuf module.
//
// That is, errors.Is(err, Error) reports whether an error is produced
// by this module.
var Error error

func init() {
	Error = errors.Error
}
