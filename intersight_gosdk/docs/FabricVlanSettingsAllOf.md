# FabricVlanSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanSettings"]
**AllowedVlans** | Pointer to **string** | Allowed VLAN IDs of the virtual interface. | [optional] 
**NativeVlan** | Pointer to **int64** | Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface. | [optional] [default to 0]

## Methods

### NewFabricVlanSettingsAllOf

`func NewFabricVlanSettingsAllOf(classId string, objectType string, ) *FabricVlanSettingsAllOf`

NewFabricVlanSettingsAllOf instantiates a new FabricVlanSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSettingsAllOfWithDefaults

`func NewFabricVlanSettingsAllOfWithDefaults() *FabricVlanSettingsAllOf`

NewFabricVlanSettingsAllOfWithDefaults instantiates a new FabricVlanSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedVlans

`func (o *FabricVlanSettingsAllOf) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *FabricVlanSettingsAllOf) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *FabricVlanSettingsAllOf) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *FabricVlanSettingsAllOf) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetNativeVlan

`func (o *FabricVlanSettingsAllOf) GetNativeVlan() int64`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *FabricVlanSettingsAllOf) GetNativeVlanOk() (*int64, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *FabricVlanSettingsAllOf) SetNativeVlan(v int64)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *FabricVlanSettingsAllOf) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


