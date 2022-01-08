package instance

import (
	"context"
	"encoding/json"
	"strconv"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	logs "github.com/sirupsen/logrus"
)

// Instance Object
func NewInstanceData() InstanceDataI {
	return &instanceData{}
}

//Instance Interfae
type InstanceDataI interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Instance Request Object
type instanceData struct {
}

// List Instance
// TODO - Documentation
func (id *instanceData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneId := d.Get("zone_id").(string)
	uuid := d.Get("uuid").(string)
	logs.Info("Instance list initiated...!")
	response, err := instanceApiObj.ListInstances(zoneId, uuid, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	logs.Info(response)
	jsonRes := make(map[string]interface{})
	if err := json.Unmarshal([]byte(response), &jsonRes); err != nil {
		return diag.FromErr(err)
	}
	count := jsonRes["count"].(float64)
	if err := d.Set("length", count); err != nil {
		return diag.FromErr(err)
	}
	datas := jsonRes["listInstanceResponse"].([]interface{})
	output := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		output = append(output, i)
	}
	if err := d.Set("instances", output); err != nil {
		return diag.FromErr(err)
	}
	logs.Info("List instances successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}