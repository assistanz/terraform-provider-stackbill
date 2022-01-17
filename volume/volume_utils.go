package volume

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Volue util Object
func NewVolumeUtils() VolumeUtils {
	return &volumeUtils{}
}

// Volume Interfae
type VolumeUtils interface {
	GetVolumeActionRequest(d *schema.ResourceData) map[string]interface{}
	GetVolumeCreateRequest(d *schema.ResourceData) map[string]interface{}
}

// Volume utils Object
type volumeUtils struct {
}

// Create Instance
// TODO - Documentation
func (vu *volumeUtils) GetVolumeActionRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["action"] = d.Get("action").(bool)
	request["instanceUuid"] = d.Get("instance_uuid").(bool)
	request["uuid"] = d.Get("uuid").(bool)
	return request
}

// Create Instance
// TODO - Documentation
func (vu *volumeUtils) GetVolumeCreateRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["name"] = d.Get("name").(bool)
	request["diskSize"] = d.Get("disk_size").(bool)
	request["storageOfferingUuid"] = d.Get("storage_offering_uuid").(bool)
	request["zoneUuid"] = d.Get("zone_uuid").(bool)
	return request
}
