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

type MessageStoragePolicyInitParameters struct {

	// A list of IDs of GCP regions where messages that are published to
	// the topic may be persisted in storage. Messages published by
	// publishers running in non-allowed GCP regions (or running outside
	// of GCP altogether) will be routed for storage in one of the
	// allowed regions. An empty list means that no regions are allowed,
	// and is not a valid configuration.
	AllowedPersistenceRegions []*string `json:"allowedPersistenceRegions,omitempty" tf:"allowed_persistence_regions,omitempty"`
}

type MessageStoragePolicyObservation struct {

	// A list of IDs of GCP regions where messages that are published to
	// the topic may be persisted in storage. Messages published by
	// publishers running in non-allowed GCP regions (or running outside
	// of GCP altogether) will be routed for storage in one of the
	// allowed regions. An empty list means that no regions are allowed,
	// and is not a valid configuration.
	AllowedPersistenceRegions []*string `json:"allowedPersistenceRegions,omitempty" tf:"allowed_persistence_regions,omitempty"`
}

type MessageStoragePolicyParameters struct {

	// A list of IDs of GCP regions where messages that are published to
	// the topic may be persisted in storage. Messages published by
	// publishers running in non-allowed GCP regions (or running outside
	// of GCP altogether) will be routed for storage in one of the
	// allowed regions. An empty list means that no regions are allowed,
	// and is not a valid configuration.
	// +kubebuilder:validation:Optional
	AllowedPersistenceRegions []*string `json:"allowedPersistenceRegions" tf:"allowed_persistence_regions,omitempty"`
}

type SchemaSettingsInitParameters struct {

	// The encoding of messages validated against schema.
	// Default value is ENCODING_UNSPECIFIED.
	// Possible values are: ENCODING_UNSPECIFIED, JSON, BINARY.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The name of the schema that messages published should be
	// validated against. Format is projects/{project}/schemas/{schema}.
	// The value of this field will be deleted-schema
	// if the schema has been deleted.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type SchemaSettingsObservation struct {

	// The encoding of messages validated against schema.
	// Default value is ENCODING_UNSPECIFIED.
	// Possible values are: ENCODING_UNSPECIFIED, JSON, BINARY.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The name of the schema that messages published should be
	// validated against. Format is projects/{project}/schemas/{schema}.
	// The value of this field will be deleted-schema
	// if the schema has been deleted.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type SchemaSettingsParameters struct {

	// The encoding of messages validated against schema.
	// Default value is ENCODING_UNSPECIFIED.
	// Possible values are: ENCODING_UNSPECIFIED, JSON, BINARY.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The name of the schema that messages published should be
	// validated against. Format is projects/{project}/schemas/{schema}.
	// The value of this field will be deleted-schema
	// if the schema has been deleted.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema" tf:"schema,omitempty"`
}

type TopicInitParameters struct {

	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Indicates the minimum duration to retain a message after it is published
	// to the topic. If this field is set, messages published to the topic in
	// the last messageRetentionDuration are always available to subscribers.
	// For instance, it allows any attached subscription to seek to a timestamp
	// that is up to messageRetentionDuration in the past. If this field is not
	// set, message retention is controlled by settings on individual subscriptions.
	// Cannot be more than 31 days or less than 10 minutes.
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty" tf:"message_retention_duration,omitempty"`

	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy []MessageStoragePolicyInitParameters `json:"messageStoragePolicy,omitempty" tf:"message_storage_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Settings for validating messages published against a schema.
	// Structure is documented below.
	SchemaSettings []SchemaSettingsInitParameters `json:"schemaSettings,omitempty" tf:"schema_settings,omitempty"`
}

type TopicObservation struct {

	// an identifier for the resource with format projects/{{project}}/topics/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// The expected format is projects/*/locations/*/keyRings/*/cryptoKeys/*
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Indicates the minimum duration to retain a message after it is published
	// to the topic. If this field is set, messages published to the topic in
	// the last messageRetentionDuration are always available to subscribers.
	// For instance, it allows any attached subscription to seek to a timestamp
	// that is up to messageRetentionDuration in the past. If this field is not
	// set, message retention is controlled by settings on individual subscriptions.
	// Cannot be more than 31 days or less than 10 minutes.
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty" tf:"message_retention_duration,omitempty"`

	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy []MessageStoragePolicyObservation `json:"messageStoragePolicy,omitempty" tf:"message_storage_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Settings for validating messages published against a schema.
	// Structure is documented below.
	SchemaSettings []SchemaSettingsObservation `json:"schemaSettings,omitempty" tf:"schema_settings,omitempty"`
}

type TopicParameters struct {

	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// The expected format is projects/*/locations/*/keyRings/*/cryptoKeys/*
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to this Topic.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Indicates the minimum duration to retain a message after it is published
	// to the topic. If this field is set, messages published to the topic in
	// the last messageRetentionDuration are always available to subscribers.
	// For instance, it allows any attached subscription to seek to a timestamp
	// that is up to messageRetentionDuration in the past. If this field is not
	// set, message retention is controlled by settings on individual subscriptions.
	// Cannot be more than 31 days or less than 10 minutes.
	// +kubebuilder:validation:Optional
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty" tf:"message_retention_duration,omitempty"`

	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MessageStoragePolicy []MessageStoragePolicyParameters `json:"messageStoragePolicy,omitempty" tf:"message_storage_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Settings for validating messages published against a schema.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SchemaSettings []SchemaSettingsParameters `json:"schemaSettings,omitempty" tf:"schema_settings,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
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
	InitProvider TopicInitParameters `json:"initProvider,omitempty"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API. A named resource to which messages are sent by publishers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec"`
	Status            TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
