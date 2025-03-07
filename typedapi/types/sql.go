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

// Sql type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L368-L371
type Sql struct {
	Available bool             `json:"available"`
	Enabled   bool             `json:"enabled"`
	Features  map[string]int   `json:"features"`
	Queries   map[string]Query `json:"queries"`
}

// SqlBuilder holds Sql struct and provides a builder API.
type SqlBuilder struct {
	v *Sql
}

// NewSql provides a builder for the Sql struct.
func NewSqlBuilder() *SqlBuilder {
	r := SqlBuilder{
		&Sql{
			Features: make(map[string]int, 0),
			Queries:  make(map[string]Query, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Sql struct
func (rb *SqlBuilder) Build() Sql {
	return *rb.v
}

func (rb *SqlBuilder) Available(available bool) *SqlBuilder {
	rb.v.Available = available
	return rb
}

func (rb *SqlBuilder) Enabled(enabled bool) *SqlBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *SqlBuilder) Features(value map[string]int) *SqlBuilder {
	rb.v.Features = value
	return rb
}

func (rb *SqlBuilder) Queries(values map[string]*QueryBuilder) *SqlBuilder {
	tmp := make(map[string]Query, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Queries = tmp
	return rb
}
