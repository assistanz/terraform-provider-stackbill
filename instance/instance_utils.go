package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance util Object
func NewInstanceUtils() InstanceUtils {
	return &instanceUtils{}
}

//Instance Interfae
type InstanceUtils interface {
	GetCreateInstanceRequest(d *schema.ResourceData) CreateRequest
	GetUpdateInstanceNameRequest(d *schema.ResourceData) UpdateNameRequest
	GetResetSshKeyInstanceRequest(d *schema.ResourceData) ResetSshkeyRequest
	GetResizeInstanceRequest(d *schema.ResourceData) ResizeRequest
	GetInstanceActionRequest(d *schema.ResourceData) InstanceActionRequest
	GetInstanceIsoActionRequest(d *schema.ResourceData) InstanceIsoActionRequest
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
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	updateNameRequest.Uuid = uuid
	return updateNameRequest
}

// Update name request
// TODO - Documentation
func (vs *instanceUtils) GetResetSshKeyInstanceRequest(d *schema.ResourceData) ResetSshkeyRequest {
	var resetSshkeyRequest ResetSshkeyRequest
	resetSshkeyRequest.SshkeyId = d.Get("ssh_key_uuid").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	resetSshkeyRequest.Uuid = uuid
	return resetSshkeyRequest
}

// Resize request
// TODO - Documentation
func (vs *instanceUtils) GetResizeInstanceRequest(d *schema.ResourceData) ResizeRequest {
	var resizeRequest ResizeRequest
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	resizeRequest.Uuid = uuid
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

// Action request
// TODO - Documentation
func (vs *instanceUtils) GetInstanceIsoActionRequest(d *schema.ResourceData) InstanceIsoActionRequest {
	var instanceIsoActionRequest InstanceIsoActionRequest
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	instanceIsoActionRequest.Uuid = uuid
	instanceIsoActionRequest.IsoUuid = d.Get("iso_uuid").(string)
	instanceIsoActionRequest.Action = d.Get("action").(string)
	return instanceIsoActionRequest
}
