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


package explaindataframeanalytics

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package explaindataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/explain_data_frame_analytics/MlExplainDataFrameAnalyticsRequest.ts#L30-L107
type Request struct {

	// AllowLazyStart Specifies whether this job can start when there is insufficient machine
	// learning node capacity for it to be immediately assigned to a node.
	AllowLazyStart *bool `json:"allow_lazy_start,omitempty"`

	// Analysis The analysis configuration, which contains the information necessary to
	// perform one of the following types of analysis: classification, outlier
	// detection, or regression.
	Analysis *types.DataframeAnalysisContainer `json:"analysis,omitempty"`

	// AnalyzedFields Specify includes and/or excludes patterns to select which fields will be
	// included in the analysis. The patterns specified in excludes are applied
	// last, therefore excludes takes precedence. In other words, if the same
	// field is specified in both includes and excludes, then the field will not
	// be included in the analysis.
	AnalyzedFields *types.DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`

	// Description A description of the job.
	Description *string `json:"description,omitempty"`

	// Dest The destination configuration, consisting of index and optionally
	// results_field (ml by default).
	Dest *types.DataframeAnalyticsDestination `json:"dest,omitempty"`

	// MaxNumThreads The maximum number of threads to be used by the analysis. Using more
	// threads may decrease the time necessary to complete the analysis at the
	// cost of using more CPU. Note that the process may use additional threads
	// for operational functionality other than the analysis itself.
	MaxNumThreads *int `json:"max_num_threads,omitempty"`

	// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for
	// analytical processing. If your `elasticsearch.yml` file contains an
	// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try to
	// create data frame analytics jobs that have `model_memory_limit` values
	// greater than that setting.
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`

	// Source The configuration of how to source the analysis data. It requires an
	// index. Optionally, query and _source may be specified.
	Source *types.DataframeAnalyticsSource `json:"source,omitempty"`
}

// RequestBuilder is the builder API for the explaindataframeanalytics.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Explaindataframeanalytics request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AllowLazyStart(allowlazystart bool) *RequestBuilder {
	rb.v.AllowLazyStart = &allowlazystart
	return rb
}

func (rb *RequestBuilder) Analysis(analysis *types.DataframeAnalysisContainerBuilder) *RequestBuilder {
	v := analysis.Build()
	rb.v.Analysis = &v
	return rb
}

func (rb *RequestBuilder) AnalyzedFields(analyzedfields *types.DataframeAnalysisAnalyzedFieldsBuilder) *RequestBuilder {
	v := analyzedfields.Build()
	rb.v.AnalyzedFields = &v
	return rb
}

func (rb *RequestBuilder) Description(description string) *RequestBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *RequestBuilder) Dest(dest *types.DataframeAnalyticsDestinationBuilder) *RequestBuilder {
	v := dest.Build()
	rb.v.Dest = &v
	return rb
}

func (rb *RequestBuilder) MaxNumThreads(maxnumthreads int) *RequestBuilder {
	rb.v.MaxNumThreads = &maxnumthreads
	return rb
}

func (rb *RequestBuilder) ModelMemoryLimit(modelmemorylimit string) *RequestBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

func (rb *RequestBuilder) Source(source *types.DataframeAnalyticsSourceBuilder) *RequestBuilder {
	v := source.Build()
	rb.v.Source = &v
	return rb
}
