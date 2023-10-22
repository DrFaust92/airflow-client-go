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

// +build !appengine

package internal

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

func Main() {
	MainPath = filepath.Dir(findMainPath())
	installHealthChecker(http.DefaultServeMux)

	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	host := ""
	if IsDevAppServer() {
		host = "127.0.0.1"
	}
	if err := http.ListenAndServe(host+":"+port, http.HandlerFunc(handleHTTP)); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}

// Find the path to package main by looking at the root Caller.
func findMainPath() string {
	pc := make([]uintptr, 100)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		// Tests won't have package main, instead they have testing.tRunner
		if frame.Function == "main.main" || frame.Function == "testing.tRunner" {
			return frame.File
		}
		if !more {
			break
		}
	}
	return ""
}

func installHealthChecker(mux *http.ServeMux) {
	// If no health check handler has been installed by this point, add a trivial one.
	const healthPath = "/_ah/health"
	hreq := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: healthPath,
		},
	}
	if _, pat := mux.Handler(hreq); pat != healthPath {
		mux.HandleFunc(healthPath, func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "ok")
		})
	}
}
