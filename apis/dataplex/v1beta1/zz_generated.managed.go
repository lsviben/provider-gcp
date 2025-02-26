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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Asset.
func (mg *Asset) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Asset.
func (mg *Asset) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Asset.
func (mg *Asset) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Asset.
func (mg *Asset) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Asset.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Asset) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Asset.
func (mg *Asset) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Asset.
func (mg *Asset) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Asset.
func (mg *Asset) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Asset.
func (mg *Asset) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Asset.
func (mg *Asset) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Asset.
func (mg *Asset) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Asset.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Asset) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Asset.
func (mg *Asset) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Asset.
func (mg *Asset) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Lake.
func (mg *Lake) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Lake.
func (mg *Lake) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Lake.
func (mg *Lake) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Lake.
func (mg *Lake) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Lake.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Lake) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Lake.
func (mg *Lake) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Lake.
func (mg *Lake) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Lake.
func (mg *Lake) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Lake.
func (mg *Lake) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Lake.
func (mg *Lake) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Lake.
func (mg *Lake) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Lake.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Lake) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Lake.
func (mg *Lake) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Lake.
func (mg *Lake) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Zone.
func (mg *Zone) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Zone.
func (mg *Zone) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this Zone.
func (mg *Zone) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this Zone.
func (mg *Zone) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Zone.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Zone) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Zone.
func (mg *Zone) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Zone.
func (mg *Zone) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Zone.
func (mg *Zone) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Zone.
func (mg *Zone) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this Zone.
func (mg *Zone) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this Zone.
func (mg *Zone) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Zone.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Zone) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Zone.
func (mg *Zone) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Zone.
func (mg *Zone) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
