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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


// Package shardstoreallocation
package shardstoreallocation

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/shard_stores/types.ts#L40-L44
type ShardStoreAllocation struct {
	name string
}

var (
	Primary = ShardStoreAllocation{"primary"}

	Replica = ShardStoreAllocation{"replica"}

	Unused = ShardStoreAllocation{"unused"}
)

func (s ShardStoreAllocation) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShardStoreAllocation) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "primary":
		*s = Primary
	case "replica":
		*s = Replica
	case "unused":
		*s = Unused
	default:
		*s = ShardStoreAllocation{string(text)}
	}

	return nil
}

func (s ShardStoreAllocation) String() string {
	return s.name
}
