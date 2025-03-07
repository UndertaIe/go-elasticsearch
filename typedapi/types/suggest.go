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

// Suggest holds the union for the following types:
//
//	CompletionSuggest
//	PhraseSuggest
//	TermSuggest
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/search/_types/suggester.ts#L34-L40
type Suggest interface{}

// SuggestBuilder holds Suggest struct and provides a builder API.
type SuggestBuilder struct {
	v Suggest
}

// NewSuggest provides a builder for the Suggest struct.
func NewSuggestBuilder() *SuggestBuilder {
	return &SuggestBuilder{}
}

// Build finalize the chain and returns the Suggest struct
func (u *SuggestBuilder) Build() Suggest {
	return u.v
}

func (u *SuggestBuilder) CompletionSuggest(completionsuggest *CompletionSuggestBuilder) *SuggestBuilder {
	v := completionsuggest.Build()
	u.v = &v
	return u
}

func (u *SuggestBuilder) PhraseSuggest(phrasesuggest *PhraseSuggestBuilder) *SuggestBuilder {
	v := phrasesuggest.Build()
	u.v = &v
	return u
}

func (u *SuggestBuilder) TermSuggest(termsuggest *TermSuggestBuilder) *SuggestBuilder {
	v := termsuggest.Build()
	u.v = &v
	return u
}
