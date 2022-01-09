package stackbill

import (
	"context"
	"terraform-provider-stackbill/auth"
	"terraform-provider-stackbill/computeoffering"
	"terraform-provider-stackbill/instance"
	"terraform-provider-stackbill/iso"
	"terraform-provider-stackbill/network"
	"terraform-provider-stackbill/sshkey"
	"terraform-provider-stackbill/storageoffering"
	"terraform-provider-stackbill/template"
	"terraform-provider-stackbill/zone"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	authClient auth.IAuthClient = auth.NewAuthClient()
)

// Provider Interface
type ProviderI interface {
	Provider() *schema.Provider
	configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics)
}

// Provider struct for implementation
type provider struct{}

// Provider Object
func NewProvider() ProviderI {
	return &provider{}
}

// Provider
func (p *provider) Provider() *schema.Provider {
	pro := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"stackbill_instance":              instance.InstanceProvider(),
			"stackbill_instance_actions":      instance.InstanceActionsProvider(),
			"stackbill_instance_update_name":  instance.InstanceUpdateNameProvider(),
			"stackbill_instance_resize":       instance.InstanceResizeProvider(),
			"stackbill_instance_reset_sshkey": instance.InstanceResetSshKeyProvider(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stackbill_compute_offerings_list": computeoffering.ComputeOfferingListProvider(),
			"stackbill_instance_list":          instance.InstanceListProvider(),
			"stackbill_network_list":           network.NetworkListProvider(),
			"stackbill_storage_offerings_list": storageoffering.StorageOfferingListProvider(),
			"stackbill_zone_list":              zone.ZoneListProvider(),
			"stackbill_template_list":          template.TemplateListProvider(),
			"stackbill_sshkey_list":            sshkey.SshkeyListProvider(),
			"stackbill_iso_list":               iso.IsoListProvider(),
		},
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stackbill Api key",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Stackbill Secret Key",
			},
		},
	}
	pro.ConfigureContextFunc = p.configure
	return pro
}

// Configure function
func (p *provider) configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	secretKey := d.Get("secret_key").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := authClient.New(&apiKey, &secretKey)
	return client, diags
}
