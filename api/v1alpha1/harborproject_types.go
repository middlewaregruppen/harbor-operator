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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HarborProjectSpec defines the desired state of HarborProject
type HarborProjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	IsPrivate      bool            `json:"isPrivate,omitempty"`
	ProxyCacheSpec *ProxyCacheSpec `json:"proxyCache,omitempty"`
	Harbor         string          `json:"harbor,omitempty"`
}

// HarborProjectStatus defines the observed state of HarborProject
type HarborProjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

type ProxyCacheSpec struct {
	Registry string `json:"registry,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HarborProject is the Schema for the harborprojects API
type HarborProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HarborProjectSpec   `json:"spec,omitempty"`
	Status HarborProjectStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HarborProjectList contains a list of HarborProject
type HarborProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HarborProject{}, &HarborProjectList{})
}
