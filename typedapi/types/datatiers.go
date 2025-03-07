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

// DataTiers type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L324-L331
type DataTiers struct {
	Available   bool                     `json:"available"`
	DataCold    DataTierPhaseStatistics  `json:"data_cold"`
	DataContent DataTierPhaseStatistics  `json:"data_content"`
	DataFrozen  *DataTierPhaseStatistics `json:"data_frozen,omitempty"`
	DataHot     DataTierPhaseStatistics  `json:"data_hot"`
	DataWarm    DataTierPhaseStatistics  `json:"data_warm"`
	Enabled     bool                     `json:"enabled"`
}

// DataTiersBuilder holds DataTiers struct and provides a builder API.
type DataTiersBuilder struct {
	v *DataTiers
}

// NewDataTiers provides a builder for the DataTiers struct.
func NewDataTiersBuilder() *DataTiersBuilder {
	r := DataTiersBuilder{
		&DataTiers{},
	}

	return &r
}

// Build finalize the chain and returns the DataTiers struct
func (rb *DataTiersBuilder) Build() DataTiers {
	return *rb.v
}

func (rb *DataTiersBuilder) Available(available bool) *DataTiersBuilder {
	rb.v.Available = available
	return rb
}

func (rb *DataTiersBuilder) DataCold(datacold *DataTierPhaseStatisticsBuilder) *DataTiersBuilder {
	v := datacold.Build()
	rb.v.DataCold = v
	return rb
}

func (rb *DataTiersBuilder) DataContent(datacontent *DataTierPhaseStatisticsBuilder) *DataTiersBuilder {
	v := datacontent.Build()
	rb.v.DataContent = v
	return rb
}

func (rb *DataTiersBuilder) DataFrozen(datafrozen *DataTierPhaseStatisticsBuilder) *DataTiersBuilder {
	v := datafrozen.Build()
	rb.v.DataFrozen = &v
	return rb
}

func (rb *DataTiersBuilder) DataHot(datahot *DataTierPhaseStatisticsBuilder) *DataTiersBuilder {
	v := datahot.Build()
	rb.v.DataHot = v
	return rb
}

func (rb *DataTiersBuilder) DataWarm(datawarm *DataTierPhaseStatisticsBuilder) *DataTiersBuilder {
	v := datawarm.Build()
	rb.v.DataWarm = v
	return rb
}

func (rb *DataTiersBuilder) Enabled(enabled bool) *DataTiersBuilder {
	rb.v.Enabled = enabled
	return rb
}
