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

// HyperflexClusterAllOf Definition of the list of properties defined in 'hyperflex.Cluster', excluding properties defined in parent classes.
type HyperflexClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                        `json:"ObjectType"`
	AlarmSummary NullableHyperflexAlarmSummary `json:"AlarmSummary,omitempty"`
	// The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.
	CapacityRunway *int64 `json:"CapacityRunway,omitempty"`
	// The name of this HyperFlex cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The storage type of this cluster (All Flash or Hybrid).
	ClusterType *int64 `json:"ClusterType,omitempty"`
	// The unique identifier for this HyperFlex cluster.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// The number of compute nodes that belong to this cluster.
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty"`
	// The number of converged nodes that belong to this cluster.
	ConvergedNodeCount *int64 `json:"ConvergedNodeCount,omitempty"`
	// The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as 'NA' (not available). * `NA` - The deployment type of the HyperFlex cluster is not available. * `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * `Edge` - The deployment type of a HyperFlex cluster consisting of 2-4 standalone nodes.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.
	FltAggr *int64 `json:"FltAggr,omitempty"`
	// The HyperFlex Data Platform version of this cluster.
	HxVersion *string `json:"HxVersion,omitempty"`
	// The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version.
	HxdpBuildVersion *string `json:"HxdpBuildVersion,omitempty"`
	// The type of hypervisor running on this cluster. * `ESXi` - A Vmware ESXi hypervisor of any version. * `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The version of hypervisor running on this cluster.
	HypervisorVersion *string                  `json:"HypervisorVersion,omitempty"`
	Summary           NullableHyperflexSummary `json:"Summary,omitempty"`
	// The storage utilization percentage is computed based on total capacity and current capacity utilization.
	UtilizationPercentage *float32 `json:"UtilizationPercentage,omitempty"`
	// The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.
	UtilizationTrendPercentage *float32 `json:"UtilizationTrendPercentage,omitempty"`
	// The number of virtual machines present on this cluster.
	VmCount *int64 `json:"VmCount,omitempty"`
	// An array of relationships to hyperflexAlarm resources.
	Alarm  []HyperflexAlarmRelationship `json:"Alarm,omitempty"`
	Health *HyperflexHealthRelationship `json:"Health,omitempty"`
	// An array of relationships to hyperflexNode resources.
	Nodes            []HyperflexNodeRelationship          `json:"Nodes,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageHyperFlexStorageContainer resources.
	StorageContainers []StorageHyperFlexStorageContainerRelationship `json:"StorageContainers,omitempty"`
	// An array of relationships to storageHyperFlexVolume resources.
	Volumes              []StorageHyperFlexVolumeRelationship `json:"Volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterAllOf HyperflexClusterAllOf

// NewHyperflexClusterAllOf instantiates a new HyperflexClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterAllOf(classId string, objectType string) *HyperflexClusterAllOf {
	this := HyperflexClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewHyperflexClusterAllOfWithDefaults instantiates a new HyperflexClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterAllOfWithDefaults() *HyperflexClusterAllOf {
	this := HyperflexClusterAllOf{}
	var classId string = "hyperflex.Cluster"
	this.ClassId = classId
	var objectType string = "hyperflex.Cluster"
	this.ObjectType = objectType
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetAlarmSummary() HyperflexAlarmSummary {
	if o == nil || o.AlarmSummary.Get() == nil {
		var ret HyperflexAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetAlarmSummaryOk() (*HyperflexAlarmSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableHyperflexAlarmSummary and assigns it to the AlarmSummary field.
func (o *HyperflexClusterAllOf) SetAlarmSummary(v HyperflexAlarmSummary) {
	o.AlarmSummary.Set(&v)
}

// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *HyperflexClusterAllOf) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *HyperflexClusterAllOf) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetCapacityRunway() int64 {
	if o == nil || o.CapacityRunway == nil {
		var ret int64
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetCapacityRunwayOk() (*int64, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given int64 and assigns it to the CapacityRunway field.
func (o *HyperflexClusterAllOf) SetCapacityRunway(v int64) {
	o.CapacityRunway = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HyperflexClusterAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetClusterType() int64 {
	if o == nil || o.ClusterType == nil {
		var ret int64
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetClusterTypeOk() (*int64, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasClusterType() bool {
	if o != nil && o.ClusterType != nil {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given int64 and assigns it to the ClusterType field.
func (o *HyperflexClusterAllOf) SetClusterType(v int64) {
	o.ClusterType = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexClusterAllOf) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetComputeNodeCount returns the ComputeNodeCount field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetComputeNodeCount() int64 {
	if o == nil || o.ComputeNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ComputeNodeCount
}

// GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetComputeNodeCountOk() (*int64, bool) {
	if o == nil || o.ComputeNodeCount == nil {
		return nil, false
	}
	return o.ComputeNodeCount, true
}

// HasComputeNodeCount returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasComputeNodeCount() bool {
	if o != nil && o.ComputeNodeCount != nil {
		return true
	}

	return false
}

// SetComputeNodeCount gets a reference to the given int64 and assigns it to the ComputeNodeCount field.
func (o *HyperflexClusterAllOf) SetComputeNodeCount(v int64) {
	o.ComputeNodeCount = &v
}

// GetConvergedNodeCount returns the ConvergedNodeCount field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetConvergedNodeCount() int64 {
	if o == nil || o.ConvergedNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ConvergedNodeCount
}

// GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetConvergedNodeCountOk() (*int64, bool) {
	if o == nil || o.ConvergedNodeCount == nil {
		return nil, false
	}
	return o.ConvergedNodeCount, true
}

// HasConvergedNodeCount returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasConvergedNodeCount() bool {
	if o != nil && o.ConvergedNodeCount != nil {
		return true
	}

	return false
}

// SetConvergedNodeCount gets a reference to the given int64 and assigns it to the ConvergedNodeCount field.
func (o *HyperflexClusterAllOf) SetConvergedNodeCount(v int64) {
	o.ConvergedNodeCount = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *HyperflexClusterAllOf) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *HyperflexClusterAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetFltAggr returns the FltAggr field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetFltAggr() int64 {
	if o == nil || o.FltAggr == nil {
		var ret int64
		return ret
	}
	return *o.FltAggr
}

// GetFltAggrOk returns a tuple with the FltAggr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetFltAggrOk() (*int64, bool) {
	if o == nil || o.FltAggr == nil {
		return nil, false
	}
	return o.FltAggr, true
}

// HasFltAggr returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasFltAggr() bool {
	if o != nil && o.FltAggr != nil {
		return true
	}

	return false
}

// SetFltAggr gets a reference to the given int64 and assigns it to the FltAggr field.
func (o *HyperflexClusterAllOf) SetFltAggr(v int64) {
	o.FltAggr = &v
}

// GetHxVersion returns the HxVersion field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetHxVersion() string {
	if o == nil || o.HxVersion == nil {
		var ret string
		return ret
	}
	return *o.HxVersion
}

// GetHxVersionOk returns a tuple with the HxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetHxVersionOk() (*string, bool) {
	if o == nil || o.HxVersion == nil {
		return nil, false
	}
	return o.HxVersion, true
}

// HasHxVersion returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasHxVersion() bool {
	if o != nil && o.HxVersion != nil {
		return true
	}

	return false
}

// SetHxVersion gets a reference to the given string and assigns it to the HxVersion field.
func (o *HyperflexClusterAllOf) SetHxVersion(v string) {
	o.HxVersion = &v
}

// GetHxdpBuildVersion returns the HxdpBuildVersion field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetHxdpBuildVersion() string {
	if o == nil || o.HxdpBuildVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpBuildVersion
}

// GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetHxdpBuildVersionOk() (*string, bool) {
	if o == nil || o.HxdpBuildVersion == nil {
		return nil, false
	}
	return o.HxdpBuildVersion, true
}

// HasHxdpBuildVersion returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasHxdpBuildVersion() bool {
	if o != nil && o.HxdpBuildVersion != nil {
		return true
	}

	return false
}

// SetHxdpBuildVersion gets a reference to the given string and assigns it to the HxdpBuildVersion field.
func (o *HyperflexClusterAllOf) SetHxdpBuildVersion(v string) {
	o.HxdpBuildVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexClusterAllOf) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexClusterAllOf) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetSummary() HyperflexSummary {
	if o == nil || o.Summary.Get() == nil {
		var ret HyperflexSummary
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetSummaryOk() (*HyperflexSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableHyperflexSummary and assigns it to the Summary field.
func (o *HyperflexClusterAllOf) SetSummary(v HyperflexSummary) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *HyperflexClusterAllOf) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *HyperflexClusterAllOf) UnsetSummary() {
	o.Summary.Unset()
}

// GetUtilizationPercentage returns the UtilizationPercentage field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetUtilizationPercentage() float32 {
	if o == nil || o.UtilizationPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationPercentage
}

// GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetUtilizationPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationPercentage == nil {
		return nil, false
	}
	return o.UtilizationPercentage, true
}

// HasUtilizationPercentage returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasUtilizationPercentage() bool {
	if o != nil && o.UtilizationPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationPercentage gets a reference to the given float32 and assigns it to the UtilizationPercentage field.
func (o *HyperflexClusterAllOf) SetUtilizationPercentage(v float32) {
	o.UtilizationPercentage = &v
}

// GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetUtilizationTrendPercentage() float32 {
	if o == nil || o.UtilizationTrendPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationTrendPercentage
}

// GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetUtilizationTrendPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationTrendPercentage == nil {
		return nil, false
	}
	return o.UtilizationTrendPercentage, true
}

// HasUtilizationTrendPercentage returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasUtilizationTrendPercentage() bool {
	if o != nil && o.UtilizationTrendPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationTrendPercentage gets a reference to the given float32 and assigns it to the UtilizationTrendPercentage field.
func (o *HyperflexClusterAllOf) SetUtilizationTrendPercentage(v float32) {
	o.UtilizationTrendPercentage = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *HyperflexClusterAllOf) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetAlarm returns the Alarm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetAlarm() []HyperflexAlarmRelationship {
	if o == nil {
		var ret []HyperflexAlarmRelationship
		return ret
	}
	return o.Alarm
}

// GetAlarmOk returns a tuple with the Alarm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool) {
	if o == nil || o.Alarm == nil {
		return nil, false
	}
	return &o.Alarm, true
}

// HasAlarm returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasAlarm() bool {
	if o != nil && o.Alarm != nil {
		return true
	}

	return false
}

// SetAlarm gets a reference to the given []HyperflexAlarmRelationship and assigns it to the Alarm field.
func (o *HyperflexClusterAllOf) SetAlarm(v []HyperflexAlarmRelationship) {
	o.Alarm = v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetHealth() HyperflexHealthRelationship {
	if o == nil || o.Health == nil {
		var ret HyperflexHealthRelationship
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetHealthOk() (*HyperflexHealthRelationship, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given HyperflexHealthRelationship and assigns it to the Health field.
func (o *HyperflexClusterAllOf) SetHealth(v HyperflexHealthRelationship) {
	o.Health = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetNodes() []HyperflexNodeRelationship {
	if o == nil {
		var ret []HyperflexNodeRelationship
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetNodesOk() (*[]HyperflexNodeRelationship, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []HyperflexNodeRelationship and assigns it to the Nodes field.
func (o *HyperflexClusterAllOf) SetNodes(v []HyperflexNodeRelationship) {
	o.Nodes = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageContainers returns the StorageContainers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetStorageContainers() []StorageHyperFlexStorageContainerRelationship {
	if o == nil {
		var ret []StorageHyperFlexStorageContainerRelationship
		return ret
	}
	return o.StorageContainers
}

// GetStorageContainersOk returns a tuple with the StorageContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetStorageContainersOk() (*[]StorageHyperFlexStorageContainerRelationship, bool) {
	if o == nil || o.StorageContainers == nil {
		return nil, false
	}
	return &o.StorageContainers, true
}

// HasStorageContainers returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasStorageContainers() bool {
	if o != nil && o.StorageContainers != nil {
		return true
	}

	return false
}

// SetStorageContainers gets a reference to the given []StorageHyperFlexStorageContainerRelationship and assigns it to the StorageContainers field.
func (o *HyperflexClusterAllOf) SetStorageContainers(v []StorageHyperFlexStorageContainerRelationship) {
	o.StorageContainers = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterAllOf) GetVolumes() []StorageHyperFlexVolumeRelationship {
	if o == nil {
		var ret []StorageHyperFlexVolumeRelationship
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterAllOf) GetVolumesOk() (*[]StorageHyperFlexVolumeRelationship, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *HyperflexClusterAllOf) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []StorageHyperFlexVolumeRelationship and assigns it to the Volumes field.
func (o *HyperflexClusterAllOf) SetVolumes(v []StorageHyperFlexVolumeRelationship) {
	o.Volumes = v
}

func (o HyperflexClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ClusterType != nil {
		toSerialize["ClusterType"] = o.ClusterType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.ComputeNodeCount != nil {
		toSerialize["ComputeNodeCount"] = o.ComputeNodeCount
	}
	if o.ConvergedNodeCount != nil {
		toSerialize["ConvergedNodeCount"] = o.ConvergedNodeCount
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.FltAggr != nil {
		toSerialize["FltAggr"] = o.FltAggr
	}
	if o.HxVersion != nil {
		toSerialize["HxVersion"] = o.HxVersion
	}
	if o.HxdpBuildVersion != nil {
		toSerialize["HxdpBuildVersion"] = o.HxdpBuildVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.Summary.IsSet() {
		toSerialize["Summary"] = o.Summary.Get()
	}
	if o.UtilizationPercentage != nil {
		toSerialize["UtilizationPercentage"] = o.UtilizationPercentage
	}
	if o.UtilizationTrendPercentage != nil {
		toSerialize["UtilizationTrendPercentage"] = o.UtilizationTrendPercentage
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.Alarm != nil {
		toSerialize["Alarm"] = o.Alarm
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageContainers != nil {
		toSerialize["StorageContainers"] = o.StorageContainers
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterAllOf := _HyperflexClusterAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterAllOf); err == nil {
		*o = HyperflexClusterAllOf(varHyperflexClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "CapacityRunway")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ClusterType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "ComputeNodeCount")
		delete(additionalProperties, "ConvergedNodeCount")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "FltAggr")
		delete(additionalProperties, "HxVersion")
		delete(additionalProperties, "HxdpBuildVersion")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "UtilizationPercentage")
		delete(additionalProperties, "UtilizationTrendPercentage")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "Alarm")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "Nodes")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageContainers")
		delete(additionalProperties, "Volumes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterAllOf struct {
	value *HyperflexClusterAllOf
	isSet bool
}

func (v NullableHyperflexClusterAllOf) Get() *HyperflexClusterAllOf {
	return v.value
}

func (v *NullableHyperflexClusterAllOf) Set(val *HyperflexClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterAllOf(val *HyperflexClusterAllOf) *NullableHyperflexClusterAllOf {
	return &NullableHyperflexClusterAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
