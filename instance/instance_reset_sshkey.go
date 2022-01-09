package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Resest ssh key Object
func NewInstanceResetSshKey() InstanceResetSshkeyI {
	return &instanceResetSshkey{}
}

//Instance Interfae
type InstanceResetSshkeyI interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceResetSshkey struct {
}

// Create Instance
// TODO - Documentation
func (irs *instanceResetSshkey) ResetSshkey(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	resetSshkeyRequest := instanceUtilsObj.GetResetSshKeyInstanceRequest(d)
	response, err := instanceApiObj.ResetSshKey(resetSshkeyRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Create Instance
// TODO - Documentation
func (irs *instanceResetSshkey) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance Reset sshkey initiated...!")
	response, err := irs.ResetSshkey(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance Reset sshkey successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (irs *instanceResetSshkey) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (irs *instanceResetSshkey) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	logs.Info("Instance Reset sshkey initiated...!")
	response, err := irs.ResetSshkey(ctx, d, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance Reset sshkey successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (irs *instanceResetSshkey) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
