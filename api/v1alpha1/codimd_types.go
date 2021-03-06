/*

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

// CodiMDSpec defines the desired state of CodiMD.
type CodiMDSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	// URL is the url of a codiMD markdown file.
	URL string `json:"url,omitempty"`
}

// CodiMDStatus defines the observed state of CodiMD.
type CodiMDStatus struct {
	// Target is the deployment created by the codimd operator.
	// +optional
	Target CodiMDStatusTarget `json:"target,omitempty"`
}

// CodiMDStatusTarget defines the observed state of a Deployment from CodiMD.
type CodiMDStatusTarget struct {
	// Name is the pod created by the codimd operator.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace is the pod created by the codimd operator.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="URL",type=string,JSONPath=`.spec.url`
// CodiMD is the Schema for the codimds API.
type CodiMD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodiMDSpec   `json:"spec,omitempty"`
	Status CodiMDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// CodiMDList contains a list of CodiMD.
type CodiMDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodiMD `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CodiMD{}, &CodiMDList{})
}
