package stackbill

import (
	"context"
	"terraform-provider-stackbill/auth"
	"terraform-provider-stackbill/instance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	authClient auth.IAuthClient = auth.NewAuthClient()
)

// Provider Interface
type IProvider interface {
	Provider() *schema.Provider
	configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics)
}

// Provider struct for implementation
type provider struct{}

// Provider Object
func NewProvider() IProvider {
	return &provider{}
}

// Provider
func (p *provider) Provider() *schema.Provider {
	pro := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"stackbill_instance": instance.InstanceProvider(),
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
