package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	newActionInstance IInstanceActions = NewInstanceActions()
)

func InstanceActionsProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: newActionInstance.Create,
		ReadContext:   newActionInstance.Read,
		UpdateContext: newActionInstance.Update,
		DeleteContext: newActionInstance.Delete,
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
