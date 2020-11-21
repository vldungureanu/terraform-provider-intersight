/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowCustomArrayItemAllOf Definition of the list of properties defined in 'workflow.CustomArrayItem', excluding properties defined in parent classes.
type WorkflowCustomArrayItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Properties           NullableWorkflowCustomDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomArrayItemAllOf WorkflowCustomArrayItemAllOf

// NewWorkflowCustomArrayItemAllOf instantiates a new WorkflowCustomArrayItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomArrayItemAllOf(classId string, objectType string) *WorkflowCustomArrayItemAllOf {
	this := WorkflowCustomArrayItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCustomArrayItemAllOfWithDefaults instantiates a new WorkflowCustomArrayItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomArrayItemAllOfWithDefaults() *WorkflowCustomArrayItemAllOf {
	this := WorkflowCustomArrayItemAllOf{}
	var classId string = "workflow.CustomArrayItem"
	this.ClassId = classId
	var objectType string = "workflow.CustomArrayItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomArrayItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomArrayItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomArrayItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomArrayItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomArrayItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomArrayItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomArrayItemAllOf) GetProperties() WorkflowCustomDataProperty {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowCustomDataProperty
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomArrayItemAllOf) GetPropertiesOk() (*WorkflowCustomDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowCustomArrayItemAllOf) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowCustomDataProperty and assigns it to the Properties field.
func (o *WorkflowCustomArrayItemAllOf) SetProperties(v WorkflowCustomDataProperty) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowCustomArrayItemAllOf) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowCustomArrayItemAllOf) UnsetProperties() {
	o.Properties.Unset()
}

func (o WorkflowCustomArrayItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Properties.IsSet() {
		toSerialize["Properties"] = o.Properties.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCustomArrayItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCustomArrayItemAllOf := _WorkflowCustomArrayItemAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowCustomArrayItemAllOf); err == nil {
		*o = WorkflowCustomArrayItemAllOf(varWorkflowCustomArrayItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCustomArrayItemAllOf struct {
	value *WorkflowCustomArrayItemAllOf
	isSet bool
}

func (v NullableWorkflowCustomArrayItemAllOf) Get() *WorkflowCustomArrayItemAllOf {
	return v.value
}

func (v *NullableWorkflowCustomArrayItemAllOf) Set(val *WorkflowCustomArrayItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomArrayItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomArrayItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomArrayItemAllOf(val *WorkflowCustomArrayItemAllOf) *NullableWorkflowCustomArrayItemAllOf {
	return &NullableWorkflowCustomArrayItemAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCustomArrayItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomArrayItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
