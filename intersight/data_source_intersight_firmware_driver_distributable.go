package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirmwareDriverDistributable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareDriverDistributableRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bundle_type": {
				Description: "The bundle type of the image, as published on cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"catalog": {
				Description: "A reference to a softwarerepositoryCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"category": {
				Description: "The device type on which the driver is installable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"component_meta": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"component_label": {
							Description: "The name of the component in the compressed HSU bundle.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"component_type": {
							Description: "The type of component image within the distributable.\n* `ALL` - This represents all the components.\n* `ALL,HDD` - This represents all the components plus the HDDs.\n* `None` - This represents none of the components.\n* `NXOS` - This represents NXOS components.\n* `IOM` - This represents IOM components.\n* `PSU` - This represents PSU components.\n* `CIMC` - This represents CIMC components.\n* `BIOS` - This represents BIOS components.\n* `PCIE` - This represents PCIE components.\n* `Drive` - This represents Storage components.\n* `DIMM` - This represents DIMM components.\n* `BoardController` - This represents Board Controller components.\n* `StorageController` - This represents Storage Controller components.\n* `HBA` - This represents HBA components.\n* `GPU` - This represents GPU components.\n* `SasExpander` - This represents SasExpander components.\n* `MSwitch` - This represents mSwitch components.\n* `CMC` - This represents CMC components.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": {
							Description: "This shows the description of component image within the distributable.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"disruption": {
							Description: "The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.\n* `None` - Indicates that the component did not receive a disruption request.\n* `HostReboot` - Indicates that the component received a host reboot request.\n* `EndpointReboot` - Indicates that the component received an end point reboot request.\n* `ManualPowerCycle` - Indicates that the component received a manual power cycle request.\n* `AutomaticPowerCycle` - Indicates that the component received an automatic power cycle request.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"image_path": {
							Description: "This shows the path of component image within the distributable.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_oob_supported": {
							Description: "If set, the component can be updated through out-of-band management, else, is updated through host service utility boot.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"model": {
							Description: "The model of the component image in the distributable.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"oob_manageability": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"packed_version": {
							Description: "The image version of components packaged in the distributable.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"redfish_url": {
							Description: "The redfish target for each component.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vendor": {
							Description: "The version of component image in the distributable.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"directory": {
				Description: "Indicates in which directory path this driver will be added.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"distributable_metas": {
				Description: "An array of relationships to firmwareDistributableMeta resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"download_count": {
				Description: "The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "The unique identifier for an image in a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"import_action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.\n* `None` - No action should be taken on the imported file.\n* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.\n* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.\n* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.\n* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.\n* `PreCache` - Cache the file into the Intersight Appliance.\n* `Cancel` - The cancel import process for the file into the repository.\n* `Extract` - The action to extract the file in the external repository.\n* `Evict` - Evict the cached file from the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_state": {
				Description: "The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.\n* `ReadyForImport` - The image is ready to be imported into the repository.\n* `Importing` - The image is being imported into the repository.\n* `Imported` - The image has been extracted and imported into the repository.\n* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.\n* `Extracting` - Indicates that the image is being extracted into the repository.\n* `Extracted` - Indicates that the image has been extracted into the repository.\n* `Failed` - The image import from an external source to the repository has failed.\n* `MetaOnly` - The image is present in an external repository.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.\n* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.\n* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.\n* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"md5e_tag": {
				Description: "The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"md5sum": {
				Description: "The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mdfid": {
				Description: "The mdfid of the image provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "The endpoint model for which this firmware image is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the file. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"osname": {
				Description: "The operating system name to which this driver is compatible.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"osversion": {
				Description: "OS Version. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"platform_type": {
				Description: "The platform type of the image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"recommended_build": {
				Description: "The build which is recommended by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"release": {
				Description: "A reference to a softwarerepositoryRelease resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"release_notes_url": {
				Description: "The url for the release notes of this image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sha512sum": {
				Description: "The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"software_advisory_url": {
				Description: "The software advisory, if any, provided by the vendor for this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_type_id": {
				Description: "The software type id provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_source": {
				Description: "Location of the file in an external repository.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"supported_models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"vendor": {
				Description: "The vendor or publisher of this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Vendor provided version for the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceFirmwareDriverDistributableRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareDriverDistributable{}
	if v, ok := d.GetOk("bundle_type"); ok {
		x := (v.(string))
		o.SetBundleType(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("directory"); ok {
		x := (v.(string))
		o.SetDirectory(x)
	}
	if v, ok := d.GetOk("download_count"); ok {
		x := int64(v.(int))
		o.SetDownloadCount(x)
	}
	if v, ok := d.GetOk("guid"); ok {
		x := (v.(string))
		o.SetGuid(x)
	}
	if v, ok := d.GetOk("import_action"); ok {
		x := (v.(string))
		o.SetImportAction(x)
	}
	if v, ok := d.GetOk("import_state"); ok {
		x := (v.(string))
		o.SetImportState(x)
	}
	if v, ok := d.GetOk("md5e_tag"); ok {
		x := (v.(string))
		o.SetMd5eTag(x)
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}
	if v, ok := d.GetOk("mdfid"); ok {
		x := (v.(string))
		o.SetMdfid(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("osname"); ok {
		x := (v.(string))
		o.SetOsname(x)
	}
	if v, ok := d.GetOk("osversion"); ok {
		x := (v.(string))
		o.SetOsversion(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("recommended_build"); ok {
		x := (v.(string))
		o.SetRecommendedBuild(x)
	}
	if v, ok := d.GetOk("release_notes_url"); ok {
		x := (v.(string))
		o.SetReleaseNotesUrl(x)
	}
	if v, ok := d.GetOk("sha512sum"); ok {
		x := (v.(string))
		o.SetSha512sum(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("software_advisory_url"); ok {
		x := (v.(string))
		o.SetSoftwareAdvisoryUrl(x)
	}
	if v, ok := d.GetOk("software_type_id"); ok {
		x := (v.(string))
		o.SetSoftwareTypeId(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareDriverDistributable object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareDriverDistributableList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching FirmwareDriverDistributable: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for FirmwareDriverDistributable list: %s", err.Error())
	}
	var s = &models.FirmwareDriverDistributableList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to FirmwareDriverDistributable list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for FirmwareDriverDistributable data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for FirmwareDriverDistributable data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.FirmwareDriverDistributable{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("bundle_type", (s.GetBundleType())); err != nil {
				return diag.Errorf("error occurred while setting property BundleType: %s", err.Error())
			}

			if err := d.Set("catalog", flattenMapSoftwarerepositoryCatalogRelationship(s.GetCatalog(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Catalog: %s", err.Error())
			}
			if err := d.Set("category", (s.GetCategory())); err != nil {
				return diag.Errorf("error occurred while setting property Category: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("component_meta", flattenListFirmwareComponentMeta(s.GetComponentMeta(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ComponentMeta: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("directory", (s.GetDirectory())); err != nil {
				return diag.Errorf("error occurred while setting property Directory: %s", err.Error())
			}

			if err := d.Set("distributable_metas", flattenListFirmwareDistributableMetaRelationship(s.GetDistributableMetas(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DistributableMetas: %s", err.Error())
			}
			if err := d.Set("download_count", (s.GetDownloadCount())); err != nil {
				return diag.Errorf("error occurred while setting property DownloadCount: %s", err.Error())
			}
			if err := d.Set("guid", (s.GetGuid())); err != nil {
				return diag.Errorf("error occurred while setting property Guid: %s", err.Error())
			}
			if err := d.Set("import_action", (s.GetImportAction())); err != nil {
				return diag.Errorf("error occurred while setting property ImportAction: %s", err.Error())
			}
			if err := d.Set("import_state", (s.GetImportState())); err != nil {
				return diag.Errorf("error occurred while setting property ImportState: %s", err.Error())
			}
			if err := d.Set("md5e_tag", (s.GetMd5eTag())); err != nil {
				return diag.Errorf("error occurred while setting property Md5eTag: %s", err.Error())
			}
			if err := d.Set("md5sum", (s.GetMd5sum())); err != nil {
				return diag.Errorf("error occurred while setting property Md5sum: %s", err.Error())
			}
			if err := d.Set("mdfid", (s.GetMdfid())); err != nil {
				return diag.Errorf("error occurred while setting property Mdfid: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("osname", (s.GetOsname())); err != nil {
				return diag.Errorf("error occurred while setting property Osname: %s", err.Error())
			}
			if err := d.Set("osversion", (s.GetOsversion())); err != nil {
				return diag.Errorf("error occurred while setting property Osversion: %s", err.Error())
			}
			if err := d.Set("platform_type", (s.GetPlatformType())); err != nil {
				return diag.Errorf("error occurred while setting property PlatformType: %s", err.Error())
			}
			if err := d.Set("recommended_build", (s.GetRecommendedBuild())); err != nil {
				return diag.Errorf("error occurred while setting property RecommendedBuild: %s", err.Error())
			}

			if err := d.Set("release", flattenMapSoftwarerepositoryReleaseRelationship(s.GetRelease(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Release: %s", err.Error())
			}
			if err := d.Set("release_notes_url", (s.GetReleaseNotesUrl())); err != nil {
				return diag.Errorf("error occurred while setting property ReleaseNotesUrl: %s", err.Error())
			}
			if err := d.Set("sha512sum", (s.GetSha512sum())); err != nil {
				return diag.Errorf("error occurred while setting property Sha512sum: %s", err.Error())
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return diag.Errorf("error occurred while setting property Size: %s", err.Error())
			}
			if err := d.Set("software_advisory_url", (s.GetSoftwareAdvisoryUrl())); err != nil {
				return diag.Errorf("error occurred while setting property SoftwareAdvisoryUrl: %s", err.Error())
			}
			if err := d.Set("software_type_id", (s.GetSoftwareTypeId())); err != nil {
				return diag.Errorf("error occurred while setting property SoftwareTypeId: %s", err.Error())
			}

			if err := d.Set("nr_source", flattenMapSoftwarerepositoryFileServer(s.GetSource(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Source: %s", err.Error())
			}
			if err := d.Set("supported_models", (s.GetSupportedModels())); err != nil {
				return diag.Errorf("error occurred while setting property SupportedModels: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return diag.Errorf("error occurred while setting property Vendor: %s", err.Error())
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return diag.Errorf("error occurred while setting property Version: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
