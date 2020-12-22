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

// StorageBaseDiskPool A disk pool is a set of drives that is logically grouped together in the storage array. Some storage vendors use pools extensively.
type StorageBaseDiskPool struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Human readable name of the pool, limited to 64 characters.
	Name *string `json:"Name,omitempty"`
	// Object ID for the pool. Platforms that use a number should convert it to string.
	PoolId *string `json:"PoolId,omitempty"`
	// Human readable status of the pool, indicating the current health. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
	Status             *string                     `json:"Status,omitempty"`
	StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	// Human readable type of the pool, such as thin, tiered, active-flash, etc.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseDiskPool StorageBaseDiskPool

// NewStorageBaseDiskPool instantiates a new StorageBaseDiskPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseDiskPool(classId string, objectType string) *StorageBaseDiskPool {
	this := StorageBaseDiskPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewStorageBaseDiskPoolWithDefaults instantiates a new StorageBaseDiskPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseDiskPoolWithDefaults() *StorageBaseDiskPool {
	this := StorageBaseDiskPool{}
	var classId string = "storage.HitachiPool"
	this.ClassId = classId
	var objectType string = "storage.HitachiPool"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseDiskPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseDiskPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseDiskPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseDiskPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseDiskPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseDiskPool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseDiskPool) SetName(v string) {
	o.Name = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *StorageBaseDiskPool) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *StorageBaseDiskPool) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *StorageBaseDiskPool) SetPoolId(v string) {
	o.PoolId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageBaseDiskPool) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageBaseDiskPool) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageBaseDiskPool) SetStatus(v string) {
	o.Status = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseDiskPool) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseDiskPool) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseDiskPool) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseDiskPool) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseDiskPool) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseDiskPool) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageBaseDiskPool) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseDiskPool) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageBaseDiskPool) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageBaseDiskPool) SetType(v string) {
	o.Type = &v
}

func (o StorageBaseDiskPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PoolId != nil {
		toSerialize["PoolId"] = o.PoolId
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseDiskPool) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseDiskPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Human readable name of the pool, limited to 64 characters.
		Name *string `json:"Name,omitempty"`
		// Object ID for the pool. Platforms that use a number should convert it to string.
		PoolId *string `json:"PoolId,omitempty"`
		// Human readable status of the pool, indicating the current health. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
		Status             *string                     `json:"Status,omitempty"`
		StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
		// Human readable type of the pool, such as thin, tiered, active-flash, etc.
		Type *string `json:"Type,omitempty"`
	}

	varStorageBaseDiskPoolWithoutEmbeddedStruct := StorageBaseDiskPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseDiskPoolWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseDiskPool := _StorageBaseDiskPool{}
		varStorageBaseDiskPool.ClassId = varStorageBaseDiskPoolWithoutEmbeddedStruct.ClassId
		varStorageBaseDiskPool.ObjectType = varStorageBaseDiskPoolWithoutEmbeddedStruct.ObjectType
		varStorageBaseDiskPool.Name = varStorageBaseDiskPoolWithoutEmbeddedStruct.Name
		varStorageBaseDiskPool.PoolId = varStorageBaseDiskPoolWithoutEmbeddedStruct.PoolId
		varStorageBaseDiskPool.Status = varStorageBaseDiskPoolWithoutEmbeddedStruct.Status
		varStorageBaseDiskPool.StorageUtilization = varStorageBaseDiskPoolWithoutEmbeddedStruct.StorageUtilization
		varStorageBaseDiskPool.Type = varStorageBaseDiskPoolWithoutEmbeddedStruct.Type
		*o = StorageBaseDiskPool(varStorageBaseDiskPool)
	} else {
		return err
	}

	varStorageBaseDiskPool := _StorageBaseDiskPool{}

	err = json.Unmarshal(bytes, &varStorageBaseDiskPool)
	if err == nil {
		o.MoBaseMo = varStorageBaseDiskPool.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PoolId")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableStorageBaseDiskPool struct {
	value *StorageBaseDiskPool
	isSet bool
}

func (v NullableStorageBaseDiskPool) Get() *StorageBaseDiskPool {
	return v.value
}

func (v *NullableStorageBaseDiskPool) Set(val *StorageBaseDiskPool) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseDiskPool) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseDiskPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseDiskPool(val *StorageBaseDiskPool) *NullableStorageBaseDiskPool {
	return &NullableStorageBaseDiskPool{value: val, isSet: true}
}

func (v NullableStorageBaseDiskPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseDiskPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
