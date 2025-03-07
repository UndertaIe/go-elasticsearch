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

// PassThroughInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/inference.ts#L209-L215
type PassThroughInferenceOptions struct {
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

// PassThroughInferenceOptionsBuilder holds PassThroughInferenceOptions struct and provides a builder API.
type PassThroughInferenceOptionsBuilder struct {
	v *PassThroughInferenceOptions
}

// NewPassThroughInferenceOptions provides a builder for the PassThroughInferenceOptions struct.
func NewPassThroughInferenceOptionsBuilder() *PassThroughInferenceOptionsBuilder {
	r := PassThroughInferenceOptionsBuilder{
		&PassThroughInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the PassThroughInferenceOptions struct
func (rb *PassThroughInferenceOptionsBuilder) Build() PassThroughInferenceOptions {
	return *rb.v
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *PassThroughInferenceOptionsBuilder) ResultsField(resultsfield string) *PassThroughInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options

func (rb *PassThroughInferenceOptionsBuilder) Tokenization(tokenization *TokenizationConfigContainerBuilder) *PassThroughInferenceOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
