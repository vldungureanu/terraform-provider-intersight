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
	"reflect"
	"strings"
)

// VmediaPolicy Policy to configure virtual media settings on endpoint.
type VmediaPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Virtual Media service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// If enabled, allows encryption of all Virtual Media communications.
	Encryption *bool `json:"Encryption,omitempty"`
	// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
	LowPowerUsb  *bool                                 `json:"LowPowerUsb,omitempty"`
	Mappings     []VmediaMapping                       `json:"Mappings,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmediaPolicy VmediaPolicy

// NewVmediaPolicy instantiates a new VmediaPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmediaPolicy(classId string, objectType string) *VmediaPolicy {
	this := VmediaPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var lowPowerUsb bool = true
	this.LowPowerUsb = &lowPowerUsb
	return &this
}

// NewVmediaPolicyWithDefaults instantiates a new VmediaPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmediaPolicyWithDefaults() *VmediaPolicy {
	this := VmediaPolicy{}
	var classId string = "vmedia.Policy"
	this.ClassId = classId
	var objectType string = "vmedia.Policy"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var lowPowerUsb bool = true
	this.LowPowerUsb = &lowPowerUsb
	return &this
}

// GetClassId returns the ClassId field value
func (o *VmediaPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VmediaPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VmediaPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VmediaPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VmediaPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VmediaPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VmediaPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *VmediaPolicy) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *VmediaPolicy) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *VmediaPolicy) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetLowPowerUsb returns the LowPowerUsb field value if set, zero value otherwise.
func (o *VmediaPolicy) GetLowPowerUsb() bool {
	if o == nil || o.LowPowerUsb == nil {
		var ret bool
		return ret
	}
	return *o.LowPowerUsb
}

// GetLowPowerUsbOk returns a tuple with the LowPowerUsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetLowPowerUsbOk() (*bool, bool) {
	if o == nil || o.LowPowerUsb == nil {
		return nil, false
	}
	return o.LowPowerUsb, true
}

// HasLowPowerUsb returns a boolean if a field has been set.
func (o *VmediaPolicy) HasLowPowerUsb() bool {
	if o != nil && o.LowPowerUsb != nil {
		return true
	}

	return false
}

// SetLowPowerUsb gets a reference to the given bool and assigns it to the LowPowerUsb field.
func (o *VmediaPolicy) SetLowPowerUsb(v bool) {
	o.LowPowerUsb = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmediaPolicy) GetMappings() []VmediaMapping {
	if o == nil {
		var ret []VmediaMapping
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmediaPolicy) GetMappingsOk() (*[]VmediaMapping, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *VmediaPolicy) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []VmediaMapping and assigns it to the Mappings field.
func (o *VmediaPolicy) SetMappings(v []VmediaMapping) {
	o.Mappings = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VmediaPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmediaPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VmediaPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VmediaPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmediaPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmediaPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *VmediaPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *VmediaPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o VmediaPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.LowPowerUsb != nil {
		toSerialize["LowPowerUsb"] = o.LowPowerUsb
	}
	if o.Mappings != nil {
		toSerialize["Mappings"] = o.Mappings
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

func (o *VmediaPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VmediaPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the Virtual Media service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// If enabled, allows encryption of all Virtual Media communications.
		Encryption *bool `json:"Encryption,omitempty"`
		// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
		LowPowerUsb  *bool                                 `json:"LowPowerUsb,omitempty"`
		Mappings     []VmediaMapping                       `json:"Mappings,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varVmediaPolicyWithoutEmbeddedStruct := VmediaPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVmediaPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVmediaPolicy := _VmediaPolicy{}
		varVmediaPolicy.ClassId = varVmediaPolicyWithoutEmbeddedStruct.ClassId
		varVmediaPolicy.ObjectType = varVmediaPolicyWithoutEmbeddedStruct.ObjectType
		varVmediaPolicy.Enabled = varVmediaPolicyWithoutEmbeddedStruct.Enabled
		varVmediaPolicy.Encryption = varVmediaPolicyWithoutEmbeddedStruct.Encryption
		varVmediaPolicy.LowPowerUsb = varVmediaPolicyWithoutEmbeddedStruct.LowPowerUsb
		varVmediaPolicy.Mappings = varVmediaPolicyWithoutEmbeddedStruct.Mappings
		varVmediaPolicy.Organization = varVmediaPolicyWithoutEmbeddedStruct.Organization
		varVmediaPolicy.Profiles = varVmediaPolicyWithoutEmbeddedStruct.Profiles
		*o = VmediaPolicy(varVmediaPolicy)
	} else {
		return err
	}

	varVmediaPolicy := _VmediaPolicy{}

	err = json.Unmarshal(bytes, &varVmediaPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVmediaPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "LowPowerUsb")
		delete(additionalProperties, "Mappings")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVmediaPolicy struct {
	value *VmediaPolicy
	isSet bool
}

func (v NullableVmediaPolicy) Get() *VmediaPolicy {
	return v.value
}

func (v *NullableVmediaPolicy) Set(val *VmediaPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVmediaPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVmediaPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmediaPolicy(val *VmediaPolicy) *NullableVmediaPolicy {
	return &NullableVmediaPolicy{value: val, isSet: true}
}

func (v NullableVmediaPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmediaPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
