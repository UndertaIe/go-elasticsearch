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

// LinearInterpolationSmoothingModel type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/search/_types/suggester.ts#L216-L220
type LinearInterpolationSmoothingModel struct {
	BigramLambda  float64 `json:"bigram_lambda"`
	TrigramLambda float64 `json:"trigram_lambda"`
	UnigramLambda float64 `json:"unigram_lambda"`
}

// LinearInterpolationSmoothingModelBuilder holds LinearInterpolationSmoothingModel struct and provides a builder API.
type LinearInterpolationSmoothingModelBuilder struct {
	v *LinearInterpolationSmoothingModel
}

// NewLinearInterpolationSmoothingModel provides a builder for the LinearInterpolationSmoothingModel struct.
func NewLinearInterpolationSmoothingModelBuilder() *LinearInterpolationSmoothingModelBuilder {
	r := LinearInterpolationSmoothingModelBuilder{
		&LinearInterpolationSmoothingModel{},
	}

	return &r
}

// Build finalize the chain and returns the LinearInterpolationSmoothingModel struct
func (rb *LinearInterpolationSmoothingModelBuilder) Build() LinearInterpolationSmoothingModel {
	return *rb.v
}

func (rb *LinearInterpolationSmoothingModelBuilder) BigramLambda(bigramlambda float64) *LinearInterpolationSmoothingModelBuilder {
	rb.v.BigramLambda = bigramlambda
	return rb
}

func (rb *LinearInterpolationSmoothingModelBuilder) TrigramLambda(trigramlambda float64) *LinearInterpolationSmoothingModelBuilder {
	rb.v.TrigramLambda = trigramlambda
	return rb
}

func (rb *LinearInterpolationSmoothingModelBuilder) UnigramLambda(unigramlambda float64) *LinearInterpolationSmoothingModelBuilder {
	rb.v.UnigramLambda = unigramlambda
	return rb
}
