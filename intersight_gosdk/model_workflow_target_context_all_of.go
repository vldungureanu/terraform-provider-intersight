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

// WorkflowTargetContextAllOf Definition of the list of properties defined in 'workflow.TargetContext', excluding properties defined in parent classes.
type WorkflowTargetContextAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid of the target Intersight managed object.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// Name of the target instance.
	TargetName *string `json:"TargetName,omitempty"`
	// Object type of the target Intersight managed object.
	TargetType           *string `json:"TargetType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetContextAllOf WorkflowTargetContextAllOf

// NewWorkflowTargetContextAllOf instantiates a new WorkflowTargetContextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetContextAllOf(classId string, objectType string) *WorkflowTargetContextAllOf {
	this := WorkflowTargetContextAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTargetContextAllOfWithDefaults instantiates a new WorkflowTargetContextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetContextAllOfWithDefaults() *WorkflowTargetContextAllOf {
	this := WorkflowTargetContextAllOf{}
	var classId string = "workflow.TargetContext"
	this.ClassId = classId
	var objectType string = "workflow.TargetContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTargetContextAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContextAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTargetContextAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTargetContextAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContextAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTargetContextAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *WorkflowTargetContextAllOf) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContextAllOf) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *WorkflowTargetContextAllOf) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *WorkflowTargetContextAllOf) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *WorkflowTargetContextAllOf) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContextAllOf) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *WorkflowTargetContextAllOf) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *WorkflowTargetContextAllOf) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowTargetContextAllOf) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContextAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowTargetContextAllOf) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowTargetContextAllOf) SetTargetType(v string) {
	o.TargetType = &v
}

func (o WorkflowTargetContextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetContextAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetContextAllOf := _WorkflowTargetContextAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetContextAllOf); err == nil {
		*o = WorkflowTargetContextAllOf(varWorkflowTargetContextAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "TargetType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTargetContextAllOf struct {
	value *WorkflowTargetContextAllOf
	isSet bool
}

func (v NullableWorkflowTargetContextAllOf) Get() *WorkflowTargetContextAllOf {
	return v.value
}

func (v *NullableWorkflowTargetContextAllOf) Set(val *WorkflowTargetContextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetContextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetContextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetContextAllOf(val *WorkflowTargetContextAllOf) *NullableWorkflowTargetContextAllOf {
	return &NullableWorkflowTargetContextAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTargetContextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetContextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
