package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiaapiDcnmCcoPost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiDcnmCcoPostRead,
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"post_date": {
				Description: "The date when this new release notice is posted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"post_detail": {
				Description: "Detail of this post including the content and the date it was posted.",
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
						"description": {
							Description: "Description of this new verison release post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"link": {
							Description: "Link of downloading the release file.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"release_note_link": {
							Description: "The link used to provide a gateway for customer to review the release note.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"release_note_link_title": {
							Description: "The link title used to provide a gateway for customer to review the release note.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"software_download_link": {
							Description: "The link used to provide a gateway for customer to download the release.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"software_download_link_title": {
							Description: "The link title used to provide a gateway for customer to download the release.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"title": {
							Description: "Title of the new verison release post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_version": {
							Description: "Version number Associate with this Post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"post_type": {
				Description: "The document type of this post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"postid": {
				Description: "Identificator of this inbox post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"revision": {
				Description: "Revision number of this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
		},
	}
}

func dataSourceNiaapiDcnmCcoPostRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiDcnmCcoPost{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("post_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetPostDate(x)
	}
	if v, ok := d.GetOk("post_type"); ok {
		x := (v.(string))
		o.SetPostType(x)
	}
	if v, ok := d.GetOk("postid"); ok {
		x := (v.(string))
		o.SetPostid(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiDcnmCcoPost object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmCcoPostList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching NiaapiDcnmCcoPost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NiaapiDcnmCcoPost list: %s", err.Error())
	}
	var s = &models.NiaapiDcnmCcoPostList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NiaapiDcnmCcoPost list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NiaapiDcnmCcoPost data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NiaapiDcnmCcoPost data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiaapiDcnmCcoPost{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("post_date", (s.GetPostDate()).String()); err != nil {
				return diag.Errorf("error occurred while setting property PostDate: %s", err.Error())
			}

			if err := d.Set("post_detail", flattenMapNiaapiNewReleaseDetail(s.GetPostDetail(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PostDetail: %s", err.Error())
			}
			if err := d.Set("post_type", (s.GetPostType())); err != nil {
				return diag.Errorf("error occurred while setting property PostType: %s", err.Error())
			}
			if err := d.Set("postid", (s.GetPostid())); err != nil {
				return diag.Errorf("error occurred while setting property Postid: %s", err.Error())
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return diag.Errorf("error occurred while setting property Revision: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
