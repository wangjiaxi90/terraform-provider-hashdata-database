package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "name of the warehouse or catalog",
			},
		},
		ResourcesMap: map[string]*schema.Resource{ //CRUD
			"hashdata_warehouse": resourceHashDataWareHouse(),
			"hashdata_catalog":   resourceHashDataWareHouse(),// TODO
			"hashdata_computing": resourceHashDataWareHouse(),// TODO
		},
	}
}
