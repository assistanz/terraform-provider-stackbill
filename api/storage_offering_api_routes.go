package api

var (
	storageOfferingList = "storage/storageOfferingList"
)

// Storage Offering List Api
func GetStorageOfferingListApi(zoneUuid string) string {
	return EndPoint + "/" + storageOfferingList + "?zoneUuid=" + zoneUuid
}
