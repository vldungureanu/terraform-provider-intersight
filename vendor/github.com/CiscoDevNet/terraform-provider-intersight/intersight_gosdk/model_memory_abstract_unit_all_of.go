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

// MemoryAbstractUnitAllOf Definition of the list of properties defined in 'memory.AbstractUnit', excluding properties defined in parent classes.
type MemoryAbstractUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// This represents the administrative state of the memory unit on a server.
	AdminState *string `json:"AdminState,omitempty"`
	// This represents the memory array to which the memory unit belongs to.
	ArrayId *int64 `json:"ArrayId,omitempty"`
	// This represents the memory bank of the memory unit on a server.
	Bank *int64 `json:"Bank,omitempty"`
	// This represents the memory capacity in MiB of the memory unit on a server.
	Capacity *string `json:"Capacity,omitempty"`
	// This represents the clock of the memory unit on a server.
	Clock *string `json:"Clock,omitempty"`
	// This represents the form factor of the memory unit on a server.
	FormFactor *string `json:"FormFactor,omitempty"`
	// This represents the latency of the memory unit on a server.
	Latency *string `json:"Latency,omitempty"`
	// This represents the location of the memory unit on a server.
	Location *string `json:"Location,omitempty"`
	// This represents the operational power state of the memory unit on a server.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	// This represents the operational state of the memory unit on a server.
	OperState *string `json:"OperState,omitempty"`
	// This represents the operability of the memory unit on a server.
	Operability *string `json:"Operability,omitempty"`
	// This represents the presence state of the memory unit on a server.
	Presence *string `json:"Presence,omitempty"`
	// This represents the set of the memory unit on a server.
	Set *int64 `json:"Set,omitempty"`
	// This represents the speed of the memory unit on a server.
	Speed *string `json:"Speed,omitempty"`
	// This represents the thremal state of the memory unit on a server.
	Thermal *string `json:"Thermal,omitempty"`
	// This represents the memory type of the memory unit on a server.
	Type *string `json:"Type,omitempty"`
	// This represents the visibility of the memory unit on a server.
	Visibility *string `json:"Visibility,omitempty"`
	// This represents the width of the memory unit on a server.
	Width                *string `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryAbstractUnitAllOf MemoryAbstractUnitAllOf

// NewMemoryAbstractUnitAllOf instantiates a new MemoryAbstractUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryAbstractUnitAllOf(classId string, objectType string) *MemoryAbstractUnitAllOf {
	this := MemoryAbstractUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryAbstractUnitAllOfWithDefaults instantiates a new MemoryAbstractUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryAbstractUnitAllOfWithDefaults() *MemoryAbstractUnitAllOf {
	this := MemoryAbstractUnitAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryAbstractUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryAbstractUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryAbstractUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryAbstractUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *MemoryAbstractUnitAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetArrayId() int64 {
	if o == nil || o.ArrayId == nil {
		var ret int64
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetArrayIdOk() (*int64, bool) {
	if o == nil || o.ArrayId == nil {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasArrayId() bool {
	if o != nil && o.ArrayId != nil {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given int64 and assigns it to the ArrayId field.
func (o *MemoryAbstractUnitAllOf) SetArrayId(v int64) {
	o.ArrayId = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetBank() int64 {
	if o == nil || o.Bank == nil {
		var ret int64
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetBankOk() (*int64, bool) {
	if o == nil || o.Bank == nil {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasBank() bool {
	if o != nil && o.Bank != nil {
		return true
	}

	return false
}

// SetBank gets a reference to the given int64 and assigns it to the Bank field.
func (o *MemoryAbstractUnitAllOf) SetBank(v int64) {
	o.Bank = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *MemoryAbstractUnitAllOf) SetCapacity(v string) {
	o.Capacity = &v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetClock() string {
	if o == nil || o.Clock == nil {
		var ret string
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetClockOk() (*string, bool) {
	if o == nil || o.Clock == nil {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasClock() bool {
	if o != nil && o.Clock != nil {
		return true
	}

	return false
}

// SetClock gets a reference to the given string and assigns it to the Clock field.
func (o *MemoryAbstractUnitAllOf) SetClock(v string) {
	o.Clock = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetFormFactor() string {
	if o == nil || o.FormFactor == nil {
		var ret string
		return ret
	}
	return *o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetFormFactorOk() (*string, bool) {
	if o == nil || o.FormFactor == nil {
		return nil, false
	}
	return o.FormFactor, true
}

// HasFormFactor returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasFormFactor() bool {
	if o != nil && o.FormFactor != nil {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given string and assigns it to the FormFactor field.
func (o *MemoryAbstractUnitAllOf) SetFormFactor(v string) {
	o.FormFactor = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetLatency() string {
	if o == nil || o.Latency == nil {
		var ret string
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetLatencyOk() (*string, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasLatency() bool {
	if o != nil && o.Latency != nil {
		return true
	}

	return false
}

// SetLatency gets a reference to the given string and assigns it to the Latency field.
func (o *MemoryAbstractUnitAllOf) SetLatency(v string) {
	o.Latency = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *MemoryAbstractUnitAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *MemoryAbstractUnitAllOf) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *MemoryAbstractUnitAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *MemoryAbstractUnitAllOf) SetOperability(v string) {
	o.Operability = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *MemoryAbstractUnitAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetSet() int64 {
	if o == nil || o.Set == nil {
		var ret int64
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetSetOk() (*int64, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasSet() bool {
	if o != nil && o.Set != nil {
		return true
	}

	return false
}

// SetSet gets a reference to the given int64 and assigns it to the Set field.
func (o *MemoryAbstractUnitAllOf) SetSet(v int64) {
	o.Set = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *MemoryAbstractUnitAllOf) SetSpeed(v string) {
	o.Speed = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *MemoryAbstractUnitAllOf) SetThermal(v string) {
	o.Thermal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MemoryAbstractUnitAllOf) SetType(v string) {
	o.Type = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *MemoryAbstractUnitAllOf) SetVisibility(v string) {
	o.Visibility = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MemoryAbstractUnitAllOf) GetWidth() string {
	if o == nil || o.Width == nil {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnitAllOf) GetWidthOk() (*string, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MemoryAbstractUnitAllOf) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *MemoryAbstractUnitAllOf) SetWidth(v string) {
	o.Width = &v
}

func (o MemoryAbstractUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ArrayId != nil {
		toSerialize["ArrayId"] = o.ArrayId
	}
	if o.Bank != nil {
		toSerialize["Bank"] = o.Bank
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Clock != nil {
		toSerialize["Clock"] = o.Clock
	}
	if o.FormFactor != nil {
		toSerialize["FormFactor"] = o.FormFactor
	}
	if o.Latency != nil {
		toSerialize["Latency"] = o.Latency
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Set != nil {
		toSerialize["Set"] = o.Set
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Visibility != nil {
		toSerialize["Visibility"] = o.Visibility
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryAbstractUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryAbstractUnitAllOf := _MemoryAbstractUnitAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryAbstractUnitAllOf); err == nil {
		*o = MemoryAbstractUnitAllOf(varMemoryAbstractUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ArrayId")
		delete(additionalProperties, "Bank")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Clock")
		delete(additionalProperties, "FormFactor")
		delete(additionalProperties, "Latency")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Set")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Visibility")
		delete(additionalProperties, "Width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryAbstractUnitAllOf struct {
	value *MemoryAbstractUnitAllOf
	isSet bool
}

func (v NullableMemoryAbstractUnitAllOf) Get() *MemoryAbstractUnitAllOf {
	return v.value
}

func (v *NullableMemoryAbstractUnitAllOf) Set(val *MemoryAbstractUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryAbstractUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryAbstractUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryAbstractUnitAllOf(val *MemoryAbstractUnitAllOf) *NullableMemoryAbstractUnitAllOf {
	return &NullableMemoryAbstractUnitAllOf{value: val, isSet: true}
}

func (v NullableMemoryAbstractUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryAbstractUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
