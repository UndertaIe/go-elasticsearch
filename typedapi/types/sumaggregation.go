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

// SumAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/metric.ts#L151-L151
type SumAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// SumAggregationBuilder holds SumAggregation struct and provides a builder API.
type SumAggregationBuilder struct {
	v *SumAggregation
}

// NewSumAggregation provides a builder for the SumAggregation struct.
func NewSumAggregationBuilder() *SumAggregationBuilder {
	r := SumAggregationBuilder{
		&SumAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SumAggregation struct
func (rb *SumAggregationBuilder) Build() SumAggregation {
	return *rb.v
}

func (rb *SumAggregationBuilder) Field(field Field) *SumAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *SumAggregationBuilder) Format(format string) *SumAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *SumAggregationBuilder) Missing(missing *MissingBuilder) *SumAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *SumAggregationBuilder) Script(script *ScriptBuilder) *SumAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
