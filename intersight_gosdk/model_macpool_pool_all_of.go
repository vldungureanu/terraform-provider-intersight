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

// MacpoolPoolAllOf Definition of the list of properties defined in 'macpool.Pool', excluding properties defined in parent classes.
type MacpoolPoolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string         `json:"ObjectType"`
	MacBlocks  []MacpoolBlock `json:"MacBlocks,omitempty"`
	// An array of relationships to macpoolIdBlock resources.
	BlockHeads           []MacpoolIdBlockRelationship          `json:"BlockHeads,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolPoolAllOf MacpoolPoolAllOf

// NewMacpoolPoolAllOf instantiates a new MacpoolPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolPoolAllOf(classId string, objectType string) *MacpoolPoolAllOf {
	this := MacpoolPoolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolPoolAllOfWithDefaults instantiates a new MacpoolPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolPoolAllOfWithDefaults() *MacpoolPoolAllOf {
	this := MacpoolPoolAllOf{}
	var classId string = "macpool.Pool"
	this.ClassId = classId
	var objectType string = "macpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolPoolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolPoolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolPoolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolPoolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacBlocks returns the MacBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolAllOf) GetMacBlocks() []MacpoolBlock {
	if o == nil {
		var ret []MacpoolBlock
		return ret
	}
	return o.MacBlocks
}

// GetMacBlocksOk returns a tuple with the MacBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolAllOf) GetMacBlocksOk() (*[]MacpoolBlock, bool) {
	if o == nil || o.MacBlocks == nil {
		return nil, false
	}
	return &o.MacBlocks, true
}

// HasMacBlocks returns a boolean if a field has been set.
func (o *MacpoolPoolAllOf) HasMacBlocks() bool {
	if o != nil && o.MacBlocks != nil {
		return true
	}

	return false
}

// SetMacBlocks gets a reference to the given []MacpoolBlock and assigns it to the MacBlocks field.
func (o *MacpoolPoolAllOf) SetMacBlocks(v []MacpoolBlock) {
	o.MacBlocks = v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolPoolAllOf) GetBlockHeads() []MacpoolIdBlockRelationship {
	if o == nil {
		var ret []MacpoolIdBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolPoolAllOf) GetBlockHeadsOk() (*[]MacpoolIdBlockRelationship, bool) {
	if o == nil || o.BlockHeads == nil {
		return nil, false
	}
	return &o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *MacpoolPoolAllOf) HasBlockHeads() bool {
	if o != nil && o.BlockHeads != nil {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []MacpoolIdBlockRelationship and assigns it to the BlockHeads field.
func (o *MacpoolPoolAllOf) SetBlockHeads(v []MacpoolIdBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *MacpoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *MacpoolPoolAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *MacpoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o MacpoolPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacBlocks != nil {
		toSerialize["MacBlocks"] = o.MacBlocks
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolPoolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMacpoolPoolAllOf := _MacpoolPoolAllOf{}

	if err = json.Unmarshal(bytes, &varMacpoolPoolAllOf); err == nil {
		*o = MacpoolPoolAllOf(varMacpoolPoolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacBlocks")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMacpoolPoolAllOf struct {
	value *MacpoolPoolAllOf
	isSet bool
}

func (v NullableMacpoolPoolAllOf) Get() *MacpoolPoolAllOf {
	return v.value
}

func (v *NullableMacpoolPoolAllOf) Set(val *MacpoolPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolPoolAllOf(val *MacpoolPoolAllOf) *NullableMacpoolPoolAllOf {
	return &NullableMacpoolPoolAllOf{value: val, isSet: true}
}

func (v NullableMacpoolPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
