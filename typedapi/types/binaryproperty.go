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

// BinaryProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/mapping/core.ts#L49-L51
type BinaryProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string         `json:"meta,omitempty"`
	Properties map[PropertyName]Property `json:"properties,omitempty"`
	Similarity *string                   `json:"similarity,omitempty"`
	Store      *bool                     `json:"store,omitempty"`
	Type       string                    `json:"type,omitempty"`
}

// BinaryPropertyBuilder holds BinaryProperty struct and provides a builder API.
type BinaryPropertyBuilder struct {
	v *BinaryProperty
}

// NewBinaryProperty provides a builder for the BinaryProperty struct.
func NewBinaryPropertyBuilder() *BinaryPropertyBuilder {
	r := BinaryPropertyBuilder{
		&BinaryProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "binary"

	return &r
}

// Build finalize the chain and returns the BinaryProperty struct
func (rb *BinaryPropertyBuilder) Build() BinaryProperty {
	return *rb.v
}

func (rb *BinaryPropertyBuilder) CopyTo(copyto *FieldsBuilder) *BinaryPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *BinaryPropertyBuilder) DocValues(docvalues bool) *BinaryPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *BinaryPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *BinaryPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *BinaryPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *BinaryPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *BinaryPropertyBuilder) IgnoreAbove(ignoreabove int) *BinaryPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *BinaryPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *BinaryPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *BinaryPropertyBuilder) Meta(value map[string]string) *BinaryPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *BinaryPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *BinaryPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *BinaryPropertyBuilder) Similarity(similarity string) *BinaryPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *BinaryPropertyBuilder) Store(store bool) *BinaryPropertyBuilder {
	rb.v.Store = &store
	return rb
}
