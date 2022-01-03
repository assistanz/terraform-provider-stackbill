package instance

import (
	"context"
	"errors"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// var (
// 	httpClient http.HttpClient = http.NewHttpClient()
// )

const (
	START = "Start"
	STOP  = "Stop"
)

// Instance Object
func NewInstanceActions() IInstanceActions {
	return &instanceActions{}
}

//Instance Interfae
type IInstanceActions interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Action(string, string, interface{}) error
}

//Instance Request Object
type instanceActions struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceActions) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	m := meta.(*auth.AuthKeys)
	logs.Info("Getting the api and secret keys...!")
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Check the api key and secret key whether not empty
	if apiKey == "" || secretKey == "" {
		logs.Info("ApiKey or SecretKey does not exist...!")
		return diag.FromErr(errors.New("ApiKey or SecretKey does not exist...!"))
	}
	// Action
	err := vs.Action(uuid, action, meta)
	if err != nil {
		return diag.FromErr(err)
	}
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
	// Meta information
	m := meta.(*auth.AuthKeys)
	logs.Info("Getting the api and secret keys...!")
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Check the api key and secret key whether not empty
	if apiKey == "" || secretKey == "" {
		logs.Info("ApiKey or SecretKey does not exist...!")
		return diag.FromErr(errors.New("ApiKey or SecretKey does not exist...!"))
	}
	// Action
	err := vs.Action(uuid, action, meta)
	if err != nil {
		return diag.FromErr(err)
	}
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

// Action Instance
// TODO - Documentation
func (vs *instanceActions) Action(uuid string, action string, meta interface{}) error {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := ""
	switch action {
	case START:
		endPoint = api.GetInstanceStartApi() + "?uuid=" + uuid
	case STOP:
		endPoint = api.GetInstanceStopApi() + "?uuid=" + uuid + "&forceStop=true"
	default:
		return errors.New("Invalid action provided...!")
	}
	logs.Info(`Instance action ` + action + ` initiated...!`)
	// endPoint += "?uuid=" + uuid
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		logs.Info(err.Error())
		return err
	}
	logs.Info(string(response))
	logs.Info(`Instance action ` + action + ` successful...!`)
	return nil
}
