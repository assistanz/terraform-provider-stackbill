package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance util Object
func NewInstanceUtils() InstanceUtilsI {
	return &instanceUtils{}
}

//Instance Interfae
type InstanceUtilsI interface {
	GetCreateInstanceRequest(d *schema.ResourceData) CreateRequest
	GetUpdateInstanceNameRequest(d *schema.ResourceData) UpdateNameRequest
	GetResizeInstanceRequest(d *schema.ResourceData) ResizeRequest
	GetInstanceActionRequest(d *schema.ResourceData) InstanceActionRequest
}

//Instance utils Object
type instanceUtils struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceUtils) GetCreateInstanceRequest(d *schema.ResourceData) CreateRequest {
	var createRequest CreateRequest
	createRequest.ComputeOfferingUuid = d.Get("compute_offering_uuid").(string)
	createRequest.CpuCore = d.Get("cpu_core").(string)
	createRequest.DiskSize = d.Get("disk_size").(int)
	createRequest.HypervisorName = d.Get("hypervisor_name").(string)
	createRequest.Memory = d.Get("memory").(string)
	createRequest.Name = d.Get("name").(string)
	createRequest.NetworkUuid = d.Get("network_uuid").(string)
	createRequest.SecuritygroupName = d.Get("security_group_name").(string)
	createRequest.SshKeyName = d.Get("ssh_key_name").(string)
	createRequest.StorageOfferingUuid = d.Get("storage_offering_uuid").(string)
	createRequest.TemplateUuid = d.Get("template_uuid").(string)
	createRequest.ZoneUuid = d.Get("zone_uuid").(string)
	return createRequest
}

// Update name request
// TODO - Documentation
func (vs *instanceUtils) GetUpdateInstanceNameRequest(d *schema.ResourceData) UpdateNameRequest {
	var updateNameRequest UpdateNameRequest
	updateNameRequest.Name = d.Get("name").(string)
	updateNameRequest.Uuid = d.Id()
	return updateNameRequest
}

// Resize request
// TODO - Documentation
func (vs *instanceUtils) GetResizeInstanceRequest(d *schema.ResourceData) ResizeRequest {
	var resizeRequest ResizeRequest
	resizeRequest.Uuid = d.Id()
	resizeRequest.OfferingUuid = d.Get("compute_offering_uuid").(string)
	resizeRequest.CpuCore = d.Get("cpu_core").(string)
	resizeRequest.Memory = d.Get("memory").(string)
	return resizeRequest
}

// Action request
// TODO - Documentation
func (vs *instanceUtils) GetInstanceActionRequest(d *schema.ResourceData) InstanceActionRequest {
	var instanceActionRequest InstanceActionRequest
	instanceActionRequest.Uuid = d.Id()
	instanceActionRequest.Action = d.Get("action").(string)
	return instanceActionRequest
}
