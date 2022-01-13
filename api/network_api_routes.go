package api

import "strconv"

var (
	networkList            = "network/networkList"
	networkOfferingList    = "networkoffering/networkOfferingList"
	vpcNetworkOfferingList = "networkoffering/vpcNetworkOfferingList"
	networkCreate          = "network/createNetwork"
	networkDelete          = "network/deleteNetwork"
	vmNetworkAdd           = "instance/attachNetwork"
	vmNetworkDelete        = "instance/detachNetwork"
)

// Network List api
func GetNewNetworkListApi(zoneUuid string) string {
	return EndPoint + "/" + networkList + "?zoneUuid=" + zoneUuid
}

// Network Offering List api
func GetNewNetworkOfferingListApi(zoneUuid string) string {
	return EndPoint + "/" + networkOfferingList + "?zoneUuid=" + zoneUuid
}

// Network Offering List api
func GetNewVpcNetworkOfferingListApi() string {
	return EndPoint + "/" + vpcNetworkOfferingList
}

// Network Offering List api
func GetNetworkCreateApi(isPublic bool, securityGroupId string) string {
	secrityGroup := ""
	if securityGroupId != "" {
		secrityGroup = "&securityGroupId=" + securityGroupId
	}
	return EndPoint + "/" + networkCreate + "?isPublic=" + strconv.FormatBool(isPublic) + secrityGroup
}

// Network Offering List api
func GetNetworkDeleteApi(uuid string) string {
	return EndPoint + "/" + networkDelete + "/" + uuid
}

// Network Attach api
func GetVmNetworkAddApi() string {
	return EndPoint + "/" + vmNetworkAdd
}

// Network Attach api
func GetVmNetworkDeleteApi() string {
	return EndPoint + "/" + vmNetworkDelete
}
