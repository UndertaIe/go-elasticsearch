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

// IcuTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/icu-plugin.ts#L30-L33
type IcuTokenizer struct {
	RuleFiles string         `json:"rule_files"`
	Type      string         `json:"type,omitempty"`
	Version   *VersionString `json:"version,omitempty"`
}

// IcuTokenizerBuilder holds IcuTokenizer struct and provides a builder API.
type IcuTokenizerBuilder struct {
	v *IcuTokenizer
}

// NewIcuTokenizer provides a builder for the IcuTokenizer struct.
func NewIcuTokenizerBuilder() *IcuTokenizerBuilder {
	r := IcuTokenizerBuilder{
		&IcuTokenizer{},
	}

	r.v.Type = "icu_tokenizer"

	return &r
}

// Build finalize the chain and returns the IcuTokenizer struct
func (rb *IcuTokenizerBuilder) Build() IcuTokenizer {
	return *rb.v
}

func (rb *IcuTokenizerBuilder) RuleFiles(rulefiles string) *IcuTokenizerBuilder {
	rb.v.RuleFiles = rulefiles
	return rb
}

func (rb *IcuTokenizerBuilder) Version(version VersionString) *IcuTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
