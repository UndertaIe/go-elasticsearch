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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationstrength"
)

// IcuCollationTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/icu-plugin.ts#L51-L65
type IcuCollationTokenFilter struct {
	Alternate              *icucollationalternate.IcuCollationAlternate         `json:"alternate,omitempty"`
	CaseFirst              *icucollationcasefirst.IcuCollationCaseFirst         `json:"caseFirst,omitempty"`
	CaseLevel              *bool                                                `json:"caseLevel,omitempty"`
	Country                *string                                              `json:"country,omitempty"`
	Decomposition          *icucollationdecomposition.IcuCollationDecomposition `json:"decomposition,omitempty"`
	HiraganaQuaternaryMode *bool                                                `json:"hiraganaQuaternaryMode,omitempty"`
	Language               *string                                              `json:"language,omitempty"`
	Numeric                *bool                                                `json:"numeric,omitempty"`
	Rules                  *string                                              `json:"rules,omitempty"`
	Strength               *icucollationstrength.IcuCollationStrength           `json:"strength,omitempty"`
	Type                   string                                               `json:"type,omitempty"`
	VariableTop            *string                                              `json:"variableTop,omitempty"`
	Variant                *string                                              `json:"variant,omitempty"`
	Version                *VersionString                                       `json:"version,omitempty"`
}

// IcuCollationTokenFilterBuilder holds IcuCollationTokenFilter struct and provides a builder API.
type IcuCollationTokenFilterBuilder struct {
	v *IcuCollationTokenFilter
}

// NewIcuCollationTokenFilter provides a builder for the IcuCollationTokenFilter struct.
func NewIcuCollationTokenFilterBuilder() *IcuCollationTokenFilterBuilder {
	r := IcuCollationTokenFilterBuilder{
		&IcuCollationTokenFilter{},
	}

	r.v.Type = "icu_collation"

	return &r
}

// Build finalize the chain and returns the IcuCollationTokenFilter struct
func (rb *IcuCollationTokenFilterBuilder) Build() IcuCollationTokenFilter {
	return *rb.v
}

func (rb *IcuCollationTokenFilterBuilder) Alternate(alternate icucollationalternate.IcuCollationAlternate) *IcuCollationTokenFilterBuilder {
	rb.v.Alternate = &alternate
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) CaseFirst(casefirst icucollationcasefirst.IcuCollationCaseFirst) *IcuCollationTokenFilterBuilder {
	rb.v.CaseFirst = &casefirst
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) CaseLevel(caselevel bool) *IcuCollationTokenFilterBuilder {
	rb.v.CaseLevel = &caselevel
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Country(country string) *IcuCollationTokenFilterBuilder {
	rb.v.Country = &country
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Decomposition(decomposition icucollationdecomposition.IcuCollationDecomposition) *IcuCollationTokenFilterBuilder {
	rb.v.Decomposition = &decomposition
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) HiraganaQuaternaryMode(hiraganaquaternarymode bool) *IcuCollationTokenFilterBuilder {
	rb.v.HiraganaQuaternaryMode = &hiraganaquaternarymode
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Language(language string) *IcuCollationTokenFilterBuilder {
	rb.v.Language = &language
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Numeric(numeric bool) *IcuCollationTokenFilterBuilder {
	rb.v.Numeric = &numeric
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Rules(rules string) *IcuCollationTokenFilterBuilder {
	rb.v.Rules = &rules
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Strength(strength icucollationstrength.IcuCollationStrength) *IcuCollationTokenFilterBuilder {
	rb.v.Strength = &strength
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) VariableTop(variabletop string) *IcuCollationTokenFilterBuilder {
	rb.v.VariableTop = &variabletop
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Variant(variant string) *IcuCollationTokenFilterBuilder {
	rb.v.Variant = &variant
	return rb
}

func (rb *IcuCollationTokenFilterBuilder) Version(version VersionString) *IcuCollationTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
