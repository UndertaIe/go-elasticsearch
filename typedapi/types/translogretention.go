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

// TranslogRetention type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L373-L392
type TranslogRetention struct {
	// Age This controls the maximum duration for which translog files are kept by each
	// shard. Keeping more
	// translog files increases the chance of performing an operation based sync
	// when recovering replicas. If
	// the translog files are not sufficient, replica recovery will fall back to a
	// file based sync. This setting
	// is ignored, and should not be set, if soft deletes are enabled. Soft deletes
	// are enabled by default in
	// indices created in Elasticsearch versions 7.0.0 and later.
	Age *Duration `json:"age,omitempty"`
	// Size This controls the total size of translog files to keep for each shard.
	// Keeping more translog files increases
	// the chance of performing an operation based sync when recovering a replica.
	// If the translog files are not
	// sufficient, replica recovery will fall back to a file based sync. This
	// setting is ignored, and should not be
	// set, if soft deletes are enabled. Soft deletes are enabled by default in
	// indices created in Elasticsearch
	// versions 7.0.0 and later.
	Size *ByteSize `json:"size,omitempty"`
}

// TranslogRetentionBuilder holds TranslogRetention struct and provides a builder API.
type TranslogRetentionBuilder struct {
	v *TranslogRetention
}

// NewTranslogRetention provides a builder for the TranslogRetention struct.
func NewTranslogRetentionBuilder() *TranslogRetentionBuilder {
	r := TranslogRetentionBuilder{
		&TranslogRetention{},
	}

	return &r
}

// Build finalize the chain and returns the TranslogRetention struct
func (rb *TranslogRetentionBuilder) Build() TranslogRetention {
	return *rb.v
}

// Age This controls the maximum duration for which translog files are kept by each
// shard. Keeping more
// translog files increases the chance of performing an operation based sync
// when recovering replicas. If
// the translog files are not sufficient, replica recovery will fall back to a
// file based sync. This setting
// is ignored, and should not be set, if soft deletes are enabled. Soft deletes
// are enabled by default in
// indices created in Elasticsearch versions 7.0.0 and later.

func (rb *TranslogRetentionBuilder) Age(age *DurationBuilder) *TranslogRetentionBuilder {
	v := age.Build()
	rb.v.Age = &v
	return rb
}

// Size This controls the total size of translog files to keep for each shard.
// Keeping more translog files increases
// the chance of performing an operation based sync when recovering a replica.
// If the translog files are not
// sufficient, replica recovery will fall back to a file based sync. This
// setting is ignored, and should not be
// set, if soft deletes are enabled. Soft deletes are enabled by default in
// indices created in Elasticsearch
// versions 7.0.0 and later.

func (rb *TranslogRetentionBuilder) Size(size *ByteSizeBuilder) *TranslogRetentionBuilder {
	v := size.Build()
	rb.v.Size = &v
	return rb
}
