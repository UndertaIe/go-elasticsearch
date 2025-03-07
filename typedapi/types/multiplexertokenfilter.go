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

// MultiplexerTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/token_filters.ts#L259-L263
type MultiplexerTokenFilter struct {
	Filters          []string       `json:"filters"`
	PreserveOriginal *bool          `json:"preserve_original,omitempty"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// MultiplexerTokenFilterBuilder holds MultiplexerTokenFilter struct and provides a builder API.
type MultiplexerTokenFilterBuilder struct {
	v *MultiplexerTokenFilter
}

// NewMultiplexerTokenFilter provides a builder for the MultiplexerTokenFilter struct.
func NewMultiplexerTokenFilterBuilder() *MultiplexerTokenFilterBuilder {
	r := MultiplexerTokenFilterBuilder{
		&MultiplexerTokenFilter{},
	}

	r.v.Type = "multiplexer"

	return &r
}

// Build finalize the chain and returns the MultiplexerTokenFilter struct
func (rb *MultiplexerTokenFilterBuilder) Build() MultiplexerTokenFilter {
	return *rb.v
}

func (rb *MultiplexerTokenFilterBuilder) Filters(filters ...string) *MultiplexerTokenFilterBuilder {
	rb.v.Filters = filters
	return rb
}

func (rb *MultiplexerTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *MultiplexerTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

func (rb *MultiplexerTokenFilterBuilder) Version(version VersionString) *MultiplexerTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
