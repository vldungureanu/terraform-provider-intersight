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
	"reflect"
	"strings"
)

// HyperflexHxapHost A HyperFlex Application Platform compute host entity that is part of HyperFlex compute cluster and probably runs VMs.
type HyperflexHxapHost struct {
	VirtualizationBaseHost
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Denotes age or life time of the Host in nano seconds.
	Age *string `json:"Age,omitempty"`
	// The UUID of the cluster to which this Host belongs to.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// Reason of the failure when host is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweredOn` - The entity is powered on. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state.
	HwPowerState *string `json:"HwPowerState,omitempty"`
	// Internal IP Address of the Host.
	InternalIpAddress *string `json:"InternalIpAddress,omitempty"`
	// Management IP Address of the Host.
	ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
	// Is the role of this host is master in the cluster? true or false.
	MasterRole *bool `json:"MasterRole,omitempty"`
	// Product version of the Host.
	Version              *string                           `json:"Version,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	ClusterMember        *AssetClusterMemberRelationship   `json:"ClusterMember,omitempty"`
	PhysicalServer       *ComputePhysicalRelationship      `json:"PhysicalServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapHost HyperflexHxapHost

// NewHyperflexHxapHost instantiates a new HyperflexHxapHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapHost(classId string, objectType string) *HyperflexHxapHost {
	this := HyperflexHxapHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// NewHyperflexHxapHostWithDefaults instantiates a new HyperflexHxapHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapHostWithDefaults() *HyperflexHxapHost {
	this := HyperflexHxapHost{}
	var classId string = "hyperflex.HxapHost"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapHost"
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *HyperflexHxapHost) SetAge(v string) {
	o.Age = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexHxapHost) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *HyperflexHxapHost) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHwPowerState returns the HwPowerState field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetHwPowerState() string {
	if o == nil || o.HwPowerState == nil {
		var ret string
		return ret
	}
	return *o.HwPowerState
}

// GetHwPowerStateOk returns a tuple with the HwPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetHwPowerStateOk() (*string, bool) {
	if o == nil || o.HwPowerState == nil {
		return nil, false
	}
	return o.HwPowerState, true
}

// HasHwPowerState returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasHwPowerState() bool {
	if o != nil && o.HwPowerState != nil {
		return true
	}

	return false
}

// SetHwPowerState gets a reference to the given string and assigns it to the HwPowerState field.
func (o *HyperflexHxapHost) SetHwPowerState(v string) {
	o.HwPowerState = &v
}

// GetInternalIpAddress returns the InternalIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetInternalIpAddress() string {
	if o == nil || o.InternalIpAddress == nil {
		var ret string
		return ret
	}
	return *o.InternalIpAddress
}

// GetInternalIpAddressOk returns a tuple with the InternalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetInternalIpAddressOk() (*string, bool) {
	if o == nil || o.InternalIpAddress == nil {
		return nil, false
	}
	return o.InternalIpAddress, true
}

// HasInternalIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasInternalIpAddress() bool {
	if o != nil && o.InternalIpAddress != nil {
		return true
	}

	return false
}

// SetInternalIpAddress gets a reference to the given string and assigns it to the InternalIpAddress field.
func (o *HyperflexHxapHost) SetInternalIpAddress(v string) {
	o.InternalIpAddress = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *HyperflexHxapHost) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetMasterRole returns the MasterRole field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetMasterRole() bool {
	if o == nil || o.MasterRole == nil {
		var ret bool
		return ret
	}
	return *o.MasterRole
}

// GetMasterRoleOk returns a tuple with the MasterRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetMasterRoleOk() (*bool, bool) {
	if o == nil || o.MasterRole == nil {
		return nil, false
	}
	return o.MasterRole, true
}

// HasMasterRole returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasMasterRole() bool {
	if o != nil && o.MasterRole != nil {
		return true
	}

	return false
}

// SetMasterRole gets a reference to the given bool and assigns it to the MasterRole field.
func (o *HyperflexHxapHost) SetMasterRole(v bool) {
	o.MasterRole = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHxapHost) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapHost) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *HyperflexHxapHost) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *HyperflexHxapHost) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHost) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *HyperflexHxapHost) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *HyperflexHxapHost) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o HyperflexHxapHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHost, errVirtualizationBaseHost := json.Marshal(o.VirtualizationBaseHost)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	errVirtualizationBaseHost = json.Unmarshal([]byte(serializedVirtualizationBaseHost), &toSerialize)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.HwPowerState != nil {
		toSerialize["HwPowerState"] = o.HwPowerState
	}
	if o.InternalIpAddress != nil {
		toSerialize["InternalIpAddress"] = o.InternalIpAddress
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.MasterRole != nil {
		toSerialize["MasterRole"] = o.MasterRole
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapHost) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Denotes age or life time of the Host in nano seconds.
		Age *string `json:"Age,omitempty"`
		// The UUID of the cluster to which this Host belongs to.
		ClusterUuid *string `json:"ClusterUuid,omitempty"`
		// Reason of the failure when host is in failed state.
		FailureReason *string `json:"FailureReason,omitempty"`
		// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweredOn` - The entity is powered on. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state.
		HwPowerState *string `json:"HwPowerState,omitempty"`
		// Internal IP Address of the Host.
		InternalIpAddress *string `json:"InternalIpAddress,omitempty"`
		// Management IP Address of the Host.
		ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
		// Is the role of this host is master in the cluster? true or false.
		MasterRole *bool `json:"MasterRole,omitempty"`
		// Product version of the Host.
		Version        *string                           `json:"Version,omitempty"`
		Cluster        *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
		ClusterMember  *AssetClusterMemberRelationship   `json:"ClusterMember,omitempty"`
		PhysicalServer *ComputePhysicalRelationship      `json:"PhysicalServer,omitempty"`
	}

	varHyperflexHxapHostWithoutEmbeddedStruct := HyperflexHxapHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHostWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapHost := _HyperflexHxapHost{}
		varHyperflexHxapHost.ClassId = varHyperflexHxapHostWithoutEmbeddedStruct.ClassId
		varHyperflexHxapHost.ObjectType = varHyperflexHxapHostWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapHost.Age = varHyperflexHxapHostWithoutEmbeddedStruct.Age
		varHyperflexHxapHost.ClusterUuid = varHyperflexHxapHostWithoutEmbeddedStruct.ClusterUuid
		varHyperflexHxapHost.FailureReason = varHyperflexHxapHostWithoutEmbeddedStruct.FailureReason
		varHyperflexHxapHost.HwPowerState = varHyperflexHxapHostWithoutEmbeddedStruct.HwPowerState
		varHyperflexHxapHost.InternalIpAddress = varHyperflexHxapHostWithoutEmbeddedStruct.InternalIpAddress
		varHyperflexHxapHost.ManagementIpAddress = varHyperflexHxapHostWithoutEmbeddedStruct.ManagementIpAddress
		varHyperflexHxapHost.MasterRole = varHyperflexHxapHostWithoutEmbeddedStruct.MasterRole
		varHyperflexHxapHost.Version = varHyperflexHxapHostWithoutEmbeddedStruct.Version
		varHyperflexHxapHost.Cluster = varHyperflexHxapHostWithoutEmbeddedStruct.Cluster
		varHyperflexHxapHost.ClusterMember = varHyperflexHxapHostWithoutEmbeddedStruct.ClusterMember
		varHyperflexHxapHost.PhysicalServer = varHyperflexHxapHostWithoutEmbeddedStruct.PhysicalServer
		*o = HyperflexHxapHost(varHyperflexHxapHost)
	} else {
		return err
	}

	varHyperflexHxapHost := _HyperflexHxapHost{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHost)
	if err == nil {
		o.VirtualizationBaseHost = varHyperflexHxapHost.VirtualizationBaseHost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Age")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "HwPowerState")
		delete(additionalProperties, "InternalIpAddress")
		delete(additionalProperties, "ManagementIpAddress")
		delete(additionalProperties, "MasterRole")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "PhysicalServer")

		// remove fields from embedded structs
		reflectVirtualizationBaseHost := reflect.ValueOf(o.VirtualizationBaseHost)
		for i := 0; i < reflectVirtualizationBaseHost.Type().NumField(); i++ {
			t := reflectVirtualizationBaseHost.Type().Field(i)

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

type NullableHyperflexHxapHost struct {
	value *HyperflexHxapHost
	isSet bool
}

func (v NullableHyperflexHxapHost) Get() *HyperflexHxapHost {
	return v.value
}

func (v *NullableHyperflexHxapHost) Set(val *HyperflexHxapHost) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapHost(val *HyperflexHxapHost) *NullableHyperflexHxapHost {
	return &NullableHyperflexHxapHost{value: val, isSet: true}
}

func (v NullableHyperflexHxapHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
