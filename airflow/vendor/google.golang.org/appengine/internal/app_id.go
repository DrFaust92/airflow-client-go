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
	"strings"
)

func parseFullAppID(appid string) (partition, domain, displayID string) {
	if i := strings.Index(appid, "~"); i != -1 {
		partition, appid = appid[:i], appid[i+1:]
	}
	if i := strings.Index(appid, ":"); i != -1 {
		domain, appid = appid[:i], appid[i+1:]
	}
	return partition, domain, appid
}

// appID returns "appid" or "domain.com:appid".
func appID(fullAppID string) string {
	_, dom, dis := parseFullAppID(fullAppID)
	if dom != "" {
		return dom + ":" + dis
	}
	return dis
}
