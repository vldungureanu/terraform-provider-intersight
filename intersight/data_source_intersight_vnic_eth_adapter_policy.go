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

func dataSourceVnicEthAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicEthAdapterPolicyRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"advanced_filter": {
				Description: "Enables advanced filtering on the interface.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"arfs_settings": {
				Description: "Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.",
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
						"enabled": {
							Description: "Status of Accelerated Receive Flow Steering on the virtual ethernet interface.",
							Type:        schema.TypeBool,
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"completion_queue_settings": {
				Description: "Completion Queue resource settings.",
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
						"nr_count": {
							Description: "The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each completion queue.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"interrupt_scaling": {
				Description: "Enables Interrupt Scaling on the interface.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"interrupt_settings": {
				Description: "Interrupt Settings for the virtual ethernet interface.",
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
						"coalescing_time": {
							Description: "The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"coalescing_type": {
							Description: "Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.\n* `MIN` - The system waits for the time specified in the Coalescing Time field before sending another interrupt event.\n* `IDLE` - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_count": {
							Description: "The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mode": {
							Description: "Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.\n* `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option.\n* `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts.\n* `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nvgre_settings": {
				Description: "Network Virtualization using Generic Routing Encapsulation Settings.",
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
						"enabled": {
							Description: "Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.",
							Type:        schema.TypeBool,
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"roce_settings": {
				Description: "Settings for RDMA over Converged Ethernet.",
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
						"class_of_service": {
							Description: "The Class of Service for RoCE on this virtual interface.\n* `5` - RDMA CoS Service Level 5.\n* `1` - RDMA CoS Service Level 1.\n* `2` - RDMA CoS Service Level 2.\n* `4` - RDMA CoS Service Level 4.\n* `6` - RDMA CoS Service Level 6.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"enabled": {
							Description: "If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"memory_regions": {
							Description: "The number of memory regions per adapter. Recommended value = integer power of 2.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"queue_pairs": {
							Description: "The number of queue pairs per adapter. Recommended value = integer power of 2.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"resource_groups": {
							Description: "The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"nr_version": {
							Description: "Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters.\n* `1` - RDMA over Converged Ethernet Protocol Version 1.\n* `2` - RDMA over Converged Ethernet Protocol Version 2.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"rss_hash_settings": {
				Description: "Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.",
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
						"ipv4_hash": {
							Description: "When enabled, the IPv4 address is used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"ipv6_ext_hash": {
							Description: "When enabled, the IPv6 extensions are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"ipv6_hash": {
							Description: "When enabled, the IPv6 address is used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"tcp_ipv4_hash": {
							Description: "When enabled, both the IPv4 address and TCP port number are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tcp_ipv6_ext_hash": {
							Description: "When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tcp_ipv6_hash": {
							Description: "When enabled, both the IPv6 address and TCP port number are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"udp_ipv4_hash": {
							Description: "When enabled, both the IPv4 address and UDP port number are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"udp_ipv6_hash": {
							Description: "When enabled, both the IPv6 address and UDP port number are used for traffic distribution.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"rss_settings": {
				Description: "Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"rx_queue_settings": {
				Description: "Receive Queue resouce settings.",
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
						"nr_count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
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
			"tcp_offload_settings": {
				Description: "The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.",
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
						"large_receive": {
							Description: "Enables the reassembly of segmented packets in hardware before sending them to the CPU.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"large_send": {
							Description: "Enables the CPU to send large packets to the hardware for segmentation.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"rx_checksum": {
							Description: "When enabled, the CPU sends all packet checksums to the hardware for validation.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tx_checksum": {
							Description: "When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"tx_queue_settings": {
				Description: "Transmit Queue resource settings.",
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
						"nr_count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"uplink_failback_timeout": {
				Description: "Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vxlan_settings": {
				Description: "Virtual Extensible LAN Protocol Settings.",
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
						"enabled": {
							Description: "Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface.",
							Type:        schema.TypeBool,
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
		},
	}
}

func dataSourceVnicEthAdapterPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicEthAdapterPolicy{}
	if v, ok := d.GetOk("advanced_filter"); ok {
		x := (v.(bool))
		o.SetAdvancedFilter(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("interrupt_scaling"); ok {
		x := (v.(bool))
		o.SetInterruptScaling(x)
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
	if v, ok := d.GetOk("rss_settings"); ok {
		x := (v.(bool))
		o.SetRssSettings(x)
	}
	if v, ok := d.GetOk("uplink_failback_timeout"); ok {
		x := int64(v.(int))
		o.SetUplinkFailbackTimeout(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicEthAdapterPolicy object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthAdapterPolicyList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching VnicEthAdapterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for VnicEthAdapterPolicy list: %s", err.Error())
	}
	var s = &models.VnicEthAdapterPolicyList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to VnicEthAdapterPolicy list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for VnicEthAdapterPolicy data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for VnicEthAdapterPolicy data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.VnicEthAdapterPolicy{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("advanced_filter", (s.GetAdvancedFilter())); err != nil {
				return diag.Errorf("error occurred while setting property AdvancedFilter: %s", err.Error())
			}

			if err := d.Set("arfs_settings", flattenMapVnicArfsSettings(s.GetArfsSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ArfsSettings: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("completion_queue_settings", flattenMapVnicCompletionQueueSettings(s.GetCompletionQueueSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property CompletionQueueSettings: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("interrupt_scaling", (s.GetInterruptScaling())); err != nil {
				return diag.Errorf("error occurred while setting property InterruptScaling: %s", err.Error())
			}

			if err := d.Set("interrupt_settings", flattenMapVnicEthInterruptSettings(s.GetInterruptSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InterruptSettings: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}

			if err := d.Set("nvgre_settings", flattenMapVnicNvgreSettings(s.GetNvgreSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NvgreSettings: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}

			if err := d.Set("roce_settings", flattenMapVnicRoceSettings(s.GetRoceSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RoceSettings: %s", err.Error())
			}

			if err := d.Set("rss_hash_settings", flattenMapVnicRssHashSettings(s.GetRssHashSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RssHashSettings: %s", err.Error())
			}
			if err := d.Set("rss_settings", (s.GetRssSettings())); err != nil {
				return diag.Errorf("error occurred while setting property RssSettings: %s", err.Error())
			}

			if err := d.Set("rx_queue_settings", flattenMapVnicEthRxQueueSettings(s.GetRxQueueSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RxQueueSettings: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("tcp_offload_settings", flattenMapVnicTcpOffloadSettings(s.GetTcpOffloadSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property TcpOffloadSettings: %s", err.Error())
			}

			if err := d.Set("tx_queue_settings", flattenMapVnicEthTxQueueSettings(s.GetTxQueueSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property TxQueueSettings: %s", err.Error())
			}
			if err := d.Set("uplink_failback_timeout", (s.GetUplinkFailbackTimeout())); err != nil {
				return diag.Errorf("error occurred while setting property UplinkFailbackTimeout: %s", err.Error())
			}

			if err := d.Set("vxlan_settings", flattenMapVnicVxlanSettings(s.GetVxlanSettings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property VxlanSettings: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
