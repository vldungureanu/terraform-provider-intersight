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

// CapabilitySectionAllOf Definition of the list of properties defined in 'capability.Section', excluding properties defined in parent classes.
type CapabilitySectionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrative action to initialize/load the catalog section from a particular source. * `None` - Nil value to indicate that no action is required. * `LoadLocal` - Load the catalog file packaged with the service. * `LoadIntersightRepository` - Load a catalog file hosted in the Intersight Repository.
	Action *string `json:"Action,omitempty"`
	// The catalog name reference.
	CatalogName *string `json:"CatalogName,omitempty"`
	// A unique name for the section inside a catalog.
	Name *string `json:"Name,omitempty"`
	// The configured source for this section of the catalog. * `Local` - The catalog file is packaged with the service. * `IntersightRepository` - The catalog file is hosted in the Intersight Repository.
	Source *string `json:"Source,omitempty"`
	// Version of the section inside a catalog.
	Version              *string                        `json:"Version,omitempty"`
	Catalog              *CapabilityCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySectionAllOf CapabilitySectionAllOf

// NewCapabilitySectionAllOf instantiates a new CapabilitySectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySectionAllOf(classId string, objectType string) *CapabilitySectionAllOf {
	this := CapabilitySectionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// NewCapabilitySectionAllOfWithDefaults instantiates a new CapabilitySectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySectionAllOfWithDefaults() *CapabilitySectionAllOf {
	this := CapabilitySectionAllOf{}
	var classId string = "capability.Section"
	this.ClassId = classId
	var objectType string = "capability.Section"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySectionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySectionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySectionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySectionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CapabilitySectionAllOf) SetAction(v string) {
	o.Action = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetCatalogName() string {
	if o == nil || o.CatalogName == nil {
		var ret string
		return ret
	}
	return *o.CatalogName
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetCatalogNameOk() (*string, bool) {
	if o == nil || o.CatalogName == nil {
		return nil, false
	}
	return o.CatalogName, true
}

// HasCatalogName returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasCatalogName() bool {
	if o != nil && o.CatalogName != nil {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given string and assigns it to the CatalogName field.
func (o *CapabilitySectionAllOf) SetCatalogName(v string) {
	o.CatalogName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CapabilitySectionAllOf) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CapabilitySectionAllOf) SetSource(v string) {
	o.Source = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CapabilitySectionAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *CapabilitySectionAllOf) GetCatalog() CapabilityCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret CapabilityCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySectionAllOf) GetCatalogOk() (*CapabilityCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *CapabilitySectionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given CapabilityCatalogRelationship and assigns it to the Catalog field.
func (o *CapabilitySectionAllOf) SetCatalog(v CapabilityCatalogRelationship) {
	o.Catalog = &v
}

func (o CapabilitySectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.CatalogName != nil {
		toSerialize["CatalogName"] = o.CatalogName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySectionAllOf := _CapabilitySectionAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySectionAllOf); err == nil {
		*o = CapabilitySectionAllOf(varCapabilitySectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "CatalogName")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySectionAllOf struct {
	value *CapabilitySectionAllOf
	isSet bool
}

func (v NullableCapabilitySectionAllOf) Get() *CapabilitySectionAllOf {
	return v.value
}

func (v *NullableCapabilitySectionAllOf) Set(val *CapabilitySectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySectionAllOf(val *CapabilitySectionAllOf) *NullableCapabilitySectionAllOf {
	return &NullableCapabilitySectionAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
