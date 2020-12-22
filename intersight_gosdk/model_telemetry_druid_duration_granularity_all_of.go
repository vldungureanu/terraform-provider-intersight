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
	"time"
)

// TelemetryDruidDurationGranularityAllOf struct for TelemetryDruidDurationGranularityAllOf
type TelemetryDruidDurationGranularityAllOf struct {
	// The duration in milliseconds.
	Duration int64 `json:"duration"`
	// An optional value specifying when to start counting time buckets from. The default value is 1970-01-01T00:00:00Z.
	Origin               *time.Time `json:"origin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidDurationGranularityAllOf TelemetryDruidDurationGranularityAllOf

// NewTelemetryDruidDurationGranularityAllOf instantiates a new TelemetryDruidDurationGranularityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDurationGranularityAllOf(duration int64) *TelemetryDruidDurationGranularityAllOf {
	this := TelemetryDruidDurationGranularityAllOf{}
	this.Duration = duration
	return &this
}

// NewTelemetryDruidDurationGranularityAllOfWithDefaults instantiates a new TelemetryDruidDurationGranularityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDurationGranularityAllOfWithDefaults() *TelemetryDruidDurationGranularityAllOf {
	this := TelemetryDruidDurationGranularityAllOf{}
	return &this
}

// GetDuration returns the Duration field value
func (o *TelemetryDruidDurationGranularityAllOf) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDurationGranularityAllOf) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *TelemetryDruidDurationGranularityAllOf) SetDuration(v int64) {
	o.Duration = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *TelemetryDruidDurationGranularityAllOf) GetOrigin() time.Time {
	if o == nil || o.Origin == nil {
		var ret time.Time
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDurationGranularityAllOf) GetOriginOk() (*time.Time, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *TelemetryDruidDurationGranularityAllOf) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given time.Time and assigns it to the Origin field.
func (o *TelemetryDruidDurationGranularityAllOf) SetOrigin(v time.Time) {
	o.Origin = &v
}

func (o TelemetryDruidDurationGranularityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidDurationGranularityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidDurationGranularityAllOf := _TelemetryDruidDurationGranularityAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidDurationGranularityAllOf); err == nil {
		*o = TelemetryDruidDurationGranularityAllOf(varTelemetryDruidDurationGranularityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "origin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidDurationGranularityAllOf struct {
	value *TelemetryDruidDurationGranularityAllOf
	isSet bool
}

func (v NullableTelemetryDruidDurationGranularityAllOf) Get() *TelemetryDruidDurationGranularityAllOf {
	return v.value
}

func (v *NullableTelemetryDruidDurationGranularityAllOf) Set(val *TelemetryDruidDurationGranularityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDurationGranularityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDurationGranularityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDurationGranularityAllOf(val *TelemetryDruidDurationGranularityAllOf) *NullableTelemetryDruidDurationGranularityAllOf {
	return &NullableTelemetryDruidDurationGranularityAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidDurationGranularityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDurationGranularityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
