// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// NodeInfoOSCPU type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/info/types.ts#L331-L340
type NodeInfoOSCPU struct {
	CacheSize        string `json:"cache_size"`
	CacheSizeInBytes int    `json:"cache_size_in_bytes"`
	CoresPerSocket   int    `json:"cores_per_socket"`
	Mhz              int    `json:"mhz"`
	Model            string `json:"model"`
	TotalCores       int    `json:"total_cores"`
	TotalSockets     int    `json:"total_sockets"`
	Vendor           string `json:"vendor"`
}

// NewNodeInfoOSCPU returns a NodeInfoOSCPU.
func NewNodeInfoOSCPU() *NodeInfoOSCPU {
	r := &NodeInfoOSCPU{}

	return r
}
