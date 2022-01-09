package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceUpdateNameProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: instanceUpdateNameObj.Create,
		ReadContext:   instanceUpdateNameObj.Read,
		UpdateContext: instanceUpdateNameObj.Update,
		DeleteContext: instanceUpdateNameObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
