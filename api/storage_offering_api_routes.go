package api

var (
	storageOfferingList = "storage/storageOfferingList"
)

// Storage Offering List Api
func GetStorageOfferingListApi(zoneId string) string {
	return EndPoint + "/" + storageOfferingList + "?zoneUuid=" + zoneId
}
