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

// Vertex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/graph/_types/Vertex.ts#L23-L28
type Vertex struct {
	Depth  int64   `json:"depth"`
	Field  Field   `json:"field"`
	Term   string  `json:"term"`
	Weight float64 `json:"weight"`
}

// VertexBuilder holds Vertex struct and provides a builder API.
type VertexBuilder struct {
	v *Vertex
}

// NewVertex provides a builder for the Vertex struct.
func NewVertexBuilder() *VertexBuilder {
	r := VertexBuilder{
		&Vertex{},
	}

	return &r
}

// Build finalize the chain and returns the Vertex struct
func (rb *VertexBuilder) Build() Vertex {
	return *rb.v
}

func (rb *VertexBuilder) Depth(depth int64) *VertexBuilder {
	rb.v.Depth = depth
	return rb
}

func (rb *VertexBuilder) Field(field Field) *VertexBuilder {
	rb.v.Field = field
	return rb
}

func (rb *VertexBuilder) Term(term string) *VertexBuilder {
	rb.v.Term = term
	return rb
}

func (rb *VertexBuilder) Weight(weight float64) *VertexBuilder {
	rb.v.Weight = weight
	return rb
}
