package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceResizeProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Resize Api's",
		CreateContext: instanceResizeObj.Create,
		ReadContext:   instanceResizeObj.Read,
		UpdateContext: instanceResizeObj.Update,
		DeleteContext: instanceResizeObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_offering_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
