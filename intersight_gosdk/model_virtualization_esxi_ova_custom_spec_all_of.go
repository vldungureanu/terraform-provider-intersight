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

// VirtualizationEsxiOvaCustomSpecAllOf Definition of the list of properties defined in 'virtualization.EsxiOvaCustomSpec', excluding properties defined in parent classes.
type VirtualizationEsxiOvaCustomSpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the OVA Environment specification which can be configured on the virtual machine.
	OvaEnvSpec           interface{} `json:"OvaEnvSpec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiOvaCustomSpecAllOf VirtualizationEsxiOvaCustomSpecAllOf

// NewVirtualizationEsxiOvaCustomSpecAllOf instantiates a new VirtualizationEsxiOvaCustomSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiOvaCustomSpecAllOf(classId string, objectType string) *VirtualizationEsxiOvaCustomSpecAllOf {
	this := VirtualizationEsxiOvaCustomSpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiOvaCustomSpecAllOfWithDefaults instantiates a new VirtualizationEsxiOvaCustomSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiOvaCustomSpecAllOfWithDefaults() *VirtualizationEsxiOvaCustomSpecAllOf {
	this := VirtualizationEsxiOvaCustomSpecAllOf{}
	var classId string = "virtualization.EsxiOvaCustomSpec"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiOvaCustomSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOvaEnvSpec returns the OvaEnvSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetOvaEnvSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OvaEnvSpec
}

// GetOvaEnvSpecOk returns a tuple with the OvaEnvSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetOvaEnvSpecOk() (*interface{}, bool) {
	if o == nil || o.OvaEnvSpec == nil {
		return nil, false
	}
	return &o.OvaEnvSpec, true
}

// HasOvaEnvSpec returns a boolean if a field has been set.
func (o *VirtualizationEsxiOvaCustomSpecAllOf) HasOvaEnvSpec() bool {
	if o != nil && o.OvaEnvSpec != nil {
		return true
	}

	return false
}

// SetOvaEnvSpec gets a reference to the given interface{} and assigns it to the OvaEnvSpec field.
func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetOvaEnvSpec(v interface{}) {
	o.OvaEnvSpec = v
}

func (o VirtualizationEsxiOvaCustomSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OvaEnvSpec != nil {
		toSerialize["OvaEnvSpec"] = o.OvaEnvSpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiOvaCustomSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationEsxiOvaCustomSpecAllOf := _VirtualizationEsxiOvaCustomSpecAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationEsxiOvaCustomSpecAllOf); err == nil {
		*o = VirtualizationEsxiOvaCustomSpecAllOf(varVirtualizationEsxiOvaCustomSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OvaEnvSpec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationEsxiOvaCustomSpecAllOf struct {
	value *VirtualizationEsxiOvaCustomSpecAllOf
	isSet bool
}

func (v NullableVirtualizationEsxiOvaCustomSpecAllOf) Get() *VirtualizationEsxiOvaCustomSpecAllOf {
	return v.value
}

func (v *NullableVirtualizationEsxiOvaCustomSpecAllOf) Set(val *VirtualizationEsxiOvaCustomSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiOvaCustomSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiOvaCustomSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiOvaCustomSpecAllOf(val *VirtualizationEsxiOvaCustomSpecAllOf) *NullableVirtualizationEsxiOvaCustomSpecAllOf {
	return &NullableVirtualizationEsxiOvaCustomSpecAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiOvaCustomSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiOvaCustomSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
