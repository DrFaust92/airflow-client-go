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

// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// +build appengine

package internal

import (
	"appengine"

	netcontext "golang.org/x/net/context"
)

func init() {
	appengineStandard = true
}

func DefaultVersionHostname(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.DefaultVersionHostname(c)
}

func Datacenter(_ netcontext.Context) string { return appengine.Datacenter() }
func ServerSoftware() string                 { return appengine.ServerSoftware() }
func InstanceID() string                     { return appengine.InstanceID() }
func IsDevAppServer() bool                   { return appengine.IsDevAppServer() }

func RequestID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.RequestID(c)
}

func ModuleName(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.ModuleName(c)
}
func VersionID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.VersionID(c)
}

func fullyQualifiedAppID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return c.FullyQualifiedAppID()
}
