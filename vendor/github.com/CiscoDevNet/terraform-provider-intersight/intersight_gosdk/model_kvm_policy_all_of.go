/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-22T00:49:18Z.
 *
 * API version: 1.0.9-3127
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KvmPolicyAllOf Definition of the list of properties defined in 'kvm.Policy', excluding properties defined in parent classes.
type KvmPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, displays KVM session on any monitor attached to the server.
	EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
	// If enabled, encrypts all video information sent through KVM.
	EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
	// State of the vKVM service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The maximum number of concurrent KVM sessions allowed.
	MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
	// The port used for KVM communication.
	RemotePort   *int64                                `json:"RemotePort,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmPolicyAllOf KvmPolicyAllOf

// NewKvmPolicyAllOf instantiates a new KvmPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmPolicyAllOf(classId string, objectType string) *KvmPolicyAllOf {
	this := KvmPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	return &this
}

// NewKvmPolicyAllOfWithDefaults instantiates a new KvmPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmPolicyAllOfWithDefaults() *KvmPolicyAllOf {
	this := KvmPolicyAllOf{}
	var classId string = "kvm.Policy"
	this.ClassId = classId
	var objectType string = "kvm.Policy"
	this.ObjectType = objectType
	var enableLocalServerVideo bool = true
	this.EnableLocalServerVideo = &enableLocalServerVideo
	var enableVideoEncryption bool = true
	this.EnableVideoEncryption = &enableVideoEncryption
	var enabled bool = true
	this.Enabled = &enabled
	var maximumSessions int64 = 4
	this.MaximumSessions = &maximumSessions
	var remotePort int64 = 2068
	this.RemotePort = &remotePort
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableLocalServerVideo returns the EnableLocalServerVideo field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetEnableLocalServerVideo() bool {
	if o == nil || o.EnableLocalServerVideo == nil {
		var ret bool
		return ret
	}
	return *o.EnableLocalServerVideo
}

// GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetEnableLocalServerVideoOk() (*bool, bool) {
	if o == nil || o.EnableLocalServerVideo == nil {
		return nil, false
	}
	return o.EnableLocalServerVideo, true
}

// HasEnableLocalServerVideo returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasEnableLocalServerVideo() bool {
	if o != nil && o.EnableLocalServerVideo != nil {
		return true
	}

	return false
}

// SetEnableLocalServerVideo gets a reference to the given bool and assigns it to the EnableLocalServerVideo field.
func (o *KvmPolicyAllOf) SetEnableLocalServerVideo(v bool) {
	o.EnableLocalServerVideo = &v
}

// GetEnableVideoEncryption returns the EnableVideoEncryption field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetEnableVideoEncryption() bool {
	if o == nil || o.EnableVideoEncryption == nil {
		var ret bool
		return ret
	}
	return *o.EnableVideoEncryption
}

// GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetEnableVideoEncryptionOk() (*bool, bool) {
	if o == nil || o.EnableVideoEncryption == nil {
		return nil, false
	}
	return o.EnableVideoEncryption, true
}

// HasEnableVideoEncryption returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasEnableVideoEncryption() bool {
	if o != nil && o.EnableVideoEncryption != nil {
		return true
	}

	return false
}

// SetEnableVideoEncryption gets a reference to the given bool and assigns it to the EnableVideoEncryption field.
func (o *KvmPolicyAllOf) SetEnableVideoEncryption(v bool) {
	o.EnableVideoEncryption = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KvmPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumSessions returns the MaximumSessions field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetMaximumSessions() int64 {
	if o == nil || o.MaximumSessions == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSessions
}

// GetMaximumSessionsOk returns a tuple with the MaximumSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetMaximumSessionsOk() (*int64, bool) {
	if o == nil || o.MaximumSessions == nil {
		return nil, false
	}
	return o.MaximumSessions, true
}

// HasMaximumSessions returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasMaximumSessions() bool {
	if o != nil && o.MaximumSessions != nil {
		return true
	}

	return false
}

// SetMaximumSessions gets a reference to the given int64 and assigns it to the MaximumSessions field.
func (o *KvmPolicyAllOf) SetMaximumSessions(v int64) {
	o.MaximumSessions = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *KvmPolicyAllOf) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KvmPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KvmPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KvmPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *KvmPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o KvmPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableLocalServerVideo != nil {
		toSerialize["EnableLocalServerVideo"] = o.EnableLocalServerVideo
	}
	if o.EnableVideoEncryption != nil {
		toSerialize["EnableVideoEncryption"] = o.EnableVideoEncryption
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MaximumSessions != nil {
		toSerialize["MaximumSessions"] = o.MaximumSessions
	}
	if o.RemotePort != nil {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKvmPolicyAllOf := _KvmPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKvmPolicyAllOf); err == nil {
		*o = KvmPolicyAllOf(varKvmPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableLocalServerVideo")
		delete(additionalProperties, "EnableVideoEncryption")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MaximumSessions")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmPolicyAllOf struct {
	value *KvmPolicyAllOf
	isSet bool
}

func (v NullableKvmPolicyAllOf) Get() *KvmPolicyAllOf {
	return v.value
}

func (v *NullableKvmPolicyAllOf) Set(val *KvmPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmPolicyAllOf(val *KvmPolicyAllOf) *NullableKvmPolicyAllOf {
	return &NullableKvmPolicyAllOf{value: val, isSet: true}
}

func (v NullableKvmPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
