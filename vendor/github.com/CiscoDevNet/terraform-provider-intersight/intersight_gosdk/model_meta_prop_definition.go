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

// MetaPropDefinition Definitions for the properties in a meta.
type MetaPropDefinition struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. * `NoAccess` - The property is not accessible from the API. * `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * `ReadWrite` - The property has read/write access. * `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.
	ApiAccess *string `json:"ApiAccess,omitempty"`
	// The name of the property.
	Name *string `json:"Name,omitempty"`
	// The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. * `ClearText` - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable. * `Encrypted` - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable. * `Pbkdf2` - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive. * `Bcrypt` - The value of the property is hashed using the bcrypt key derivation function. * `Sha512crypt` - The value of the property is hashed using the sha512crypt key derivation function. * `Argon2id` - The value of the property is hashed using the argon2id key derivation function.
	OpSecurity *string `json:"OpSecurity,omitempty"`
	// Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.
	SearchWeight         *float32 `json:"SearchWeight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaPropDefinition MetaPropDefinition

// NewMetaPropDefinition instantiates a new MetaPropDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaPropDefinition(classId string, objectType string) *MetaPropDefinition {
	this := MetaPropDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var apiAccess string = "NoAccess"
	this.ApiAccess = &apiAccess
	var opSecurity string = "ClearText"
	this.OpSecurity = &opSecurity
	return &this
}

// NewMetaPropDefinitionWithDefaults instantiates a new MetaPropDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaPropDefinitionWithDefaults() *MetaPropDefinition {
	this := MetaPropDefinition{}
	var classId string = "meta.PropDefinition"
	this.ClassId = classId
	var objectType string = "meta.PropDefinition"
	this.ObjectType = objectType
	var apiAccess string = "NoAccess"
	this.ApiAccess = &apiAccess
	var opSecurity string = "ClearText"
	this.OpSecurity = &opSecurity
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaPropDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaPropDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaPropDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaPropDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiAccess returns the ApiAccess field value if set, zero value otherwise.
func (o *MetaPropDefinition) GetApiAccess() string {
	if o == nil || o.ApiAccess == nil {
		var ret string
		return ret
	}
	return *o.ApiAccess
}

// GetApiAccessOk returns a tuple with the ApiAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetApiAccessOk() (*string, bool) {
	if o == nil || o.ApiAccess == nil {
		return nil, false
	}
	return o.ApiAccess, true
}

// HasApiAccess returns a boolean if a field has been set.
func (o *MetaPropDefinition) HasApiAccess() bool {
	if o != nil && o.ApiAccess != nil {
		return true
	}

	return false
}

// SetApiAccess gets a reference to the given string and assigns it to the ApiAccess field.
func (o *MetaPropDefinition) SetApiAccess(v string) {
	o.ApiAccess = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaPropDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaPropDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaPropDefinition) SetName(v string) {
	o.Name = &v
}

// GetOpSecurity returns the OpSecurity field value if set, zero value otherwise.
func (o *MetaPropDefinition) GetOpSecurity() string {
	if o == nil || o.OpSecurity == nil {
		var ret string
		return ret
	}
	return *o.OpSecurity
}

// GetOpSecurityOk returns a tuple with the OpSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetOpSecurityOk() (*string, bool) {
	if o == nil || o.OpSecurity == nil {
		return nil, false
	}
	return o.OpSecurity, true
}

// HasOpSecurity returns a boolean if a field has been set.
func (o *MetaPropDefinition) HasOpSecurity() bool {
	if o != nil && o.OpSecurity != nil {
		return true
	}

	return false
}

// SetOpSecurity gets a reference to the given string and assigns it to the OpSecurity field.
func (o *MetaPropDefinition) SetOpSecurity(v string) {
	o.OpSecurity = &v
}

// GetSearchWeight returns the SearchWeight field value if set, zero value otherwise.
func (o *MetaPropDefinition) GetSearchWeight() float32 {
	if o == nil || o.SearchWeight == nil {
		var ret float32
		return ret
	}
	return *o.SearchWeight
}

// GetSearchWeightOk returns a tuple with the SearchWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaPropDefinition) GetSearchWeightOk() (*float32, bool) {
	if o == nil || o.SearchWeight == nil {
		return nil, false
	}
	return o.SearchWeight, true
}

// HasSearchWeight returns a boolean if a field has been set.
func (o *MetaPropDefinition) HasSearchWeight() bool {
	if o != nil && o.SearchWeight != nil {
		return true
	}

	return false
}

// SetSearchWeight gets a reference to the given float32 and assigns it to the SearchWeight field.
func (o *MetaPropDefinition) SetSearchWeight(v float32) {
	o.SearchWeight = &v
}

func (o MetaPropDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApiAccess != nil {
		toSerialize["ApiAccess"] = o.ApiAccess
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OpSecurity != nil {
		toSerialize["OpSecurity"] = o.OpSecurity
	}
	if o.SearchWeight != nil {
		toSerialize["SearchWeight"] = o.SearchWeight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaPropDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type MetaPropDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc. * `NoAccess` - The property is not accessible from the API. * `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * `ReadWrite` - The property has read/write access. * `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.
		ApiAccess *string `json:"ApiAccess,omitempty"`
		// The name of the property.
		Name *string `json:"Name,omitempty"`
		// The data-at-rest security protection applied to this property when a Managed Object is persisted. For example, Encrypted or Cleartext. * `ClearText` - Data at rest is not encrypted using account specific keys.Note that data is always protected using volume encryption. ClearText propertiesare queryable and searchable. * `Encrypted` - The value of the property is encrypted with account-specific cryptographic keys.This setting is used for properties that contain sensitive data. Encrypted propertiesare not queryable and searchable. * `Pbkdf2` - The value of the property is hashed using the pbkdf2 key derivation functions thattakes a password, a salt, and a cost factor as inputs then generates a password hash.Its purpose is to make each password guessing trial by an attacker who has obtaineda password hash file expensive and therefore the cost of a guessing attack high or prohibitive. * `Bcrypt` - The value of the property is hashed using the bcrypt key derivation function. * `Sha512crypt` - The value of the property is hashed using the sha512crypt key derivation function. * `Argon2id` - The value of the property is hashed using the argon2id key derivation function.
		OpSecurity *string `json:"OpSecurity,omitempty"`
		// Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.
		SearchWeight *float32 `json:"SearchWeight,omitempty"`
	}

	varMetaPropDefinitionWithoutEmbeddedStruct := MetaPropDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMetaPropDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varMetaPropDefinition := _MetaPropDefinition{}
		varMetaPropDefinition.ClassId = varMetaPropDefinitionWithoutEmbeddedStruct.ClassId
		varMetaPropDefinition.ObjectType = varMetaPropDefinitionWithoutEmbeddedStruct.ObjectType
		varMetaPropDefinition.ApiAccess = varMetaPropDefinitionWithoutEmbeddedStruct.ApiAccess
		varMetaPropDefinition.Name = varMetaPropDefinitionWithoutEmbeddedStruct.Name
		varMetaPropDefinition.OpSecurity = varMetaPropDefinitionWithoutEmbeddedStruct.OpSecurity
		varMetaPropDefinition.SearchWeight = varMetaPropDefinitionWithoutEmbeddedStruct.SearchWeight
		*o = MetaPropDefinition(varMetaPropDefinition)
	} else {
		return err
	}

	varMetaPropDefinition := _MetaPropDefinition{}

	err = json.Unmarshal(bytes, &varMetaPropDefinition)
	if err == nil {
		o.MoBaseComplexType = varMetaPropDefinition.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiAccess")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OpSecurity")
		delete(additionalProperties, "SearchWeight")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableMetaPropDefinition struct {
	value *MetaPropDefinition
	isSet bool
}

func (v NullableMetaPropDefinition) Get() *MetaPropDefinition {
	return v.value
}

func (v *NullableMetaPropDefinition) Set(val *MetaPropDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaPropDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaPropDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaPropDefinition(val *MetaPropDefinition) *NullableMetaPropDefinition {
	return &NullableMetaPropDefinition{value: val, isSet: true}
}

func (v NullableMetaPropDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaPropDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
