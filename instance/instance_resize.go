package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Resize Object
func NewInstanceResize() InstanceResizeI {
	return &instanceResize{}
}

// Resize Interfae
type InstanceResizeI interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Resize Object
type instanceResize struct {
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResize) Resize(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	resizeRequest := instanceUtilsObj.GetResizeInstanceRequest(d)
	response, err := instanceApiObj.ResizeInstance(resizeRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResize) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance resize initiated...!")
	response, err := ir.Resize(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance resize successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResize) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResize) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance resize initiated...!")
	response, err := ir.Resize(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance resize successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (ir *instanceResize) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
