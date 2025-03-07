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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditiontype"
)

// ExecutionResultCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/watcher/_types/Execution.ts#L68-L72
type ExecutionResultCondition struct {
	Met    bool                                    `json:"met"`
	Status actionstatusoptions.ActionStatusOptions `json:"status"`
	Type   conditiontype.ConditionType             `json:"type"`
}

// ExecutionResultConditionBuilder holds ExecutionResultCondition struct and provides a builder API.
type ExecutionResultConditionBuilder struct {
	v *ExecutionResultCondition
}

// NewExecutionResultCondition provides a builder for the ExecutionResultCondition struct.
func NewExecutionResultConditionBuilder() *ExecutionResultConditionBuilder {
	r := ExecutionResultConditionBuilder{
		&ExecutionResultCondition{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionResultCondition struct
func (rb *ExecutionResultConditionBuilder) Build() ExecutionResultCondition {
	return *rb.v
}

func (rb *ExecutionResultConditionBuilder) Met(met bool) *ExecutionResultConditionBuilder {
	rb.v.Met = met
	return rb
}

func (rb *ExecutionResultConditionBuilder) Status(status actionstatusoptions.ActionStatusOptions) *ExecutionResultConditionBuilder {
	rb.v.Status = status
	return rb
}

func (rb *ExecutionResultConditionBuilder) Type_(type_ conditiontype.ConditionType) *ExecutionResultConditionBuilder {
	rb.v.Type = type_
	return rb
}
