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

// KubernetesEsxiVirtualMachineInfraConfig Infrastructure provider allocation configuration for ESXi virtual machine Kubernetes nodes.
type KubernetesEsxiVirtualMachineInfraConfig struct {
	KubernetesBaseVirtualMachineInfraConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the vSphere cluster on which the virtual machines are created.
	Cluster *string `json:"Cluster,omitempty"`
	// Name of the datasore on which the virtual machine disks are created.
	Datastore *string `json:"Datastore,omitempty"`
	// Indicates whether the value of the 'passphrase' property has been set.
	IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
	// Passphrase for the vcenter user.
	Passphrase *string `json:"Passphrase,omitempty"`
	// Name of the vSphere resource pool on which the virtual machines are created.
	ResourcePool         *string `json:"ResourcePool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEsxiVirtualMachineInfraConfig KubernetesEsxiVirtualMachineInfraConfig

// NewKubernetesEsxiVirtualMachineInfraConfig instantiates a new KubernetesEsxiVirtualMachineInfraConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEsxiVirtualMachineInfraConfig(classId string, objectType string) *KubernetesEsxiVirtualMachineInfraConfig {
	this := KubernetesEsxiVirtualMachineInfraConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPassphraseSet bool = false
	this.IsPassphraseSet = &isPassphraseSet
	return &this
}

// NewKubernetesEsxiVirtualMachineInfraConfigWithDefaults instantiates a new KubernetesEsxiVirtualMachineInfraConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEsxiVirtualMachineInfraConfigWithDefaults() *KubernetesEsxiVirtualMachineInfraConfig {
	this := KubernetesEsxiVirtualMachineInfraConfig{}
	var classId string = "kubernetes.EsxiVirtualMachineInfraConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.EsxiVirtualMachineInfraConfig"
	this.ObjectType = objectType
	var isPassphraseSet bool = false
	this.IsPassphraseSet = &isPassphraseSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetCluster(v string) {
	o.Cluster = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetDatastore(v string) {
	o.Datastore = &v
}

// GetIsPassphraseSet returns the IsPassphraseSet field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetIsPassphraseSet() bool {
	if o == nil || o.IsPassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPassphraseSet
}

// GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetIsPassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPassphraseSet == nil {
		return nil, false
	}
	return o.IsPassphraseSet, true
}

// HasIsPassphraseSet returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) HasIsPassphraseSet() bool {
	if o != nil && o.IsPassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPassphraseSet gets a reference to the given bool and assigns it to the IsPassphraseSet field.
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetIsPassphraseSet(v bool) {
	o.IsPassphraseSet = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfig) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *KubernetesEsxiVirtualMachineInfraConfig) SetResourcePool(v string) {
	o.ResourcePool = &v
}

func (o KubernetesEsxiVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesBaseVirtualMachineInfraConfig, errKubernetesBaseVirtualMachineInfraConfig := json.Marshal(o.KubernetesBaseVirtualMachineInfraConfig)
	if errKubernetesBaseVirtualMachineInfraConfig != nil {
		return []byte{}, errKubernetesBaseVirtualMachineInfraConfig
	}
	errKubernetesBaseVirtualMachineInfraConfig = json.Unmarshal([]byte(serializedKubernetesBaseVirtualMachineInfraConfig), &toSerialize)
	if errKubernetesBaseVirtualMachineInfraConfig != nil {
		return []byte{}, errKubernetesBaseVirtualMachineInfraConfig
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Datastore != nil {
		toSerialize["Datastore"] = o.Datastore
	}
	if o.IsPassphraseSet != nil {
		toSerialize["IsPassphraseSet"] = o.IsPassphraseSet
	}
	if o.Passphrase != nil {
		toSerialize["Passphrase"] = o.Passphrase
	}
	if o.ResourcePool != nil {
		toSerialize["ResourcePool"] = o.ResourcePool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesEsxiVirtualMachineInfraConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the vSphere cluster on which the virtual machines are created.
		Cluster *string `json:"Cluster,omitempty"`
		// Name of the datasore on which the virtual machine disks are created.
		Datastore *string `json:"Datastore,omitempty"`
		// Indicates whether the value of the 'passphrase' property has been set.
		IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
		// Passphrase for the vcenter user.
		Passphrase *string `json:"Passphrase,omitempty"`
		// Name of the vSphere resource pool on which the virtual machines are created.
		ResourcePool *string `json:"ResourcePool,omitempty"`
	}

	varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct := KubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesEsxiVirtualMachineInfraConfig := _KubernetesEsxiVirtualMachineInfraConfig{}
		varKubernetesEsxiVirtualMachineInfraConfig.ClassId = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.ClassId
		varKubernetesEsxiVirtualMachineInfraConfig.ObjectType = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesEsxiVirtualMachineInfraConfig.Cluster = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.Cluster
		varKubernetesEsxiVirtualMachineInfraConfig.Datastore = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.Datastore
		varKubernetesEsxiVirtualMachineInfraConfig.IsPassphraseSet = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.IsPassphraseSet
		varKubernetesEsxiVirtualMachineInfraConfig.Passphrase = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.Passphrase
		varKubernetesEsxiVirtualMachineInfraConfig.ResourcePool = varKubernetesEsxiVirtualMachineInfraConfigWithoutEmbeddedStruct.ResourcePool
		*o = KubernetesEsxiVirtualMachineInfraConfig(varKubernetesEsxiVirtualMachineInfraConfig)
	} else {
		return err
	}

	varKubernetesEsxiVirtualMachineInfraConfig := _KubernetesEsxiVirtualMachineInfraConfig{}

	err = json.Unmarshal(bytes, &varKubernetesEsxiVirtualMachineInfraConfig)
	if err == nil {
		o.KubernetesBaseVirtualMachineInfraConfig = varKubernetesEsxiVirtualMachineInfraConfig.KubernetesBaseVirtualMachineInfraConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Datastore")
		delete(additionalProperties, "IsPassphraseSet")
		delete(additionalProperties, "Passphrase")
		delete(additionalProperties, "ResourcePool")

		// remove fields from embedded structs
		reflectKubernetesBaseVirtualMachineInfraConfig := reflect.ValueOf(o.KubernetesBaseVirtualMachineInfraConfig)
		for i := 0; i < reflectKubernetesBaseVirtualMachineInfraConfig.Type().NumField(); i++ {
			t := reflectKubernetesBaseVirtualMachineInfraConfig.Type().Field(i)

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

type NullableKubernetesEsxiVirtualMachineInfraConfig struct {
	value *KubernetesEsxiVirtualMachineInfraConfig
	isSet bool
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfig) Get() *KubernetesEsxiVirtualMachineInfraConfig {
	return v.value
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfig) Set(val *KubernetesEsxiVirtualMachineInfraConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEsxiVirtualMachineInfraConfig(val *KubernetesEsxiVirtualMachineInfraConfig) *NullableKubernetesEsxiVirtualMachineInfraConfig {
	return &NullableKubernetesEsxiVirtualMachineInfraConfig{value: val, isSet: true}
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
