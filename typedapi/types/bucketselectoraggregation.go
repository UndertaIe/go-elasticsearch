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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// BucketSelectorAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/pipeline.ts#L75-L77
type BucketSelectorAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Script      *Script              `json:"script,omitempty"`
}

// BucketSelectorAggregationBuilder holds BucketSelectorAggregation struct and provides a builder API.
type BucketSelectorAggregationBuilder struct {
	v *BucketSelectorAggregation
}

// NewBucketSelectorAggregation provides a builder for the BucketSelectorAggregation struct.
func NewBucketSelectorAggregationBuilder() *BucketSelectorAggregationBuilder {
	r := BucketSelectorAggregationBuilder{
		&BucketSelectorAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BucketSelectorAggregation struct
func (rb *BucketSelectorAggregationBuilder) Build() BucketSelectorAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *BucketSelectorAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *BucketSelectorAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *BucketSelectorAggregationBuilder) Format(format string) *BucketSelectorAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *BucketSelectorAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *BucketSelectorAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *BucketSelectorAggregationBuilder) Meta(meta *MetadataBuilder) *BucketSelectorAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *BucketSelectorAggregationBuilder) Name(name string) *BucketSelectorAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *BucketSelectorAggregationBuilder) Script(script *ScriptBuilder) *BucketSelectorAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
