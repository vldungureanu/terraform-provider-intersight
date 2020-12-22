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

// FirmwareUpgradeImpactStatus Captures the impact for an upgrade.
type FirmwareUpgradeImpactStatus struct {
	FirmwareUpgradeImpactBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	Upgrade              *FirmwareUpgradeBaseRelationship `json:"Upgrade,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeImpactStatus FirmwareUpgradeImpactStatus

// NewFirmwareUpgradeImpactStatus instantiates a new FirmwareUpgradeImpactStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeImpactStatus(classId string, objectType string) *FirmwareUpgradeImpactStatus {
	this := FirmwareUpgradeImpactStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeImpactStatusWithDefaults instantiates a new FirmwareUpgradeImpactStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeImpactStatusWithDefaults() *FirmwareUpgradeImpactStatus {
	this := FirmwareUpgradeImpactStatus{}
	var classId string = "firmware.UpgradeImpactStatus"
	this.ClassId = classId
	var objectType string = "firmware.UpgradeImpactStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeImpactStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeImpactStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeImpactStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeImpactStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactStatus) GetUpgrade() FirmwareUpgradeBaseRelationship {
	if o == nil || o.Upgrade == nil {
		var ret FirmwareUpgradeBaseRelationship
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactStatus) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactStatus) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given FirmwareUpgradeBaseRelationship and assigns it to the Upgrade field.
func (o *FirmwareUpgradeImpactStatus) SetUpgrade(v FirmwareUpgradeBaseRelationship) {
	o.Upgrade = &v
}

func (o FirmwareUpgradeImpactStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeImpactBase, errFirmwareUpgradeImpactBase := json.Marshal(o.FirmwareUpgradeImpactBase)
	if errFirmwareUpgradeImpactBase != nil {
		return []byte{}, errFirmwareUpgradeImpactBase
	}
	errFirmwareUpgradeImpactBase = json.Unmarshal([]byte(serializedFirmwareUpgradeImpactBase), &toSerialize)
	if errFirmwareUpgradeImpactBase != nil {
		return []byte{}, errFirmwareUpgradeImpactBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Upgrade != nil {
		toSerialize["Upgrade"] = o.Upgrade
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeImpactStatus) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeImpactStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		Upgrade    *FirmwareUpgradeBaseRelationship `json:"Upgrade,omitempty"`
	}

	varFirmwareUpgradeImpactStatusWithoutEmbeddedStruct := FirmwareUpgradeImpactStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactStatusWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgradeImpactStatus := _FirmwareUpgradeImpactStatus{}
		varFirmwareUpgradeImpactStatus.ClassId = varFirmwareUpgradeImpactStatusWithoutEmbeddedStruct.ClassId
		varFirmwareUpgradeImpactStatus.ObjectType = varFirmwareUpgradeImpactStatusWithoutEmbeddedStruct.ObjectType
		varFirmwareUpgradeImpactStatus.Upgrade = varFirmwareUpgradeImpactStatusWithoutEmbeddedStruct.Upgrade
		*o = FirmwareUpgradeImpactStatus(varFirmwareUpgradeImpactStatus)
	} else {
		return err
	}

	varFirmwareUpgradeImpactStatus := _FirmwareUpgradeImpactStatus{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactStatus)
	if err == nil {
		o.FirmwareUpgradeImpactBase = varFirmwareUpgradeImpactStatus.FirmwareUpgradeImpactBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Upgrade")

		// remove fields from embedded structs
		reflectFirmwareUpgradeImpactBase := reflect.ValueOf(o.FirmwareUpgradeImpactBase)
		for i := 0; i < reflectFirmwareUpgradeImpactBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeImpactBase.Type().Field(i)

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

type NullableFirmwareUpgradeImpactStatus struct {
	value *FirmwareUpgradeImpactStatus
	isSet bool
}

func (v NullableFirmwareUpgradeImpactStatus) Get() *FirmwareUpgradeImpactStatus {
	return v.value
}

func (v *NullableFirmwareUpgradeImpactStatus) Set(val *FirmwareUpgradeImpactStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeImpactStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeImpactStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeImpactStatus(val *FirmwareUpgradeImpactStatus) *NullableFirmwareUpgradeImpactStatus {
	return &NullableFirmwareUpgradeImpactStatus{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeImpactStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeImpactStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
