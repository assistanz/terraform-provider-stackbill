package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceActionsProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: actionInstanceObj.Create,
		ReadContext:   actionInstanceObj.Read,
		UpdateContext: actionInstanceObj.Update,
		DeleteContext: actionInstanceObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
