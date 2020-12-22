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

// OsDistributionAllOf Definition of the list of properties defined in 'os.Distribution', excluding properties defined in parent classes.
type OsDistributionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the OS distribution such as ESXi, CentOS.
	Name                 *string                         `json:"Name,omitempty"`
	SupportedEditions    []string                        `json:"SupportedEditions,omitempty"`
	Catalog              *OsCatalogRelationship          `json:"Catalog,omitempty"`
	Version              *HclOperatingSystemRelationship `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsDistributionAllOf OsDistributionAllOf

// NewOsDistributionAllOf instantiates a new OsDistributionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsDistributionAllOf(classId string, objectType string) *OsDistributionAllOf {
	this := OsDistributionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsDistributionAllOfWithDefaults instantiates a new OsDistributionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsDistributionAllOfWithDefaults() *OsDistributionAllOf {
	this := OsDistributionAllOf{}
	var classId string = "os.Distribution"
	this.ClassId = classId
	var objectType string = "os.Distribution"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsDistributionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsDistributionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsDistributionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsDistributionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsDistributionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsDistributionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsDistributionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistributionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsDistributionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsDistributionAllOf) SetName(v string) {
	o.Name = &v
}

// GetSupportedEditions returns the SupportedEditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsDistributionAllOf) GetSupportedEditions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedEditions
}

// GetSupportedEditionsOk returns a tuple with the SupportedEditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsDistributionAllOf) GetSupportedEditionsOk() (*[]string, bool) {
	if o == nil || o.SupportedEditions == nil {
		return nil, false
	}
	return &o.SupportedEditions, true
}

// HasSupportedEditions returns a boolean if a field has been set.
func (o *OsDistributionAllOf) HasSupportedEditions() bool {
	if o != nil && o.SupportedEditions != nil {
		return true
	}

	return false
}

// SetSupportedEditions gets a reference to the given []string and assigns it to the SupportedEditions field.
func (o *OsDistributionAllOf) SetSupportedEditions(v []string) {
	o.SupportedEditions = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *OsDistributionAllOf) GetCatalog() OsCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret OsCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistributionAllOf) GetCatalogOk() (*OsCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *OsDistributionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given OsCatalogRelationship and assigns it to the Catalog field.
func (o *OsDistributionAllOf) SetCatalog(v OsCatalogRelationship) {
	o.Catalog = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OsDistributionAllOf) GetVersion() HclOperatingSystemRelationship {
	if o == nil || o.Version == nil {
		var ret HclOperatingSystemRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistributionAllOf) GetVersionOk() (*HclOperatingSystemRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OsDistributionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given HclOperatingSystemRelationship and assigns it to the Version field.
func (o *OsDistributionAllOf) SetVersion(v HclOperatingSystemRelationship) {
	o.Version = &v
}

func (o OsDistributionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SupportedEditions != nil {
		toSerialize["SupportedEditions"] = o.SupportedEditions
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsDistributionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsDistributionAllOf := _OsDistributionAllOf{}

	if err = json.Unmarshal(bytes, &varOsDistributionAllOf); err == nil {
		*o = OsDistributionAllOf(varOsDistributionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SupportedEditions")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsDistributionAllOf struct {
	value *OsDistributionAllOf
	isSet bool
}

func (v NullableOsDistributionAllOf) Get() *OsDistributionAllOf {
	return v.value
}

func (v *NullableOsDistributionAllOf) Set(val *OsDistributionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsDistributionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsDistributionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsDistributionAllOf(val *OsDistributionAllOf) *NullableOsDistributionAllOf {
	return &NullableOsDistributionAllOf{value: val, isSet: true}
}

func (v NullableOsDistributionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsDistributionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
