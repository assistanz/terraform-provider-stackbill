package volume

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Volue util Object
func NewVolumeUtils() VolumeUtils {
	return &volumeUtils{}
}

// Volume Interfae
type VolumeUtils interface {
	GetVolumeActionRequest(d *schema.ResourceData) VolumeActionRequest
	GetVolumeCreateRequest(d *schema.ResourceData) VolumeCreateRequest
}

// Volume utils Object
type volumeUtils struct {
}

// Create Instance
// TODO - Documentation
func (vu *volumeUtils) GetVolumeActionRequest(d *schema.ResourceData) VolumeActionRequest {
	var volumeActionRequest VolumeActionRequest
	volumeActionRequest.Action = d.Get("action").(string)
	volumeActionRequest.InstanceUuid = d.Get("instance_uuid").(string)
	volumeActionRequest.Uuid = d.Get("uuid").(string)
	return volumeActionRequest
}

// Create Instance
// TODO - Documentation
func (vu *volumeUtils) GetVolumeCreateRequest(d *schema.ResourceData) VolumeCreateRequest {
	var volumeCreateRequest VolumeCreateRequest
	volumeCreateRequest.Name = d.Get("name").(string)
	volumeCreateRequest.DiskSize = d.Get("disk_size").(int)
	volumeCreateRequest.StorageOfferingUuid = d.Get("storage_offering_uuid").(string)
	volumeCreateRequest.ZoneUuid = d.Get("zone_uuid").(string)
	return volumeCreateRequest
}
