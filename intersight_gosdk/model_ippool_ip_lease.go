/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IppoolIpLease IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
type IppoolIpLease struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType *string `json:"IpType,omitempty"`
	// IPv4 Address given as a lease to an external entity like server profiles.
	IpV4Address *string                  `json:"IpV4Address,omitempty"`
	IpV4Config  NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
	// IPv6 Address given as a lease to an external entity like server profiles.
	IpV6Address                   *string                                          `json:"IpV6Address,omitempty"`
	IpV6Config                    NullableIppoolIpV6Config                         `json:"IpV6Config,omitempty"`
	Var0ClusterProfile            *KubernetesClusterProfileRelationship            `json:"_0_ClusterProfile,omitempty"`
	Var1ClusterProfile            *KubernetesClusterProfileRelationship            `json:"_1_ClusterProfile,omitempty"`
	Var2VirtualMachineNodeProfile *KubernetesVirtualMachineNodeProfileRelationship `json:"_2_VirtualMachineNodeProfile,omitempty"`
	Var3VirtualMachineNodeProfile *KubernetesVirtualMachineNodeProfileRelationship `json:"_3_VirtualMachineNodeProfile,omitempty"`
	AssignedToEntity              *MoBaseMoRelationship                            `json:"AssignedToEntity,omitempty"`
	Pool                          *IppoolPoolRelationship                          `json:"Pool,omitempty"`
	PoolMember                    *IppoolPoolMemberRelationship                    `json:"PoolMember,omitempty"`
	Universe                      *IppoolUniverseRelationship                      `json:"Universe,omitempty"`
	Vrf                           *VrfVrfRelationship                              `json:"Vrf,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _IppoolIpLease IppoolIpLease

// NewIppoolIpLease instantiates a new IppoolIpLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpLease(classId string, objectType string) *IppoolIpLease {
	this := IppoolIpLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolIpLeaseWithDefaults instantiates a new IppoolIpLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpLeaseWithDefaults() *IppoolIpLease {
	this := IppoolIpLease{}
	var classId string = "ippool.IpLease"
	this.ClassId = classId
	var objectType string = "ippool.IpLease"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpLease) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolIpLease) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolIpLease) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolIpLease) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolIpLease) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolIpLease) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config.Get() == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config.Get()
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolIpLease) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV4Config.Get(), o.IpV4Config.IsSet()
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV4Config() bool {
	if o != nil && o.IpV4Config.IsSet() {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolIpLease) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config.Set(&v)
}

// SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil
func (o *IppoolIpLease) SetIpV4ConfigNil() {
	o.IpV4Config.Set(nil)
}

// UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
func (o *IppoolIpLease) UnsetIpV4Config() {
	o.IpV4Config.Unset()
}

// GetIpV6Address returns the IpV6Address field value if set, zero value otherwise.
func (o *IppoolIpLease) GetIpV6Address() string {
	if o == nil || o.IpV6Address == nil {
		var ret string
		return ret
	}
	return *o.IpV6Address
}

// GetIpV6AddressOk returns a tuple with the IpV6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetIpV6AddressOk() (*string, bool) {
	if o == nil || o.IpV6Address == nil {
		return nil, false
	}
	return o.IpV6Address, true
}

// HasIpV6Address returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV6Address() bool {
	if o != nil && o.IpV6Address != nil {
		return true
	}

	return false
}

// SetIpV6Address gets a reference to the given string and assigns it to the IpV6Address field.
func (o *IppoolIpLease) SetIpV6Address(v string) {
	o.IpV6Address = &v
}

// GetIpV6Config returns the IpV6Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolIpLease) GetIpV6Config() IppoolIpV6Config {
	if o == nil || o.IpV6Config.Get() == nil {
		var ret IppoolIpV6Config
		return ret
	}
	return *o.IpV6Config.Get()
}

// GetIpV6ConfigOk returns a tuple with the IpV6Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolIpLease) GetIpV6ConfigOk() (*IppoolIpV6Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV6Config.Get(), o.IpV6Config.IsSet()
}

// HasIpV6Config returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV6Config() bool {
	if o != nil && o.IpV6Config.IsSet() {
		return true
	}

	return false
}

// SetIpV6Config gets a reference to the given NullableIppoolIpV6Config and assigns it to the IpV6Config field.
func (o *IppoolIpLease) SetIpV6Config(v IppoolIpV6Config) {
	o.IpV6Config.Set(&v)
}

// SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil
func (o *IppoolIpLease) SetIpV6ConfigNil() {
	o.IpV6Config.Set(nil)
}

// UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
func (o *IppoolIpLease) UnsetIpV6Config() {
	o.IpV6Config.Unset()
}

// GetVar0ClusterProfile returns the Var0ClusterProfile field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVar0ClusterProfile() KubernetesClusterProfileRelationship {
	if o == nil || o.Var0ClusterProfile == nil {
		var ret KubernetesClusterProfileRelationship
		return ret
	}
	return *o.Var0ClusterProfile
}

// GetVar0ClusterProfileOk returns a tuple with the Var0ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVar0ClusterProfileOk() (*KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.Var0ClusterProfile == nil {
		return nil, false
	}
	return o.Var0ClusterProfile, true
}

// HasVar0ClusterProfile returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVar0ClusterProfile() bool {
	if o != nil && o.Var0ClusterProfile != nil {
		return true
	}

	return false
}

// SetVar0ClusterProfile gets a reference to the given KubernetesClusterProfileRelationship and assigns it to the Var0ClusterProfile field.
func (o *IppoolIpLease) SetVar0ClusterProfile(v KubernetesClusterProfileRelationship) {
	o.Var0ClusterProfile = &v
}

// GetVar1ClusterProfile returns the Var1ClusterProfile field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVar1ClusterProfile() KubernetesClusterProfileRelationship {
	if o == nil || o.Var1ClusterProfile == nil {
		var ret KubernetesClusterProfileRelationship
		return ret
	}
	return *o.Var1ClusterProfile
}

// GetVar1ClusterProfileOk returns a tuple with the Var1ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVar1ClusterProfileOk() (*KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.Var1ClusterProfile == nil {
		return nil, false
	}
	return o.Var1ClusterProfile, true
}

// HasVar1ClusterProfile returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVar1ClusterProfile() bool {
	if o != nil && o.Var1ClusterProfile != nil {
		return true
	}

	return false
}

// SetVar1ClusterProfile gets a reference to the given KubernetesClusterProfileRelationship and assigns it to the Var1ClusterProfile field.
func (o *IppoolIpLease) SetVar1ClusterProfile(v KubernetesClusterProfileRelationship) {
	o.Var1ClusterProfile = &v
}

// GetVar2VirtualMachineNodeProfile returns the Var2VirtualMachineNodeProfile field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVar2VirtualMachineNodeProfile() KubernetesVirtualMachineNodeProfileRelationship {
	if o == nil || o.Var2VirtualMachineNodeProfile == nil {
		var ret KubernetesVirtualMachineNodeProfileRelationship
		return ret
	}
	return *o.Var2VirtualMachineNodeProfile
}

// GetVar2VirtualMachineNodeProfileOk returns a tuple with the Var2VirtualMachineNodeProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVar2VirtualMachineNodeProfileOk() (*KubernetesVirtualMachineNodeProfileRelationship, bool) {
	if o == nil || o.Var2VirtualMachineNodeProfile == nil {
		return nil, false
	}
	return o.Var2VirtualMachineNodeProfile, true
}

// HasVar2VirtualMachineNodeProfile returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVar2VirtualMachineNodeProfile() bool {
	if o != nil && o.Var2VirtualMachineNodeProfile != nil {
		return true
	}

	return false
}

// SetVar2VirtualMachineNodeProfile gets a reference to the given KubernetesVirtualMachineNodeProfileRelationship and assigns it to the Var2VirtualMachineNodeProfile field.
func (o *IppoolIpLease) SetVar2VirtualMachineNodeProfile(v KubernetesVirtualMachineNodeProfileRelationship) {
	o.Var2VirtualMachineNodeProfile = &v
}

// GetVar3VirtualMachineNodeProfile returns the Var3VirtualMachineNodeProfile field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVar3VirtualMachineNodeProfile() KubernetesVirtualMachineNodeProfileRelationship {
	if o == nil || o.Var3VirtualMachineNodeProfile == nil {
		var ret KubernetesVirtualMachineNodeProfileRelationship
		return ret
	}
	return *o.Var3VirtualMachineNodeProfile
}

// GetVar3VirtualMachineNodeProfileOk returns a tuple with the Var3VirtualMachineNodeProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVar3VirtualMachineNodeProfileOk() (*KubernetesVirtualMachineNodeProfileRelationship, bool) {
	if o == nil || o.Var3VirtualMachineNodeProfile == nil {
		return nil, false
	}
	return o.Var3VirtualMachineNodeProfile, true
}

// HasVar3VirtualMachineNodeProfile returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVar3VirtualMachineNodeProfile() bool {
	if o != nil && o.Var3VirtualMachineNodeProfile != nil {
		return true
	}

	return false
}

// SetVar3VirtualMachineNodeProfile gets a reference to the given KubernetesVirtualMachineNodeProfileRelationship and assigns it to the Var3VirtualMachineNodeProfile field.
func (o *IppoolIpLease) SetVar3VirtualMachineNodeProfile(v KubernetesVirtualMachineNodeProfileRelationship) {
	o.Var3VirtualMachineNodeProfile = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IppoolIpLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolIpLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolIpLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolIpLease) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolIpLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolIpLease) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *IppoolIpLease) GetPoolMember() IppoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret IppoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IppoolIpLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given IppoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IppoolIpLease) SetPoolMember(v IppoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IppoolIpLease) GetUniverse() IppoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolIpLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolIpLease) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolIpLease) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolIpLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.IpV4Config.IsSet() {
		toSerialize["IpV4Config"] = o.IpV4Config.Get()
	}
	if o.IpV6Address != nil {
		toSerialize["IpV6Address"] = o.IpV6Address
	}
	if o.IpV6Config.IsSet() {
		toSerialize["IpV6Config"] = o.IpV6Config.Get()
	}
	if o.Var0ClusterProfile != nil {
		toSerialize["_0_ClusterProfile"] = o.Var0ClusterProfile
	}
	if o.Var1ClusterProfile != nil {
		toSerialize["_1_ClusterProfile"] = o.Var1ClusterProfile
	}
	if o.Var2VirtualMachineNodeProfile != nil {
		toSerialize["_2_VirtualMachineNodeProfile"] = o.Var2VirtualMachineNodeProfile
	}
	if o.Var3VirtualMachineNodeProfile != nil {
		toSerialize["_3_VirtualMachineNodeProfile"] = o.Var3VirtualMachineNodeProfile
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolIpLease) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolIpLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
		IpType *string `json:"IpType,omitempty"`
		// IPv4 Address given as a lease to an external entity like server profiles.
		IpV4Address *string                  `json:"IpV4Address,omitempty"`
		IpV4Config  NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
		// IPv6 Address given as a lease to an external entity like server profiles.
		IpV6Address                   *string                                          `json:"IpV6Address,omitempty"`
		IpV6Config                    NullableIppoolIpV6Config                         `json:"IpV6Config,omitempty"`
		Var0ClusterProfile            *KubernetesClusterProfileRelationship            `json:"_0_ClusterProfile,omitempty"`
		Var1ClusterProfile            *KubernetesClusterProfileRelationship            `json:"_1_ClusterProfile,omitempty"`
		Var2VirtualMachineNodeProfile *KubernetesVirtualMachineNodeProfileRelationship `json:"_2_VirtualMachineNodeProfile,omitempty"`
		Var3VirtualMachineNodeProfile *KubernetesVirtualMachineNodeProfileRelationship `json:"_3_VirtualMachineNodeProfile,omitempty"`
		AssignedToEntity              *MoBaseMoRelationship                            `json:"AssignedToEntity,omitempty"`
		Pool                          *IppoolPoolRelationship                          `json:"Pool,omitempty"`
		PoolMember                    *IppoolPoolMemberRelationship                    `json:"PoolMember,omitempty"`
		Universe                      *IppoolUniverseRelationship                      `json:"Universe,omitempty"`
		Vrf                           *VrfVrfRelationship                              `json:"Vrf,omitempty"`
	}

	varIppoolIpLeaseWithoutEmbeddedStruct := IppoolIpLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolIpLeaseWithoutEmbeddedStruct)
	if err == nil {
		varIppoolIpLease := _IppoolIpLease{}
		varIppoolIpLease.ClassId = varIppoolIpLeaseWithoutEmbeddedStruct.ClassId
		varIppoolIpLease.ObjectType = varIppoolIpLeaseWithoutEmbeddedStruct.ObjectType
		varIppoolIpLease.IpType = varIppoolIpLeaseWithoutEmbeddedStruct.IpType
		varIppoolIpLease.IpV4Address = varIppoolIpLeaseWithoutEmbeddedStruct.IpV4Address
		varIppoolIpLease.IpV4Config = varIppoolIpLeaseWithoutEmbeddedStruct.IpV4Config
		varIppoolIpLease.IpV6Address = varIppoolIpLeaseWithoutEmbeddedStruct.IpV6Address
		varIppoolIpLease.IpV6Config = varIppoolIpLeaseWithoutEmbeddedStruct.IpV6Config
		varIppoolIpLease.Var0ClusterProfile = varIppoolIpLeaseWithoutEmbeddedStruct.Var0ClusterProfile
		varIppoolIpLease.Var1ClusterProfile = varIppoolIpLeaseWithoutEmbeddedStruct.Var1ClusterProfile
		varIppoolIpLease.Var2VirtualMachineNodeProfile = varIppoolIpLeaseWithoutEmbeddedStruct.Var2VirtualMachineNodeProfile
		varIppoolIpLease.Var3VirtualMachineNodeProfile = varIppoolIpLeaseWithoutEmbeddedStruct.Var3VirtualMachineNodeProfile
		varIppoolIpLease.AssignedToEntity = varIppoolIpLeaseWithoutEmbeddedStruct.AssignedToEntity
		varIppoolIpLease.Pool = varIppoolIpLeaseWithoutEmbeddedStruct.Pool
		varIppoolIpLease.PoolMember = varIppoolIpLeaseWithoutEmbeddedStruct.PoolMember
		varIppoolIpLease.Universe = varIppoolIpLeaseWithoutEmbeddedStruct.Universe
		varIppoolIpLease.Vrf = varIppoolIpLeaseWithoutEmbeddedStruct.Vrf
		*o = IppoolIpLease(varIppoolIpLease)
	} else {
		return err
	}

	varIppoolIpLease := _IppoolIpLease{}

	err = json.Unmarshal(bytes, &varIppoolIpLease)
	if err == nil {
		o.MoBaseMo = varIppoolIpLease.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "IpV6Address")
		delete(additionalProperties, "IpV6Config")
		delete(additionalProperties, "_0_ClusterProfile")
		delete(additionalProperties, "_1_ClusterProfile")
		delete(additionalProperties, "_2_VirtualMachineNodeProfile")
		delete(additionalProperties, "_3_VirtualMachineNodeProfile")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		delete(additionalProperties, "Vrf")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIppoolIpLease struct {
	value *IppoolIpLease
	isSet bool
}

func (v NullableIppoolIpLease) Get() *IppoolIpLease {
	return v.value
}

func (v *NullableIppoolIpLease) Set(val *IppoolIpLease) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpLease) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpLease(val *IppoolIpLease) *NullableIppoolIpLease {
	return &NullableIppoolIpLease{value: val, isSet: true}
}

func (v NullableIppoolIpLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
