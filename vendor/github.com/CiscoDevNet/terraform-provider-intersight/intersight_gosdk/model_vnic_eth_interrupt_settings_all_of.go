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

// VnicEthInterruptSettingsAllOf Definition of the list of properties defined in 'vnic.EthInterruptSettings', excluding properties defined in parent classes.
type VnicEthInterruptSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.
	CoalescingTime *int64 `json:"CoalescingTime,omitempty"`
	// Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field. * `MIN` - The system waits for the time specified in the Coalescing Time field before sending another interrupt event. * `IDLE` - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.
	CoalescingType *string `json:"CoalescingType,omitempty"`
	// The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.
	Count *int64 `json:"Count,omitempty"`
	// Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthInterruptSettingsAllOf VnicEthInterruptSettingsAllOf

// NewVnicEthInterruptSettingsAllOf instantiates a new VnicEthInterruptSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthInterruptSettingsAllOf(classId string, objectType string) *VnicEthInterruptSettingsAllOf {
	this := VnicEthInterruptSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var coalescingTime int64 = 125
	this.CoalescingTime = &coalescingTime
	var coalescingType string = "MIN"
	this.CoalescingType = &coalescingType
	var count int64 = 8
	this.Count = &count
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// NewVnicEthInterruptSettingsAllOfWithDefaults instantiates a new VnicEthInterruptSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthInterruptSettingsAllOfWithDefaults() *VnicEthInterruptSettingsAllOf {
	this := VnicEthInterruptSettingsAllOf{}
	var classId string = "vnic.EthInterruptSettings"
	this.ClassId = classId
	var objectType string = "vnic.EthInterruptSettings"
	this.ObjectType = objectType
	var coalescingTime int64 = 125
	this.CoalescingTime = &coalescingTime
	var coalescingType string = "MIN"
	this.CoalescingType = &coalescingType
	var count int64 = 8
	this.Count = &count
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthInterruptSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthInterruptSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthInterruptSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthInterruptSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCoalescingTime returns the CoalescingTime field value if set, zero value otherwise.
func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTime() int64 {
	if o == nil || o.CoalescingTime == nil {
		var ret int64
		return ret
	}
	return *o.CoalescingTime
}

// GetCoalescingTimeOk returns a tuple with the CoalescingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTimeOk() (*int64, bool) {
	if o == nil || o.CoalescingTime == nil {
		return nil, false
	}
	return o.CoalescingTime, true
}

// HasCoalescingTime returns a boolean if a field has been set.
func (o *VnicEthInterruptSettingsAllOf) HasCoalescingTime() bool {
	if o != nil && o.CoalescingTime != nil {
		return true
	}

	return false
}

// SetCoalescingTime gets a reference to the given int64 and assigns it to the CoalescingTime field.
func (o *VnicEthInterruptSettingsAllOf) SetCoalescingTime(v int64) {
	o.CoalescingTime = &v
}

// GetCoalescingType returns the CoalescingType field value if set, zero value otherwise.
func (o *VnicEthInterruptSettingsAllOf) GetCoalescingType() string {
	if o == nil || o.CoalescingType == nil {
		var ret string
		return ret
	}
	return *o.CoalescingType
}

// GetCoalescingTypeOk returns a tuple with the CoalescingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTypeOk() (*string, bool) {
	if o == nil || o.CoalescingType == nil {
		return nil, false
	}
	return o.CoalescingType, true
}

// HasCoalescingType returns a boolean if a field has been set.
func (o *VnicEthInterruptSettingsAllOf) HasCoalescingType() bool {
	if o != nil && o.CoalescingType != nil {
		return true
	}

	return false
}

// SetCoalescingType gets a reference to the given string and assigns it to the CoalescingType field.
func (o *VnicEthInterruptSettingsAllOf) SetCoalescingType(v string) {
	o.CoalescingType = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthInterruptSettingsAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthInterruptSettingsAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicEthInterruptSettingsAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicEthInterruptSettingsAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettingsAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicEthInterruptSettingsAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicEthInterruptSettingsAllOf) SetMode(v string) {
	o.Mode = &v
}

func (o VnicEthInterruptSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CoalescingTime != nil {
		toSerialize["CoalescingTime"] = o.CoalescingTime
	}
	if o.CoalescingType != nil {
		toSerialize["CoalescingType"] = o.CoalescingType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthInterruptSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthInterruptSettingsAllOf := _VnicEthInterruptSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthInterruptSettingsAllOf); err == nil {
		*o = VnicEthInterruptSettingsAllOf(varVnicEthInterruptSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CoalescingTime")
		delete(additionalProperties, "CoalescingType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthInterruptSettingsAllOf struct {
	value *VnicEthInterruptSettingsAllOf
	isSet bool
}

func (v NullableVnicEthInterruptSettingsAllOf) Get() *VnicEthInterruptSettingsAllOf {
	return v.value
}

func (v *NullableVnicEthInterruptSettingsAllOf) Set(val *VnicEthInterruptSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthInterruptSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthInterruptSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthInterruptSettingsAllOf(val *VnicEthInterruptSettingsAllOf) *NullableVnicEthInterruptSettingsAllOf {
	return &NullableVnicEthInterruptSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicEthInterruptSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthInterruptSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
