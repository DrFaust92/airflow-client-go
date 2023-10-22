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

// Copyright 2014 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package internal

// This file implements a network dialer that limits the number of concurrent connections.
// It is only used for API calls.

import (
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

var limitSem = make(chan int, 100) // TODO(dsymonds): Use environment variable.

func limitRelease() {
	// non-blocking
	select {
	case <-limitSem:
	default:
		// This should not normally happen.
		log.Print("appengine: unbalanced limitSem release!")
	}
}

func limitDial(network, addr string) (net.Conn, error) {
	limitSem <- 1

	// Dial with a timeout in case the API host is MIA.
	// The connection should normally be very fast.
	conn, err := net.DialTimeout(network, addr, 10*time.Second)
	if err != nil {
		limitRelease()
		return nil, err
	}
	lc := &limitConn{Conn: conn}
	runtime.SetFinalizer(lc, (*limitConn).Close) // shouldn't usually be required
	return lc, nil
}

type limitConn struct {
	close sync.Once
	net.Conn
}

func (lc *limitConn) Close() error {
	defer lc.close.Do(func() {
		limitRelease()
		runtime.SetFinalizer(lc, nil)
	})
	return lc.Conn.Close()
}
