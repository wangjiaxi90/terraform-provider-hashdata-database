package provider
import (
	"github.com/hashicorp/terraform/helper/schema"
	"internal/provider/cloudmgr/model_core_create_warehouse_request"
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
			resourceName: {
				Type:     schema.TypeString,
				Optional: true,
			},
			resourceDescription: {
				Type:     schema.TypeString,
				Optional: true,
			},
			resourceEipBandwidth: {
				Type:     schema.TypeInt,
				Required: true,
			},
			resourceEipBillMode: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "bandwidth",
				ValidateFunc: withinArrayString("traffic", "bandwidth"),
			},
			resourceEipNeedIcp: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: withinArrayInt(0, 1),
			},
			resourceTagIds:   tagIdsSchema(),
			resourceTagNames: tagNamesSchema(),
			resourceEipAddr: {
				Type:     schema.TypeString,
				Computed: true,
			},
			resourceEipResource: {
				Type:         schema.TypeMap,
				Computed:     true,
				ComputedWhen: []string{"id"},
			},
		},
	}
}

func resourceHashDataWareHouseCreate(d *schema.ResourceData, meta interface{}) error {
	req := NewCoreCreateWarehouseRequest()


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

