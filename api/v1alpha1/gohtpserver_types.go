/*
Copyright 2022 dlbppx.

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

// gohttpserverSpec defines the desired state of gohttpserver
type gohttpserverSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of gohttpserver. Edit gohttpserver_types.go to remove/update
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Image     string `json:"image,omitempty"`
	Replicas  *int32 `json:"replicas,omitempty"`
}

// gohttpserverStatus defines the observed state of gohttpserver
type gohttpserverStatus struct {
	AvailableReplicas int `json:"availablereplicas,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="replicas",type="integer",JSONPath=".spec.replicas",description="Replicas of Gohttpserver"
//+kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.image",description="Use Image"
// gohttpserver is the Schema for the gohttpservers API
type gohttpserver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   gohttpserverSpec   `json:"spec,omitempty"`
	Status gohttpserverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// gohttpserverList contains a list of gohttpserver
type gohttpserverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []gohttpserver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&gohttpserver{}, &gohttpserverList{})
}
