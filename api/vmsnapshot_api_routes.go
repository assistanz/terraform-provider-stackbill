package api

var (
	vmSnapShotList   = "vmsnapshot/vmsnapshotList"
	vmSnapshot       = "vmsnapshot/createVmSnapshot"
	vmSnapshotDelete = "vmsnapshot/deleteVmSnapshot"
)

// Instance Snapshot api
func GetVmSnapshotCreateApi() string {
	return EndPoint + "/" + vmSnapshot
}

// Snapshot List api
func GetVmSnapshotListApi(zoneUuid string) string {
	return EndPoint + "/" + vmSnapShotList + "?zoneUuid=" + zoneUuid
}

// Zone List api
func GetVmSnapshotDeleteApi(uuid string) string {
	return EndPoint + "/" + vmSnapshotDelete + "/" + uuid
}
