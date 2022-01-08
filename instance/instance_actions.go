package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Object
func NewInstanceActions() InstanceActionsI {
	return &instanceActions{}
}

//Instance Interfae
type InstanceActionsI interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceActions struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceActions) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance " + action + " initiated...!")
	actionRequest := instanceUtilsObj.GetInstanceActionRequest(d)
	response, err := instanceApiObj.InstanceActions(actionRequest, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (vs *instanceActions) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (vs *instanceActions) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance " + action + " initiated...!")
	actionRequest := instanceUtilsObj.GetInstanceActionRequest(d)
	response, err := instanceApiObj.InstanceActions(actionRequest, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instanceActions) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
