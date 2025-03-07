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

// TopMetricsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L700-L703
type TopMetricsAggregate struct {
	Meta *Metadata    `json:"meta,omitempty"`
	Top  []TopMetrics `json:"top"`
}

// TopMetricsAggregateBuilder holds TopMetricsAggregate struct and provides a builder API.
type TopMetricsAggregateBuilder struct {
	v *TopMetricsAggregate
}

// NewTopMetricsAggregate provides a builder for the TopMetricsAggregate struct.
func NewTopMetricsAggregateBuilder() *TopMetricsAggregateBuilder {
	r := TopMetricsAggregateBuilder{
		&TopMetricsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the TopMetricsAggregate struct
func (rb *TopMetricsAggregateBuilder) Build() TopMetricsAggregate {
	return *rb.v
}

func (rb *TopMetricsAggregateBuilder) Meta(meta *MetadataBuilder) *TopMetricsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *TopMetricsAggregateBuilder) Top(top []TopMetricsBuilder) *TopMetricsAggregateBuilder {
	tmp := make([]TopMetrics, len(top))
	for _, value := range top {
		tmp = append(tmp, value.Build())
	}
	rb.v.Top = tmp
	return rb
}
