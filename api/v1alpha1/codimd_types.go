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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CodiMDSpec defines the desired state of CodiMD
type CodiMDSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// URL is an example field of CodiMD. Edit CodiMD_types.go to remove/update
	URL string `json:"url,omitempty"`
}

// CodiMDStatus defines the observed state of CodiMD
type CodiMDStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Target is the pod created by the codimd operator.
	// +optional
	Target corev1.ObjectReference `json:"target,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CodiMD is the Schema for the codimds API
type CodiMD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodiMDSpec   `json:"spec,omitempty"`
	Status CodiMDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodiMDList contains a list of CodiMD
type CodiMDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodiMD `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CodiMD{}, &CodiMDList{})
}
