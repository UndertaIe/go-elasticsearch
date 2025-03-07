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

// SpanFieldMaskingQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/query_dsl/span.ts#L30-L33
type SpanFieldMaskingQuery struct {
	Boost      *float32   `json:"boost,omitempty"`
	Field      Field      `json:"field"`
	Query      *SpanQuery `json:"query,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

// SpanFieldMaskingQueryBuilder holds SpanFieldMaskingQuery struct and provides a builder API.
type SpanFieldMaskingQueryBuilder struct {
	v *SpanFieldMaskingQuery
}

// NewSpanFieldMaskingQuery provides a builder for the SpanFieldMaskingQuery struct.
func NewSpanFieldMaskingQueryBuilder() *SpanFieldMaskingQueryBuilder {
	r := SpanFieldMaskingQueryBuilder{
		&SpanFieldMaskingQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanFieldMaskingQuery struct
func (rb *SpanFieldMaskingQueryBuilder) Build() SpanFieldMaskingQuery {
	return *rb.v
}

func (rb *SpanFieldMaskingQueryBuilder) Boost(boost float32) *SpanFieldMaskingQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SpanFieldMaskingQueryBuilder) Field(field Field) *SpanFieldMaskingQueryBuilder {
	rb.v.Field = field
	return rb
}

func (rb *SpanFieldMaskingQueryBuilder) Query(query *SpanQueryBuilder) *SpanFieldMaskingQueryBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *SpanFieldMaskingQueryBuilder) QueryName_(queryname_ string) *SpanFieldMaskingQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
