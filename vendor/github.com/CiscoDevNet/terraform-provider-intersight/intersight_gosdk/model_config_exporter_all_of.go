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

// ConfigExporterAllOf Definition of the list of properties defined in 'config.Exporter', excluding properties defined in parent classes.
type ConfigExporterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired.
	DownloadPath *string       `json:"DownloadPath,omitempty"`
	Items        []ConfigMoRef `json:"Items,omitempty"`
	// An identifier for the exporter instance.
	Name *string `json:"Name,omitempty"`
	// Status of the export operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// An array of relationships to configExportedItem resources.
	ExportedItems        []ConfigExportedItemRelationship      `json:"ExportedItems,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigExporterAllOf ConfigExporterAllOf

// NewConfigExporterAllOf instantiates a new ConfigExporterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigExporterAllOf(classId string, objectType string) *ConfigExporterAllOf {
	this := ConfigExporterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigExporterAllOfWithDefaults instantiates a new ConfigExporterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigExporterAllOfWithDefaults() *ConfigExporterAllOf {
	this := ConfigExporterAllOf{}
	var classId string = "config.Exporter"
	this.ClassId = classId
	var objectType string = "config.Exporter"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConfigExporterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConfigExporterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConfigExporterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConfigExporterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadPath returns the DownloadPath field value if set, zero value otherwise.
func (o *ConfigExporterAllOf) GetDownloadPath() string {
	if o == nil || o.DownloadPath == nil {
		var ret string
		return ret
	}
	return *o.DownloadPath
}

// GetDownloadPathOk returns a tuple with the DownloadPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetDownloadPathOk() (*string, bool) {
	if o == nil || o.DownloadPath == nil {
		return nil, false
	}
	return o.DownloadPath, true
}

// HasDownloadPath returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasDownloadPath() bool {
	if o != nil && o.DownloadPath != nil {
		return true
	}

	return false
}

// SetDownloadPath gets a reference to the given string and assigns it to the DownloadPath field.
func (o *ConfigExporterAllOf) SetDownloadPath(v string) {
	o.DownloadPath = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigExporterAllOf) GetItems() []ConfigMoRef {
	if o == nil {
		var ret []ConfigMoRef
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigExporterAllOf) GetItemsOk() (*[]ConfigMoRef, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConfigMoRef and assigns it to the Items field.
func (o *ConfigExporterAllOf) SetItems(v []ConfigMoRef) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigExporterAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigExporterAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigExporterAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigExporterAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigExporterAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigExporterAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetExportedItems returns the ExportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigExporterAllOf) GetExportedItems() []ConfigExportedItemRelationship {
	if o == nil {
		var ret []ConfigExportedItemRelationship
		return ret
	}
	return o.ExportedItems
}

// GetExportedItemsOk returns a tuple with the ExportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigExporterAllOf) GetExportedItemsOk() (*[]ConfigExportedItemRelationship, bool) {
	if o == nil || o.ExportedItems == nil {
		return nil, false
	}
	return &o.ExportedItems, true
}

// HasExportedItems returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasExportedItems() bool {
	if o != nil && o.ExportedItems != nil {
		return true
	}

	return false
}

// SetExportedItems gets a reference to the given []ConfigExportedItemRelationship and assigns it to the ExportedItems field.
func (o *ConfigExporterAllOf) SetExportedItems(v []ConfigExportedItemRelationship) {
	o.ExportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConfigExporterAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporterAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConfigExporterAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ConfigExporterAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ConfigExporterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DownloadPath != nil {
		toSerialize["DownloadPath"] = o.DownloadPath
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.ExportedItems != nil {
		toSerialize["ExportedItems"] = o.ExportedItems
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigExporterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigExporterAllOf := _ConfigExporterAllOf{}

	if err = json.Unmarshal(bytes, &varConfigExporterAllOf); err == nil {
		*o = ConfigExporterAllOf(varConfigExporterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadPath")
		delete(additionalProperties, "Items")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "ExportedItems")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigExporterAllOf struct {
	value *ConfigExporterAllOf
	isSet bool
}

func (v NullableConfigExporterAllOf) Get() *ConfigExporterAllOf {
	return v.value
}

func (v *NullableConfigExporterAllOf) Set(val *ConfigExporterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExporterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExporterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExporterAllOf(val *ConfigExporterAllOf) *NullableConfigExporterAllOf {
	return &NullableConfigExporterAllOf{value: val, isSet: true}
}

func (v NullableConfigExporterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExporterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
