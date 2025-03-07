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

// ExistsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/query_dsl/term.ts#L36-L38
type ExistsQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	Field      Field    `json:"field"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// ExistsQueryBuilder holds ExistsQuery struct and provides a builder API.
type ExistsQueryBuilder struct {
	v *ExistsQuery
}

// NewExistsQuery provides a builder for the ExistsQuery struct.
func NewExistsQueryBuilder() *ExistsQueryBuilder {
	r := ExistsQueryBuilder{
		&ExistsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ExistsQuery struct
func (rb *ExistsQueryBuilder) Build() ExistsQuery {
	return *rb.v
}

func (rb *ExistsQueryBuilder) Boost(boost float32) *ExistsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *ExistsQueryBuilder) Field(field Field) *ExistsQueryBuilder {
	rb.v.Field = field
	return rb
}

func (rb *ExistsQueryBuilder) QueryName_(queryname_ string) *ExistsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
