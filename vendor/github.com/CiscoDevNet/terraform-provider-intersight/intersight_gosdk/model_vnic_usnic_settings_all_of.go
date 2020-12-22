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

// VnicUsnicSettingsAllOf Definition of the list of properties defined in 'vnic.UsnicSettings', excluding properties defined in parent classes.
type VnicUsnicSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Class of Service to be used for traffic on the usNIC.
	Cos *int64 `json:"Cos,omitempty"`
	// Number of usNIC interfaces to be created.
	Count *int64 `json:"Count,omitempty"`
	// Ethernet Adapter policy to be associated with the usNICs.
	UsnicAdapterPolicy   *string `json:"UsnicAdapterPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicUsnicSettingsAllOf VnicUsnicSettingsAllOf

// NewVnicUsnicSettingsAllOf instantiates a new VnicUsnicSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicUsnicSettingsAllOf(classId string, objectType string) *VnicUsnicSettingsAllOf {
	this := VnicUsnicSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cos int64 = 5
	this.Cos = &cos
	return &this
}

// NewVnicUsnicSettingsAllOfWithDefaults instantiates a new VnicUsnicSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicUsnicSettingsAllOfWithDefaults() *VnicUsnicSettingsAllOf {
	this := VnicUsnicSettingsAllOf{}
	var classId string = "vnic.UsnicSettings"
	this.ClassId = classId
	var objectType string = "vnic.UsnicSettings"
	this.ObjectType = objectType
	var cos int64 = 5
	this.Cos = &cos
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicUsnicSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicUsnicSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicUsnicSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicUsnicSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicUsnicSettingsAllOf) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettingsAllOf) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicUsnicSettingsAllOf) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicUsnicSettingsAllOf) SetCos(v int64) {
	o.Cos = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicUsnicSettingsAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettingsAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicUsnicSettingsAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicUsnicSettingsAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetUsnicAdapterPolicy returns the UsnicAdapterPolicy field value if set, zero value otherwise.
func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicy() string {
	if o == nil || o.UsnicAdapterPolicy == nil {
		var ret string
		return ret
	}
	return *o.UsnicAdapterPolicy
}

// GetUsnicAdapterPolicyOk returns a tuple with the UsnicAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicyOk() (*string, bool) {
	if o == nil || o.UsnicAdapterPolicy == nil {
		return nil, false
	}
	return o.UsnicAdapterPolicy, true
}

// HasUsnicAdapterPolicy returns a boolean if a field has been set.
func (o *VnicUsnicSettingsAllOf) HasUsnicAdapterPolicy() bool {
	if o != nil && o.UsnicAdapterPolicy != nil {
		return true
	}

	return false
}

// SetUsnicAdapterPolicy gets a reference to the given string and assigns it to the UsnicAdapterPolicy field.
func (o *VnicUsnicSettingsAllOf) SetUsnicAdapterPolicy(v string) {
	o.UsnicAdapterPolicy = &v
}

func (o VnicUsnicSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.UsnicAdapterPolicy != nil {
		toSerialize["UsnicAdapterPolicy"] = o.UsnicAdapterPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicUsnicSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicUsnicSettingsAllOf := _VnicUsnicSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicUsnicSettingsAllOf); err == nil {
		*o = VnicUsnicSettingsAllOf(varVnicUsnicSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "UsnicAdapterPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicUsnicSettingsAllOf struct {
	value *VnicUsnicSettingsAllOf
	isSet bool
}

func (v NullableVnicUsnicSettingsAllOf) Get() *VnicUsnicSettingsAllOf {
	return v.value
}

func (v *NullableVnicUsnicSettingsAllOf) Set(val *VnicUsnicSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicUsnicSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicUsnicSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicUsnicSettingsAllOf(val *VnicUsnicSettingsAllOf) *NullableVnicUsnicSettingsAllOf {
	return &NullableVnicUsnicSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicUsnicSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicUsnicSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
