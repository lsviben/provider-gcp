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

type EnvgroupInitParameters struct {

	// Hostnames of the environment group.
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`
}

type EnvgroupObservation struct {

	// Hostnames of the environment group.
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// an identifier for the resource with format {{org_id}}/envgroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Apigee Organization associated with the Apigee environment group,
	// in the format organizations/{{org_name}}.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type EnvgroupParameters struct {

	// Hostnames of the environment group.
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// The Apigee Organization associated with the Apigee environment group,
	// in the format organizations/{{org_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`
}

// EnvgroupSpec defines the desired state of Envgroup
type EnvgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvgroupParameters `json:"forProvider"`
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
	InitProvider EnvgroupInitParameters `json:"initProvider,omitempty"`
}

// EnvgroupStatus defines the observed state of Envgroup.
type EnvgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Envgroup is the Schema for the Envgroups API. An
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Envgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvgroupSpec   `json:"spec"`
	Status            EnvgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvgroupList contains a list of Envgroups
type EnvgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Envgroup `json:"items"`
}

// Repository type metadata.
var (
	Envgroup_Kind             = "Envgroup"
	Envgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Envgroup_Kind}.String()
	Envgroup_KindAPIVersion   = Envgroup_Kind + "." + CRDGroupVersion.String()
	Envgroup_GroupVersionKind = CRDGroupVersion.WithKind(Envgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Envgroup{}, &EnvgroupList{})
}
