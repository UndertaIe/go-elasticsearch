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

// IpRangeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/bucket.ts#L249-L252
type IpRangeAggregation struct {
	Field  *Field                    `json:"field,omitempty"`
	Meta   *Metadata                 `json:"meta,omitempty"`
	Name   *string                   `json:"name,omitempty"`
	Ranges []IpRangeAggregationRange `json:"ranges,omitempty"`
}

// IpRangeAggregationBuilder holds IpRangeAggregation struct and provides a builder API.
type IpRangeAggregationBuilder struct {
	v *IpRangeAggregation
}

// NewIpRangeAggregation provides a builder for the IpRangeAggregation struct.
func NewIpRangeAggregationBuilder() *IpRangeAggregationBuilder {
	r := IpRangeAggregationBuilder{
		&IpRangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the IpRangeAggregation struct
func (rb *IpRangeAggregationBuilder) Build() IpRangeAggregation {
	return *rb.v
}

func (rb *IpRangeAggregationBuilder) Field(field Field) *IpRangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *IpRangeAggregationBuilder) Meta(meta *MetadataBuilder) *IpRangeAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *IpRangeAggregationBuilder) Name(name string) *IpRangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *IpRangeAggregationBuilder) Ranges(ranges []IpRangeAggregationRangeBuilder) *IpRangeAggregationBuilder {
	tmp := make([]IpRangeAggregationRange, len(ranges))
	for _, value := range ranges {
		tmp = append(tmp, value.Build())
	}
	rb.v.Ranges = tmp
	return rb
}
