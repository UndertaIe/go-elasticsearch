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

// UnitFloatMillis type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/Time.ts#L37-L38
type UnitFloatMillis float64

// UnitFloatMillisBuilder holds UnitFloatMillis struct and provides a builder API.
type UnitFloatMillisBuilder struct {
	v UnitFloatMillis
}

// NewUnitFloatMillis provides a builder for the UnitFloatMillis struct.
func NewUnitFloatMillisBuilder() *UnitFloatMillisBuilder {
	return &UnitFloatMillisBuilder{}
}

// Build finalize the chain and returns the UnitFloatMillis struct
func (b *UnitFloatMillisBuilder) Build() UnitFloatMillis {
	return b.v
}

func (b *UnitFloatMillisBuilder) UnitFloatMillis(value UnitFloatMillis) *UnitFloatMillisBuilder {
	b.v = value
	return b
}
