package provider

import (
	"context"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//func init() {
//	// Set descriptions to support markdown syntax, this will be used in document generation
//	// and the language server.
//	schema.DescriptionKind = schema.StringMarkdown
//
//	// Customize the content of descriptions when output. For example you can add defaults on
//	// to the exported descriptions if present.
//	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
//	// 	desc := s.Description
//	// 	if s.Default != nil {
//	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
//	// 	}
//	// 	return strings.TrimSpace(desc)
//	// }
//}

//func New(version string) func() *schema.Provider {
//	return func() *schema.Provider {
//		p := &schema.Provider{
//			DataSourcesMap: map[string]*schema.Resource{
//				"scaffolding_data_source": dataSourceScaffolding(),
//			},
//			ResourcesMap: map[string]*schema.Resource{
//				"scaffolding_resource": resourceScaffolding(),
//			},
//		}
//
//		p.ConfigureContextFunc = configure(version, p)
//
//		return p
//	}
//}
Catalog** | Pointer to **string** |  | [optional]
ChargeChannel** | Pointwarehouseer to **string** |  | [optional]
ChargeType** | Pointer to **string** |  | [optional]
Description** | Pointer to **string** |  | [optional]
ExpiredAt** | Pointer to **string** |  | [optional]
Extras** | Pointer to [**CoreCreateServiceIaasExtraRequest**](CoreCreateServiceIaasExtraRequest.md) |  | [optional]
Features** | Pointer to [**CoreCreateServiceFeatureRequest**](CoreCreateServiceFeatureRequest.md) |  | [optional]
Master** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]
Metadata** | Pointer to **map[string]interface{}** |  | [optional]
Name** | Pointer to **string** |  | [optional]
Oss** | Pointer to [**CoreCreateServiceOssZoneRequest**](CoreCreateServiceOssZoneRequest.md) |  | [optional]
Segment** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]
Standby** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional]




func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"catalog": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["access_key"],
			},
			"ChargeChannel": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["secret_key"],
			},
			"ChargeType": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["zone"],
			},
			"Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "ExpiredAt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			}, "Description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["endpoint"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"qingcloud_vpn_cert": dataSourceQingcloudVpnCert(), //Only Read
		},
		ResourcesMap: map[string]*schema.Resource{ //CRUD
			"hashdata_warehouse": ,
			"hashdata_catalog":,
			"hashdata_computing",
			//"qingcloud_eip":                   resourceQingcloudEip(),
			//"qingcloud_keypair":               resourceQingcloudKeypair(),
			//"qingcloud_security_group":        resourceQingcloudSecurityGroup(),
			//"qingcloud_security_group_rule":   resourceQingcloudSecurityGroupRule(),
			//"qingcloud_vxnet":                 resourceQingcloudVxnet(),
			//"qingcloud_vpc":                   resourceQingcloudVpc(),
			//"qingcloud_instance":              resourceQingcloudInstance(),
			//"qingcloud_volume":                resourceQingcloudVolume(),
			//"qingcloud_tag":                   resourceQingcloudTag(),
			//"qingcloud_vpc_static":            resourceQingcloudVpcStatic(),
			//"qingcloud_loadbalancer":          resourceQingcloudLoadBalancer(),
			//"qingcloud_loadbalancer_listener": resourceQingcloudLoadBalancerListener(),
			//"qingcloud_loadbalancer_backend":  resourceQingcloudLoadBalancerBackend(),
			//"qingcloud_server_certificate":    resourceQingcloudServerCertificate(),
		},
		// ConfigureFunc: providerConfigure, 默认配置
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent

		return &apiClient{}, nil
	}
}
