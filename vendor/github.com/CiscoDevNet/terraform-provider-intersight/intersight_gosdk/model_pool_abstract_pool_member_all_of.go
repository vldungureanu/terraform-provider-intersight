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

// PoolAbstractPoolMemberAllOf Definition of the list of properties defined in 'pool.AbstractPoolMember', excluding properties defined in parent classes.
type PoolAbstractPoolMemberAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Boolean to represent whether the ID is assigned or not.
	Assigned             *bool `json:"Assigned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractPoolMemberAllOf PoolAbstractPoolMemberAllOf

// NewPoolAbstractPoolMemberAllOf instantiates a new PoolAbstractPoolMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractPoolMemberAllOf(classId string, objectType string) *PoolAbstractPoolMemberAllOf {
	this := PoolAbstractPoolMemberAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// NewPoolAbstractPoolMemberAllOfWithDefaults instantiates a new PoolAbstractPoolMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractPoolMemberAllOfWithDefaults() *PoolAbstractPoolMemberAllOf {
	this := PoolAbstractPoolMemberAllOf{}
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractPoolMemberAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMemberAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractPoolMemberAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractPoolMemberAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMemberAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractPoolMemberAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PoolAbstractPoolMemberAllOf) GetAssigned() bool {
	if o == nil || o.Assigned == nil {
		var ret bool
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolMemberAllOf) GetAssignedOk() (*bool, bool) {
	if o == nil || o.Assigned == nil {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PoolAbstractPoolMemberAllOf) HasAssigned() bool {
	if o != nil && o.Assigned != nil {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given bool and assigns it to the Assigned field.
func (o *PoolAbstractPoolMemberAllOf) SetAssigned(v bool) {
	o.Assigned = &v
}

func (o PoolAbstractPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Assigned != nil {
		toSerialize["Assigned"] = o.Assigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractPoolMemberAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractPoolMemberAllOf := _PoolAbstractPoolMemberAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractPoolMemberAllOf); err == nil {
		*o = PoolAbstractPoolMemberAllOf(varPoolAbstractPoolMemberAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Assigned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractPoolMemberAllOf struct {
	value *PoolAbstractPoolMemberAllOf
	isSet bool
}

func (v NullablePoolAbstractPoolMemberAllOf) Get() *PoolAbstractPoolMemberAllOf {
	return v.value
}

func (v *NullablePoolAbstractPoolMemberAllOf) Set(val *PoolAbstractPoolMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractPoolMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractPoolMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractPoolMemberAllOf(val *PoolAbstractPoolMemberAllOf) *NullablePoolAbstractPoolMemberAllOf {
	return &NullablePoolAbstractPoolMemberAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractPoolMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
