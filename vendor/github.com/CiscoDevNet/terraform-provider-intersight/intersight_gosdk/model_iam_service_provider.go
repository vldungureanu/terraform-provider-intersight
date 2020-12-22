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

// IamServiceProvider SAML Service provider related information in Intersight.
type IamServiceProvider struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
	EntityId *string `json:"EntityId,omitempty"`
	// Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication.
	Metadata *string `json:"Metadata,omitempty"`
	// Name of the Intersight Service Provider.
	Name                 *string                `json:"Name,omitempty"`
	System               *IamSystemRelationship `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamServiceProvider IamServiceProvider

// NewIamServiceProvider instantiates a new IamServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamServiceProvider(classId string, objectType string) *IamServiceProvider {
	this := IamServiceProvider{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamServiceProviderWithDefaults instantiates a new IamServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamServiceProviderWithDefaults() *IamServiceProvider {
	this := IamServiceProvider{}
	var classId string = "iam.ServiceProvider"
	this.ClassId = classId
	var objectType string = "iam.ServiceProvider"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamServiceProvider) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamServiceProvider) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamServiceProvider) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamServiceProvider) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *IamServiceProvider) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *IamServiceProvider) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *IamServiceProvider) SetEntityId(v string) {
	o.EntityId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamServiceProvider) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamServiceProvider) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *IamServiceProvider) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamServiceProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamServiceProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamServiceProvider) SetName(v string) {
	o.Name = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamServiceProvider) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceProvider) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamServiceProvider) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamServiceProvider) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamServiceProvider) MarshalJSON() ([]byte, error) {
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
	if o.EntityId != nil {
		toSerialize["EntityId"] = o.EntityId
	}
	if o.Metadata != nil {
		toSerialize["Metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamServiceProvider) UnmarshalJSON(bytes []byte) (err error) {
	type IamServiceProviderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
		EntityId *string `json:"EntityId,omitempty"`
		// Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication.
		Metadata *string `json:"Metadata,omitempty"`
		// Name of the Intersight Service Provider.
		Name   *string                `json:"Name,omitempty"`
		System *IamSystemRelationship `json:"System,omitempty"`
	}

	varIamServiceProviderWithoutEmbeddedStruct := IamServiceProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamServiceProviderWithoutEmbeddedStruct)
	if err == nil {
		varIamServiceProvider := _IamServiceProvider{}
		varIamServiceProvider.ClassId = varIamServiceProviderWithoutEmbeddedStruct.ClassId
		varIamServiceProvider.ObjectType = varIamServiceProviderWithoutEmbeddedStruct.ObjectType
		varIamServiceProvider.EntityId = varIamServiceProviderWithoutEmbeddedStruct.EntityId
		varIamServiceProvider.Metadata = varIamServiceProviderWithoutEmbeddedStruct.Metadata
		varIamServiceProvider.Name = varIamServiceProviderWithoutEmbeddedStruct.Name
		varIamServiceProvider.System = varIamServiceProviderWithoutEmbeddedStruct.System
		*o = IamServiceProvider(varIamServiceProvider)
	} else {
		return err
	}

	varIamServiceProvider := _IamServiceProvider{}

	err = json.Unmarshal(bytes, &varIamServiceProvider)
	if err == nil {
		o.MoBaseMo = varIamServiceProvider.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EntityId")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "System")

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

type NullableIamServiceProvider struct {
	value *IamServiceProvider
	isSet bool
}

func (v NullableIamServiceProvider) Get() *IamServiceProvider {
	return v.value
}

func (v *NullableIamServiceProvider) Set(val *IamServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIamServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIamServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamServiceProvider(val *IamServiceProvider) *NullableIamServiceProvider {
	return &NullableIamServiceProvider{value: val, isSet: true}
}

func (v NullableIamServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
