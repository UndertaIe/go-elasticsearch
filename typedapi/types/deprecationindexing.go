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


package types

// DeprecationIndexing type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/nodes/info/types.ts#L140-L142
type DeprecationIndexing struct {
	Enabled string `json:"enabled"`
}

// DeprecationIndexingBuilder holds DeprecationIndexing struct and provides a builder API.
type DeprecationIndexingBuilder struct {
	v *DeprecationIndexing
}

// NewDeprecationIndexing provides a builder for the DeprecationIndexing struct.
func NewDeprecationIndexingBuilder() *DeprecationIndexingBuilder {
	r := DeprecationIndexingBuilder{
		&DeprecationIndexing{},
	}

	return &r
}

// Build finalize the chain and returns the DeprecationIndexing struct
func (rb *DeprecationIndexingBuilder) Build() DeprecationIndexing {
	return *rb.v
}

func (rb *DeprecationIndexingBuilder) Enabled(arg string) *DeprecationIndexingBuilder {
	rb.v.Enabled = arg
	return rb
}
