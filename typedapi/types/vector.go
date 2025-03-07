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

// Vector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L436-L440
type Vector struct {
	Available               bool `json:"available"`
	DenseVectorDimsAvgCount int  `json:"dense_vector_dims_avg_count"`
	DenseVectorFieldsCount  int  `json:"dense_vector_fields_count"`
	Enabled                 bool `json:"enabled"`
	SparseVectorFieldsCount *int `json:"sparse_vector_fields_count,omitempty"`
}

// VectorBuilder holds Vector struct and provides a builder API.
type VectorBuilder struct {
	v *Vector
}

// NewVector provides a builder for the Vector struct.
func NewVectorBuilder() *VectorBuilder {
	r := VectorBuilder{
		&Vector{},
	}

	return &r
}

// Build finalize the chain and returns the Vector struct
func (rb *VectorBuilder) Build() Vector {
	return *rb.v
}

func (rb *VectorBuilder) Available(available bool) *VectorBuilder {
	rb.v.Available = available
	return rb
}

func (rb *VectorBuilder) DenseVectorDimsAvgCount(densevectordimsavgcount int) *VectorBuilder {
	rb.v.DenseVectorDimsAvgCount = densevectordimsavgcount
	return rb
}

func (rb *VectorBuilder) DenseVectorFieldsCount(densevectorfieldscount int) *VectorBuilder {
	rb.v.DenseVectorFieldsCount = densevectorfieldscount
	return rb
}

func (rb *VectorBuilder) Enabled(enabled bool) *VectorBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *VectorBuilder) SparseVectorFieldsCount(sparsevectorfieldscount int) *VectorBuilder {
	rb.v.SparseVectorFieldsCount = &sparsevectorfieldscount
	return rb
}
