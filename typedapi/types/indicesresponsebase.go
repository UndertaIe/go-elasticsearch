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

// IndicesResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/Base.ts#L77-L79
type IndicesResponseBase struct {
	// Acknowledged For a successful response, this value is always true. On failure, an
	// exception is returned instead.
	Acknowledged bool             `json:"acknowledged"`
	Shards_      *ShardStatistics `json:"_shards,omitempty"`
}

// IndicesResponseBaseBuilder holds IndicesResponseBase struct and provides a builder API.
type IndicesResponseBaseBuilder struct {
	v *IndicesResponseBase
}

// NewIndicesResponseBase provides a builder for the IndicesResponseBase struct.
func NewIndicesResponseBaseBuilder() *IndicesResponseBaseBuilder {
	r := IndicesResponseBaseBuilder{
		&IndicesResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the IndicesResponseBase struct
func (rb *IndicesResponseBaseBuilder) Build() IndicesResponseBase {
	return *rb.v
}

// Acknowledged For a successful response, this value is always true. On failure, an
// exception is returned instead.

func (rb *IndicesResponseBaseBuilder) Acknowledged(acknowledged bool) *IndicesResponseBaseBuilder {
	rb.v.Acknowledged = acknowledged
	return rb
}

func (rb *IndicesResponseBaseBuilder) Shards_(shards_ *ShardStatisticsBuilder) *IndicesResponseBaseBuilder {
	v := shards_.Build()
	rb.v.Shards_ = &v
	return rb
}
