package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

//Catalog string
//ChargeChannel string
//ChargeType string
//Description string
//ExpiredAt string
//Extras | Pointer to [**CoreCreateServiceIaasExtraRequest**](CoreCreateServiceIaasExtraRequest.md) |  | [optional]
//Features | Pointer to [**CoreCreateServiceFeatureRequest**](CoreCreateServiceFeatureRequest.md) |  | [optional]
//Master | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]
//Metadata map[string]interface{}
//Name string
//Oss | Pointer to [**CoreCreateServiceOssZoneRequest**](CoreCreateServiceOssZoneRequest.md) |  | [optional]
//Segment | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]
//Standby | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]

const (
	Catalog       = "catalog"
	ChargeChannel = "charge_channel"
	ChargeType    = "charge_type"
	Description   = "description"
	ExpiredAt     = "expired_at"
	Extras        = "extras"
	Features      = "features"
	Master        = "master"
	Metadata      = "metadata"
	Name          = "name"
	Oss           = "oss"
	Segment       = "segment"
	Standby       = "standby"
)

func resourceHashDataWareHouse() *schema.Resource {
	return &schema.Resource{
		Create: resourceHashDataWareHouseCreate,
		Read:   resourceHashDataWareHouseRead,
		Update: resourceHashDataWareHouseUpdate,
		Delete: resourceHashDataWareHouseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			Catalog: {
				Type:     schema.TypeString,
				Optional: true,
			},
			ChargeChannel: {
				Type:     schema.TypeString,
				Optional: true,
			}, ChargeType: {
				Type:     schema.TypeString,
				Optional: true,
			}, Description: {
				Type:     schema.TypeString,
				Optional: true,
			}, ExpiredAt: {
				Type:     schema.TypeString,
				Optional: true,
			}, Extras: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			}, Features: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			}, Master: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			}, Metadata: {
				Type:     schema.TypeMap, //TODO 应该是map
				Optional: true,
			}, Name: {
				Type:     schema.TypeString,
				Optional: true,
			}, Oss: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			}, Segment: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			}, Standby: {
				Type:     schema.TypeString, //TODO
				Optional: true,
			},
		},
	}
}

func resourceHashDataWareHouseCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("====================")
	fmt.Println("start create warehouse")
	fmt.Println(&d)
	fmt.Println(meta)
	fmt.Println("end create warehouse")
	fmt.Println("====================")
	//req := cloudmgr.NewCoreCreateWarehouseRequest()
	//master := cloudmgr.NewCoreCreateServiceComponentRequest()
	//master.Iaas = &cloudmgr.CloudmgrcoreIaasResource{
	//	Count:        d.Get("count").(*int32), //TODO 这里不太对
	//	InstanceType: d.Get("instance_type").(*string),
	//	VolumeType:   d.Get("volume_type").(*string),
	//	VolumeSize:   d.Get("volume_size").(*int32),
	//	Image:        d.Get("image").(*string),
	//	Zone:         d.Get("zone").(*string),
	//}
	//
	//segment := cloudmgr.NewCoreCreateServiceComponentRequestWithDefaults()
	//segment.Iaas = &cloudmgr.CloudmgrcoreIaasResource{
	//	Count:        d.Get("count").(*int32), //TODO 这里不太对
	//	InstanceType: d.Get("instance_type").(*string),
	//	VolumeType:   d.Get("volume_type").(*string),
	//	VolumeSize:   d.Get("volume_size").(*int32),
	//	Image:        d.Get("image").(*string),
	//	Zone:         d.Get("zone").(*string),
	//}
	//
	//extra := cloudmgr.NewCoreCreateServiceIaasExtraRequest()
	//extra.Vpc = d.Get("vpc").(*string)
	//extra.Subnet = d.Get("subnet").(*string)
	//extra.Keypair = d.Get("keypair").(*string)
	//
	//var metadata map[string]string
	//metadata["default_database"] = d.Get("defaultDatabase").(string) //TODO 这里不太对
	//metadata["default_user"] = d.Get("defaultUser").(string)
	//metadata["default_password"] = d.Get("defaultPassword").(string)
	//metadata["logic_part"] = d.Get("logicPart").(string)

	return nil
	//clt := meta.(*QingCloudClient).eip
	//input := new(qc.AllocateEIPsInput)
	//input.Bandwidth = qc.Int(d.Get(resourceEipBandwidth).(int))
	//input.BillingMode = qc.String(d.Get(resourceEipBillMode).(string))
	//input.NeedICP = qc.Int(d.Get(resourceEipNeedIcp).(int))
	//input.Count = qc.Int(1)
	//input.EIPName, _ = getNamePointer(d)
	//var output *qc.AllocateEIPsOutput
	//var err error
	//simpleRetry(func() error {
	//	output, err = clt.AllocateEIPs(input)
	//	return isServerBusy(err)
	//})
	//if err != nil {
	//	return err
	//}
	//d.SetId(qc.StringValue(output.EIPs[0]))
	//if _, err := EIPTransitionStateRefresh(clt, d.Id()); err != nil {
	//	return nil
	//}
	//return resourceQingcloudEipUpdate(d, meta)
}

func resourceHashDataWareHouseRead(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("====================")
	fmt.Println("start read warehouse")
	fmt.Println(&d)
	fmt.Println(meta)
	fmt.Println("end read warehouse")
	fmt.Println("====================")
	return nil
}

func resourceHashDataWareHouseUpdate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("====================")
	fmt.Println("start update warehouse")
	fmt.Println(&d)
	fmt.Println(meta)
	fmt.Println("end update warehouse")
	fmt.Println("====================")
	return nil
}

func resourceHashDataWareHouseDelete(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("====================")
	fmt.Println("start delete warehouse")
	fmt.Println(&d)
	fmt.Println(meta)
	fmt.Println("end delete warehouse")
	fmt.Println("====================")
	return nil
}
