package api

var (
	instancePostApi       = "instance/createInstance"
	instanceNameUpdateApi = "instance/updateInstanceName"
	instanceStart         = "instance/startInstance"
	instanceStop          = "instance/stopInstance"
	instanceDelete        = "instance/destroyInstance"
	instanceResize        = "instance/resizeVm"
	instanceList          = "instance/instanceList"
	instanceAttachIso     = "instance/attachIso"
	instanceDetachIso     = "instance/detachIso"
	instanceStatus        = "instance/vmStatus"
)

// Instance create api
func GetInstanceCreateApi() string {
	return EndPoint + "/" + instancePostApi
}

// Instance name update api
func GetInstanceNameUpdateApi() string {
	return EndPoint + "/" + instanceNameUpdateApi
}

// Start instance api
func GetInstanceStartApi() string {
	return EndPoint + "/" + instanceStart
}

// Stop instance api
func GetInstanceStopApi() string {
	return EndPoint + "/" + instanceStop
}

// Delete instance api
func GetInstanceDeleteApi() string {
	return EndPoint + "/" + instanceDelete
}

// Instance Resize api
func GetInstanceResizeApi() string {
	return EndPoint + "/" + instanceResize
}

// Instance List api
func GetInstanceListApi(zoneId string) string {
	return EndPoint + "/" + instanceList + "?zoneUuid=" + zoneId
}

// Instace Attach Iso
func GetInstanceAttachIsoApi(uuid string, isoUUid string) string {
	return EndPoint + "/" + instanceAttachIso + "?uuid=" + uuid + "&isoUuid=" + isoUUid
}

// Instace Detach Iso
func GetInstanceDetachIsoApi(uuid string, isoUUid string) string {
	return EndPoint + "/" + instanceDetachIso + "?uuid=" + uuid + "&isoUuid=" + isoUUid
}

// Instance status api
func GetInstanceStatusApi(uuid string) string {
	return EndPoint + "/" + instanceStatus + "?uuid=" + uuid
}
