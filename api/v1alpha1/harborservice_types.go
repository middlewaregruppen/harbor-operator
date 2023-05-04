/*
Copyright 2023.

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

import (
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HarborServiceSpec defines the desired state of HarborService
type HarborServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of HarborService. Edit harborservice_types.go to remove/update
	External *ExternalBackend `json:"externalBackend,omitempty"`
	Internal *InternalBackend `json:"internalBackend,omitempty"`
	// +kubebuilder:default=false
	Insecure bool `json:"insecure,omitempty"`
	// +kubebuilder:default=https
	Scheme    string     `json:"scheme,omitempty"`
	SecretRef *SecretRef `json:"secretRef,omitempty"`
}

type ExternalBackend struct {
	Host string `json:"host,omitempty"`
	Port int32  `json:"port,omitempty"`
}

type SecretRef struct {
	Name string `json:"name,omitempty"`
}

type InternalBackend struct {
	Name string                `json:"name"`
	Port v1.ServiceBackendPort `json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`
}

// HarborServiceStatus defines the observed state of HarborService
type HarborServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// HarborService is the Schema for the harborservices API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type HarborService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HarborServiceSpec   `json:"spec,omitempty"`
	Status HarborServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HarborServiceList contains a list of HarborService
type HarborServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HarborService{}, &HarborServiceList{})
}
