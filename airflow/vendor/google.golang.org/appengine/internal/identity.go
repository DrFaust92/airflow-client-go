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

// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package internal

import (
	"os"

	netcontext "golang.org/x/net/context"
)

var (
	// This is set to true in identity_classic.go, which is behind the appengine build tag.
	// The appengine build tag is set for the first generation runtimes (<= Go 1.9) but not
	// the second generation runtimes (>= Go 1.11), so this indicates whether we're on a
	// first-gen runtime. See IsStandard below for the second-gen check.
	appengineStandard bool

	// This is set to true in identity_flex.go, which is behind the appenginevm build tag.
	appengineFlex bool
)

// AppID is the implementation of the wrapper function of the same name in
// ../identity.go. See that file for commentary.
func AppID(c netcontext.Context) string {
	return appID(FullyQualifiedAppID(c))
}

// IsStandard is the implementation of the wrapper function of the same name in
// ../appengine.go. See that file for commentary.
func IsStandard() bool {
	// appengineStandard will be true for first-gen runtimes (<= Go 1.9) but not
	// second-gen (>= Go 1.11).
	return appengineStandard || IsSecondGen()
}

// IsStandard is the implementation of the wrapper function of the same name in
// ../appengine.go. See that file for commentary.
func IsSecondGen() bool {
	// Second-gen runtimes set $GAE_ENV so we use that to check if we're on a second-gen runtime.
	return os.Getenv("GAE_ENV") == "standard"
}

// IsFlex is the implementation of the wrapper function of the same name in
// ../appengine.go. See that file for commentary.
func IsFlex() bool {
	return appengineFlex
}

// IsAppEngine is the implementation of the wrapper function of the same name in
// ../appengine.go. See that file for commentary.
func IsAppEngine() bool {
	return IsStandard() || IsFlex()
}
