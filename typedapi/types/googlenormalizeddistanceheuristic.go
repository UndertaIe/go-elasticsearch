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

// GoogleNormalizedDistanceHeuristic type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/bucket.ts#L326-L328
type GoogleNormalizedDistanceHeuristic struct {
	BackgroundIsSuperset *bool `json:"background_is_superset,omitempty"`
}

// GoogleNormalizedDistanceHeuristicBuilder holds GoogleNormalizedDistanceHeuristic struct and provides a builder API.
type GoogleNormalizedDistanceHeuristicBuilder struct {
	v *GoogleNormalizedDistanceHeuristic
}

// NewGoogleNormalizedDistanceHeuristic provides a builder for the GoogleNormalizedDistanceHeuristic struct.
func NewGoogleNormalizedDistanceHeuristicBuilder() *GoogleNormalizedDistanceHeuristicBuilder {
	r := GoogleNormalizedDistanceHeuristicBuilder{
		&GoogleNormalizedDistanceHeuristic{},
	}

	return &r
}

// Build finalize the chain and returns the GoogleNormalizedDistanceHeuristic struct
func (rb *GoogleNormalizedDistanceHeuristicBuilder) Build() GoogleNormalizedDistanceHeuristic {
	return *rb.v
}

func (rb *GoogleNormalizedDistanceHeuristicBuilder) BackgroundIsSuperset(backgroundissuperset bool) *GoogleNormalizedDistanceHeuristicBuilder {
	rb.v.BackgroundIsSuperset = &backgroundissuperset
	return rb
}
