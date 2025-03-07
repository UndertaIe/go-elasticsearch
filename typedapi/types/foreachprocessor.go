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

// ForeachProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ingest/_types/Processors.ts#L215-L219
type ForeachProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Processor     *ProcessorContainer  `json:"processor,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// ForeachProcessorBuilder holds ForeachProcessor struct and provides a builder API.
type ForeachProcessorBuilder struct {
	v *ForeachProcessor
}

// NewForeachProcessor provides a builder for the ForeachProcessor struct.
func NewForeachProcessorBuilder() *ForeachProcessorBuilder {
	r := ForeachProcessorBuilder{
		&ForeachProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the ForeachProcessor struct
func (rb *ForeachProcessorBuilder) Build() ForeachProcessor {
	return *rb.v
}

func (rb *ForeachProcessorBuilder) Field(field Field) *ForeachProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *ForeachProcessorBuilder) If_(if_ string) *ForeachProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *ForeachProcessorBuilder) IgnoreFailure(ignorefailure bool) *ForeachProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *ForeachProcessorBuilder) IgnoreMissing(ignoremissing bool) *ForeachProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *ForeachProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *ForeachProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *ForeachProcessorBuilder) Processor(processor *ProcessorContainerBuilder) *ForeachProcessorBuilder {
	v := processor.Build()
	rb.v.Processor = &v
	return rb
}

func (rb *ForeachProcessorBuilder) Tag(tag string) *ForeachProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
