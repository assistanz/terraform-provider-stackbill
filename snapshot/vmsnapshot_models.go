package snapshot

// Instance Snapshot Request
type VmSnapshotRequest struct {
	Description        string `json:"description" validate:"required"`
	Name               string `json:"name" validate:"required"`
	SnapshotMemory     bool   `json:"snapshotMemory" validate:"required"`
	VirtualmachineUuid string `json:"virtualmachineUuid" validate:"required"`
	ZoneUuid           string `json:"zoneUuid" validate:"required"`
}
