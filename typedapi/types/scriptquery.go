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

// ScriptQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/query_dsl/specialized.ts#L164-L166
type ScriptQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
	Script     Script   `json:"script"`
}

// ScriptQueryBuilder holds ScriptQuery struct and provides a builder API.
type ScriptQueryBuilder struct {
	v *ScriptQuery
}

// NewScriptQuery provides a builder for the ScriptQuery struct.
func NewScriptQueryBuilder() *ScriptQueryBuilder {
	r := ScriptQueryBuilder{
		&ScriptQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ScriptQuery struct
func (rb *ScriptQueryBuilder) Build() ScriptQuery {
	return *rb.v
}

func (rb *ScriptQueryBuilder) Boost(boost float32) *ScriptQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *ScriptQueryBuilder) QueryName_(queryname_ string) *ScriptQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *ScriptQueryBuilder) Script(script *ScriptBuilder) *ScriptQueryBuilder {
	v := script.Build()
	rb.v.Script = v
	return rb
}
