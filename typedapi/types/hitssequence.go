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

// HitsSequence type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/eql/_types/EqlHits.ts#L51-L59
type HitsSequence struct {
	// Events Contains events matching the query. Each object represents a matching event.
	Events []HitsEvent `json:"events"`
	// JoinKeys Shared field values used to constrain matches in the sequence. These are
	// defined using the by keyword in the EQL query syntax.
	JoinKeys []interface{} `json:"join_keys"`
}

// HitsSequenceBuilder holds HitsSequence struct and provides a builder API.
type HitsSequenceBuilder struct {
	v *HitsSequence
}

// NewHitsSequence provides a builder for the HitsSequence struct.
func NewHitsSequenceBuilder() *HitsSequenceBuilder {
	r := HitsSequenceBuilder{
		&HitsSequence{},
	}

	return &r
}

// Build finalize the chain and returns the HitsSequence struct
func (rb *HitsSequenceBuilder) Build() HitsSequence {
	return *rb.v
}

// Events Contains events matching the query. Each object represents a matching event.

func (rb *HitsSequenceBuilder) Events(events []HitsEventBuilder) *HitsSequenceBuilder {
	tmp := make([]HitsEvent, len(events))
	for _, value := range events {
		tmp = append(tmp, value.Build())
	}
	rb.v.Events = tmp
	return rb
}

// JoinKeys Shared field values used to constrain matches in the sequence. These are
// defined using the by keyword in the EQL query syntax.

func (rb *HitsSequenceBuilder) JoinKeys(join_keys ...interface{}) *HitsSequenceBuilder {
	rb.v.JoinKeys = join_keys
	return rb
}
