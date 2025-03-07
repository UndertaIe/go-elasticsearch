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

// AnalysisMemoryLimit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/Analysis.ts#L117-L122
type AnalysisMemoryLimit struct {
	// ModelMemoryLimit Limits can be applied for the resources required to hold the mathematical
	// models in memory. These limits are approximate and can be set per job. They
	// do not control the memory used by other processes, for example the
	// Elasticsearch Java processes.
	ModelMemoryLimit string `json:"model_memory_limit"`
}

// AnalysisMemoryLimitBuilder holds AnalysisMemoryLimit struct and provides a builder API.
type AnalysisMemoryLimitBuilder struct {
	v *AnalysisMemoryLimit
}

// NewAnalysisMemoryLimit provides a builder for the AnalysisMemoryLimit struct.
func NewAnalysisMemoryLimitBuilder() *AnalysisMemoryLimitBuilder {
	r := AnalysisMemoryLimitBuilder{
		&AnalysisMemoryLimit{},
	}

	return &r
}

// Build finalize the chain and returns the AnalysisMemoryLimit struct
func (rb *AnalysisMemoryLimitBuilder) Build() AnalysisMemoryLimit {
	return *rb.v
}

// ModelMemoryLimit Limits can be applied for the resources required to hold the mathematical
// models in memory. These limits are approximate and can be set per job. They
// do not control the memory used by other processes, for example the
// Elasticsearch Java processes.

func (rb *AnalysisMemoryLimitBuilder) ModelMemoryLimit(modelmemorylimit string) *AnalysisMemoryLimitBuilder {
	rb.v.ModelMemoryLimit = modelmemorylimit
	return rb
}
