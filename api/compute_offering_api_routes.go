package api

var (
	computeOfferingList = "compute/computeOfferingList"
)

// Compute Offering List Api
func GetComputeOfferingListApi(zoneId string) string {
	return EndPoint + "/" + computeOfferingList + "?zoneUuid=" + zoneId
}
