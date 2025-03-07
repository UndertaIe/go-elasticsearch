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

// MappingLimitSettingsNestedFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L436-L444
type MappingLimitSettingsNestedFields struct {
	// Limit The maximum number of distinct nested mappings in an index. The nested type
	// should only be used in special cases, when
	// arrays of objects need to be queried independently of each other. To
	// safeguard against poorly designed mappings, this
	// setting limits the number of unique nested types per index.
	Limit *int `json:"limit,omitempty"`
}

// MappingLimitSettingsNestedFieldsBuilder holds MappingLimitSettingsNestedFields struct and provides a builder API.
type MappingLimitSettingsNestedFieldsBuilder struct {
	v *MappingLimitSettingsNestedFields
}

// NewMappingLimitSettingsNestedFields provides a builder for the MappingLimitSettingsNestedFields struct.
func NewMappingLimitSettingsNestedFieldsBuilder() *MappingLimitSettingsNestedFieldsBuilder {
	r := MappingLimitSettingsNestedFieldsBuilder{
		&MappingLimitSettingsNestedFields{},
	}

	return &r
}

// Build finalize the chain and returns the MappingLimitSettingsNestedFields struct
func (rb *MappingLimitSettingsNestedFieldsBuilder) Build() MappingLimitSettingsNestedFields {
	return *rb.v
}

// Limit The maximum number of distinct nested mappings in an index. The nested type
// should only be used in special cases, when
// arrays of objects need to be queried independently of each other. To
// safeguard against poorly designed mappings, this
// setting limits the number of unique nested types per index.

func (rb *MappingLimitSettingsNestedFieldsBuilder) Limit(limit int) *MappingLimitSettingsNestedFieldsBuilder {
	rb.v.Limit = &limit
	return rb
}
