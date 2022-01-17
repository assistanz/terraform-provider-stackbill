package api

var (
	volumeList   = "volume/volumeList"
	attachVolume = "volume/attachVolume"
	detachVolume = "volume/detachVolume"
	createVolume = "volume/createVolume"
	deleteVolume = "volume/deleteVolume"
)

// Volume List api
func GetVolumeListApi(zoneUuid string) string {
	return EndPoint + "/" + volumeList + "?zoneUuid=" + zoneUuid
}

// Volume Attach api
func GetVolumeAttachApi(uuid string) string {
	return EndPoint + "/" + attachVolume + "?uuid=" + uuid
}

// Volume Detach api
func GetVolumeDetachApi(uuid string) string {
	return EndPoint + "/" + detachVolume + "?uuid=" + uuid
}

// Volume Create api
func GetVolumeCreateApi() string {
	return EndPoint + "/" + createVolume
}

// Volume Delete api
func GetVolumeDeleteApi(uuid string) string {
	return EndPoint + "/" + deleteVolume + "/" + uuid
}
