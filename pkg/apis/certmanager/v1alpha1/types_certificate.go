/*
Copyright 2018 The Jetstack cert-manager contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:resource:path=certificates
// Certificate is a type to represent a Certificate from ACME
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec,omitempty"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateList is a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Certificate `json:"items"`
}

type KeyAlgorithm string

const (
	RSAKeyAlgorithm   KeyAlgorithm = "rsa"
	ECDSAKeyAlgorithm KeyAlgorithm = "ecdsa"
)

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	// CommonName is a common name to be used on the Certificate
	CommonName string `json:"commonName,omitempty"`

	// Organization is the organization to be used on the Certificate
	Organization []string `json:"organization,omitempty"`

	// DNSNames is a list of subject alt names to be used on the Certificate
	DNSNames []string `json:"dnsNames,omitempty"`

	// SecretName is the name of the secret resource to store this secret in
	SecretName string `json:"secretName"`

	// IssuerRef is a reference to the issuer for this certificate.
	// If the 'kind' field is not set, or set to 'Issuer', an Issuer resource
	// with the given name in the same namespace as the Certificate will be used.
	// If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the
	// provided name will be used.
	// The 'name' field in this stanza is required at all times.
	IssuerRef ObjectReference `json:"issuerRef"`

	// IsCA will mark this Certificate as valid for signing.
	// This implies that the 'signing' usage is set
	IsCA bool `json:"isCA,omitempty"`

	// ACME contains configuration specific to ACME Certificates.
	// Notably, this contains details on how the domain names listed on this
	// Certificate resource should be 'solved', i.e. mapping HTTP01 and DNS01
	// providers to DNS names.
	ACME *ACMECertificateConfig `json:"acme,omitempty"`

	// KeySize is the key bit size of the corresponding private key for this certificate.
	// If provided, value must be between 2048 and 8192 inclusive when KeyAlgorithm is
	// empty or is set to "rsa", and value must be one of (256, 384, 521) when
	// KeyAlgorithm is set to "ecdsa".
	KeySize int `json:"keySize,omitempty"`
	// KeyAlgorithm is the private key algorithm of the corresponding private key
	// for this certificate. If provided, allowed values are either "rsa" or "ecdsa"
	// If KeyAlgorithm is specified and KeySize is not provided,
	// key size of 256 will be used for "ecdsa" key algorithm and
	// key size of 2048 will be used for "rsa" key algorithm.
	KeyAlgorithm KeyAlgorithm `json:"keyAlgorithm,omitempty"`
}

// ACMECertificateConfig contains the configuration for the ACME certificate provider
type ACMECertificateConfig struct {
	Config []DomainSolverConfig `json:"config"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	Conditions      []CertificateCondition `json:"conditions,omitempty"`
	LastFailureTime *metav1.Time           `json:"lastFailureTime,omitempty"`
}

// CertificateCondition contains condition information for an Certificate.
type CertificateCondition struct {
	// Type of the condition, currently ('Ready').
	Type CertificateConditionType `json:"type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status ConditionStatus `json:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message"`
}

// CertificateConditionType represents an Certificate condition value.
type CertificateConditionType string

const (
	// CertificateConditionReady represents the fact that a given Certificate condition
	// is in ready state.
	CertificateConditionReady CertificateConditionType = "Ready"

	// CertificateConditionValidationFailed is used to indicate whether a
	// validation for a Certificate has failed.
	// This is currently used by the ACME issuer to track when the last
	// validation was attempted.
	CertificateConditionValidationFailed CertificateConditionType = "ValidateFailed"
)
