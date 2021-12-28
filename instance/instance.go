package instance

import (
	"context"
	"encoding/json"
	"errors"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
	"terraform-provider-stackbill/http"

	"github.com/go-playground/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

var (
	httpClient http.HttpClient = http.NewHttpClient()
)

// Instance request object
type InstanceRequest struct {
	ComputeOfferingUuid string `json:"computeOfferingUuid" validate:"required"`
	CpuCore             string `json:"cpuCore"`
	DiskSize            int    `json:"diskSize"`
	HypervisorName      string `json:"hypervisorName"`
	Memory              string `json:"memory"`
	Name                string `json:"name" validate:"required"`
	NetworkUuid         string `json:"networkUuid" validate:"required"`
	SecuritygroupName   string `json:"securitygroupName"`
	SshKeyName          string `json:"sshKeyName"`
	StorageOfferingUuid string `json:"storageOfferingUuid"`
	TemplateUuid        string `json:"templateUuid" validate:"required"`
	ZoneUuid            string `json:"zoneUuid" validate:"required"`
}

// Instance Object
func NewInstance() IInstance {
	return &instance{}
}

//Instance Interfae
type IInstance interface {
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
	// Meta information
	m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	logs.Info("Getting the api and secret keys...!")
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	// Check the api key and secret key whether not empty
	if apiKey == "" || secretKey == "" {
		logs.Info("ApiKey or SecretKey does not exist...!")
		return diag.FromErr(errors.New("ApiKey or SecretKey does not exist...!"))
	}
	// Set all the values
	logs.Info("Creating instance request object...!")
	var instanceReq InstanceRequest
	instanceReq.ComputeOfferingUuid = d.Get("compute_offering_uuid").(string)
	instanceReq.CpuCore = d.Get("cpu_core").(string)
	instanceReq.DiskSize = d.Get("disk_size").(int)
	instanceReq.HypervisorName = d.Get("hypervisor_name").(string)
	instanceReq.Memory = d.Get("memory").(string)
	instanceReq.Name = d.Get("name").(string)
	instanceReq.NetworkUuid = d.Get("network_uuid").(string)
	instanceReq.SecuritygroupName = d.Get("security_group_name").(string)
	instanceReq.SshKeyName = d.Get("ssh_key_name").(string)
	instanceReq.StorageOfferingUuid = d.Get("storage_offering_uuid").(string)
	instanceReq.TemplateUuid = d.Get("template_uuid").(string)
	instanceReq.ZoneUuid = d.Get("zone_uuid").(string)
	logs.Info("Created instance request object...!")
	// Cheking validation
	logs.Info("Validating the data...!")
	validate := validator.New()
	err := validate.Struct(instanceReq)
	if err != nil {
		logs.Info(err.Error())
		return diag.FromErr(err)
	}
	logs.Info("Validation successful...!")
	logs.Info("Instance create request initiated...!")
	endPoint := api.GetInstanceCreateApi()
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, instanceReq)
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
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instance) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
