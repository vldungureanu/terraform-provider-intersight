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

// FabricUplinkRole Configuration object sent by user to create a uplink port.
type FabricUplinkRole struct {
	FabricTransceiverRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured state for UDLD for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	UdldAdminState       *string `json:"UdldAdminState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricUplinkRole FabricUplinkRole

// NewFabricUplinkRole instantiates a new FabricUplinkRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUplinkRole(classId string, objectType string) *FabricUplinkRole {
	this := FabricUplinkRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// NewFabricUplinkRoleWithDefaults instantiates a new FabricUplinkRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUplinkRoleWithDefaults() *FabricUplinkRole {
	this := FabricUplinkRole{}
	var classId string = "fabric.UplinkRole"
	this.ClassId = classId
	var objectType string = "fabric.UplinkRole"
	this.ObjectType = objectType
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUplinkRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUplinkRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricUplinkRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUplinkRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUdldAdminState returns the UdldAdminState field value if set, zero value otherwise.
func (o *FabricUplinkRole) GetUdldAdminState() string {
	if o == nil || o.UdldAdminState == nil {
		var ret string
		return ret
	}
	return *o.UdldAdminState
}

// GetUdldAdminStateOk returns a tuple with the UdldAdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkRole) GetUdldAdminStateOk() (*string, bool) {
	if o == nil || o.UdldAdminState == nil {
		return nil, false
	}
	return o.UdldAdminState, true
}

// HasUdldAdminState returns a boolean if a field has been set.
func (o *FabricUplinkRole) HasUdldAdminState() bool {
	if o != nil && o.UdldAdminState != nil {
		return true
	}

	return false
}

// SetUdldAdminState gets a reference to the given string and assigns it to the UdldAdminState field.
func (o *FabricUplinkRole) SetUdldAdminState(v string) {
	o.UdldAdminState = &v
}

func (o FabricUplinkRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricTransceiverRole, errFabricTransceiverRole := json.Marshal(o.FabricTransceiverRole)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	errFabricTransceiverRole = json.Unmarshal([]byte(serializedFabricTransceiverRole), &toSerialize)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UdldAdminState != nil {
		toSerialize["UdldAdminState"] = o.UdldAdminState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricUplinkRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricUplinkRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured state for UDLD for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		UdldAdminState *string `json:"UdldAdminState,omitempty"`
	}

	varFabricUplinkRoleWithoutEmbeddedStruct := FabricUplinkRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricUplinkRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricUplinkRole := _FabricUplinkRole{}
		varFabricUplinkRole.ClassId = varFabricUplinkRoleWithoutEmbeddedStruct.ClassId
		varFabricUplinkRole.ObjectType = varFabricUplinkRoleWithoutEmbeddedStruct.ObjectType
		varFabricUplinkRole.UdldAdminState = varFabricUplinkRoleWithoutEmbeddedStruct.UdldAdminState
		*o = FabricUplinkRole(varFabricUplinkRole)
	} else {
		return err
	}

	varFabricUplinkRole := _FabricUplinkRole{}

	err = json.Unmarshal(bytes, &varFabricUplinkRole)
	if err == nil {
		o.FabricTransceiverRole = varFabricUplinkRole.FabricTransceiverRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UdldAdminState")

		// remove fields from embedded structs
		reflectFabricTransceiverRole := reflect.ValueOf(o.FabricTransceiverRole)
		for i := 0; i < reflectFabricTransceiverRole.Type().NumField(); i++ {
			t := reflectFabricTransceiverRole.Type().Field(i)

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

type NullableFabricUplinkRole struct {
	value *FabricUplinkRole
	isSet bool
}

func (v NullableFabricUplinkRole) Get() *FabricUplinkRole {
	return v.value
}

func (v *NullableFabricUplinkRole) Set(val *FabricUplinkRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUplinkRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUplinkRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUplinkRole(val *FabricUplinkRole) *NullableFabricUplinkRole {
	return &NullableFabricUplinkRole{value: val, isSet: true}
}

func (v NullableFabricUplinkRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUplinkRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
