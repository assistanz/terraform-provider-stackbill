package api

var (
	networkList = "network/networkList"
)

// Network List api
func GetNewNetworkListApi(zoneId string) string {
	return EndPoint + "/" + networkList + "?zoneUuid=" + zoneId
}
