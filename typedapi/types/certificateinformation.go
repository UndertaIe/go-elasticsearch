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

// CertificateInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ssl/certificates/types.ts#L22-L31
type CertificateInformation struct {
	Alias         string   `json:"alias,omitempty"`
	Expiry        DateTime `json:"expiry"`
	Format        string   `json:"format"`
	HasPrivateKey bool     `json:"has_private_key"`
	Issuer        *string  `json:"issuer,omitempty"`
	Path          string   `json:"path"`
	SerialNumber  string   `json:"serial_number"`
	SubjectDn     string   `json:"subject_dn"`
}

// CertificateInformationBuilder holds CertificateInformation struct and provides a builder API.
type CertificateInformationBuilder struct {
	v *CertificateInformation
}

// NewCertificateInformation provides a builder for the CertificateInformation struct.
func NewCertificateInformationBuilder() *CertificateInformationBuilder {
	r := CertificateInformationBuilder{
		&CertificateInformation{},
	}

	return &r
}

// Build finalize the chain and returns the CertificateInformation struct
func (rb *CertificateInformationBuilder) Build() CertificateInformation {
	return *rb.v
}

func (rb *CertificateInformationBuilder) Alias(alias string) *CertificateInformationBuilder {
	rb.v.Alias = alias
	return rb
}

func (rb *CertificateInformationBuilder) Expiry(expiry *DateTimeBuilder) *CertificateInformationBuilder {
	v := expiry.Build()
	rb.v.Expiry = v
	return rb
}

func (rb *CertificateInformationBuilder) Format(format string) *CertificateInformationBuilder {
	rb.v.Format = format
	return rb
}

func (rb *CertificateInformationBuilder) HasPrivateKey(hasprivatekey bool) *CertificateInformationBuilder {
	rb.v.HasPrivateKey = hasprivatekey
	return rb
}

func (rb *CertificateInformationBuilder) Issuer(issuer string) *CertificateInformationBuilder {
	rb.v.Issuer = &issuer
	return rb
}

func (rb *CertificateInformationBuilder) Path(path string) *CertificateInformationBuilder {
	rb.v.Path = path
	return rb
}

func (rb *CertificateInformationBuilder) SerialNumber(serialnumber string) *CertificateInformationBuilder {
	rb.v.SerialNumber = serialnumber
	return rb
}

func (rb *CertificateInformationBuilder) SubjectDn(subjectdn string) *CertificateInformationBuilder {
	rb.v.SubjectDn = subjectdn
	return rb
}
