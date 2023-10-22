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

// This file has code for accessing metadata.
//
// References:
//	https://cloud.google.com/compute/docs/metadata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	metadataHost = "metadata"
	metadataPath = "/computeMetadata/v1/"
)

var (
	metadataRequestHeaders = http.Header{
		"Metadata-Flavor": []string{"Google"},
	}
)

// TODO(dsymonds): Do we need to support default values, like Python?
func mustGetMetadata(key string) []byte {
	b, err := getMetadata(key)
	if err != nil {
		panic(fmt.Sprintf("Metadata fetch failed for '%s': %v", key, err))
	}
	return b
}

func getMetadata(key string) ([]byte, error) {
	// TODO(dsymonds): May need to use url.Parse to support keys with query args.
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   metadataHost,
			Path:   metadataPath + key,
		},
		Header: metadataRequestHeaders,
		Host:   metadataHost,
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("metadata server returned HTTP %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
