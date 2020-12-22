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

// TelemetryDruidMinMaxAggregator Computes the min or max of values.
type TelemetryDruidMinMaxAggregator struct {
	// The aggregator type.
	Type string `json:"type"`
	// Output name for the min/max value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidMinMaxAggregator TelemetryDruidMinMaxAggregator

// NewTelemetryDruidMinMaxAggregator instantiates a new TelemetryDruidMinMaxAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidMinMaxAggregator(type_ string, name string, fieldName string) *TelemetryDruidMinMaxAggregator {
	this := TelemetryDruidMinMaxAggregator{}
	this.Type = type_
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidMinMaxAggregatorWithDefaults instantiates a new TelemetryDruidMinMaxAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidMinMaxAggregatorWithDefaults() *TelemetryDruidMinMaxAggregator {
	this := TelemetryDruidMinMaxAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidMinMaxAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidMinMaxAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidMinMaxAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TelemetryDruidMinMaxAggregator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidMinMaxAggregator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidMinMaxAggregator) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidMinMaxAggregator) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidMinMaxAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidMinMaxAggregator) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidMinMaxAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidMinMaxAggregator) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidMinMaxAggregator := _TelemetryDruidMinMaxAggregator{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidMinMaxAggregator); err == nil {
		*o = TelemetryDruidMinMaxAggregator(varTelemetryDruidMinMaxAggregator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidMinMaxAggregator struct {
	value *TelemetryDruidMinMaxAggregator
	isSet bool
}

func (v NullableTelemetryDruidMinMaxAggregator) Get() *TelemetryDruidMinMaxAggregator {
	return v.value
}

func (v *NullableTelemetryDruidMinMaxAggregator) Set(val *TelemetryDruidMinMaxAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidMinMaxAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidMinMaxAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidMinMaxAggregator(val *TelemetryDruidMinMaxAggregator) *NullableTelemetryDruidMinMaxAggregator {
	return &NullableTelemetryDruidMinMaxAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidMinMaxAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidMinMaxAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
