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

// PercentilesBucketAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L169-L170
type PercentilesBucketAggregate struct {
	Meta   *Metadata   `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

// PercentilesBucketAggregateBuilder holds PercentilesBucketAggregate struct and provides a builder API.
type PercentilesBucketAggregateBuilder struct {
	v *PercentilesBucketAggregate
}

// NewPercentilesBucketAggregate provides a builder for the PercentilesBucketAggregate struct.
func NewPercentilesBucketAggregateBuilder() *PercentilesBucketAggregateBuilder {
	r := PercentilesBucketAggregateBuilder{
		&PercentilesBucketAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the PercentilesBucketAggregate struct
func (rb *PercentilesBucketAggregateBuilder) Build() PercentilesBucketAggregate {
	return *rb.v
}

func (rb *PercentilesBucketAggregateBuilder) Meta(meta *MetadataBuilder) *PercentilesBucketAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *PercentilesBucketAggregateBuilder) Values(values *PercentilesBuilder) *PercentilesBucketAggregateBuilder {
	v := values.Build()
	rb.v.Values = v
	return rb
}
