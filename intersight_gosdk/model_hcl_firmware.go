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

// HclFirmware Model which holds the details of firmware version and driver version.
type HclFirmware struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3.
	DriverName *string `json:"DriverName,omitempty"`
	// Version of the Driver supported.
	DriverVersion *string `json:"DriverVersion,omitempty"`
	// Error code for the support status. * `Success` - The input combination is valid. * `Unknown` - Unknown API request to the service. * `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database. * `InvalidUcsVersion` - UCS Version is not in expected format. * `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * `OSUnknown` - OS vendor or version is not known as per the HCL database. * `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported. * `ProductUnknown` - Product is not known as per the HCL database. * `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version. * `DriverNameNotSupported` - Driver protocol or name is not supported for the given product. * `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination. * `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination. * `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * `InternalError` - API requests to the service have either failed or timed out. * `MarshallingError` - Error in JSON marshalling. * `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// Firmware version of the product/adapter supported.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// Identifier of the firmware.
	Id *string `json:"Id,omitempty"`
	// True if the driver is latest recommended driver.
	LatestDriver *bool `json:"LatestDriver,omitempty"`
	// True if the firmware is latest recommended firmware.
	LatestFirmware       *bool `json:"LatestFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclFirmware HclFirmware

// NewHclFirmware instantiates a new HclFirmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclFirmware(classId string, objectType string) *HclFirmware {
	this := HclFirmware{}
	this.ClassId = classId
	this.ObjectType = objectType
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// NewHclFirmwareWithDefaults instantiates a new HclFirmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclFirmwareWithDefaults() *HclFirmware {
	this := HclFirmware{}
	var classId string = "hcl.Firmware"
	this.ClassId = classId
	var objectType string = "hcl.Firmware"
	this.ObjectType = objectType
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclFirmware) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclFirmware) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclFirmware) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclFirmware) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverName returns the DriverName field value if set, zero value otherwise.
func (o *HclFirmware) GetDriverName() string {
	if o == nil || o.DriverName == nil {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetDriverNameOk() (*string, bool) {
	if o == nil || o.DriverName == nil {
		return nil, false
	}
	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *HclFirmware) HasDriverName() bool {
	if o != nil && o.DriverName != nil {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *HclFirmware) SetDriverName(v string) {
	o.DriverName = &v
}

// GetDriverVersion returns the DriverVersion field value if set, zero value otherwise.
func (o *HclFirmware) GetDriverVersion() string {
	if o == nil || o.DriverVersion == nil {
		var ret string
		return ret
	}
	return *o.DriverVersion
}

// GetDriverVersionOk returns a tuple with the DriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetDriverVersionOk() (*string, bool) {
	if o == nil || o.DriverVersion == nil {
		return nil, false
	}
	return o.DriverVersion, true
}

// HasDriverVersion returns a boolean if a field has been set.
func (o *HclFirmware) HasDriverVersion() bool {
	if o != nil && o.DriverVersion != nil {
		return true
	}

	return false
}

// SetDriverVersion gets a reference to the given string and assigns it to the DriverVersion field.
func (o *HclFirmware) SetDriverVersion(v string) {
	o.DriverVersion = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclFirmware) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclFirmware) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclFirmware) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *HclFirmware) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *HclFirmware) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *HclFirmware) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclFirmware) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclFirmware) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclFirmware) SetId(v string) {
	o.Id = &v
}

// GetLatestDriver returns the LatestDriver field value if set, zero value otherwise.
func (o *HclFirmware) GetLatestDriver() bool {
	if o == nil || o.LatestDriver == nil {
		var ret bool
		return ret
	}
	return *o.LatestDriver
}

// GetLatestDriverOk returns a tuple with the LatestDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetLatestDriverOk() (*bool, bool) {
	if o == nil || o.LatestDriver == nil {
		return nil, false
	}
	return o.LatestDriver, true
}

// HasLatestDriver returns a boolean if a field has been set.
func (o *HclFirmware) HasLatestDriver() bool {
	if o != nil && o.LatestDriver != nil {
		return true
	}

	return false
}

// SetLatestDriver gets a reference to the given bool and assigns it to the LatestDriver field.
func (o *HclFirmware) SetLatestDriver(v bool) {
	o.LatestDriver = &v
}

// GetLatestFirmware returns the LatestFirmware field value if set, zero value otherwise.
func (o *HclFirmware) GetLatestFirmware() bool {
	if o == nil || o.LatestFirmware == nil {
		var ret bool
		return ret
	}
	return *o.LatestFirmware
}

// GetLatestFirmwareOk returns a tuple with the LatestFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmware) GetLatestFirmwareOk() (*bool, bool) {
	if o == nil || o.LatestFirmware == nil {
		return nil, false
	}
	return o.LatestFirmware, true
}

// HasLatestFirmware returns a boolean if a field has been set.
func (o *HclFirmware) HasLatestFirmware() bool {
	if o != nil && o.LatestFirmware != nil {
		return true
	}

	return false
}

// SetLatestFirmware gets a reference to the given bool and assigns it to the LatestFirmware field.
func (o *HclFirmware) SetLatestFirmware(v bool) {
	o.LatestFirmware = &v
}

func (o HclFirmware) MarshalJSON() ([]byte, error) {
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
	if o.DriverName != nil {
		toSerialize["DriverName"] = o.DriverName
	}
	if o.DriverVersion != nil {
		toSerialize["DriverVersion"] = o.DriverVersion
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.LatestDriver != nil {
		toSerialize["LatestDriver"] = o.LatestDriver
	}
	if o.LatestFirmware != nil {
		toSerialize["LatestFirmware"] = o.LatestFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclFirmware) UnmarshalJSON(bytes []byte) (err error) {
	type HclFirmwareWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3.
		DriverName *string `json:"DriverName,omitempty"`
		// Version of the Driver supported.
		DriverVersion *string `json:"DriverVersion,omitempty"`
		// Error code for the support status. * `Success` - The input combination is valid. * `Unknown` - Unknown API request to the service. * `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database. * `InvalidUcsVersion` - UCS Version is not in expected format. * `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * `OSUnknown` - OS vendor or version is not known as per the HCL database. * `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported. * `ProductUnknown` - Product is not known as per the HCL database. * `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version. * `DriverNameNotSupported` - Driver protocol or name is not supported for the given product. * `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination. * `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination. * `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * `InternalError` - API requests to the service have either failed or timed out. * `MarshallingError` - Error in JSON marshalling. * `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.
		ErrorCode *string `json:"ErrorCode,omitempty"`
		// Firmware version of the product/adapter supported.
		FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
		// Identifier of the firmware.
		Id *string `json:"Id,omitempty"`
		// True if the driver is latest recommended driver.
		LatestDriver *bool `json:"LatestDriver,omitempty"`
		// True if the firmware is latest recommended firmware.
		LatestFirmware *bool `json:"LatestFirmware,omitempty"`
	}

	varHclFirmwareWithoutEmbeddedStruct := HclFirmwareWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclFirmwareWithoutEmbeddedStruct)
	if err == nil {
		varHclFirmware := _HclFirmware{}
		varHclFirmware.ClassId = varHclFirmwareWithoutEmbeddedStruct.ClassId
		varHclFirmware.ObjectType = varHclFirmwareWithoutEmbeddedStruct.ObjectType
		varHclFirmware.DriverName = varHclFirmwareWithoutEmbeddedStruct.DriverName
		varHclFirmware.DriverVersion = varHclFirmwareWithoutEmbeddedStruct.DriverVersion
		varHclFirmware.ErrorCode = varHclFirmwareWithoutEmbeddedStruct.ErrorCode
		varHclFirmware.FirmwareVersion = varHclFirmwareWithoutEmbeddedStruct.FirmwareVersion
		varHclFirmware.Id = varHclFirmwareWithoutEmbeddedStruct.Id
		varHclFirmware.LatestDriver = varHclFirmwareWithoutEmbeddedStruct.LatestDriver
		varHclFirmware.LatestFirmware = varHclFirmwareWithoutEmbeddedStruct.LatestFirmware
		*o = HclFirmware(varHclFirmware)
	} else {
		return err
	}

	varHclFirmware := _HclFirmware{}

	err = json.Unmarshal(bytes, &varHclFirmware)
	if err == nil {
		o.MoBaseComplexType = varHclFirmware.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverName")
		delete(additionalProperties, "DriverVersion")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "LatestDriver")
		delete(additionalProperties, "LatestFirmware")

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

type NullableHclFirmware struct {
	value *HclFirmware
	isSet bool
}

func (v NullableHclFirmware) Get() *HclFirmware {
	return v.value
}

func (v *NullableHclFirmware) Set(val *HclFirmware) {
	v.value = val
	v.isSet = true
}

func (v NullableHclFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableHclFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclFirmware(val *HclFirmware) *NullableHclFirmware {
	return &NullableHclFirmware{value: val, isSet: true}
}

func (v NullableHclFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
