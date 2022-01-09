package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Update name Object
func NewInstanceUpdaeName() InstanceUpdateNameI {
	return &instanceUpdateName{}
}

//Instance Interfae
type InstanceUpdateNameI interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceUpdateName struct {
}

// Create Instance
// TODO - Documentation
func (iu *instanceUpdateName) UpdateName(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	updateNameRequest := instanceUtilsObj.GetUpdateInstanceNameRequest(d)
	response, err := instanceApiObj.UpdateInstanceName(updateNameRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Create Instance
// TODO - Documentation
func (iu *instanceUpdateName) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Update instance name initiated...!")
	response, err := iu.UpdateName(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Update instance name successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (iu *instanceUpdateName) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (iu *instanceUpdateName) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Update instance name initiated...!")
	response, err := iu.UpdateName(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Update instance name successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (iu *instanceUpdateName) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
