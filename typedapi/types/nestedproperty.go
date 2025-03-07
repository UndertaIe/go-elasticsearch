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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// NestedProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/mapping/complex.ts#L39-L44
type NestedProperty struct {
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled         *bool                          `json:"enabled,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IncludeInParent *bool                          `json:"include_in_parent,omitempty"`
	IncludeInRoot   *bool                          `json:"include_in_root,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string         `json:"meta,omitempty"`
	Properties map[PropertyName]Property `json:"properties,omitempty"`
	Similarity *string                   `json:"similarity,omitempty"`
	Store      *bool                     `json:"store,omitempty"`
	Type       string                    `json:"type,omitempty"`
}

// NestedPropertyBuilder holds NestedProperty struct and provides a builder API.
type NestedPropertyBuilder struct {
	v *NestedProperty
}

// NewNestedProperty provides a builder for the NestedProperty struct.
func NewNestedPropertyBuilder() *NestedPropertyBuilder {
	r := NestedPropertyBuilder{
		&NestedProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "nested"

	return &r
}

// Build finalize the chain and returns the NestedProperty struct
func (rb *NestedPropertyBuilder) Build() NestedProperty {
	return *rb.v
}

func (rb *NestedPropertyBuilder) CopyTo(copyto *FieldsBuilder) *NestedPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *NestedPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *NestedPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *NestedPropertyBuilder) Enabled(enabled bool) *NestedPropertyBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *NestedPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *NestedPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *NestedPropertyBuilder) IgnoreAbove(ignoreabove int) *NestedPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *NestedPropertyBuilder) IncludeInParent(includeinparent bool) *NestedPropertyBuilder {
	rb.v.IncludeInParent = &includeinparent
	return rb
}

func (rb *NestedPropertyBuilder) IncludeInRoot(includeinroot bool) *NestedPropertyBuilder {
	rb.v.IncludeInRoot = &includeinroot
	return rb
}

func (rb *NestedPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *NestedPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *NestedPropertyBuilder) Meta(value map[string]string) *NestedPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *NestedPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *NestedPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *NestedPropertyBuilder) Similarity(similarity string) *NestedPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *NestedPropertyBuilder) Store(store bool) *NestedPropertyBuilder {
	rb.v.Store = &store
	return rb
}
