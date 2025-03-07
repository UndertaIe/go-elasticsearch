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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/templateformat"
)

// RoleTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/get_role/types.ts#L47-L50
type RoleTemplate struct {
	Format   *templateformat.TemplateFormat `json:"format,omitempty"`
	Template Script                         `json:"template"`
}

// RoleTemplateBuilder holds RoleTemplate struct and provides a builder API.
type RoleTemplateBuilder struct {
	v *RoleTemplate
}

// NewRoleTemplate provides a builder for the RoleTemplate struct.
func NewRoleTemplateBuilder() *RoleTemplateBuilder {
	r := RoleTemplateBuilder{
		&RoleTemplate{},
	}

	return &r
}

// Build finalize the chain and returns the RoleTemplate struct
func (rb *RoleTemplateBuilder) Build() RoleTemplate {
	return *rb.v
}

func (rb *RoleTemplateBuilder) Format(format templateformat.TemplateFormat) *RoleTemplateBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *RoleTemplateBuilder) Template(template *ScriptBuilder) *RoleTemplateBuilder {
	v := template.Build()
	rb.v.Template = v
	return rb
}
