/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DefaultVersionInitParameters struct {

	// The name specified for the version when it was created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DefaultVersionObservation struct {

	// The name specified for the version when it was created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DefaultVersionParameters struct {

	// The name specified for the version when it was created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ModelInitParameters struct {

	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.
	// Structure is documented below.
	DefaultVersion []DefaultVersionInitParameters `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// The description specified for the model when it was created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more labels that you can add, to organize your models.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name specified for the model.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging *bool `json:"onlinePredictionConsoleLogging,omitempty" tf:"online_prediction_console_logging,omitempty"`

	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging *bool `json:"onlinePredictionLogging,omitempty" tf:"online_prediction_logging,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ModelObservation struct {

	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.
	// Structure is documented below.
	DefaultVersion []DefaultVersionObservation `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// The description specified for the model when it was created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/models/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more labels that you can add, to organize your models.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name specified for the model.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging *bool `json:"onlinePredictionConsoleLogging,omitempty" tf:"online_prediction_console_logging,omitempty"`

	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging *bool `json:"onlinePredictionLogging,omitempty" tf:"online_prediction_logging,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ModelParameters struct {

	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DefaultVersion []DefaultVersionParameters `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// The description specified for the model when it was created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more labels that you can add, to organize your models.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name specified for the model.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	// +kubebuilder:validation:Optional
	OnlinePredictionConsoleLogging *bool `json:"onlinePredictionConsoleLogging,omitempty" tf:"online_prediction_console_logging,omitempty"`

	// If true, online prediction access logs are sent to StackDriver Logging.
	// +kubebuilder:validation:Optional
	OnlinePredictionLogging *bool `json:"onlinePredictionLogging,omitempty" tf:"online_prediction_logging,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ModelInitParameters `json:"initProvider,omitempty"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Model is the Schema for the Models API. Represents a machine learning solution.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ModelSpec   `json:"spec"`
	Status ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	Model_Kind             = "Model"
	Model_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Model_Kind}.String()
	Model_KindAPIVersion   = Model_Kind + "." + CRDGroupVersion.String()
	Model_GroupVersionKind = CRDGroupVersion.WithKind(Model_Kind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}
