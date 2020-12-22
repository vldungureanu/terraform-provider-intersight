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

// ConnectorCommandTerminalStream Holds the i/o of a terminal command session.
type ConnectorCommandTerminalStream struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of data this message contains.
	MsgType *string `json:"MsgType,omitempty"`
	// Sequence of the message within a session to handle out-of-order delivery.
	Sequence *int64 `json:"Sequence,omitempty"`
	// The input/output payload to/from the pseudo terminal session. When sent from the cloud service if the msgType is CommandInput stream is piped to stdin of the command or a resize message if msgType is CommandResize. From the device connector value is always the combined output of stdout & stderr.
	Stream               *string `json:"Stream,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorCommandTerminalStream ConnectorCommandTerminalStream

// NewConnectorCommandTerminalStream instantiates a new ConnectorCommandTerminalStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorCommandTerminalStream(classId string, objectType string) *ConnectorCommandTerminalStream {
	this := ConnectorCommandTerminalStream{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorCommandTerminalStreamWithDefaults instantiates a new ConnectorCommandTerminalStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorCommandTerminalStreamWithDefaults() *ConnectorCommandTerminalStream {
	this := ConnectorCommandTerminalStream{}
	var classId string = "connector.CommandTerminalStream"
	this.ClassId = classId
	var objectType string = "connector.CommandTerminalStream"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorCommandTerminalStream) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorCommandTerminalStream) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorCommandTerminalStream) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorCommandTerminalStream) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorCommandTerminalStream) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorCommandTerminalStream) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *ConnectorCommandTerminalStream) GetMsgType() string {
	if o == nil || o.MsgType == nil {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandTerminalStream) GetMsgTypeOk() (*string, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *ConnectorCommandTerminalStream) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *ConnectorCommandTerminalStream) SetMsgType(v string) {
	o.MsgType = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *ConnectorCommandTerminalStream) GetSequence() int64 {
	if o == nil || o.Sequence == nil {
		var ret int64
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandTerminalStream) GetSequenceOk() (*int64, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *ConnectorCommandTerminalStream) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int64 and assigns it to the Sequence field.
func (o *ConnectorCommandTerminalStream) SetSequence(v int64) {
	o.Sequence = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ConnectorCommandTerminalStream) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandTerminalStream) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ConnectorCommandTerminalStream) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *ConnectorCommandTerminalStream) SetStream(v string) {
	o.Stream = &v
}

func (o ConnectorCommandTerminalStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MsgType != nil {
		toSerialize["MsgType"] = o.MsgType
	}
	if o.Sequence != nil {
		toSerialize["Sequence"] = o.Sequence
	}
	if o.Stream != nil {
		toSerialize["Stream"] = o.Stream
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorCommandTerminalStream) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorCommandTerminalStreamWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The type of data this message contains.
		MsgType *string `json:"MsgType,omitempty"`
		// Sequence of the message within a session to handle out-of-order delivery.
		Sequence *int64 `json:"Sequence,omitempty"`
		// The input/output payload to/from the pseudo terminal session. When sent from the cloud service if the msgType is CommandInput stream is piped to stdin of the command or a resize message if msgType is CommandResize. From the device connector value is always the combined output of stdout & stderr.
		Stream *string `json:"Stream,omitempty"`
	}

	varConnectorCommandTerminalStreamWithoutEmbeddedStruct := ConnectorCommandTerminalStreamWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorCommandTerminalStreamWithoutEmbeddedStruct)
	if err == nil {
		varConnectorCommandTerminalStream := _ConnectorCommandTerminalStream{}
		varConnectorCommandTerminalStream.ClassId = varConnectorCommandTerminalStreamWithoutEmbeddedStruct.ClassId
		varConnectorCommandTerminalStream.ObjectType = varConnectorCommandTerminalStreamWithoutEmbeddedStruct.ObjectType
		varConnectorCommandTerminalStream.MsgType = varConnectorCommandTerminalStreamWithoutEmbeddedStruct.MsgType
		varConnectorCommandTerminalStream.Sequence = varConnectorCommandTerminalStreamWithoutEmbeddedStruct.Sequence
		varConnectorCommandTerminalStream.Stream = varConnectorCommandTerminalStreamWithoutEmbeddedStruct.Stream
		*o = ConnectorCommandTerminalStream(varConnectorCommandTerminalStream)
	} else {
		return err
	}

	varConnectorCommandTerminalStream := _ConnectorCommandTerminalStream{}

	err = json.Unmarshal(bytes, &varConnectorCommandTerminalStream)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorCommandTerminalStream.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MsgType")
		delete(additionalProperties, "Sequence")
		delete(additionalProperties, "Stream")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorCommandTerminalStream struct {
	value *ConnectorCommandTerminalStream
	isSet bool
}

func (v NullableConnectorCommandTerminalStream) Get() *ConnectorCommandTerminalStream {
	return v.value
}

func (v *NullableConnectorCommandTerminalStream) Set(val *ConnectorCommandTerminalStream) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorCommandTerminalStream) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorCommandTerminalStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorCommandTerminalStream(val *ConnectorCommandTerminalStream) *NullableConnectorCommandTerminalStream {
	return &NullableConnectorCommandTerminalStream{value: val, isSet: true}
}

func (v NullableConnectorCommandTerminalStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorCommandTerminalStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
