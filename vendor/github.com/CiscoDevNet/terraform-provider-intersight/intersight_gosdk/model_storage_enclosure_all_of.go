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

// StorageEnclosureAllOf Definition of the list of properties defined in 'storage.Enclosure', excluding properties defined in parent classes.
type StorageEnclosureAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This represent the chassis-ID that houses the storage enclosure.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// This represnets the description for the storage enclosure.
	Description *string `json:"Description,omitempty"`
	// This represnets the Identifier for the storage enclosure.
	EnclosureId *int64 `json:"EnclosureId,omitempty"`
	// This represent the number of slots present in storage enclosure.
	NumSlots *int64 `json:"NumSlots,omitempty"`
	// This represent the availability of storage enclosure.
	Presence *string `json:"Presence,omitempty"`
	// This represent the server-ID that houses the storage enclosure.
	ServerId *int64 `json:"ServerId,omitempty"`
	// This represent the type of storage enclosure.
	Type            *string                      `json:"Type,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to storageEnclosureDiskSlotEp resources.
	EnclosureDiskSlots []StorageEnclosureDiskSlotEpRelationship `json:"EnclosureDiskSlots,omitempty"`
	// An array of relationships to storageEnclosureDisk resources.
	EnclosureDisks      []StorageEnclosureDiskRelationship `json:"EnclosureDisks,omitempty"`
	EquipmentChassis    *EquipmentChassisRelationship      `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship   `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to storagePhysicalDisk resources.
	PhysicalDisks        []StoragePhysicalDiskRelationship    `json:"PhysicalDisks,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageEnclosureAllOf StorageEnclosureAllOf

// NewStorageEnclosureAllOf instantiates a new StorageEnclosureAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEnclosureAllOf(classId string, objectType string) *StorageEnclosureAllOf {
	this := StorageEnclosureAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageEnclosureAllOfWithDefaults instantiates a new StorageEnclosureAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEnclosureAllOfWithDefaults() *StorageEnclosureAllOf {
	this := StorageEnclosureAllOf{}
	var classId string = "storage.Enclosure"
	this.ClassId = classId
	var objectType string = "storage.Enclosure"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageEnclosureAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageEnclosureAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageEnclosureAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageEnclosureAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *StorageEnclosureAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageEnclosureAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEnclosureId returns the EnclosureId field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetEnclosureId() int64 {
	if o == nil || o.EnclosureId == nil {
		var ret int64
		return ret
	}
	return *o.EnclosureId
}

// GetEnclosureIdOk returns a tuple with the EnclosureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetEnclosureIdOk() (*int64, bool) {
	if o == nil || o.EnclosureId == nil {
		return nil, false
	}
	return o.EnclosureId, true
}

// HasEnclosureId returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasEnclosureId() bool {
	if o != nil && o.EnclosureId != nil {
		return true
	}

	return false
}

// SetEnclosureId gets a reference to the given int64 and assigns it to the EnclosureId field.
func (o *StorageEnclosureAllOf) SetEnclosureId(v int64) {
	o.EnclosureId = &v
}

// GetNumSlots returns the NumSlots field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetNumSlots() int64 {
	if o == nil || o.NumSlots == nil {
		var ret int64
		return ret
	}
	return *o.NumSlots
}

// GetNumSlotsOk returns a tuple with the NumSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetNumSlotsOk() (*int64, bool) {
	if o == nil || o.NumSlots == nil {
		return nil, false
	}
	return o.NumSlots, true
}

// HasNumSlots returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasNumSlots() bool {
	if o != nil && o.NumSlots != nil {
		return true
	}

	return false
}

// SetNumSlots gets a reference to the given int64 and assigns it to the NumSlots field.
func (o *StorageEnclosureAllOf) SetNumSlots(v int64) {
	o.NumSlots = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *StorageEnclosureAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetServerId() int64 {
	if o == nil || o.ServerId == nil {
		var ret int64
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetServerIdOk() (*int64, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int64 and assigns it to the ServerId field.
func (o *StorageEnclosureAllOf) SetServerId(v int64) {
	o.ServerId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageEnclosureAllOf) SetType(v string) {
	o.Type = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *StorageEnclosureAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *StorageEnclosureAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEnclosureDiskSlots returns the EnclosureDiskSlots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageEnclosureAllOf) GetEnclosureDiskSlots() []StorageEnclosureDiskSlotEpRelationship {
	if o == nil {
		var ret []StorageEnclosureDiskSlotEpRelationship
		return ret
	}
	return o.EnclosureDiskSlots
}

// GetEnclosureDiskSlotsOk returns a tuple with the EnclosureDiskSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageEnclosureAllOf) GetEnclosureDiskSlotsOk() (*[]StorageEnclosureDiskSlotEpRelationship, bool) {
	if o == nil || o.EnclosureDiskSlots == nil {
		return nil, false
	}
	return &o.EnclosureDiskSlots, true
}

// HasEnclosureDiskSlots returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasEnclosureDiskSlots() bool {
	if o != nil && o.EnclosureDiskSlots != nil {
		return true
	}

	return false
}

// SetEnclosureDiskSlots gets a reference to the given []StorageEnclosureDiskSlotEpRelationship and assigns it to the EnclosureDiskSlots field.
func (o *StorageEnclosureAllOf) SetEnclosureDiskSlots(v []StorageEnclosureDiskSlotEpRelationship) {
	o.EnclosureDiskSlots = v
}

// GetEnclosureDisks returns the EnclosureDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageEnclosureAllOf) GetEnclosureDisks() []StorageEnclosureDiskRelationship {
	if o == nil {
		var ret []StorageEnclosureDiskRelationship
		return ret
	}
	return o.EnclosureDisks
}

// GetEnclosureDisksOk returns a tuple with the EnclosureDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageEnclosureAllOf) GetEnclosureDisksOk() (*[]StorageEnclosureDiskRelationship, bool) {
	if o == nil || o.EnclosureDisks == nil {
		return nil, false
	}
	return &o.EnclosureDisks, true
}

// HasEnclosureDisks returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasEnclosureDisks() bool {
	if o != nil && o.EnclosureDisks != nil {
		return true
	}

	return false
}

// SetEnclosureDisks gets a reference to the given []StorageEnclosureDiskRelationship and assigns it to the EnclosureDisks field.
func (o *StorageEnclosureAllOf) SetEnclosureDisks(v []StorageEnclosureDiskRelationship) {
	o.EnclosureDisks = v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *StorageEnclosureAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageEnclosureAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDisks returns the PhysicalDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageEnclosureAllOf) GetPhysicalDisks() []StoragePhysicalDiskRelationship {
	if o == nil {
		var ret []StoragePhysicalDiskRelationship
		return ret
	}
	return o.PhysicalDisks
}

// GetPhysicalDisksOk returns a tuple with the PhysicalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageEnclosureAllOf) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisks == nil {
		return nil, false
	}
	return &o.PhysicalDisks, true
}

// HasPhysicalDisks returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasPhysicalDisks() bool {
	if o != nil && o.PhysicalDisks != nil {
		return true
	}

	return false
}

// SetPhysicalDisks gets a reference to the given []StoragePhysicalDiskRelationship and assigns it to the PhysicalDisks field.
func (o *StorageEnclosureAllOf) SetPhysicalDisks(v []StoragePhysicalDiskRelationship) {
	o.PhysicalDisks = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageEnclosureAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageEnclosureAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageEnclosureAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageEnclosureAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EnclosureId != nil {
		toSerialize["EnclosureId"] = o.EnclosureId
	}
	if o.NumSlots != nil {
		toSerialize["NumSlots"] = o.NumSlots
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.ServerId != nil {
		toSerialize["ServerId"] = o.ServerId
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EnclosureDiskSlots != nil {
		toSerialize["EnclosureDiskSlots"] = o.EnclosureDiskSlots
	}
	if o.EnclosureDisks != nil {
		toSerialize["EnclosureDisks"] = o.EnclosureDisks
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDisks != nil {
		toSerialize["PhysicalDisks"] = o.PhysicalDisks
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageEnclosureAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageEnclosureAllOf := _StorageEnclosureAllOf{}

	if err = json.Unmarshal(bytes, &varStorageEnclosureAllOf); err == nil {
		*o = StorageEnclosureAllOf(varStorageEnclosureAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EnclosureId")
		delete(additionalProperties, "NumSlots")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "ServerId")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EnclosureDiskSlots")
		delete(additionalProperties, "EnclosureDisks")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisks")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageEnclosureAllOf struct {
	value *StorageEnclosureAllOf
	isSet bool
}

func (v NullableStorageEnclosureAllOf) Get() *StorageEnclosureAllOf {
	return v.value
}

func (v *NullableStorageEnclosureAllOf) Set(val *StorageEnclosureAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEnclosureAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEnclosureAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEnclosureAllOf(val *StorageEnclosureAllOf) *NullableStorageEnclosureAllOf {
	return &NullableStorageEnclosureAllOf{value: val, isSet: true}
}

func (v NullableStorageEnclosureAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEnclosureAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
