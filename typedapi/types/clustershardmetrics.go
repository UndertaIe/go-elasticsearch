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

// ClusterShardMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/cluster/stats/types.ts#L270-L274
type ClusterShardMetrics struct {
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

// ClusterShardMetricsBuilder holds ClusterShardMetrics struct and provides a builder API.
type ClusterShardMetricsBuilder struct {
	v *ClusterShardMetrics
}

// NewClusterShardMetrics provides a builder for the ClusterShardMetrics struct.
func NewClusterShardMetricsBuilder() *ClusterShardMetricsBuilder {
	r := ClusterShardMetricsBuilder{
		&ClusterShardMetrics{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterShardMetrics struct
func (rb *ClusterShardMetricsBuilder) Build() ClusterShardMetrics {
	return *rb.v
}

func (rb *ClusterShardMetricsBuilder) Avg(avg float64) *ClusterShardMetricsBuilder {
	rb.v.Avg = avg
	return rb
}

func (rb *ClusterShardMetricsBuilder) Max(max float64) *ClusterShardMetricsBuilder {
	rb.v.Max = max
	return rb
}

func (rb *ClusterShardMetricsBuilder) Min(min float64) *ClusterShardMetricsBuilder {
	rb.v.Min = min
	return rb
}
