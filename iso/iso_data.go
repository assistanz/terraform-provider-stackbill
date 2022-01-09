package iso

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

// Template Object
func NewIsoData() IsoDataI {
	return &isoData{}
}

// Template Interface
type IsoDataI interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Iso data object
type isoData struct {
}

// List Templates
// TODO - Documentation
func (co *isoData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneId := d.Get("zone_id").(string)
	uuid := d.Get("uuid").(string)
	logs.Info("Iso list initiated...!")
	response, err := isoApiObj.ListIso(zoneId, uuid, meta)
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
	datas := jsonRes["listTemplateResponse"].([]interface{})
	output := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		format := v["format"].(string)
		if format != "ISO" {
			continue
		}
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		output = append(output, i)
	}
	if err := d.Set("iso", output); err != nil {
		return diag.FromErr(err)
	}
	logs.Info("Iso list successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
