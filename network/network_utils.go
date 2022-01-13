package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Network util Object
func NewNetworkUtils() NetworkUtils {
	return &networkUtils{}
}

//	Network Interface
type NetworkUtils interface {
	GetCreateNetworkRequest(d *schema.ResourceData) map[string]interface{}
	GetCreateNetworkOptionalRequest(d *schema.ResourceData) CreateNetworkRequestOptional
	GetNetworkActionRequest(d *schema.ResourceData) NetworkActionRequest
}

// Network utils Object
type networkUtils struct {
}

// Snapshot Request
// TODO - Documentation
func (nt *networkUtils) GetCreateNetworkRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["name"] = d.Get("name").(string)
	request["networkOfferingUuid"] = d.Get("network_offering_uuid").(string)
	if d.Get("virtual_machine_uuid").(string) != "" {
		request["virtualmachineUuid"] = d.Get("virtual_machine_uuid").(string)
	}
	request["zoneUuid"] = d.Get("zone_uuid").(string)
	return request
}

// Snapshot Request
// TODO - Documentation
func (nt *networkUtils) GetCreateNetworkOptionalRequest(d *schema.ResourceData) CreateNetworkRequestOptional {
	var createNetworkRequestOptional CreateNetworkRequestOptional
	createNetworkRequestOptional.IsPublic = d.Get("is_public").(bool)
	createNetworkRequestOptional.SecurityGroupId = d.Get("security_group_id").(string)
	return createNetworkRequestOptional
}

// Action Request
// TODO - Documentation
func (nt *networkUtils) GetNetworkActionRequest(d *schema.ResourceData) NetworkActionRequest {
	var actionNetworkRequest NetworkActionRequest
	actionNetworkRequest.NetworkUuid = d.Get("network_uuid").(string)
	actionNetworkRequest.Uuid = d.Get("virutal_machine_uuid").(string)
	return actionNetworkRequest
}
