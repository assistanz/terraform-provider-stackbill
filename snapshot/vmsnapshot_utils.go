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
	GetVmSnapshotRequest(d *schema.ResourceData) VmSnapshotRequest
}

//Vm Snapshot utils Object
type vmSnapshotUtils struct {
}

// Snapshot Request
// TODO - Documentation
func (vs *vmSnapshotUtils) GetVmSnapshotRequest(d *schema.ResourceData) VmSnapshotRequest {
	var vmSnapshotRequest VmSnapshotRequest
	vmSnapshotRequest.Description = d.Get("description").(string)
	vmSnapshotRequest.Name = d.Get("name").(string)
	vmSnapshotRequest.SnapshotMemory = d.Get("snapshot_memory").(bool)
	vmSnapshotRequest.VirtualmachineUuid = d.Get("virtual_machine_uuid").(string)
	vmSnapshotRequest.ZoneUuid = d.Get("zone_uuid").(string)
	return vmSnapshotRequest
}
