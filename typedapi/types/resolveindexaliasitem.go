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

// ResolveIndexAliasItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/resolve_index/ResolveIndexResponse.ts#L37-L40
type ResolveIndexAliasItem struct {
	Indices Indices `json:"indices"`
	Name    Name    `json:"name"`
}

// ResolveIndexAliasItemBuilder holds ResolveIndexAliasItem struct and provides a builder API.
type ResolveIndexAliasItemBuilder struct {
	v *ResolveIndexAliasItem
}

// NewResolveIndexAliasItem provides a builder for the ResolveIndexAliasItem struct.
func NewResolveIndexAliasItemBuilder() *ResolveIndexAliasItemBuilder {
	r := ResolveIndexAliasItemBuilder{
		&ResolveIndexAliasItem{},
	}

	return &r
}

// Build finalize the chain and returns the ResolveIndexAliasItem struct
func (rb *ResolveIndexAliasItemBuilder) Build() ResolveIndexAliasItem {
	return *rb.v
}

func (rb *ResolveIndexAliasItemBuilder) Indices(indices *IndicesBuilder) *ResolveIndexAliasItemBuilder {
	v := indices.Build()
	rb.v.Indices = v
	return rb
}

func (rb *ResolveIndexAliasItemBuilder) Name(name Name) *ResolveIndexAliasItemBuilder {
	rb.v.Name = name
	return rb
}
