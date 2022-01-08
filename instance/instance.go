package instance

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Object
func NewInstance() InstanceI {
	return &instance{}
}

//Instance Interfae
type InstanceI interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instance struct {
}

// Create Instance
// TODO - Documentation
func (vs *instance) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Set all the values
	logs.Info("Creating instance request object...!")
	instanceRequest := instanceUtilsObj.GetCreateInstanceRequest(d)
	logs.Info("Created instance request object...!")
	logs.Info("Instance create request initiated...!")
	response, err := instanceApiObj.CreateInstance(instanceRequest, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	instanceRes := resJson["listInstanceResponse"].([]interface{})
	if len(instanceRes) > 0 {
		instanceObj := instanceRes[0].(map[string]interface{})
		uuid := instanceObj["uuid"].(string)
		// Update the state
		d.SetId(uuid)
	}
	logs.Info("Instance creation successful...!")
	return diags
}

// Read Instance
// TODO - Documentation
func (vs *instance) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (vs *instance) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Check any changes occured in the name
	if d.HasChange("name") {
		logs.Info("Update instance name initiated...!")
		updateNameRequest := instanceUtilsObj.GetUpdateInstanceNameRequest(d)
		response, err := instanceApiObj.UpdateInstanceName(updateNameRequest, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		logs.Info(response)
		logs.Info("Update instance name successful...!")
	}
	// Check any changes occured in the compute offering
	if d.HasChange("compute_offering_uuid") {
		logs.Info("Instance resize initiated...!")
		resizeRequest := instanceUtilsObj.GetResizeInstanceRequest(d)
		response, err := instanceApiObj.ResizeInstance(resizeRequest, meta)
		if err != nil {
			logs.Info(err.Error())
			return diag.FromErr(err)
		}
		logs.Info(response)
		logs.Info("Instance resize successful...!")
	}
	// Check any changes occured in the compute offering
	if d.HasChange("storage_offering_uuid") {
		logs.Info("Changes ----- Changes " + d.Get("storage_offering_uuid").(string))
	}
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instance) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	logs.Info("Instance delete request initiated...!")
	uuid := d.Id()
	response, err := instanceApiObj.DeleteInstance(uuid, meta)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info(response)
	logs.Info("Instance resize successful...!")
	d.SetId("")
	return diags
}
