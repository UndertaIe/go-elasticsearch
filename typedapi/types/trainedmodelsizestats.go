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

// TrainedModelSizeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/TrainedModel.ts#L120-L125
type TrainedModelSizeStats struct {
	// ModelSizeBytes The size of the model in bytes.
	ModelSizeBytes ByteSize `json:"model_size_bytes"`
	// RequiredNativeMemoryBytes The amount of memory required to load the model in bytes.
	RequiredNativeMemoryBytes int `json:"required_native_memory_bytes"`
}

// TrainedModelSizeStatsBuilder holds TrainedModelSizeStats struct and provides a builder API.
type TrainedModelSizeStatsBuilder struct {
	v *TrainedModelSizeStats
}

// NewTrainedModelSizeStats provides a builder for the TrainedModelSizeStats struct.
func NewTrainedModelSizeStatsBuilder() *TrainedModelSizeStatsBuilder {
	r := TrainedModelSizeStatsBuilder{
		&TrainedModelSizeStats{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelSizeStats struct
func (rb *TrainedModelSizeStatsBuilder) Build() TrainedModelSizeStats {
	return *rb.v
}

// ModelSizeBytes The size of the model in bytes.

func (rb *TrainedModelSizeStatsBuilder) ModelSizeBytes(modelsizebytes *ByteSizeBuilder) *TrainedModelSizeStatsBuilder {
	v := modelsizebytes.Build()
	rb.v.ModelSizeBytes = v
	return rb
}

// RequiredNativeMemoryBytes The amount of memory required to load the model in bytes.

func (rb *TrainedModelSizeStatsBuilder) RequiredNativeMemoryBytes(requirednativememorybytes int) *TrainedModelSizeStatsBuilder {
	rb.v.RequiredNativeMemoryBytes = requirednativememorybytes
	return rb
}
