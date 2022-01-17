package snapshot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Vm Snapshot util Object
func NewVmSnapshotUtils() VmSnapshotUtils {
	return &vmSnapshotUtils{}
}

//Vm Snapshot Interfae
type VmSnapshotUtils interface {
	GetVmSnapshotRequest(d *schema.ResourceData) map[string]interface{}
}

//Vm Snapshot utils Object
type vmSnapshotUtils struct {
}

// Snapshot Request
// TODO - Documentation
func (vs *vmSnapshotUtils) GetVmSnapshotRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["description"] = d.Get("description").(bool)
	request["name"] = d.Get("name").(bool)
	request["snapshotMemory"] = d.Get("snapshot_memory").(bool)
	request["virtualmachineUuid"] = d.Get("virtual_machine_uuid").(bool)
	request["zoneUuid"] = d.Get("zone_uuid").(bool)
	return request
}
