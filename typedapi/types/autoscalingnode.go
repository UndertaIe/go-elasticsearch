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

// AutoscalingNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/autoscaling/get_autoscaling_capacity/GetAutoscalingCapacityResponse.ts#L48-L50
type AutoscalingNode struct {
	Name NodeName `json:"name"`
}

// AutoscalingNodeBuilder holds AutoscalingNode struct and provides a builder API.
type AutoscalingNodeBuilder struct {
	v *AutoscalingNode
}

// NewAutoscalingNode provides a builder for the AutoscalingNode struct.
func NewAutoscalingNodeBuilder() *AutoscalingNodeBuilder {
	r := AutoscalingNodeBuilder{
		&AutoscalingNode{},
	}

	return &r
}

// Build finalize the chain and returns the AutoscalingNode struct
func (rb *AutoscalingNodeBuilder) Build() AutoscalingNode {
	return *rb.v
}

func (rb *AutoscalingNodeBuilder) Name(name NodeName) *AutoscalingNodeBuilder {
	rb.v.Name = name
	return rb
}
